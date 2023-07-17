package match

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/redis/go-redis/v9"
	"time"
)

func (e *engine) Claim() {
	//TODO implement me
	panic("implement me")
}

func (e *engine) ReadPendingStream(stream types.Stream, group types.Group) (*redis.XPending, error) {
	return e.tradeManager.ReadPendingStream(stream, group)
}

func (e *engine) ClaimStream(stream types.Stream, group types.Group, consumer types.Consumer, minIdleTime time.Duration, ids []string) ([]redis.XMessage, error) {
	return e.tradeManager.ClaimStream(stream, group, consumer, minIdleTime, ids)
}