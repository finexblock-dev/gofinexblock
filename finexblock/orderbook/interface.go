package orderbook

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/safety"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

// Manager is interface for order book manager, use Service, and Service use Repository.
// Receive request and control order book
type Manager interface {
	safety.Subscriber
	LimitAskInsert(ask *grpc_order.Order) (order *grpc_order.Order, err error)
	LimitBidInsert(bid *grpc_order.Order) (order *grpc_order.Order, err error)
	MarketAskInsert(ask *grpc_order.Order) (order *grpc_order.Order, err error)
	MarketBidInsert(bid *grpc_order.Order) (order *grpc_order.Order, err error)
	CancelOrder(uuid string) (order *grpc_order.Order, err error)
	BidOrder() (bids []*grpc_order.Order, err error)
	AskOrder() (asks []*grpc_order.Order, err error)
}

// Service is interface for order book service, control repository.
// Also has market price for each order book(ask/bid).
type Service interface {
	LimitAsk(ask *grpc_order.Order) error                         // 지정가 매도 주문
	LimitBid(bid *grpc_order.Order) error                         // 지정가 매수 주문
	MarketAsk(ask *grpc_order.Order) error                        // 시장가 매도 주문
	MarketBid(bid *grpc_order.Order) error                        // 시장가 매수 주문
	CancelOrder(uuid string) (order *grpc_order.Order, err error) // 주문 취소 요청
	BidOrder() (bids []*grpc_order.Order, err error)              // 매수 주문 리스트
	AskOrder() (asks []*grpc_order.Order, err error)              // 매도 주문 리스트
}

// Repository is interface for order book repository
type Repository interface {
	PushAsk(order *grpc_order.Order)
	PushBid(order *grpc_order.Order)
	PopAsk() (order *grpc_order.Order)
	PopBid() (order *grpc_order.Order)
	RemoveAsk(uuid string) (order *grpc_order.Order)
	RemoveBid(uuid string) (order *grpc_order.Order)

	BidMarketPrice() decimal.Decimal
	AskMarketPrice() decimal.Decimal

	BidOrder() []*grpc_order.Order
	AskOrder() []*grpc_order.Order
}

func NewRepository() Repository {
	return newRepository()
}

func NewService(cluster *redis.ClusterClient) Service {
	return newService(cluster)
}

func New(cluster *redis.ClusterClient) Manager {
	return newManager(cluster)
}