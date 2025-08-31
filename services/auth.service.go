package services

import (
  "main/models"
  "main/db"
  "gorm.io/gorm"
  "errors"
  "main/utils"
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


func LoginUser(email, password string) (string, error) {
  var user models.User
  err := db.DB.Where("email = ?", email).First(&user).Error
  if err != nil {
    return "", errors.New("Invalid credentials")
  }
  
  if !utils.CheckPassword(password, user.Password) {
    return "", errors.New("Invalid credentials")
  }
  token, err := utils.GenerateToken(user)
  if err != nil {
    return "", errors.New("Failed to generate token")
  }
  return token, nil
}

