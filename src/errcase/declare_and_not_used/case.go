package declare_and_not_used

import (
	"github.com/gin-gonic/gin"

	"github.com/ZyhHelen/autodebug-dataset/src/errcase/declare_and_not_used/handlers"
)

func Run() {
	r := gin.Default()

	var name string

	// 定义路由
	r.GET("/ping", handlers.Ping)
	r.POST("/submit", handlers.Submit)
	r.GET("/user/:name", handlers.GetUser)

	// 运行服务器
	r.Run(":8080")
}
