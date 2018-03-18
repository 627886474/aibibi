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
)

//用户注册
func Register(c *gin.Context){
	SendErrJSON := common.SendErrJSON
	type UserReqData struct{
		Name 	string 	`json:"name" binding:"required"`
		PassWord string	`json:"pass_word" binding:"required"`
		Mobile 	string 	`json:"mobile" binding:"required"`
	}

	var userData UserReqData
	if err := c.ShouldBindJSON(&userData);err != nil{
		SendErrJSON("参数错误",c)
		return
	}
	userData.Name = utils.AvoidXSS(userData.Name) //避免xss攻击
	userData.Name = strings.TrimSpace(userData.Name) //消除空格
	userData.Mobile = strings.TrimSpace(userData.Mobile)

	if strings.Index(userData.Name,"@") != -1{
		SendErrJSON("昵称不能包含@符号",c)
		return
	}
	if ok,err := regexp.MatchString(config.RegexpNickName,userData.Name);userData.Name == "" || !ok || err !=nil{
		SendErrJSON("昵称格式不正确",c)
		return
	}
	if ok,err := regexp.MatchString(config.RegexpPhone,userData.Mobile);userData.Mobile == "" || !ok || err !=nil{
		SendErrJSON("手机格式不正确",c)
		return
	}

	var user model.User
	if err :=model.DB.Where("name = ? OR mobile = ?",userData.Name,userData.Mobile).Find(&user).Error;err ==nil{
		if user.Name == userData.Name{
			SendErrJSON("昵称已被占用",c)
			return
		}
		if user.Mobile == userData.Mobile{
			SendErrJSON("手机号已注册",c)
			return
		}
	}

	var newUser model.User
	newUser.Name = userData.Name
	newUser.Mobile = userData.Mobile
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
	type MobileNameLogin struct {
		SigninInput 	string `json:"signin_input" binding:"required"`
		Password 	string `json:"password" binding:"required,min=6,max=20"`
	}
	type UserNameLogin struct {
		SigninInput 	string `json:"sigin_input" binding:"required"`
		PassWord 	string `json:"pass_word" binding:"required,min=6,max=20"`
	}

	var mobileLogin MobileNameLogin
	var usernameLogin UserNameLogin
	var signinInput string
	var passWord string
	var sql string

	if c.Query("loginType") =="mobile"{
		if err :=c.ShouldBindWith(&mobileLogin,binding.JSON);err !=nil{
			fmt.Println(err.Error())
			SendErrJSON("手机或密码错误",c)
			return
		}
		signinInput = mobileLogin.SigninInput
		passWord = mobileLogin.Password
		sql = "mobile = ?"
		//}else if c.Query("loginType") == "username"{  //通过前端传递的参数判断登陆方式，这里模拟接口请求直接name登陆
	}else {
		if err :=c.ShouldBindWith(&usernameLogin,binding.JSON);err !=nil{
			fmt.Println(err.Error())
			SendErrJSON("用户名或密码错误",c)
			return
		}
		signinInput = usernameLogin.SigninInput
		passWord = usernameLogin.PassWord
		sql = "name = ?"
	}

	var user model.User
	if err := model.DB.Where(sql,signinInput).First(&user).Error;err !=nil{
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