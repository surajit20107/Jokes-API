package routes

import (
  "github.com/gin-gonic/gin"
  "main/controllers"
  middleware "main/middlewares"
)

func RegisterAuthRoutes(router *gin.Engine) {
  auth := router.Group("/api/v1/auth")
  {
    auth.POST("/register", controllers.Register)
    auth.POST("/login", controllers.Login)
    auth.POST("/logout", controllers.Logout)
    auth.GET("/me", middleware.CheckAuth(), controllers.Me)
  }
}