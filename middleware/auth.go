package middleware

import (
	"log"
	"net/http"
	"webapp/config"
	"webapp/controllers"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.Abort()
		return
	}

	claims := &controllers.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().JwtSecretKey), nil
	})
	log.Println(token)
	log.Println(claims.Username)
	log.Println(token.Valid)

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		c.Abort()
		return
	}
	log.Println(claims.Username)
	c.Set("username", claims.Username)
	c.Next()
}
