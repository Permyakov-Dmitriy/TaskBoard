package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
