package controllers

import (
	"net/http"
	"webapp/models"
	"webapp/services"
	"webapp/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uc.UserService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.UserService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := utils.TransformSliceModelToResponse[models.UserResponse](users)
	c.JSON(http.StatusOK, res)
}

func (uc *UserController) GetUser(c *gin.Context) {
	user_id := c.Params.ByName("id")

	user, err := uc.UserService.GetUser(user_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	res := utils.TransformSingleModelToResponse[models.UserResponse](&user)

	c.JSON(http.StatusOK, res)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	user_id := c.Param("id")

	user, err := uc.UserService.GetUser(user_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := uc.UserService.SaveUpdatedUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	user_id := c.Param("id")

	if err := uc.UserService.DeleteUser(user_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
