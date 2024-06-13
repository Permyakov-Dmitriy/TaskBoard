package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique" binding:"required"`
	Email    string `json:"email" gorm:"unique" binding:"required" validate:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserLogin struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdatedUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
