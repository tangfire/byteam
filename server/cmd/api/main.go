package main

import (
	"log"

	"byml/server/internal/app"
)

func main() {
	cfg := app.LoadConfig()

	server, err := app.NewServer(cfg)
	if err != nil {
		log.Fatalf("failed to initialize server: %v", err)
	}

	if err := server.Run(); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
