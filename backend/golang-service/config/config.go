package config

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

var (
	AppPort    string
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
	RedisHost  string
	RedisPort  string
)

// LoadConfig loads environment variables from .env file
func LoadConfig() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, using system environment variables.")
	}

	AppPort = getEnv("APP_PORT", "8080")
	DbHost = getEnv("DB_HOST", "localhost")
	DbUser = getEnv("DB_USER", "postgres")
	DbPassword = getEnv("DB_PASSWORD", "password")
	DbName = getEnv("DB_NAME", "cloudcost")
	RedisHost = getEnv("REDIS_HOST", "localhost")
	RedisPort = getEnv("REDIS_PORT", "6379")
}

// getEnv gets the value of an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}