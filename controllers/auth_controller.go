package controllers

import (
	"log"
	"net/http"
	"time"
	"webapp/config"
	"webapp/models"
	"webapp/services"
	"webapp/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService *services.AuthService
	UserService *services.UserService
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (ac *AuthController) RegisterHandler(c *gin.Context) {
	validated_data, exists := c.Get("validatedData")
	validated_user_data := validated_data.(models.User)
	if !exists {
		log.Println("Validated data not found")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Validated data not found"})
		return
	}

	if _, err := ac.UserService.GetUserByEmail(validated_user_data.Email); err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	hashed_password, err := utils.HashPassword(validated_user_data.Password)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	validated_user_data.Password = hashed_password

	if err := ac.UserService.CreateUser(&validated_user_data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (ac *AuthController) LoginHandler(c *gin.Context) {
	validated_data, exists := c.Get("validatedData")
	validated_user_data := validated_data.(models.User)
	if !exists {
		log.Println("Validated data not found")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Validated data not found"})
		return
	}

	user, err := ac.UserService.GetUserByEmail(validated_user_data.Email)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	if check := utils.CheckPasswordHash(validated_user_data.Password, user.Password); !check {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	config := config.GetConfig()

	accessToken, err := utils.GenerateToken(user.Username, []byte(config.JwtSecretKey), 15*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate access token"})
		return
	}

	refreshToken, err := utils.GenerateToken(user.Username, []byte(config.RefreshSecretKey), 24*time.Hour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken, "refresh_token": refreshToken})
}

func (ac *AuthController) RefreshHandler(c *gin.Context) {
	var requestBody struct {
		RefreshToken string `json:"refresh_token"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	claims := &Claims{}
	config := config.GetConfig()
	refreshToken, err := jwt.ParseWithClaims(requestBody.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return config.RefreshSecretKey, nil
	})

	if err != nil || !refreshToken.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
		return
	}

	accessToken, err := utils.GenerateToken(claims.Username, []byte(config.JwtSecretKey), 15*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}

func (ac *AuthController) SecuredHandler(c *gin.Context) {
	username := c.GetString("username")
	c.JSON(http.StatusOK, gin.H{"message": "You are authorized", "username": username})
}