package dto

import "time"

type User struct {
	ID        uint   `json:"id"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"-" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
