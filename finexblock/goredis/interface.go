package goredis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type Repository interface {
	XReadGroup(ctx context.Context, args *redis.XReadGroupArgs) (result []redis.XStream, err error)
	XDel(ctx context.Context, stream, group string, id string) (err error)
	XAck(ctx context.Context, stream, group string, id string) (err error)
	XAdd(ctx context.Context, args *redis.XAddArgs) (err error)
	XGroupCreate(ctx context.Context, stream, group string) (err error)
	XGroupCreateMkStream(ctx context.Context, stream, group string) (err error)
	Get(ctx context.Context, key string) (value string, err error)
	Set(ctx context.Context, key string, value interface{}, exp time.Duration) (err error)
	SetNX(ctx context.Context, key string, value interface{}, exp time.Duration) (ok bool, err error)
	Del(ctx context.Context, key string) (err error)
}

type Service interface {
	XReadGroup(args *redis.XReadGroupArgs) (result []redis.XStream, err error)
	XDel(stream, group string, id string) (err error)
	XAck(stream, group string, id string) (err error)
	XAdd(args *redis.XAddArgs) (err error)
	XGroupCreate(stream, group string) (err error)
	XGroupCreateMkStream(stream, group string) (err error)
	Get(key string) (value string, err error)
	Set(key string, value interface{}, exp time.Duration) (err error)
	SetNX(key string, value interface{}, exp time.Duration) (ok bool, err error)
	Del(key string) (err error)
}
