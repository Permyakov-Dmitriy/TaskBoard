package models

import (
	"gorm.io/gorm"
)

func AutoMigrate[T any](db *gorm.DB, model T) {
	db.AutoMigrate(model)
}
