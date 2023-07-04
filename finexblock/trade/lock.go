package trade

import (
	"context"
	"time"
)

func (a *tradeService) AcquireLock(ctx context.Context, uuid, currency string) (bool, error) {
	var key string

	key = getAccountLockKey(uuid, currency)
	return a.redisClient.SetNX(ctx, key, lock, time.Second*10).Result()
}

func (a *tradeService) ReleaseLock(ctx context.Context, uuid, currency string) error {
	var key string

	key = getAccountLockKey(uuid, currency)
	return a.redisClient.Del(ctx, key).Err()
}