package platform

import (
	"log"
	"sync"
)

type Platform struct {
	DBClient *DatabaseClient
	Env      *Environment
}

var (
	instance *Platform
	once     sync.Once
)

func newPlatform() *Platform {
	env := NewEnvironment()

	dbConfig := DatabaseConfig{
		DatabaseURL: env.DatabaseURL,
	}

	dbClient, err := NewDatabaseClient(dbConfig)
	if err != nil {
		log.Fatalf("Failed to initialize database client: %v", err)
	}

	return &Platform{
		DBClient: dbClient,
		Env:      env,
	}
}

func GetPlatform() *Platform {
	once.Do(func() {
		instance = newPlatform()
	})
	return instance
}
