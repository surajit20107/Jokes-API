package db

import (
  "log"
  "main/models"
  "gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
  err := db.AutoMigrate(
    &models.Joke{},
  )
  
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
  }
  
}