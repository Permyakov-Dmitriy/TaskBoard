package routes

import (
	"webapp/controllers"
	"webapp/middleware"
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
	r.POST("/", middleware.ValidatorMiddleware[models.User](), userController.CreateUser)
	r.PUT("/:id", middleware.ValidatorMiddleware[models.User](), userController.UpdateUser)
	r.DELETE("/:id", userController.DeleteUser)
}
