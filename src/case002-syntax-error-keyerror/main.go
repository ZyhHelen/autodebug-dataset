package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ZyhHelen/autodebug-dataset/src/case002-syntax-error-keyerror/handlers"
)

func main() {
	r := gin.Default()

	// 定义路由
	r.GET("/ping", handlers.Ping)
	r.POST("/submit", handlers.Submit)
	r.GET("/user/:name", handlers.GetUser)

	// 运行服务器
	r.Run(":8080")
}
