package stream

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/redis/go-redis/v9"
	"time"
)

type BaseClaimer interface {
	Claim()
	Claimer(claimer types.Claimer) types.Claimer
}

type Claimer interface {
	BaseClaimer
	ReadPendingStream(stream types.Stream, group types.Group) (*redis.XPending, error)
	ClaimStream(stream types.Stream, group types.Group, claimer types.Claimer, minIdleTime time.Duration, ids []string) ([]redis.XMessage, error)
}

type BaseConsumer interface {
	Consume()
	Consumer(consumer types.Consumer) types.Consumer
}

type SingleStreamConsumer interface {
	BaseConsumer
	ReadStream(stream types.Stream, group types.Group, consumer types.Consumer, count int64, block time.Duration) ([]redis.XStream, error)
}

type MultiStreamConsumer interface {
	BaseConsumer
	ReadStreams(streams []types.Stream, group types.Group, consumer types.Consumer, count int64, block time.Duration) ([]redis.XStream, error)
}