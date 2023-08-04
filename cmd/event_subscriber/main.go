package main

import (
	"github.com/finexblock-dev/gofinexblock/cmd/event_subscriber/internal"
	"github.com/finexblock-dev/gofinexblock/pkg/safety"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed to load env file : %v", err)
	}
	log.Println("EVENT SUBSCRIBER")

	safety.GracefullyStopBootstrap(internal.Bootstrap)
}