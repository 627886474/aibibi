package user

import (
	"github.com/gin-gonic/gin"
	"github.com/zl/aibibi/controller/common"
	"strings"
	"github.com/zl/aibibi/utils"
	"regexp"
	"github.com/zl/aibibi/model"
	"github.com/zl/aibibi/config"
	"net/http"
	"github.com/gin-gonic/gin/binding"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"errors"
	"github.com/garyburd/redigo/redis"
	"crypto/md5"
	"github.com/zl/aibibi/controller/mail"
	"time"
)

const(
	resetDuration = 24*60*60
)

//邮件格式
func sendMail(action string,title string,curTime int64,user model.User,c *gin.Context){
	siteName := config.ServerConfig.SiteName
	sitURL := "http://"+config.ServerConfig.Host
	secretStr := fmt.Sprintf("%d%s%s",curTime,user.Email,user.PassWord)
	secretStr = fmt.Sprintf("%x",md5.Sum([]byte(secretStr)))
	actionURL := sitURL +action+"/%d/%s"

	actionURL = fmt.Sprintf(actionURL,user.ID,secretStr)
	fmt.Println(actionURL)

	content := "<p><b>亲爱的" + user.Name + ":</b></p>" +
		"<p>我们收到您在 " + siteName + " 的注册信息, 请点击下面的链接, 或粘贴到浏览器地址栏来激活帐号.</p>" +
		"<a href=\"" + actionURL + "\">" + actionURL + "</a>" +
		"<p>如果您没有在 " + siteName + " 填写过注册信息, 说明有人滥用了您的邮箱, 请删除此邮件, 我们对给您造成的打扰感到抱歉.</p>" +
		"<p>" + siteName + " 谨上.</p>"

	if action == "/reset" {
		content = "<p><b>亲爱的" + user.Name + ":</b></p>" +
			"<p>你的密码重设要求已经得到验证。请点击以下链接, 或粘贴到浏览器地址栏来设置新的密码: </p>" +
			"<a href=\"" + actionURL + "\">" + actionURL + "</a>" +
			"<p>感谢你对" + siteName + "的支持，希望你在" + siteName + "的体验有益且愉快。</p>" +
			"<p>(这是一封自动产生的email，请勿回复。)</p>"
	}
	content += "<p><img src=\"" + sitURL + "/images/logo.png\" style=\"height: 42px;\"/></p>"
	//fmt.Println(content)

	mail.SendMail(user.Email, title, content)
}
//校验链接地址
func verifyLink(cacheKey string,c *gin.Context)(model.User,error){
	var user model.User
	userID,err :=strconv.Atoi(c.Param("id"))
	if err !=nil ||userID <=0{
		return user,errors.New("无效链接")
	}
	secrect := c.Param("secret")
	if secrect ==""{
		return user,errors.New("无效的链接")
	}

	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()

	emailTime,redisErr := redis.Int64(RedisConn.Do("GET",cacheKey+fmt.Sprintf("%d", userID)))
	if redisErr !=nil{
		return user,errors.New("无效链接")
	}
	if err := model.DB.First(&user,userID).Error;err !=nil{
		return user,errors.New("无效链接")
	}

	secrectStr := fmt.Sprintf("%d%s",emailTime,user.Email)
	secrectStr = fmt.Sprintf("%x",md5.Sum([]byte(secrectStr)))
	if secrect != secrectStr{
		fmt.Println(secrect,secrectStr)
		return user,errors.New("无效链接")
	}
	return user,nil

}

//重置密码
func ResetPassword(c *gin.Context){
	SendErrJSON :=common.SendErrJSON
	type UserReqData struct {
		Password  string 	`json:"password" binding:"required,min=6,max=20"`
	}
	var userData UserReqData

	if err:=c.ShouldBindJSON(&userData);err !=nil{
		SendErrJSON("参数无效",c)
		return
	}
	var verifyErr error
	var user model.User
	if user,verifyErr =verifyLink(model.ResetTime,c);verifyErr !=nil{
		SendErrJSON("重置链接已失效",c)
		return
	}

	user.PassWord = user.EncryptPassword(userData.Password,user.Salt())

	if user.ID <=0{
		SendErrJSON("重置链接已失效",c)
		return
	}
	if err := model.DB.Model(&user).Update("pass_word",user.PassWord).Error;err!=nil{
		SendErrJSON("error",c)
	}

	redisConn :=model.RedisPool.Get()
	defer redisConn.Close()
	if _,err:= redisConn.Do("DEL",fmt.Sprintf("%s%d",model.ResetTime,user.ID));err!=nil{
		SendErrJSON("redis delete failed",err)
	}

	c.JSON(http.StatusOK,gin.H{
		"errNo": model.ErrorCode.SUCCESS,
		"msg":   "success",
		"data":  gin.H{},
	})
}

