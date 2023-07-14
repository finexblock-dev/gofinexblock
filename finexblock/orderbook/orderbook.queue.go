package orderbook

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
)

type queue struct {
	repo Repository

	limitAsk  chan *grpc_order.Order // limitAsk is channel for limit ask order
	marketAsk chan *grpc_order.Order // marketAsk is channel for market ask order
	limitBid  chan *grpc_order.Order // limitBid is channel for limit bid order
	marketBid chan *grpc_order.Order // marketBid is channel for market bid order

	askRemove chan string // askRemove is channel for cancel ask order
	bidRemove chan string // bidRemove is channel for cancel bid order

	// bidMarketPrice
	// askMarketPrice
}

func (q *queue) Subscribe() {

	for {
		select {
		//case ask := <-q.limitAsk:
		//case bid := <-q.limitBid:
		//case ask := <-q.marketAsk:
		//case bid := <-q.marketBid:
		//case uuid := <-q.askRemove:
		//case uuid := <-q.bidRemove:
		}
	}
}