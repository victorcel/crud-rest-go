package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// Config for databases
type Config struct {
	Username     string
	Password     string
	DatabaseName string
	URL          string
}

// GetConfig returns hardcoded config
func GetConfig() *Config {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		Username:     os.Getenv("DATABASE_USERNAME"),
		Password:     os.Getenv("DATAASE_PASSWORD"),
		DatabaseName: os.Getenv("DATABASE_NAME"),
		URL:          os.Getenv("DATABASE_URL"),
	}
}
