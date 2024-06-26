package main

import (
	"fmt"
	"gin-template/config"
	"gin-template/global"
	"gin-template/initialize"
	"gin-template/routing"
	"github.com/gin-gonic/gin"
)

func main() {

	initialize.ConnectDB()
	global.Log = initialize.InitZap() // 初始化zap日志库

	router := gin.Default()
	router.Use(gin.Logger())
	routing.Setup(router)
	err := router.Run(fmt.Sprintf(":%s", config.Config.Port))
	if err != nil {
		return
	}
}
