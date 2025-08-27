package main

import (
	"github.com/gin-gonic/gin"
	"main/routes"
)

func main() {
	router := gin.Default()

	routes.RegisterJokeRoutes(router)

	router.Run(":3000")
}
