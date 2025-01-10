package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}
}
