package match

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/stream"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/redis/go-redis/v9"
)

type Engine interface {
	stream.Consumer
	stream.Claimer

	ParseMessage(message redis.XMessage) (_case types.Case, pair *grpc_order.BidAsk, err error)
	Do(_case types.Case, pair *grpc_order.BidAsk) (err error)

	LimitAskBigger(pair *grpc_order.BidAsk) (err error)
	LimitAskEqual(pair *grpc_order.BidAsk) (err error)
	LimitAskSmaller(pair *grpc_order.BidAsk) (err error)
	LimitBidBigger(pair *grpc_order.BidAsk) (err error)
	LimitBidEqual(pair *grpc_order.BidAsk) (err error)
	LimitBidSmaller(pair *grpc_order.BidAsk) (err error)
	MarketAskBigger(pair *grpc_order.BidAsk) (err error)
	MarketAskEqual(pair *grpc_order.BidAsk) (err error)
	MarketAskSmaller(pair *grpc_order.BidAsk) (err error)
	MarketBidBigger(pair *grpc_order.BidAsk) (err error)
	MarketBidEqual(pair *grpc_order.BidAsk) (err error)
	MarketBidSmaller(pair *grpc_order.BidAsk) (err error)
}

func New(cluster *redis.ClusterClient) Engine {
	return newEngine(cluster)
}