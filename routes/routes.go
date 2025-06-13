package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/malka123456/task-manager-api/controllers"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	router.GET("/tasks", controllers.GetTasks)
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)
}
