package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zl/aibibi/controller/middleware"
	"github.com/zl/aibibi/config"
	"github.com/zl/aibibi/controller/user"
	"github.com/zl/aibibi/controller/category"
	"github.com/zl/aibibi/controller/argue"
	"github.com/zl/aibibi/controller/talk"
)

func Route(router *gin.Engine){
	apiPrefix :=config.ServerConfig.APIPrefix
	v1_api:= router.Group(apiPrefix+"/v1")
	{
		v1_api.POST("/user/register",user.Register)  //注册
		v1_api.GET("/captch",middleware.LoadVerify)//验证码
		//v1_api.GET("/captch2",middleware.CaptchaController{}.Server)
		v1_api.POST("/user/login",user.Signin)//登录
		v1_api.PUT("/user/update/:field",middleware.SigninRequired,
			user.UpdateInfo) //修改用户信息
		v1_api.PUT("/user/password/update",middleware.SigninRequired,
			user.UpdatePassword)//修改密码
		v1_api.GET("/category/info/:id", category.Info)//通过id获取一个分类的详情
		v1_api.GET("/category/list", category.List)//获取所有分类
		v1_api.POST("/argue/add",middleware.SigninRequired,
			argue.Create) //添加话题分类
		v1_api.PUT("/argue/update/:id",middleware.SigninRequired,
			argue.Update) //更新话题分类
		v1_api.POST("/talk/add",middleware.SigninRequired,
			talk.Save) //添加评论
		v1_api.PUT("/talk/up",middleware.SigninRequired,
			talk.UP) //点赞评论
	}

	v1_admin_api := router.Group(apiPrefix+"/admin/v1") //管理员访问的路由
	{
		v1_admin_api.POST("/category/add",middleware.AdminRequired,
			category.Create)
		v1_admin_api.PUT("/category/update",middleware.AdminRequired,
			category.Update)

	}
}