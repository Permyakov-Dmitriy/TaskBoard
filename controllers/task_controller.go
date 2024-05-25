package controllers

import (
	"log"
	"net/http"
	"webapp/models"
	"webapp/services"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskService *services.TaskService
}

var AllowedSortFields = map[string]bool{
	"priority":   true,
	"status":     true,
	"created_at": true,
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	validatedData, exists := c.Get("validatedData")
	if !exists {
		log.Println("Validated data not found")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Validated data not found"})
		return
	}

	task := validatedData.(models.Task)
	if err := tc.TaskService.CreateTask(&task); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.TaskService.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetOrderedTasks(c *gin.Context) {
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
	tasks, err := tc.TaskService.GetSortedTasks(orderClause)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTask(c *gin.Context) {
	task_id := c.Params.ByName("id")
	task, err := tc.TaskService.GetTask(task_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	task_id := c.Param("id")

	task, err := tc.TaskService.GetTask(task_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := tc.TaskService.SaveUpdatedTask(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	task_id := c.Param("id")

	if err := tc.TaskService.DeleteTask(task_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
