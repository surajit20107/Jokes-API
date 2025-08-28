package config

import (
	"log"
	"os"
)

type Config struct {
	DBUrl string
	Port  string
}

func LoadConfig() *Config {
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
