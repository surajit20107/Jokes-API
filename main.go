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
	routes.RegisterAuthRoutes(router)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message": "Api is running...",
			"Version": "1.0.0",
			"Author": "Surajit",
		})
	})

	error := router.Run(":" + config.Port)
	if error != nil {
		log.Fatal("Failed to start server:", error)
	}
}
