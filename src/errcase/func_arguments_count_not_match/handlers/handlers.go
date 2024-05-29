package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ZyhHelen/autodebug-dataset/src/errcase/func_arguments_count_not_match/models"
	"github.com/ZyhHelen/autodebug-dataset/src/errcase/func_arguments_count_not_match/utils"
)

// Ping handler
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// Submit handler
func Submit(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")

	user := models.User{Name: name, Email: email}
	// 模拟将用户保存到数据库或其他地方

	c.JSON(http.StatusOK, gin.H{
		"status": "submitted",
		"user":   user,
	})
}

// GetUser handler
func GetUser(c *gin.Context) {
	name := c.Param("name")
	// 模拟从数据库获取用户信息
	user := models.User{Name: name, Email: name + "@example.com"}

	userInfo := utils.FormatUserInfo(user.Name)

	c.JSON(http.StatusOK, gin.H{
		"user": userInfo,
	})
}
