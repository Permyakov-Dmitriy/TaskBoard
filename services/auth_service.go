package services

import (
	"webapp/models"

	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func (us *AuthService) CreateUser(user *models.User) error {
	return us.DB.Create(user).Error
}

func (us *AuthService) GetUsers() ([]models.User, error) {
	var users []models.User
	err := us.DB.Find(&users).Error
	return users, err
}

func (us *AuthService) SaveUpdatedUser(user *models.User) error {
	err := us.DB.Save(&user).Error
	return err
}

func (us *AuthService) DeleteUser(id int) error {
	err := us.DB.Delete(&models.User{}, id).Error
	return err
}
