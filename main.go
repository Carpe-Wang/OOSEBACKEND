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
	if err := config.DB.AutoMigrate(&models.User{}, &models.Note{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migrated successfully!")

	router := gin.New()
	router.RedirectTrailingSlash = false
	routes.AuthRoutes(router)
	routes.NotesRoutes(router)

	for _, route := range router.Routes() {
		log.Printf("Registered Route: %s %s\n", route.Method, route.Path)
	}

	// Start server
	log.Println("Server running at http://localhost:8080")
	router.Run(":8080")
}
