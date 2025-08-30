package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"main/dto"
	"main/models"
	"main/services"
	"main/validator"
	"net/http"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	validEmail := validator.ValidateEmail(user.Email)

	if !validEmail {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email",
			"data": user,
		})
		return
	}

	validPassword := validator.ValidatePassword(user.Password)

	if !validPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
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

	userResponse := dto.ToUserResponse(createdUser)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    userResponse,
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
