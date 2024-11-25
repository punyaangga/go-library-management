package main

import (
	"libraryManagement/config"
	"libraryManagement/models"
	"libraryManagement/routes"
	"libraryManagement/seeder"
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
	// migrate database
	if err := config.DB.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.ExpiredToken{},
		&models.Product{},
		&models.Stock{},
	); err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	//migrate data seeder. If you want to seed data, uncomment the line below
	seeder.SeederProduct()

	// Setup router
	router := gin.Default()
	routes.SetupRoutes(router)

	// Run server
	router.Run(":8080")
}
