package account

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
)

type FinexblockAccountService interface {
	AcquireLock(ctx context.Context, uuid, currency string) (bool, error)
	ReleaseLock(ctx context.Context, uuid, currency string) error
	GetBalance(ctx context.Context, uuid, currency string) (decimal.Decimal, error)
	SetBalance(ctx context.Context, uuid, currency string, amount decimal.Decimal) error
	PlusBalance(ctx context.Context, uuid, currency string, amount decimal.Decimal) error
	MinusBalance(ctx context.Context, uuid, currency string, amount decimal.Decimal) error
}

type accountService struct {
	redisClient *redis.ClusterClient
}

func newAccountService(redisClient *redis.ClusterClient) FinexblockAccountService {
	return &accountService{redisClient: redisClient}
}

func NewFinexblockAccountService(redisClient *redis.ClusterClient) FinexblockAccountService {
	return newAccountService(redisClient)
}
