package main

import (
	"fmt"
	"log"

	"github.com/ZyhHelen/autodebug-dataset/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New() // 创建一个新的Echo实例
	port := 8080

	log.Print("server start ..")

	// 定义路由
	e.GET("/info", handlers.Info)
	e.GET("/ping", handlers.Ping)
	e.GET("/user/:name", handlers.GetUser)

	// 注意：POST路由应该这样定义
	e.POST("/submit", handlers.Submit)

	domain := "http://localhost:" + fmt.Sprintf("%d", port)
	log.Print()
	log.Printf("GET /info        --> %s/info\n", domain)
	log.Printf("GET /ping        --> %s/ping\n", domain)
	log.Printf("GET /user/:name  --> %s/user/:name\n", domain)
	log.Printf("POST /submit     --> %s/submit\n", domain) // 添加Submit路由
	log.Print()

	// 运行服务器
	if err := e.Start(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