//重置密码的邮件
func ResetPasswordMail(c *gin.Context){
	SendErrJSON := common.SendErrJSON
	type UserReqData struct {
		Email 		string 	`json:"email" binding:"required,email"`
	}
	var userData UserReqData
	if err := c.ShouldBindJSON(&userData);err != nil{
		SendErrJSON("无效的邮箱",c)
		return
	}

	var user model.User
	if err := model.DB.Where("email=?",userData.Email).Find(&user).Error;err !=nil{
		SendErrJSON("没有邮箱为"+user.Email+"的用户",c)
		return
	}
	cureTime := time.Now().Unix()
	resetUser := fmt.Sprintf("%s%d",model.ResetTime,user.ID)

	RedisConn := model.RedisPool.Get()
	defer RedisConn.Close()

	if _,err :=RedisConn.Do("SET",resetUser,cureTime,"EX",resetDuration);err !=nil{
		SendErrJSON("redis set failed",c)
		return
	}

	go func() {
		sendMail("/ac","修改密码",cureTime,user,c)
	}()

	c.JSON(http.StatusOK,gin.H{
		"errNo":model.ErrorCode.SUCCESS,
		"message":"success",
		"data":gin.H{},
	})
}

//验证重置密码的链接是否失效
func VerifyResetPasswordLink(c *gin.Context){
	SendErrJOSN := common.SendErrJSON
	if _,err :=verifyLink(model.ResetTime,c);err !=nil{
		fmt.Println(err.Error())
		SendErrJOSN("重置链接已失效",c)
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"errNo":model.ErrorCode.SUCCESS,
		"msg":"sucess",
		"data":gin.H{},
	})
}

//用户注册
func Register(c *gin.Context){
	SendErrJSON := common.SendErrJSON
	type UserReqData struct{
		Name 	string 	`json:"name" binding:"required"`
		PassWord string	`json:"pass_word" binding:"required"`
		Email 	string 	`json:"email" binding:"required"`
		//CaptchaId string `json:"captcha_id"`
	}

	var userData UserReqData
	if err := c.ShouldBindJSON(&userData);err != nil{
		SendErrJSON("参数错误",c)
		return
	}
	userData.Name = utils.AvoidXSS(userData.Name) //避免xss攻击
	userData.Name = strings.TrimSpace(userData.Name) //消除空格
	userData.Email = strings.TrimSpace(userData.Email)

	if strings.Index(userData.Name,"@") != -1{
		SendErrJSON("昵称不能包含@符号",c)
		return
	}
	if ok,err := regexp.MatchString(config.RegexpNickName,userData.Name);userData.Name == "" || !ok || err !=nil{
		SendErrJSON("昵称格式不正确",c)
		return
	}
	if ok,err := regexp.MatchString(config.RegexpEmail,userData.Email);userData.Email == "" || !ok || err !=nil{
		SendErrJSON("邮箱不正确",c)
		return
	}

	//v, ok := c.GetSession(config.CaptchaSessionName).(string)
	//if !ok || !strings.EqualFold(v, captcha) {
	//	c.JsonResult(6001, "验证码不正确")
	//}

	var user model.User
	if err :=model.DB.Where("name = ? OR mobile = ?",userData.Name,userData.Email).Find(&user).Error;err ==nil{
		if user.Name == userData.Name{
			SendErrJSON("昵称已被占用",c)
			return
		}
		if user.Mobile == userData.Email{
			SendErrJSON("邮箱已注册",c)
			return
		}
	}

	var newUser model.User
	newUser.Name = userData.Name
	newUser.Email = userData.Email
	newUser.AvatarURL = "/root/images/avatar/default.png"   //设置一个默认头像
	newUser.Role = model.UserRoleNormal     //默认为 普通用户   1
	newUser.Sex = model.UserSexMale    	  //默认为  男性     0
	newUser.PassWord = newUser.EncryptPassword(userData.PassWord,newUser.Salt())

	if err :=model.DB.Create(&newUser).Error;err!=nil{
		SendErrJSON("error",c)
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"errNo":model.ErrorCode.SUCCESS,
		"msg":"success",
		"data":newUser,
	})
}

