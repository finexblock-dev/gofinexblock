package types

import (
	"context"
	"database/sql"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type Model interface {
	TableName() string
	Alias() string
}

type Repository interface {
	Tx(level sql.IsolationLevel) *gorm.DB
	Conn() *gorm.DB
}

type Service interface {
	Ctx() context.Context
	CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc)
	Tx(level sql.IsolationLevel) *gorm.DB
	Conn() *gorm.DB
}

type Heap[T any] interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Pop() T
	Push(x T)
}

type BaseClaimer interface {
	Claim()
	Claimer(claimer Claimer) Claimer
}

type SingleStreamClaimer interface {
	BaseClaimer
	ReadPendingStream(stream Stream, group Group) (*redis.XPending, error)
	ClaimStream(stream Stream, group Group, claimer Claimer, minIdleTime time.Duration, ids []string) ([]redis.XMessage, error)
}

type BaseConsumer interface {
	Consume()
	Consumer(consumer Consumer) Consumer
}

type SingleStreamConsumer interface {
	BaseConsumer
	ReadStream(stream Stream, group Group, consumer Consumer, count int64, block time.Duration) ([]redis.XStream, error)
}

type MultiStreamConsumer interface {
	BaseConsumer
	ReadStreams(streams []Stream, group Group, consumer Consumer, count int64, block time.Duration) ([]redis.XStream, error)
}
