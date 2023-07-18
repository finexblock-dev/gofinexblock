package cancellation

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/trade"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type engine struct {
	tradeManager trade.Manager
	eventManager grpc_order.EventClient
}

func newEngine(cluster *redis.ClusterClient, conn *grpc.ClientConn) *engine {
	return &engine{tradeManager: trade.NewManager(cluster), eventManager: grpc_order.NewEventClient(conn)}
}