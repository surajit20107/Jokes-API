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

  newUser := db.DB.Create(&models.User{
    Name: user,
    Email: email,
    Password: password,
  })

  if newUser.Error != nil {
    return models.User{},
    errors.New("Failed to create user")
  }

  return newUser, nil
  
}