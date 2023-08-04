package main

import (
	"github.com/finexblock-dev/gofinexblock/cmd/ethereum_key/internal/bootstrap"
	"github.com/finexblock-dev/gofinexblock/pkg/safety"
	"log"
)

func main() {
	log.Println("Ethereum key server start")

	safety.GracefullyStopBootstrap(bootstrap.EthereumKeyBootstrap)
}