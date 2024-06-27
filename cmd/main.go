package main

import (
	"log"

	"github.com/AlexandrKobalt/trip-track_web-server/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %s", err.Error())
	}
}
