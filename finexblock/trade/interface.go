package trade

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

type Service interface {
	AcquireLock(ctx context.Context, uuid, currency string) (bool, error)
	ReleaseLock(ctx context.Context, uuid, currency string) error
	GetBalance(ctx context.Context, uuid, currency string) (decimal.Decimal, error)
	SetBalance(ctx context.Context, uuid, currency string, amount decimal.Decimal) error
	PlusBalance(ctx context.Context, uuid, currency string, amount decimal.Decimal) error
	MinusBalance(ctx context.Context, uuid, currency string, amount decimal.Decimal) error
	SetOrder(ctx context.Context, orderUUID string, side string) error
	GetOrder(ctx context.Context, orderUUID string) (string, error)
	DeleteOrder(ctx context.Context, orderUUID string) error
}

type tradeService struct {
	redisClient *redis.ClusterClient
}

func newTradeService(redisClient *redis.ClusterClient) *tradeService {
	return &tradeService{redisClient: redisClient}
}

func NewService(redisClient *redis.ClusterClient) Service {
	return newTradeService(redisClient)
}