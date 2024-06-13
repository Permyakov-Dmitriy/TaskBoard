package routes

import (
	"webapp/controllers"
	"webapp/middleware"
	"webapp/models"
	"webapp/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TasksRoutes(r *gin.RouterGroup, db *gorm.DB) {
	taskService := &services.TaskService{DB: db}
	userService := &services.UserService{DB: db}
	taskController := &controllers.TaskController{
		TaskService: taskService,
		UserService: userService,
	}

	models.AutoMigrate(db, models.TaskModel{})

	r.Use(middleware.AuthMiddleware)

	r.GET("/", taskController.GetTasks)
	r.GET("/:id", taskController.GetTask)
	r.GET("/ordered", taskController.GetOrderedTasks)
	r.POST("/", middleware.ValidatorMiddleware[models.Task](), taskController.CreateTask)
	r.PUT("/:id", middleware.ValidatorMiddleware[models.UpdatedTask](), taskController.UpdateTask)
	r.DELETE("/:id", taskController.DeleteTask)
}
