package routes

import (
	"webapp/controllers"
	"webapp/middleware"
	"webapp/models"
	"webapp/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(r *gin.RouterGroup, db *gorm.DB) {
	userService := &services.UserService{DB: db}
	authController := &controllers.AuthController{UserService: userService}

	r.POST("/register", middleware.ValidatorMiddleware[models.User](), authController.RegisterHandler)
	r.POST("/login", middleware.ValidatorMiddleware[models.User](), authController.LoginHandler)
	r.POST("/refresh", middleware.ValidatorMiddleware[models.User](), authController.RefreshHandler)

	r.GET("/secured", middleware.AuthMiddleware, authController.SecuredHandler)
}
