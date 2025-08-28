package main

import (
	"github.com/gin-gonic/gin"
	"main/routes"
	"main/config"
	"log"
	"main/db"
)

func main() {
	config := config.LoadConfig()

	db.ConnectDatabase(config)
	
	router := gin.Default()
	
	routes.RegisterJokeRoutes(router)

	err := router.Run(":" + config.Port)
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
