package routes

import (
  "github.com/gin-gonic/gin"
  "main/controllers"
)

func RegisterAuthRoutes(router *gin.Engine) {
  auth := router.Group("/api/v1/auth")
  {
    auth.POST("/register", controllers.Register)
      
  }
}