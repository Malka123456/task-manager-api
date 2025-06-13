package main

import (
	"github.com/gin-gonic/gin"
    "github.com/malka123456/task-manager-api/config"
    "github.com/malka123456/task-manager-api/routes"


)

func main() {
	config.ConnectDB()
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run(":8080")
}
