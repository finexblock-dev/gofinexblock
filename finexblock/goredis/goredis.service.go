package goredis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type service struct {
	repository Repository
}

func (s *service) XInfoStream(stream string) (*redis.XInfoStream, error) {
	return s.repository.XInfoStream(context.Background(), stream)
}

func (s *service) XClaim(args *redis.XClaimArgs) ([]redis.XMessage, error) {
	return s.repository.XClaim(context.Background(), args)
}

func (s *service) XPending(stream, group string) (result *redis.XPending, err error) {
	return s.repository.XPending(context.Background(), stream, group)
}

func (s *service) XReadGroup(args *redis.XReadGroupArgs) (result []redis.XStream, err error) {
	return s.repository.XReadGroup(context.Background(), args)
}

func (s *service) XDel(stream, group string, id string) (err error) {
	return s.repository.XDel(context.Background(), stream, group, id)
}

func (s *service) XAck(stream, group string, id string) (err error) {
	return s.repository.XAck(context.Background(), stream, group, id)
}

func (s *service) XAdd(args *redis.XAddArgs) (err error) {
	return s.repository.XAdd(context.Background(), args)
}

func (s *service) XGroupCreate(stream, group string) (err error) {
	return s.repository.XGroupCreate(context.Background(), stream, group)
}

func (s *service) XGroupCreateConsumer(stream, group, consumer string) error {
	return s.repository.XGroupCreateConsumer(context.Background(), stream, group, consumer)
}

func (s *service) XGroupCreateMkStream(stream, group string) (err error) {
	return s.repository.XGroupCreateMkStream(context.Background(), stream, group)
}

func (s *service) Get(key string) (value string, err error) {
	return s.repository.Get(context.Background(), key)
}

func (s *service) Set(key string, value interface{}, exp time.Duration) (err error) {
	return s.repository.Set(context.Background(), key, value, exp)
}

func (s *service) SetNX(key string, value interface{}, exp time.Duration) (ok bool, err error) {
	return s.repository.SetNX(context.Background(), key, value, exp)
}

func (s *service) Del(key string) (err error) {
	return s.repository.Del(context.Background(), key)
}

func newService(cluster *redis.ClusterClient) *service {
	return &service{repository: newRepository(cluster)}
}