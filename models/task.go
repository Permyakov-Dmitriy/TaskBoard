package models

import (
	"time"
)

type TaskModel struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" binding:"required" validate:"max=20"`
	Description string `json:"description" binding:"required"`
	Priority    int    `json:"priority" binding:"required" validate:"gte=1,lte=10"`
	Status      string `json:"status" binding:"required"`

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type Task struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" binding:"required" validate:"max=20"`
	Description string `json:"description" binding:"required"`
	Priority    int    `json:"priority" binding:"required" validate:"gte=1,lte=10"`
	Status      string `json:"status" binding:"required"`

	User UserResponse

	CreatedAt time.Time
	UpdatedAt time.Time
}
