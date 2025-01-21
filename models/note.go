package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey"`
	UserID  uint   `gorm:"not null"`
	Title   string `gorm:"not null"`
	Content string `gorm:"type:text"`
}
