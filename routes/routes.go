package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api/v1")

	users := api.Group("/users")
	UsersRoutes(users, db)

	tasks := api.Group("/tasks")
	TasksRoutes(tasks, db)

	auth := api.Group("/auth")
	AuthRoutes(auth, db)
}
