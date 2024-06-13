package controllers

import (
	"log"
	"net/http"
	"strconv"
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
	uint_user_id, err := strconv.ParseUint(user_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.UserService.GetUser(uint(uint_user_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	res := utils.TransformSingleModelToResponse[models.UserResponse](&user)
	c.JSON(http.StatusOK, res)
}

// GetProfile godoc
// @Summary      get profile
// @Description  send id get user
// @Tags         User
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.UserResponse
// @Router       /users/me [get]
func (uc *UserController) GetProfile(c *gin.Context) {
	auth_user_id, exists := c.Get("auth_user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth user not found"})
		return
	}

	user, err := uc.UserService.GetUser(auth_user_id.(uint))
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
	uint_user_id, err := strconv.ParseUint(user_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	validatedData, exists := c.Get("validatedData")
	if !exists {
		log.Println("Validated data not found")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Validated data not found"})
		return
	}
	user_data := validatedData.(models.UpdatedUser)

	user, err := uc.UserService.GetUser(uint(uint_user_id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	auth_user_id, exists := c.Get("auth_user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth user not found"})
		return
	}
	uint_auth_user_id := auth_user_id.(uint)

	if uint_auth_user_id != user.ID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Аccess denied"})
		return
	}

	updated_user := utils.CopySingleModel(&user_data, &user)

	if err := uc.UserService.SaveUpdatedUser(updated_user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := utils.TransformSingleModelToResponse[models.UserResponse](updated_user)
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
	auth_user_id, exists := c.Get("auth_user_id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Auth user not found"})
		return
	}
	str_auth_user_id := auth_user_id.(string)

	if err := uc.UserService.DeleteUser(str_auth_user_id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
