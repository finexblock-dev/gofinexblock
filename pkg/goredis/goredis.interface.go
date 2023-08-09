package goredis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Repository interface {
	XClaim(ctx context.Context, args *redis.XClaimArgs) ([]redis.XMessage, error)
	XInfoStream(ctx context.Context, stream string) (*redis.XInfoStream, error)
	XPending(ctx context.Context, stream, group string) (*redis.XPending, error)
	XRange(ctx context.Context, stream, start, end string) ([]redis.XMessage, error)
	TxPipeline() redis.Pipeliner
	XReadGroup(ctx context.Context, args *redis.XReadGroupArgs) ([]redis.XStream, error)
	XDel(ctx context.Context, stream, group string, id string) error
	XAck(ctx context.Context, stream, group string, id string) error
	XAdd(ctx context.Context, args *redis.XAddArgs) error
	XGroupCreateConsumer(ctx context.Context, stream, group, consumer string) error
	XGroupCreate(ctx context.Context, stream, group string) error
	XGroupCreateMkStream(ctx context.Context, stream, group string) error
	Keys(ctx context.Context, pattern string) (result []string, err error)
	Get(ctx context.Context, key string) (value string, err error)
	SetNX(ctx context.Context, key string, value interface{}, exp time.Duration) (ok bool, err error)
	Set(ctx context.Context, key string, value string, exp time.Duration) (err error)
	Del(ctx context.Context, key string) (err error)

	XAddPipeline(tx redis.Pipeliner, ctx context.Context, args *redis.XAddArgs) error
}

type Service interface {
	XInfoStream(stream string) (*redis.XInfoStream, error)
	XRange(stream, start, end string) ([]redis.XMessage, error)
	XClaim(args *redis.XClaimArgs) ([]redis.XMessage, error)
	XPending(stream, group string) (result *redis.XPending, err error)
	XReadGroup(args *redis.XReadGroupArgs) (result []redis.XStream, err error)
	XDel(stream, group string, id string) (err error)
	XAck(stream, group string, id string) (err error)
	XAdd(args *redis.XAddArgs) (err error)
	XGroupCreate(stream, group string) (err error)
	XGroupCreateConsumer(stream, group, consumer string) error
	XGroupCreateMkStream(stream, group string) (err error)
	Get(key string) (value string, err error)
	Set(key string, value string, exp time.Duration) (err error)
	SetNX(key string, value interface{}, exp time.Duration) (ok bool, err error)
	Del(key string) (err error)
	Keys(pattern string) (result []string, err error)

	TxPipeline() redis.Pipeliner
	XAddPipeline(tx redis.Pipeliner, ctx context.Context, args *redis.XAddArgs) (err error)
}

func NewRepository(cluster *redis.ClusterClient) Repository {
	return newRepository(cluster)
}

func NewService(cluster *redis.ClusterClient) Service {
	return newService(cluster)
}