package platform

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Environment struct {
	DatabaseURL string
	Port        int
}

func NewEnvironment() *Environment {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Could not load .env file: %v", err)
	}

	databaseURL := required("DATABASE_URL")

	port, err := strconv.Atoi(required("PORT"))
	if err != nil {
		log.Fatalf("Invalid PORT value: %s", err)
	}

	return &Environment{
		DatabaseURL: databaseURL,
		Port:        port,
	}
}

func required(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s not set", key)
	}
	return value
}

func optional(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("Environment variable %s not set", key)
	}
	return value
}
