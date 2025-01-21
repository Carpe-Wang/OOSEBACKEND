package routes

import (
	"markdown-notes-backend/controllers"
	"markdown-notes-backend/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.GET("/auth/github", controllers.GitHubAuth)
	router.GET("/auth/github/callback", controllers.GitHubCallback)

	// Protected route example
	protected := router.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/test", func(c *gin.Context) {
		userID, _ := c.Get("userID")
		c.JSON(200, gin.H{
			"message": "You are authenticated",
			"userID":  userID,
		})
	})
}
