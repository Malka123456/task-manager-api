package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterUser is a placeholder for now
func RegisterUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Register route working!",
	})
}

// LoginUser is a placeholder for now
func LoginUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login route working!",
	})
}
