package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidatorMiddleware[T any]() gin.HandlerFunc {
	validate := validator.New()
	return func(c *gin.Context) {
		var data T

		if err := c.ShouldBindJSON(&data); err != nil {
			log.Println("Invalid JSON")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			c.Abort()
			return
		}

		if err := validate.Struct(data); err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("validatedData", data)
		c.Next()
	}
}
