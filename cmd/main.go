package main

import (
	"fmt"
	"log"

	"github.com/esferachill/mini/internal/platform"
	"github.com/esferachill/mini/internal/server"
)

func main() {
	platform := platform.GetPlatform()
	server := server.NewServer(platform.Env.Port)

	welcome := fmt.Sprintf("Starting mini at %s://%s:%d", platform.Env.HostScheme, platform.Env.Host, platform.Env.Port)
	log.Println(welcome)
	if err := server.Start(); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
