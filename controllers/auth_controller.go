package controllers

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"markdown-notes-backend/config"
	"markdown-notes-backend/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func GitHubAuth(c *gin.Context) {
	provider := "github"
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "provider", provider))
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func GitHubCallback(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed", "details": err.Error()})
		return
	}
	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required for GitHub login"})
		return
	}
	var dbUser models.User
	result := config.DB.Where("email = ?", user.Email).First(&dbUser)

	if result.RowsAffected == 0 {
		dbUser = models.User{
			Email:    user.Email,
			Provider: user.Provider,
		}
		if err := config.DB.Create(&dbUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user", "details": err.Error()})
			return
		}
	}
	token, err := generateJWT(dbUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	redirectBaseURL := "https://markdown-notes-frontend.vercel.app"
	redirectURL := redirectBaseURL + "?token=" + token

	// 打印日志（可选）
	log.Printf("Redirecting to: %v", redirectURL)

	// 重定向到前端页面
	c.Redirect(http.StatusTemporaryRedirect, redirectURL)
	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user": gin.H{
			"id":    dbUser.ID,
			"email": dbUser.Email,
		},
	})
}

func generateJWT(userID uint) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
		"iat":    time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
