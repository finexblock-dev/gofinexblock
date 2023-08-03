package match

import (
	"github.com/finexblock-dev/gofinexblock/pkg/trade"
	"github.com/redis/go-redis/v9"
)

type engine struct {
	tradeManager trade.Manager
}

func newEngine(cluster *redis.ClusterClient) *engine {
	return &engine{tradeManager: trade.New(cluster)}
}