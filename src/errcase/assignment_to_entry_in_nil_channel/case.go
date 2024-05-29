package assignment_to_entry_in_nil_channel

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/ZyhHelen/autodebug-dataset/src/errcase/assignment_to_entry_in_nil_channel/handlers"
)

// startLogger 启动日志记录器
func startLogger() chan string {
	var logChan chan string

	go func() {
		for {
			msg := <-logChan              // 从日志通道接收日志信息
			log.Println("[Logger]:", msg) // 打印日志信息
		}
	}()

	return logChan
}

func Run() {
	logChan := startLogger()
	logChan <- "start"

	defer func() {
		logChan <- "exit"
	}()

	r := gin.Default()

	// 定义路由
	r.GET("/ping", handlers.Ping)
	r.POST("/submit", handlers.Submit)
	r.GET("/user/:name", handlers.GetUser)

	// 运行服务器
	r.Run(":8080")

	logChan <- "exit"
}
