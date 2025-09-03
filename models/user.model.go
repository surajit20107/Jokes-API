package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null" json:"name"`
	Email    string `gorm:"unique;not null" json:"email" binding:"required,email,min=8,max=30"`
	Password string `gorm:"not null" json:"password" binding:"required,min=8,max=20"`
	Jokes    []Joke `gorm:"foreignKey:Author"`
}
