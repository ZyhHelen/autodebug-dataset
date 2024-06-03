package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/ZyhHelen/autodebug-dataset/handlers"
	"github.com/ZyhHelen/autodebug-dataset/utils"
)

func main() {
	r := gin.Default()

	log.Println(utils.GetInfoMap())

	// 定义路由
	r.GET("/ping", handlers.Ping)
	r.POST("/submit", handlers.Submit)
	r.GET("/user/:name", handlers.GetUser)

	// 运行服务器
	r.Run(":8080")
}
