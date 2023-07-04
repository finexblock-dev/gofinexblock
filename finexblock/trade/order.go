package trade

import (
	"context"
	"fmt"
)

func (a *accountService) SetOrder(ctx context.Context, orderUUID string, side string) error {
	return a.redisClient.Set(ctx, fmt.Sprintf("order:%v", orderUUID), side, 0).Err()
}

func (a *accountService) GetOrder(ctx context.Context, orderUUID string) (string, error) {
	return a.redisClient.Get(ctx, fmt.Sprintf("order:%v", orderUUID)).Result()
}

func (a *accountService) DeleteOrder(ctx context.Context, orderUUID string) error {
	return a.redisClient.Del(ctx, fmt.Sprintf("order:%v", orderUUID)).Err()
}