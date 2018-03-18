package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zl/aibibi/model"
	"github.com/zl/aibibi/config"
)

//刷新过期时间
func RefreshTokenCookie(c *gin.Context){
	tokenString,err := c.Cookie("token")
	if tokenString != ""&&err ==nil{
		c.SetCookie("token",tokenString,config.ServerConfig.TokenMaxAge,"/","",true,true)
		if user,err := getUser(c);err ==nil{
			model.UserToredis(user)
		}
	}
	c.Next()
}
