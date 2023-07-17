package stream

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/redis/go-redis/v9"
	"time"
)

type Claimer interface {
	ReadPendingStream(stream types.Stream, group types.Group) (*redis.XPending, error)
	ClaimStream(stream types.Stream, group types.Group, consumer types.Consumer, minIdleTime time.Duration, ids []string) ([]redis.XMessage, error)
}

type Consumer interface {
	ReadStream(stream types.Stream, group types.Group, consumer types.Consumer, count int64, block time.Duration) ([]redis.XStream, error)
}