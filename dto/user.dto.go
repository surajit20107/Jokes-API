package dto

import (
	"main/models"
	"time"
)

type userResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name" binding:"required,email"`
	Email     string `json:"email" binding:"required,min=8"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToUserResponse(user models.User) userResponse {
	return userResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
