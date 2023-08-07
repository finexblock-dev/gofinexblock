package server

import (
	"github.com/finexblock-dev/gofinexblock/pkg/orderbook"
	"github.com/finexblock-dev/gofinexblock/pkg/safety"
	"github.com/finexblock-dev/gofinexblock/pkg/trade"
	"log"
)

func LoadOrderBook(manager orderbook.Manager) {
	//var err error
	//if err = manager.LoadOrderBook(); err != nil {
	//	log.Fatalln(err)
	//}

	go safety.InfinitySubscribe(manager)
	//go manager.SnapshotCron(time.Minute * 5)
}

func LoadStream(manager trade.Manager) {
	var err error

	if err = manager.StreamsInit(); err != nil {
		log.Println(err)
	}
}