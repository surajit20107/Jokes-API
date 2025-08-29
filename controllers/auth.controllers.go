package controllers

import (
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Register route",
	})
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Login route",
	})
}

func Logout(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "Logout route",
  })
}

func Me(c *gin.Context) {
  c.JSON(200, gin.H{
    "message": "Current user route",
  })
}
