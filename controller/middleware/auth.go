package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zl/aibibi/model"
	jwt "github.com/dgrijalva/jwt-go"
	"errors"
	"fmt"
	"github.com/zl/aibibi/config"
	"github.com/zl/aibibi/controller/common"
)

func getUser(c *gin.Context) (model.User,error){
	var user model.User
	tokenString,cookieErr :=c.Cookie("token")

	if cookieErr !=nil{
		return user,errors.New("未登录")
	}

	token,tokenErr :=jwt.Parse(tokenString,func(token *jwt.Token)(interface{},error){
		if _,ok :=token.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil,fmt.Errorf("Unexpected signing method:%v",token.Header["alg"])
		}
		return []byte(config.ServerConfig.TokenSecret),nil
	})

	if tokenErr !=nil{
		return user,errors.New("未登录")
	}

	if claims,ok :=token.Claims.(jwt.MapClaims);ok&&token.Valid{
		userId := int(claims["id"].(float64))
		var err error
		user,err = model.UserFromRedis(userId)
		if err != nil{
			return user,errors.New("未登录")
		}
		return user,nil
	}
	return user,errors.New("未登录")
}

// 必须是登陆用户
func SigninRequired(c *gin.Context){
	SendErrJSON := common.SendErrJSON
	var user model.User
	var err error
	if user,err = getUser(c);err !=nil{
		SendErrJSON("未登录",model.ErrorCode.LoginTimeout,c)
		return
	}
	c.Set("user",user)
	c.Next()
}

//必须是管理员
func AdminRequired(c *gin.Context){
	SendErrJSON := common.SendErrJSON
	var user model.User
	var err error
	if user,err = getUser(c);err !=nil{
		SendErrJSON("未登录",model.ErrorCode.LoginTimeout,c)
		return
	}
	if user.Role == model.UserRoleAdmin {
		c.Set("user",user)
		c.Next()
	}else{
		SendErrJSON("没有权限",c)
	}
}