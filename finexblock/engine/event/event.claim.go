package event

import (
	"fmt"
	"github.com/finexblock-dev/gofinexblock/finexblock/goaws"
	"github.com/finexblock-dev/gofinexblock/finexblock/trade"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
	"time"
)

func (e *engine) Claim() {
	for {
		var xMessages []redis.XMessage
		var xPending *redis.XPending
		var event proto.Message
		var err error

		var group = trade.EventGroup
		var claimer = e.Claimer(trade.EventClaimer)

		for _, stream := range streams {
			xPending, err = e.ReadPendingStream(stream, group)
			if err != nil {
				continue
			}

			if xPending.Count == 0 {
				continue
			}

			xMessages, err = e.ClaimStream(stream, group, claimer, 0, []string{xPending.Lower})
			if err != nil {
				continue
			}

			for _, xMessage := range xMessages {
				go func(message redis.XMessage) {
					event, err = e.ParseMessage(stream, message)
					if err != nil {
						// FIXME: error handling
						return
					}

					if err = e.Do(stream, event); err != nil {
						// FIXME: error handling
						return
					}

					// FIXME: error handling
					_ = e.tradeManager.AckStream(stream, group, message.ID)
				}(xMessage)
			}
		}
	}
}

func (e *engine) Claimer(claimer types.Claimer) types.Claimer {
	var privateIP string
	var err error
	privateIP, err = goaws.OwnPrivateIP()
	if err != nil {
		panic(err)
	}
	return types.Claimer(fmt.Sprintf("%s:%s", claimer, privateIP))
}

func (e *engine) ReadPendingStream(stream types.Stream, group types.Group) (*redis.XPending, error) {
	return e.tradeManager.ReadPendingStream(stream, group)

}

func (e *engine) ClaimStream(stream types.Stream, group types.Group, claimer types.Claimer, minIdleTime time.Duration, ids []string) ([]redis.XMessage, error) {
	return e.tradeManager.ClaimStream(stream, group, types.Consumer(claimer), minIdleTime, ids)

}