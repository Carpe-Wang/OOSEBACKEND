package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=wangcarpe password=wkp159262 dbname=markdown_notes port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = database
	log.Println("Database connected successfully!")
}
