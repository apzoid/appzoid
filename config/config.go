package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config uses for getting enviromental variables
func Config(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file")
	}
	return os.Getenv(key)
}
