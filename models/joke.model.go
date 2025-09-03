package models

import (
	"gorm.io/gorm"
)

type Joke struct {
	gorm.Model
	Content string `json:"content" binding:"required"`
	Author  uint   `json:"author"`
	User    User   `gorm:"foreignKey:Author;constraint:OnDelete:CASCADE;" json:"-"`
}
