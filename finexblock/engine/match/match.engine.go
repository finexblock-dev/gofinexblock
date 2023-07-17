package match

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/trade"
	"github.com/redis/go-redis/v9"
)

type engine struct {
	tradeManager trade.Manager
}

func newEngine(cluster *redis.ClusterClient) *engine {
	return &engine{tradeManager: trade.NewManager(cluster)}
}