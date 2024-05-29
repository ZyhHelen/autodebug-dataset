package yntax_error_forget_left_double_quote

import (
	"github.com/gin-gonic/gin"

	"github.com/ZyhHelen/autodebug-dataset/src/errcase/syntax_error_forget_left_double_quote/handlers"
)

func Run() {
	r := gin.Default()

	// 定义路由
	r.GET("/ping", handlers.Ping)
	r.POST("/submit", handlers.Submit)
	r.GET("/user/:name", handlers.GetUser)

	// 运行服务器
	r.Run(":8080")
}
