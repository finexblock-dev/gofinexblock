package refund

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/stream"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type Engine interface {
	stream.SingleStreamConsumer
	stream.Claimer

	ParseMessage(message redis.XMessage) (event *grpc_order.BalanceUpdate, err error)
	Do(event *grpc_order.BalanceUpdate) (err error)
}

func New(cluster *redis.ClusterClient, conn *grpc.ClientConn) Engine {
	return newEngine(cluster, conn)
}