package main

import (
	"github.com/gin-gonic/gin"
	"main/routes"
	"main/config"
	"log"
	"main/db"
	"main/models"
)

func main() {
	config := config.LoadConfig()

	db.ConnectDatabase(config)

	err := db.DB.AutoMigrate(&models.Joke{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	
	router := gin.Default()
	
	routes.RegisterJokeRoutes(router)

	error := router.Run(":" + config.Port)
	if error != nil {
		log.Fatal("Failed to start server:", error)
	}
}
