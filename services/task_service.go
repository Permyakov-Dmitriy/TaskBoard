package services

import (
	"webapp/models"

	"gorm.io/gorm"
)

type TaskService struct {
	DB *gorm.DB
}

func (us *TaskService) CreateTask(task *models.Task) error {
	return us.DB.Create(task).Error
}

func (us *TaskService) GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := us.DB.Find(&tasks).Error
	return tasks, err
}

func (us *TaskService) GetSortedTasks(orderClause string) ([]models.Task, error) {
	var tasks []models.Task
	err := us.DB.Order(orderClause).Find(&tasks).Error
	return tasks, err
}

func (us *TaskService) GetTask(id int) (models.Task, error) {
	var task models.Task
	err := us.DB.First(&task, id).Error
	return task, err
}

func (us *TaskService) SaveUpdatedTask(task *models.Task) error {
	err := us.DB.Save(&task).Error
	return err
}

func (us *TaskService) DeleteTask(id int) error {
	err := us.DB.Delete(&models.Task{}, id).Error
	return err
}
