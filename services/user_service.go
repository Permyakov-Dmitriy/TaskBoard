package services

import (
	"log"
	"webapp/models"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (us *UserService) CreateUser(user *models.User) error {
	return us.DB.Create(user).Error
}

func (us *UserService) GetUsers() ([]models.User, error) {
	var users []models.User
	err := us.DB.Find(&users).Error
	return users, err
}

func (us *UserService) GetUser(id int) (models.User, error) {
	var user models.User
	err := us.DB.First(&user, id).Error
	log.Println(err)
	return user, err
}

func (us *UserService) SaveUpdatedUser(user *models.User) error {
	err := us.DB.Save(&user).Error
	return err
}

func (us *UserService) DeleteUser(id int) error {
	err := us.DB.Delete(&models.User{}, id).Error
	return err
}
