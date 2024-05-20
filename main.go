package main

import (
	"log"
	"webapp/config"
	"webapp/database"
	"webapp/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	db, err := database.InitDB(config.GetConfig().DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := gin.Default()

	routes.RegisterRoutes(r, db)

	r.Run(config.GetConfig().ServerPort)
}
