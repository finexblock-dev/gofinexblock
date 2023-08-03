package refund

import (
	"github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/pkg/types"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type Engine interface {
	types.SingleStreamConsumer
	types.SingleStreamClaimer

	ParseMessage(message redis.XMessage) (event *grpc_order.BalanceUpdate, err error)
	Do(event *grpc_order.BalanceUpdate) (err error)
}

func New(cluster *redis.ClusterClient, eventSubscriber, chartServer *grpc.ClientConn) Engine {
	return newEngine(cluster, eventSubscriber, chartServer)
}