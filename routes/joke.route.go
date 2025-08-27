package routes

import (
	"github.com/gin-gonic/gin"
	controller "main/controllers"
)

func RegisterJokeRoutes(router *gin.Engine) {
	jokes := router.Group("/jokes")
	{
		jokes.GET("/", controller.GetJokes)
		jokes.GET("/:id", controller.GetJokesById)
		jokes.POST("/create", controller.CreateJoke)
		jokes.PATCH("/:id", controller.UpdateJoke)
		jokes.DELETE("/:id", controller.DeleteJoke)
		jokes.GET("/random", controller.GetRandomJoke)
	}
}
