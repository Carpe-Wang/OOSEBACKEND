package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	// 替换为你的 Supabase 数据库连接信息
	// dsn := "host=localhost user=wangcarpe password=wkp159262 dbname=markdown_notes port=5432 sslmode=disable"
	dsn := "host=db.hapjzsiacynlxhwtteku.supabase.co user=postgres password=yBVzAtOpzo6yvLLw dbname=postgres port=5432 sslmode=require"

	// 连接数据库
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 自动迁移模型（如果需要）
	// err = database.AutoMigrate(&models.User{}, &models.Note{})
	// if err != nil {
	// 	log.Fatal("Failed to migrate database:", err)
	// }

	// 检查并确保唯一约束
	err = ensureUniqueConstraints(database)
	if err != nil {
		log.Fatal("Failed to ensure unique constraints:", err)
	}

	DB = database
	log.Println("Connected to Supabase database successfully!")
}

func ensureUniqueConstraints(db *gorm.DB) error {
	var exists bool

	// 检查是否已存在唯一约束
	err := db.Raw(`
		SELECT EXISTS (
			SELECT 1 
			FROM information_schema.table_constraints 
			WHERE table_name = 'users' 
			AND constraint_name = 'uni_users_email'
		);
	`).Scan(&exists).Error

	if err != nil {
		return err
	}

	if !exists {
		// 添加唯一约束
		err := db.Exec("ALTER TABLE users ADD CONSTRAINT uni_users_email UNIQUE (email);").Error
		if err != nil {
			return err
		}
		log.Println("Unique constraint 'uni_users_email' added successfully")
	} else {
		log.Println("Unique constraint 'uni_users_email' already exists")
	}

	return nil
}
