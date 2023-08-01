package main

import (
	"github.com/finexblock-dev/gofinexblock/cmd/polygon_key/internal/bootstrap"
	"github.com/finexblock-dev/gofinexblock/finexblock/safety"
	"log"
)

func main() {
	log.Println("Polygon key server start")

	safety.GracefullyStopBootstrap(bootstrap.PolygonKeyBootstrap)
}
