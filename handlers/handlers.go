package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/ZyhHelen/autodebug-dataset/models"
	"github.com/ZyhHelen/autodebug-dataset/utils"
)

// Info handler
func Info(c echo.Context) error {
	log.Printf("[INFO] <Info> receive request from: %s ", c.RealIP())

	infoMap := utils.GetInfoMap()

	// format info
	infoMsg := ""
	for k, v := range infoMap {
		infoMsg += fmt.Sprintf("%v\t\t:%v\n", k, v)
	}

	// return json response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": infoMsg,
	})
}

// Ping handler
func Ping(c echo.Context) error {
	log.Printf("[INFO] <Ping> receive request from: %s ", c.RealIP())

	// return response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "pong",
	})
}

// Submit handler
func Submit(c echo.Context) error {
	log.Printf("[INFO] <Submit> receive request from: %s ", c.RealIP())

	name := c.FormValue("name")
	email := c.FormValue("email")

	// 模拟将用户保存到数据库或其他地方
	user := models.User{Name: name, Email: email}

	// return json response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "submitted",
		"user":   user,
	})
}

// GetUser handler
func GetUser(c echo.Context) error {
	log.Printf("[INFO] <GetUser> receive request from: %s ", c.RealIP())

	name := c.Param("name")

	// 模拟从数据库获取用户信息
	user := models.User{Name: name, Email: name + "@example.com"}

	// return json response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}
