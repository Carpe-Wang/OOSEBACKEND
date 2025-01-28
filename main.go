package main

import (
	"log"
	"markdown-notes-backend/config"
	"markdown-notes-backend/models"
	"markdown-notes-backend/routes"
	"net/http"

	"github.com/gin-contrib/cors"
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

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is running!")
	})

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://markdown-notes-frontend.vercel.app/"}, // 前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.AuthRoutes(router)
	routes.NotesRoutes(router)

	for _, route := range router.Routes() {
		log.Printf("Registered Route: %s %s\n", route.Method, route.Path)
	}

	// Start server
	log.Println("Server running at http://localhost:8080")
	router.Run(":8080")
}
