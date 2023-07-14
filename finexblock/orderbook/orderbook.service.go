package orderbook

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/cache"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
)

type service struct {
	queue Queue
	cache.DefaultKeyValueStore[grpc_order.Order]
}

func (s *service) LimitAsk(ask *grpc_order.Order) {
	var order *grpc_order.Order
	var err error
	order, err = s.queue.LimitAskInsert(ask)
	if err != nil {
		return
	}
	_ = s.Set(order.OrderUUID, order)
}

func (s *service) LimitBid(bid *grpc_order.Order) {
	s.queue.LimitBidInsert(bid)
}

func (s *service) MarketAsk(ask *grpc_order.Order) {
	s.queue.MarketAskInsert(ask)
}

func (s *service) MarketBid(bid *grpc_order.Order) {
	s.queue.MarketBidInsert(bid)
}

func (s *service) CancelOrder(uuid string) (order *grpc_order.Order, err error) {
	order, err = s.Get(uuid)
	if err == cache.ErrKeyNotFound {
		return nil, ErrOrderNotFound
	}

	switch order.OrderType {
	case grpc_order.OrderType_ASK:
		return s.queue.AskRemove(uuid)
	case grpc_order.OrderType_BID:
		return s.queue.BidRemove(uuid)
	default:
		return nil, ErrOrderTypeNotFound
	}
}

func (s *service) BidOrder() (bids []*grpc_order.Order, err error) {
	return s.queue.BidOrder()
}

func (s *service) AskOrder() (asks []*grpc_order.Order, err error) {
	return s.queue.AskOrder()
}