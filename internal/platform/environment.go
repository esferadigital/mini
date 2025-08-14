package platform

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	DSN string
}

func NewEnvironment() *Environment {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := required("DSN")

	return &Environment{
		DSN: dsn,
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
