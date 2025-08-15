package main

import (
	"log"

	"github.com/esferachill/mini/internal/platform"
	"github.com/esferachill/mini/internal/server"
)

func main() {
	platform := platform.GetPlatform()
	server := server.NewServer(platform.Env.Port)

	log.Println("Starting Mini URL Shortener on :8080")
	if err := server.Start(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
