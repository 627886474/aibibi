package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zl/aibibi/model"
	"os"
	"io"
	"github.com/zl/aibibi/router"
	"github.com/zl/aibibi/config"
)

func main(){
	fmt.Println("gin.Version: ",gin.Version)
	if config.ServerConfig.Env != model.DevelopmentMode{
		//非开发模式，生成日志文件
		logFile, err := os.OpenFile(config.ServerConfig.LogFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			fmt.Printf(err.Error())
			os.Exit(-1)
		}
		gin.DefaultWriter = io.MultiWriter(logFile)
	}
	app :=gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	router.Route(app)
	app.Run(":"+fmt.Sprintf("%d",config.ServerConfig.Port))

}