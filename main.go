package main

import (
	"libraryManagement/config"
	"libraryManagement/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Connect to database
	config.Database()

	// Setup router
	router := gin.Default()
	routes.SetupRoutes(router)

	// Run server
	router.Run(":8080")
}
