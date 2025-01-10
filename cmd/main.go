package main

import (
	"todo-app/config"
	"todo-app/database"
	"todo-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database
	config.LoadEnv()

	database.InitDB()

	// Set up Gin router
	router := gin.Default()
	routes.RegisterRoutes(router)

	// Start the server
	router.Run(":8080")
}
