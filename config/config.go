package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBAddress  string
	DBPassword string
	DBName     string
}

var AppConfig = initConfig()

func initConfig() Config {
    err := godotenv.Load()
    if err != nil {
      log.Fatal("Error loading .env file")
    }
	return Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBAddress:  getEnv("DB_ADDRESS", "localhost"),
		DBPassword: getEnv("DB_PASSWORD", "09876"),
		DBName:     getEnv("DB_NAME", "go_auth"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
