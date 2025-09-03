package dto

import (
	"main/models"
	"time"
)

type UserRegistration struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email,min=8,max=30"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email,min=8,max=30"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

type userResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToUserResponse(user models.User) userResponse {
	return userResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
