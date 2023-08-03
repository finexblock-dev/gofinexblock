package event

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/trade"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type engine struct {
	chartServer     grpc_order.EventClient
	eventSubscriber grpc_order.EventClient
	tradeManager    trade.Manager
}

func newEngine(cluster *redis.ClusterClient, eventSubscriber, chartServer *grpc.ClientConn) *engine {
	return &engine{
		chartServer:     grpc_order.NewEventClient(chartServer),
		eventSubscriber: grpc_order.NewEventClient(eventSubscriber),
		tradeManager:    trade.New(cluster),
	}
}