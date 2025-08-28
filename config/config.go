package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl string
	Port  string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	dbUrl := os.Getenv("DATABASE_URI")
	if dbUrl == "" {
		log.Fatal("Database url is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Config{
		DBUrl: dbUrl,
		Port:  port,
	}

}
