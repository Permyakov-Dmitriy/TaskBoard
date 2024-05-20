package routes

import (
	"webapp/controllers"
	"webapp/models"
	"webapp/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TasksRoutes(r *gin.RouterGroup, db *gorm.DB) {
	taskService := &services.TaskService{DB: db}
	taskController := &controllers.TaskController{TaskService: taskService}

	models.AutoMigrate(db, models.Task{})

	r.GET("/", taskController.GetTasks)
	r.GET("/:id", taskController.GetTask)
	r.POST("/", taskController.CreateTask)
	r.PUT("/:id", taskController.UpdateTask)
	r.DELETE("/:id", taskController.DeleteTask)
}
