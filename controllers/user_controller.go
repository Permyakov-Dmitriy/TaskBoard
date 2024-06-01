package controllers

import (
	"log"
	"net/http"
	"webapp/models"
	"webapp/services"
	"webapp/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

// CreateUser godoc
// @Summary      Create new user
// @Description
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.UserResponse
// @Router       /users [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	validatedData, exists := c.Get("validatedData")
	if !exists {
		log.Println("Validated data not found")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Validated data not found"})
		return
	}
	user := validatedData.(models.User)

	if err := uc.UserService.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := utils.TransformSingleModelToResponse[models.UserResponse](&user)
	c.JSON(http.StatusOK, res)
}

// GetUsers godoc
// @Summary      Get list users
// @Description
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.UserResponse
// @Router       /users [get]
func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.UserService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := utils.TransformSliceModelToResponse[models.UserResponse](users)
	c.JSON(http.StatusOK, res)
}

// GetUser godoc
// @Summary      get user
// @Description	 send id get user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param id path int true "id user" example(1)
// @Success      200  {object}  models.UserResponse
// @Router       /users/{id} [get]
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

func (uc *UserController) GetProfile(c *gin.Context) {
	auth_username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth user not found"})
		return
	}

	user, err := uc.UserService.GetUserByUsername(auth_username.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	res := utils.TransformSingleModelToResponse[models.UserResponse](&user)
	c.JSON(http.StatusOK, res)
}

// UpdateUser godoc
// @Summary      update user
// @Description	 send id update user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param id path int true "id user" example(1)
// @Success      200  {object}  models.UserResponse
// @Router       /users/{id} [put]
func (uc *UserController) UpdateUser(c *gin.Context) {
	user_id := c.Param("id")

	user, err := uc.UserService.GetUser(user_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	auth_username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth user not found"})
		return
	}

	if auth_username != user.Username {
		c.JSON(http.StatusForbidden, gin.H{"error": "Аccess denied"})
		return
	}

	if err := uc.UserService.SaveUpdatedUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := utils.TransformSingleModelToResponse[models.UserResponse](&user)
	c.JSON(http.StatusOK, res)
}

// DeleteUser godoc
// @Summary      delete user
// @Description	 send id delete user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param id path int true "id user" example(1)
// @Success      200  {object}  nil
// @Router       /users/{id} [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	user_id := c.Param("id")

	if err := uc.UserService.DeleteUser(user_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
