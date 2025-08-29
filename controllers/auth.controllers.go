package controllers

import (
	"github.com/gin-gonic/gin"
	"main/models"
	"net/http"
	"main/services"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	user.Password = string(hash)
	createdUser, err := services.RegisterUser(user.Name, user.Email, user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
	return
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
