package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// GetTasks is a basic controller handler
func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Tasks fetched successfully!",
	})
}
