package models

import (
	//"time"
	"gorm.io/gorm"
)

type Joke struct {
	gorm.Model
	// Id        uint      `gorm:"primaryKey" json:"id"`
	Content   string    `json:"content" binding:"required"`
	Author    uint    `json:"author" binding:"required"`
	
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
}
