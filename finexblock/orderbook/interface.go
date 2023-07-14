package orderbook

import "github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"

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

type Queue interface {
	LimitAskInsert(ask *grpc_order.Order)
	LimitBidInsert(bid *grpc_order.Order)
	MarketAskInsert(ask *grpc_order.Order)
	MarketBidInsert(bid *grpc_order.Order)
	AskRemove(uuid string) (order *grpc_order.Order)
	BidRemove(uuid string) (order *grpc_order.Order)
}

type Service interface {
	LimitAsk(ask *grpc_order.Order)
	LimitBid(bid *grpc_order.Order)
	MarketAsk(ask *grpc_order.Order)
	MarketBid(bid *grpc_order.Order)
	CancelOrder(uuid string) (order *grpc_order.Order, err error)
}

func NewRepository() Repository {
	return newRepository()
}