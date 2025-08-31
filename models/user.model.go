package models

import (
	//"time"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	//ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null" json:"name" binding:"required,min=5,max=30"`
	Email     string `gorm:"unique;not null" json:"email" binding:"required,email,min=8,max=30"`
	Password  string `gorm:"not null" json:"password" binding:"required,min=8,max=20"`
	Jokes []Joke `gorm:"foreignKey:Author"`
	//CreatedAt time.Time
	//UpdatedAt time.Time
}
