package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required" validate:"email"`
}
