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

	//// Auto-migrate models without altering existing constraints
	//err = database.Set("gorm:table_options", "WITHOUT ROWID").AutoMigrate(&models.User{}, &models.Note{})
	//if err != nil {
	//	log.Fatal("Failed to migrate database:", err)
	//}

	// Manually ensure unique constraints
	ensureUniqueConstraints(database)

	DB = database
	log.Println("Database connected successfully!")
}

func ensureUniqueConstraints(db *gorm.DB) {
	var exists bool
	err := db.Raw(`
		SELECT EXISTS (
			SELECT 1 
			FROM information_schema.table_constraints 
			WHERE table_name = 'users' 
			AND constraint_name = 'uni_users_email'
		);
	`).Scan(&exists).Error

	if err != nil {
		log.Fatal("Failed to check unique constraint on users.email:", err)
	}
	if !exists {
		err := db.Exec("ALTER TABLE users ADD CONSTRAINT uni_users_email UNIQUE (email);").Error
		if err != nil {
			log.Fatal("Failed to add unique constraint on users.email:", err)
		}
		log.Println("Unique constraint 'uni_users_email' added successfully")
	} else {
		log.Println("Unique constraint 'uni_users_email' already exists")
	}
}
