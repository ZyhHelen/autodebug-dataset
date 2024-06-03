package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/ZyhHelen/autodebug-dataset/handlers"
)

func main() {
	r := gin.Default()
	port := 8080

	log.Print("server start ..")

	// 定义路由
	r.GET("/info", handlers.Info)
	r.GET("/ping", handlers.Ping)
	r.GET("/user/:name", handlers.GetUser)

	domain := "http://localhost:8080"
	log.Print()
	log.Printf("GET /info        --> %s/info\n", domain)
	log.Printf("GET /ping        --> %s/ping\n", domain)
	log.Printf("GET /user/:name  --> %s/user/:name\n", domain)
	log.Print()

	// 运行服务器
	r.Run(fmt.Sprintf(":%d", port))
}
