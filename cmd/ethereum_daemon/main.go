package main

import (
	"github.com/finexblock-dev/gofinexblock/cmd/ethereum_daemon/internal/bootstrap"
	"github.com/finexblock-dev/gofinexblock/pkg/safety"
)

func main() {
	safety.GracefullyStopBootstrap(bootstrap.EthereumBootstrap)
}