//更新用户信息
func UpdateInfo(c *gin.Context){
	sendErrJSON :=common.SendErrJSON
	var userReqData model.User

	if err :=c.ShouldBindWith(&userReqData,binding.JSON);err !=nil{
		sendErrJSON("参数无效",c)
	}

	userInter,_ := c.Get("user")
	user := userInter.(model.User)

	field :=c.Param("field")
	resData := make(map[string]interface{})
	resData["id"] = user.ID

	switch field {
	case "sex":
		if userReqData.Sex != model.UserSexMale && userReqData.Sex != model.UserSexFemale{
			sendErrJSON("无效的性别",c)
			return
		}
		if err := model.DB.Model(&user).Update("sex",userReqData.Sex).Error;err !=nil{
			fmt.Println(err.Error())
			sendErrJSON("error",c)
			return
		}
		resData[field] = userReqData.Sex
	case "mobile":
		userReqData.Mobile = utils.AvoidXSS(userReqData.Mobile)
		userReqData.Mobile = strings.TrimSpace(userReqData.Mobile)
		if ok,err := regexp.MatchString(config.RegexpPhone,userReqData.Mobile);userReqData.Mobile==""||!ok||err !=nil{
			sendErrJSON("手机格式不正确",c)
			return
		}
		if err := model.DB.Where("mobile = ?",userReqData.Mobile).Find(&userReqData).Error;err ==nil{
			sendErrJSON("手机号已被注册",c)
			return
		}
		if err := model.DB.Model(&user).Update("mobile",userReqData.Mobile).Error;err !=nil {
			fmt.Println(err.Error())
			sendErrJSON("error",c)
			return
		}
		resData[field] = userReqData.Mobile
	case "email":
		userReqData.Email = utils.AvoidXSS(userReqData.Email)
		userReqData.Email = strings.TrimSpace(userReqData.Mobile)
		if ok,err := regexp.MatchString(config.RegexpEmail,userReqData.Email);userReqData.Email ==""||!ok||err !=nil{
			sendErrJSON("邮箱格式不正确",c)
			return
		}
		if err :=model.DB.Where("email = ?",userReqData.Email).Find(&userReqData).Error;err ==nil{
			sendErrJSON("邮箱已被注册",c)
			return
		}
		if err :=model.DB.Model(&user).Update("email",userReqData.Email).Error;err !=nil{
			fmt.Println(err.Error())
			sendErrJSON("error",c)
			return
		}
		resData[field] = userReqData.Email
	default:
		sendErrJSON("参数无效",c)
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"errNo":model.ErrorCode.SUCCESS,
		"msg":"success",
		"data":resData,
	})
}

//修改密码
func UpdatePassword(c *gin.Context){
	sendErrJSON :=common.SendErrJSON
	type userReqData struct{
		Pass string  	`json:"pass" binding:"required,min=6,max=20"`
		NewPwd string 	`json:"new_pwd" binding:"required,min=6,max=20"`
	}
	var userData userReqData
	if err :=c.ShouldBindWith(&userData,binding.JSON);err !=nil{
		sendErrJSON("参数无效",c)
		return
	}

	userInter,_ :=c.Get("user")
	user := userInter.(model.User)

	if err := model.DB.First(&user,user.ID).Error;err !=nil{
		sendErrJSON("error",c)
		return
	}

	if user.CheckPassword(userData.Pass){
		user.PassWord =user.EncryptPassword(userData.NewPwd,user.Salt())
		if err := model.DB.Model(&user).Update("pass_word",user.PassWord).Error;err !=nil{
			sendErrJSON("原密码不正确",c)
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"errNo":model.ErrorCode.SUCCESS,
			"msg":"success",
			"data":gin.H{},
		})
	}else {
		sendErrJSON("原密码错误",c)
		return
	}
}

//用户登录
func Signin(c *gin.Context){
	SendErrJSON :=common.SendErrJSON
	type EmailLogin struct {
		SigninInput 	string `json:"sigin_input" binding:"required"`
		Password 	string `json:"password" binding:"required,min=6,max=20"`
	}
	//type UserNameLogin struct {
	//	SigninInput 	string `json:"sigin_input" binding:"required"`
	//	PassWord 	string `json:"pass_word" binding:"required,min=6,max=20"`
	//}

	var emailLogin EmailLogin
	//var usernameLogin UserNameLogin
	var signinInput string
	var passWord string
	var sql string

	//if c.Query("loginType") =="mobile"{
	if err :=c.ShouldBindWith(&emailLogin,binding.JSON);err !=nil{
		fmt.Println(err.Error())
		SendErrJSON("参数错误",c)
		return
	}

	signinInput = emailLogin.SigninInput
	passWord = emailLogin.Password
	sql = "email = ?"
		//}else if c.Query("loginType") == "username"{  //通过前端传递的参数判断登陆方式，这里模拟接口请求直接name登陆
	//}else {
	//	if err :=c.ShouldBindWith(&usernameLogin,binding.JSON);err !=nil{
	//		fmt.Println(err.Error())
	//		SendErrJSON("用户名或密码错误",c)
	//		return
	//	}
	//	signinInput = usernameLogin.SigninInput
	//	passWord = usernameLogin.PassWord
	//	sql = "name = ?"
	//}

	var user model.User
	if err := model.DB.Debug().Where(sql,signinInput).First(&user).Error;err !=nil{
		SendErrJSON("账户不存在",c)
		return
	}

	if user.CheckPassword(passWord){
		// token认证
		token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
			"id":user.ID,
		})
		tokenString ,err := token.SignedString([]byte(config.ServerConfig.TokenSecret))
		if err !=nil{
			fmt.Println(err.Error())
			SendErrJSON("未知错误",c)
			return
		}
		//存入redis
		if err := model.UserToredis(user);err !=nil{
			SendErrJSON("内部错误",c)
			return
		}

		c.SetCookie("token",tokenString,config.ServerConfig.TokenMaxAge,"/","",true,true)

		c.JSON(http.StatusOK,gin.H{
			"errNo":model.ErrorCode.SUCCESS,
			"msg":"success",
			"data":gin.H{
				"token":tokenString,
				"user":user,
			},
		})
		return
	}
	SendErrJSON("账号或密码错误",c)
}