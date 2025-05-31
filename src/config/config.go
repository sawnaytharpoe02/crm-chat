package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	value := os.Getenv(key)
	if value == "" {
		fmt.Printf("Environment variable %s not set\n", key)
		return ""
	}

	return value
}
