package main

import (
	"github.com/finexblock-dev/gofinexblock/cmd/matching_engine/internal/server"
	"github.com/finexblock-dev/gofinexblock/pkg/safety"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error occurred while load env file ... : %v", err)
	}

	safety.GracefullyStopBootstrap(server.Bootstrap)
}