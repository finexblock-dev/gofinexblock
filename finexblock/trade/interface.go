package trade

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/goredis"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

type Service interface {
	AcquireLock(uuid, currency string) (bool, error)
	ReleaseLock(uuid, currency string) error
	GetBalance(uuid, currency string) (decimal.Decimal, error)
	SetBalance(uuid, currency string, amount decimal.Decimal) error
	PlusBalance(uuid, currency string, amount decimal.Decimal) error
	MinusBalance(uuid, currency string, amount decimal.Decimal) error
	SetOrder(orderUUID string, side string) error
	GetOrder(orderUUID string) (string, error)
	DeleteOrder(orderUUID string) error
}

func NewService(redisClient *redis.ClusterClient) Service {
	return newService(goredis.NewService(redisClient))
}
