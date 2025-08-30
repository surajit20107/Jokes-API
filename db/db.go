package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"main/config"
	"fmt"
)

var DB *gorm.DB // declare a global variable for database connection

func ConnectDatabase(cfg *config.Config) *gorm.DB {
	// connect to the database
	database, err := gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	
	fmt.Println("Database connection successful ✅")
	DB = database // assign the database connection to the global variable
	return database
}
