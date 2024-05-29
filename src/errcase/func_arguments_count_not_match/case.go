package func_arguments_count_not_match

import (
	"github.com/gin-gonic/gin"

	"github.com/ZyhHelen/autodebug-dataset/src/errcase/func_arguments_count_not_match/handlers"
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
