package services

import (
	"webapp/models"

	"gorm.io/gorm"
)

type TaskService struct {
	DB *gorm.DB
}

func (ts *TaskService) CreateTask(task *models.TaskModel) error {
	return ts.DB.Create(task).Error
}

func (ts *TaskService) GetTasks() ([]models.TaskModel, error) {
	var tasks []models.TaskModel
	err := ts.DB.Find(&tasks).Error
	return tasks, err
}

func (ts *TaskService) GetSortedTasks(orderClause string) ([]models.TaskModel, error) {
	var tasks []models.TaskModel
	err := ts.DB.Order(orderClause).Find(&tasks).Error
	return tasks, err
}

func (ts *TaskService) GetTask(id string) (models.TaskModel, error) {
	var task models.TaskModel
	err := ts.DB.First(&task, id).Error
	return task, err
}

func (ts *TaskService) SaveUpdatedTask(task *models.TaskModel) error {
	err := ts.DB.Save(&task).Error
	return err
}

func (ts *TaskService) DeleteTask(id string) error {
	err := ts.DB.Delete(&models.TaskModel{}, id).Error
	return err
}
