package orderbook

import "github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"

// Repository is interface for orderbook repository
type Repository interface {
	PushAsk(order *grpc_order.Order)
	PushBid(order *grpc_order.Order)
	PopAsk() (order *grpc_order.Order)
	PopBid() (order *grpc_order.Order)
	RemoveAsk(uuid string) (order *grpc_order.Order)
	RemoveBid(uuid string) (order *grpc_order.Order)

	BidOrder() []*grpc_order.Order
	AskOrder() []*grpc_order.Order
}

// Queue is interface for orderbook queue
type Queue interface {
	LimitAskInsert(ask *grpc_order.Order) (order *grpc_order.Order, err error)
	LimitBidInsert(bid *grpc_order.Order) (order *grpc_order.Order, err error)
	MarketAskInsert(ask *grpc_order.Order) (order *grpc_order.Order, err error)
	MarketBidInsert(bid *grpc_order.Order) (order *grpc_order.Order, err error)
	AskRemove(uuid string) (order *grpc_order.Order, err error)
	BidRemove(uuid string) (order *grpc_order.Order, err error)
	BidOrder() (bids []*grpc_order.Order, err error)
	AskOrder() (asks []*grpc_order.Order, err error)
}

// Service is interface for orderbook service
type Service interface {
	LimitAsk(ask *grpc_order.Order)
	LimitBid(bid *grpc_order.Order)
	MarketAsk(ask *grpc_order.Order)
	MarketBid(bid *grpc_order.Order)
	CancelOrder(uuid string) (order *grpc_order.Order, err error)
	BidOrder() (bids []*grpc_order.Order, err error)
	AskOrder() (asks []*grpc_order.Order, err error)
}

func NewRepository() Repository {
	return newRepository()
}