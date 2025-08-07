package main

import (
	"log"

	"github.com/esferachill/mini/infra"
	"github.com/esferachill/mini/repo"
	"github.com/esferachill/mini/server"
)

func main() {
	memory := infra.NewMemory()
	linkRepo := repo.NewLinkRepository(memory)
	server := server.NewServer(8080, linkRepo)

	log.Println("Starting Mini URL Shortener on :8080")
	if err := server.Start(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
