package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ZyhHelen/autodebug-dataset/src/case001-declare-and-not-used/handlers"
)

func main() {
	r := gin.Default()

	var name string

	// 定义路由
	r.GET("/ping", handlers.Ping)
	r.POST("/submit", handlers.Submit)
	r.GET("/user/:name", handlers.GetUser)

	// 运行服务器
	r.Run(":8080")
}
