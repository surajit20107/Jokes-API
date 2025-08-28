package models

import "time"

type Joke struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Content   string    `json:"content" binding:"required"`
	Author    string    `json:"author" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
