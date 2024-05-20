package models


type Task struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Priority    int    `json:"priority" binding:"required"`
	Status      string `json:"status" binding:"required"`
}
