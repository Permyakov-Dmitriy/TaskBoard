package controllers

import (
	"log"
	"net/http"
	"webapp/models"
	"webapp/services"
	"webapp/utils"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskService *services.TaskService
	UserService *services.UserService
}

var AllowedSortFields = map[string]bool{
	"priority":   true,
	"status":     true,
	"created_at": true,
}

// CreateTask godoc
// @Summary      create task
// @Description	 send Task data create task
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param task body models.Task true " "
// @Success      200  {object}  models.Task
// @Router       /tasks [post]
func (tc *TaskController) CreateTask(c *gin.Context) {
	validatedData, exists := c.Get("validatedData")
	if !exists {
		log.Println("Validated data not found")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Validated data not found"})
		return
	}
	task := validatedData.(models.Task)

	auth_user_id, exists := c.Get("auth_user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth user not found"})
		return
	}

	user, err := tc.UserService.GetUser(auth_user_id.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	task_res := utils.TransformSingleModelToResponse[models.TaskModel](&task)
	task_res.UserID = auth_user_id.(uint)
	task_res.User = user

	if err := tc.TaskService.CreateTask(&task_res); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task_res)
}

// GetTasks godoc
// @Summary      get list tasks
// @Description
// @Tags         Task
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.Task
// @Router       /tasks [get]
func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.TaskService.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tasks_res := utils.TransformSliceModelToResponse[models.Task](tasks)

	c.JSON(http.StatusOK, tasks_res)
}

// GetOrderedTasks godoc
// @Summary      get ordered list tasks
// @Description
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param sort_field query string true "field for sort" example("priority, status, created_at")
// @Param sort_order query string true "type sort" example("asc, desc")
// @Success      200  {array}  models.Task
// @Router       /tasks/ordered [get]
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

	tasks_res := utils.TransformSliceModelToResponse[models.Task](tasks)

	c.JSON(http.StatusOK, tasks_res)
}

// GetTask godoc
// @Summary      get task
// @Description	 send id get task
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param id path int true "id task" example(1)
// @Success      200  {object}  models.Task
// @Router       /tasks/{id} [get]
func (tc *TaskController) GetTask(c *gin.Context) {
	task_id := c.Params.ByName("id")
	task, err := tc.TaskService.GetTask(task_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	task_res := utils.TransformSingleModelToResponse[models.Task](&task)

	c.JSON(http.StatusOK, task_res)
}

// UpdateTask godoc
// @Summary      update user
// @Description	 send id update task
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param id path int true "id task" example(1)
// @Success      200  {object}  models.Task
// @Router       /tasks/{id} [put]
func (tc *TaskController) UpdateTask(c *gin.Context) {
	validatedData, exists := c.Get("validatedData")
	if !exists {
		log.Println("Validated data not found")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Validated data not found"})
		return
	}
	task := validatedData.(models.UpdatedTask)

	task_id := c.Param("id")

	task_bd, err := tc.TaskService.GetTask(task_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	updated_task := utils.CopySingleModel(&task, &task_bd)

	if err := tc.TaskService.SaveUpdatedTask(updated_task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task_res := utils.TransformSingleModelToResponse[models.Task](&updated_task)

	c.JSON(http.StatusOK, task_res)
}

// DeleteTask godoc
// @Summary      delete task
// @Description
// @Tags         Task
// @Accept       json
// @Produce      json
// @Param id path int true "id task" example(1)
// @Success      200  {object}  nil
// @Router       /tasks/{id} [delete]
func (tc *TaskController) DeleteTask(c *gin.Context) {
	auth_user_id, exists := c.Get("auth_user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth user not found"})
		return
	}

	task_id := c.Param("id")

	task_bd, err := tc.TaskService.GetTask(task_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if auth_user_id != task_bd.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Аccess denied"})
		return
	}

	if err := tc.TaskService.DeleteTask(task_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
