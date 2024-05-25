package services

import (
	"webapp/models"

	"gorm.io/gorm"
)

type TaskService struct {
	DB *gorm.DB
}

func (ts *TaskService) CreateTask(task *models.Task) error {
	return ts.DB.Create(task).Error
}

func (ts *TaskService) GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	err := ts.DB.Find(&tasks).Error
	return tasks, err
}

func (ts *TaskService) GetSortedTasks(orderClause string) ([]models.Task, error) {
	var tasks []models.Task
	err := ts.DB.Order(orderClause).Find(&tasks).Error
	return tasks, err
}

func (ts *TaskService) GetTask(id string) (models.Task, error) {
	var task models.Task
	err := ts.DB.First(&task, id).Error
	return task, err
}

func (ts *TaskService) SaveUpdatedTask(task *models.Task) error {
	err := ts.DB.Save(&task).Error
	return err
}

func (ts *TaskService) DeleteTask(id string) error {
	err := ts.DB.Delete(&models.Task{}, id).Error
	return err
}
