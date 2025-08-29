package routes

import (
  "github.com/gin-gonic/gin"
  "main/controllers"
)

func RegisterAuthRoutes(router *gin.Engine) {
  auth := router.Group("/api/v1/auth")
  {
    auth.POST("/register", controllers.Register)
    auth.POST("/login", controllers.Login)
    auth.POST("/logout", controllers.Logout)
    auth.GET("/me", controllers.Me)
  }
}