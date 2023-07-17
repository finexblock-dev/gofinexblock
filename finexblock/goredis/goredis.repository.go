package goredis

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

type repository struct {
	cluster *redis.ClusterClient
}

func newRepository(cluster *redis.ClusterClient) *repository {
	return &repository{cluster: cluster}
}

func (r *repository) XClaim(ctx context.Context, args *redis.XClaimArgs) ([]redis.XMessage, error) {
	return r.cluster.XClaim(ctx, args).Result()
}

func (r *repository) XInfoStream(ctx context.Context, stream string) (*redis.XInfoStream, error) {
	return r.cluster.XInfoStream(ctx, stream).Result()
}

func (r *repository) XPending(ctx context.Context, stream, group string) (*redis.XPending, error) {
	return r.cluster.XPending(ctx, stream, group).Result()
}

func (r *repository) TxPipeline() redis.Pipeliner {
	return r.cluster.TxPipeline()
}

func (r *repository) XReadGroup(ctx context.Context, args *redis.XReadGroupArgs) ([]redis.XStream, error) {
	return r.cluster.XReadGroup(ctx, args).Result()
}

func (r *repository) XDel(ctx context.Context, stream, group string, id string) error {
	return r.cluster.XDel(ctx, stream, group, id).Err()
}

func (r *repository) XAck(ctx context.Context, stream, group string, id string) error {
	return r.cluster.XAck(ctx, stream, group, id).Err()
}

func (r *repository) XAdd(ctx context.Context, args *redis.XAddArgs) error {
	return r.cluster.XAdd(ctx, args).Err()
}

func (r *repository) XGroupCreateConsumer(ctx context.Context, stream, group, consumer string) error {
	return r.cluster.XGroupCreateConsumer(ctx, stream, group, consumer).Err()
}

func (r *repository) XGroupCreate(ctx context.Context, stream, group string) error {
	return r.cluster.XGroupCreate(ctx, stream, group, "$").Err()
}

func (r *repository) XGroupCreateMkStream(ctx context.Context, stream, group string) error {
	return r.cluster.XGroupCreateMkStream(ctx, stream, group, "$").Err()
}

func (r *repository) Get(ctx context.Context, key string) (value string, err error) {
	return r.cluster.Get(ctx, key).Result()
}

func (r *repository) Set(ctx context.Context, key string, value interface{}, exp time.Duration) (err error) {
	var bytes []byte

	bytes, err = json.Marshal(value)
	if err != nil {
		return err
	}

	return r.cluster.Set(ctx, key, string(bytes), exp).Err()
}

func (r *repository) SetNX(ctx context.Context, key string, value interface{}, exp time.Duration) (ok bool, err error) {
	var bytes []byte

	bytes, err = json.Marshal(value)
	if err != nil {
		return ok, err
	}

	return r.cluster.SetNX(ctx, key, string(bytes), exp).Result()
}

func (r *repository) Del(ctx context.Context, key string) (err error) {
	return r.cluster.Del(ctx, key).Err()
}