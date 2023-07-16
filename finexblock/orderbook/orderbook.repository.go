package orderbook

import (
	"container/heap"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/shopspring/decimal"
)

type repository struct {
	askOrderBook *askOrderBook
	bidOrderBook *bidOrderBook
}

func (r *repository) BidMarketPrice() decimal.Decimal {
	if r.bidOrderBook.Len() > 0 {
		order := r.PopBid()
		if order == nil {
			return decimal.Zero
		}
		marketPrice := decimal.NewFromFloat(order.UnitPrice)
		r.PushBid(order)
		return marketPrice
	}
	return decimal.Zero
}

func (r *repository) AskMarketPrice() decimal.Decimal {
	if r.askOrderBook.Len() > 0 {
		order := r.PopAsk()
		if order == nil {
			return decimal.Zero
		}
		marketPrice := decimal.NewFromFloat(order.UnitPrice)
		r.PushAsk(order)
		return marketPrice
	}
	return decimal.Zero
}

func (r *repository) BidOrder() []*grpc_order.Order {
	return r.bidOrderBook.Requests
}

func (r *repository) AskOrder() []*grpc_order.Order {
	return r.askOrderBook.Requests
}

func (r *repository) PushAsk(order *grpc_order.Order) {
	heap.Push(r.askOrderBook, order)
}

func (r *repository) PushBid(order *grpc_order.Order) {
	heap.Push(r.bidOrderBook, order)
}

func (r *repository) PopAsk() (order *grpc_order.Order) {
	return heap.Pop(r.askOrderBook).(*grpc_order.Order)
}

func (r *repository) PopBid() (order *grpc_order.Order) {
	return heap.Pop(r.bidOrderBook).(*grpc_order.Order)
}

func (r *repository) RemoveAsk(uuid string) (order *grpc_order.Order) {
	if index, ok := r.askOrderBook.IndexMap[uuid]; ok {
		return heap.Remove(r.askOrderBook, index).(*grpc_order.Order)
	}

	return nil
}

func (r *repository) RemoveBid(uuid string) (order *grpc_order.Order) {
	if index, ok := r.bidOrderBook.IndexMap[uuid]; ok {
		return heap.Remove(r.bidOrderBook, index).(*grpc_order.Order)
	}

	return nil
}

func newRepository() *repository {
	var ask = newAskOrderBook()
	var bid = newBidOrderBook()

	heap.Init(ask)
	heap.Init(bid)

	return &repository{askOrderBook: newAskOrderBook(), bidOrderBook: newBidOrderBook()}
}
