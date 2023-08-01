package main

import (
	"github.com/finexblock-dev/gofinexblock/cmd/proxy/internal/bootstrap"
	"github.com/finexblock-dev/gofinexblock/finexblock/safety"
	"log"
)

func main() {
	log.Println("Wallet proxy server start")

	safety.GracefullyStopBootstrap(bootstrap.ProxyBootstrap)
}
