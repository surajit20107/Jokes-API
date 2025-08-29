package models

import (
	"time"
	"gorm.io/gorm"
)

type Joke struct {
	gorm.Model // no need if manually defining ID, CreatedAt, UpdatedAt, DeletedAt
	Id        uint      `gorm:"primaryKey" json:"id"`
	Content   string    `json:"content" binding:"required"`
	Author    string    `json:"author" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
