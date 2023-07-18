package event

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/trade"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type engine struct {
	eventManager grpc_order.EventClient
	tradeManager trade.Manager
}

func newEngine(cluster *redis.ClusterClient, conn *grpc.ClientConn) *engine {
	return &engine{eventManager: grpc_order.NewEventClient(conn), tradeManager: trade.New(cluster)}
}