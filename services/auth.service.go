package services

import (
  "main/models"
  "main/db"
  "gorm.io/gorm"
  "errors"
)

func RegisterUser(user, email, password string) (models.User, error) { 
  var existingUser models.User
  err := db.DB.Where("email = ?", email).First(&existingUser).Error
  if err == nil {
    return models.User{},
    errors.New("User already exists")
  }

  if !errors.Is(err, gorm.ErrRecordNotFound) {
    return models.User{},
    errors.New("Database error")
  }

  newUser := models.User{
    Name: user,
    Email: email,
    Password: password,
  }  

  result := db.DB.Create(&newUser)

  if result.Error != nil {
    return models.User{},
    errors.New("Failed to create user")
  }

  return newUser, nil 
}