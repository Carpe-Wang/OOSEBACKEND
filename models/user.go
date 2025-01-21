package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"uniqueIndex;not null"`
	Provider string `gorm:"not null"` // OAuth provider, e.g., Google, GitHub
}
