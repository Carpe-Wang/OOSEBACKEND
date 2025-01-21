package routes

import (
	"markdown-notes-backend/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.GET("/auth/github", controllers.GitHubAuth)
	router.GET("/auth/github/callback", controllers.GitHubCallback)
}
