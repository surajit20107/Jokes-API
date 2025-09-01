package main

import (
	"github.com/gin-gonic/gin"
	"main/routes"
	"main/config"
	"log"
	"main/db"
	"main/frontend"
)

func main() {
	config := config.LoadConfig()

	dataBase := db.ConnectDatabase(config)

	db.Migrate(dataBase)
	
	router := gin.Default()
	
	routes.RegisterJokeRoutes(router)
	routes.RegisterAuthRoutes(router)

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message": "Api is running...",
			"Version": "1.0.0",
			"Author": "Surajit",
		})
	})

	frontend.FrontendRoutes(router)

	error := router.Run(":" + config.Port)
	if error != nil {
		log.Fatal("Failed to start server:", error)
	}
}
