package main

import (
	"log"

	"github.com/pu4mane/NATSOrderViewer/config"
	"github.com/pu4mane/NATSOrderViewer/internal/app/apiserver"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	if err := apiserver.Start(cfg); err != nil {
		log.Fatalf("Error starting API server: %v", err)
	}
}
