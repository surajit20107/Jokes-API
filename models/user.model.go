package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID   `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string `gorm:"not null" json:"name" binding:"required,min=5,max=30"`
	Email     string `gorm:"unique;not null" json:"email" binding:"required,email,min=8,max=30"`
	Password  string `gorm:"not null" json:"password" binding:"required,min=8,max=20"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
