package models

import (
	"time"
)

type Task struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" binding:"required" validate:"max=20"`
	Description string `json:"description" binding:"required"`
	Priority    int    `json:"priority" binding:"required" validate:"gte=1,lte=10"`
	Status      string `json:"status" binding:"required"`
	CreatedAt   time.Time
}
