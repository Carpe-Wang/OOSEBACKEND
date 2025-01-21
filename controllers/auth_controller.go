package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func GitHubAuth(c *gin.Context) {
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func GitHubCallback(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"email":    user.Email,
			"provider": user.Provider,
		},
	})
}
