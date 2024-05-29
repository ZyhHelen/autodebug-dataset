package assignment_to_entry_in_nil_map

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/ZyhHelen/autodebug-dataset/src/errcase/assignment_to_entry_in_nil_map/handlers"
	"github.com/ZyhHelen/autodebug-dataset/src/errcase/assignment_to_entry_in_nil_map/utils"
)

func Run() {
	r := gin.Default()

	log.Println(utils.GetInfoMap())

	// 定义路由
	r.GET("/ping", handlers.Ping)
	r.POST("/submit", handlers.Submit)
	r.GET("/user/:name", handlers.GetUser)

	// 运行服务器
	r.Run(":8080")
}
