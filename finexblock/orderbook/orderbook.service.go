package orderbook

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/cache"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
)

type service struct {
	repository Repository
	cache.DefaultKeyValueStore[grpc_order.Order]
}
