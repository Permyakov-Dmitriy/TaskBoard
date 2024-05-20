package routes

import (
	"webapp/controllers"
	"webapp/models"
	"webapp/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UsersRoutes(r *gin.RouterGroup, db *gorm.DB) {
	userService := &services.UserService{DB: db}
	userController := &controllers.UserController{UserService: userService}

	models.AutoMigrate(db, models.User{})

	r.GET("/", userController.GetUsers)
	r.GET("/:id", userController.GetUser)
	r.POST("/", userController.CreateUser)
	r.PUT("/:id", userController.UpdateUser)
	r.DELETE("/:id", userController.DeleteUser)
}
