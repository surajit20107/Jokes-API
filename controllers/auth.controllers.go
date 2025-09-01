package controllers

import (
	"github.com/gin-gonic/gin"
	"main/utils"
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
			"email": user.Email,
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

	hash, err := utils.HashPassword(user.Password)
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
			"email": user.Email,
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
	token, err := services.LoginUser(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.SetCookie("token", token, 60 * 60 * 24 * 7, "/", "", false, true)
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token": token,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successfully",
	})
	return
}

func Me(c *gin.Context) {
	userId, exists := c.Get("user_id")
	
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	user, err := services.GetUserByID(userId.(uint))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	userResponse := dto.ToUserResponse(user)

	c.JSON(http.StatusOK, gin.H{
		"user": userResponse,
	})
	return
}
