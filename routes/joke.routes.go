package routes

import (
	"github.com/gin-gonic/gin"
	controller "main/controllers"
	middleware "main/middlewares"
)

func RegisterJokeRoutes(router *gin.Engine) {
	jokes := router.Group("/api/v1/jokes")
	{
		jokes.GET("/", controller.GetJokes)
		jokes.POST("/create", middleware.CheckAuth(), controller.CreateJoke)
		jokes.PATCH("/:id", controller.UpdateJoke)
		jokes.DELETE("/:id", controller.DeleteJoke)
		jokes.GET("/random", controller.GetRandomJoke)
	}
}
