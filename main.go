package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"main/config"
	"main/db"
	"main/frontend"
	middleware "main/middlewares"
	"main/routes"
)

func main() {
	config := config.LoadConfig()

	dataBase := db.ConnectDatabase(config)

	db.Migrate(dataBase)

	router := gin.Default()

	router.Use(middleware.SecurityHeaders()) // custom middleware for security headers

	routes.RegisterJokeRoutes(router)
	routes.RegisterAuthRoutes(router)

	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message": "Api is running...",
			"Version": "1.0.0",
			"Author":  "Surajit",
		})
	})

	frontend.FrontendRoutes(router)

	error := router.Run(":" + config.Port)
	if error != nil {
		log.Fatal("Failed to start server:", error)
	}
}
