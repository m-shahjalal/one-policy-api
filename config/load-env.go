package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	// Log important environment variables
	log.Printf("Environment variables loaded:")
	log.Printf("JWT_SECRET length: %d", len(os.Getenv("JWT_SECRET")))
	log.Printf("JWT_ACCESS_TOKEN_EXPIRY: %s", os.Getenv("JWT_ACCESS_TOKEN_EXPIRY"))
	log.Printf("JWT_REFRESH_TOKEN_EXPIRY: %s", os.Getenv("JWT_REFRESH_TOKEN_EXPIRY"))
	log.Printf("DB_URL: %s", os.Getenv("DB_URL"))
}
