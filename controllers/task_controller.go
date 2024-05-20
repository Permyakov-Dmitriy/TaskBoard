package controllers

import (
	"net/http"
	"strconv"
	"webapp/models"
	"webapp/services"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskService *services.TaskService
}

var AllowedSortFields = map[string]bool{
	"Priority":  true,
	"Status":    true,
	"CreatedAt": true,
}

func (uc *TaskController) CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uc.TaskService.CreateTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (uc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := uc.TaskService.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (uc *TaskController) GetOrderedTasks(c *gin.Context) {
	sortField := c.DefaultQuery("sort_field", "priority")
	sortOrder := c.DefaultQuery("sort_order", "asc")

	if _, ok := AllowedSortFields[sortField]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sort field"})
		return
	}

	if sortOrder != "asc" && sortOrder != "desc" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid sort order"})
		return
	}

	orderClause := sortField + " " + sortOrder
	tasks, err := uc.TaskService.GetSortedTasks(orderClause)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (uc *TaskController) GetTask(c *gin.Context) {
	task_id := c.Params.ByName("id")
	id, err := strconv.Atoi(task_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := uc.TaskService.GetTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (uc *TaskController) UpdateTask(c *gin.Context) {
	task_id := c.Param("id")
	id, err := strconv.Atoi(task_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := uc.TaskService.GetTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := uc.TaskService.SaveUpdatedTask(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (uc *TaskController) DeleteTask(c *gin.Context) {
	task_id := c.Param("id")
	id, err := strconv.Atoi(task_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	if err := uc.TaskService.DeleteTask(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
