package routes

import (
	"webapp/controllers"
	"webapp/models"
	"webapp/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	userService := &services.UserService{DB: db}
	userController := &controllers.UserController{UserService: userService}

	models.AutoMigrate(db)

	r.GET("/users/", userController.GetUsers)
	r.GET("/users/:id", userController.GetUser)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users/:id", userController.UpdateUser)
	r.DELETE("/users/:id", userController.DeleteUser)
}
