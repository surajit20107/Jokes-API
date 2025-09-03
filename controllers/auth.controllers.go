package controllers

import (
	"github.com/gin-gonic/gin"
	"main/utils"
	"main/dto"
	"main/services"
	"main/validator"
	"net/http"
)

func Register(c *gin.Context) {
	var userDTO dto.UserRegistration

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !validator.ValidateEmail(userDTO.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email",
		})
		return
	}

	if !validator.ValidatePassword(userDTO.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid password",
		})
		return
	}

	hash, err := utils.HashPassword(userDTO.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	userDTO.Password = string(hash)

	createdUser, err := services.RegisterUser(userDTO.Name, userDTO.Email, userDTO.Password)

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
	var userDTO dto.UserLogin
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	token, err := services.LoginUser(userDTO.Email, userDTO.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	c.SetCookie("token", token, 60 * 60 * 24 * 7, "/", "", true, true)
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token": token,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", true, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successfully",
	})
	return
}

func Me(c *gin.Context) {
	user, exist := c.Get("user")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}
	
}
