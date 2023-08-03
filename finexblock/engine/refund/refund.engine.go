package refund

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/trade"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type engine struct {
	tradeManager                 trade.Manager
	eventSubscriber, chartServer grpc_order.EventClient
}

func newEngine(cluster *redis.ClusterClient, eventSubscriber, chartServer *grpc.ClientConn) *engine {
	return &engine{tradeManager: trade.New(cluster), eventSubscriber: grpc_order.NewEventClient(eventSubscriber), chartServer: grpc_order.NewEventClient(chartServer)}
}