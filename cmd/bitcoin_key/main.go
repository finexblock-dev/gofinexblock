package main

import (
	"github.com/finexblock-dev/gofinexblock/cmd/bitcoin_key/internal/bootstrap"
	"github.com/finexblock-dev/gofinexblock/pkg/safety"
	"log"
)

func main() {
	log.Println("Bitcoin key server start")

	safety.GracefullyStopBootstrap(bootstrap.BitcoinKeyBoostrap)
}