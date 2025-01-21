package main

import (
	"log"
	"markdown-notes-backend/config"
	"markdown-notes-backend/models"
	"markdown-notes-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	config.ConnectDatabase()

	// Initialize OAuth
	config.InitOAuth()

	// Auto-migrate models
	err := config.DB.AutoMigrate(&models.User{}, &models.Note{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migrated successfully!")

	// Initialize Gin
	router := gin.Default()

	// Register routes
	routes.AuthRoutes(router)

	// Start server
	log.Println("Server running at http://localhost:8080")
	router.Run(":8080")
}
