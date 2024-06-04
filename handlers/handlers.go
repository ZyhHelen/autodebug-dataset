package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ZyhHelen/autodebug-dataset/models"
	"github.com/ZyhHelen/autodebug-dataset/utils"
)

func Info(c *gin.Context) {
	log.Printf("[INFO] <Info> receive request from: %s ", c.ClientIP())

	// get info
	infoMap := utils.GetInfoMap()

	// format info
	infoMsg := ""
	for k, v := range infoMap {
		infoMsg += fmt.Sprintf("%v\t\t:%v\n", k, v)
	}

	// return json response
	c.JSON(http.StatusOK, gin.H{
		"message": infoMsg,
	})
}

// Ping handler
func Ping(c *gin.Context) {
	log.Printf("[INFO] <Ping> receive request from: %s ", c.ClientIP())

	// return response
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// Submit handler
func Submit(c *gin.Context) {
	log.Printf("[INFO] <Submit> receive request from: %s ", c.ClientIP())

	name := c.PostForm("name")
	email := c.PostForm("email")

	// 模拟将用户保存到数据库或其他地方
	user := models.User{Name: name, Email: email}

	// return json response
	c.JSON(http.StatusOK, gin.H{
		"status": "submitted",
		"user":   user,
	})
}

// GetUser handler
func GetUser(c *gin.Context) {
	log.Printf("[INFO] <GetUser> receive request from: %s ", c.ClientIP())

	name := c.Param("name")

	// 模拟从数据库获取用户信息
	user := models.User{Name: name, Email: name + "@example.com"}

	// return json response
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
