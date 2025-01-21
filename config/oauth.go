package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
)

func InitOAuth() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	goth.UseProviders(
		github.New(
			os.Getenv("GITHUB_CLIENT_ID"),
			os.Getenv("GITHUB_CLIENT_SECRET"),
			os.Getenv("GITHUB_CALLBACK_URL"),
		),
	)

	log.Println("OAuth providers initialized successfully!")
}
