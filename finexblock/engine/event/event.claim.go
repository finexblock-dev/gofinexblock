package event

import (
	"fmt"
	"github.com/finexblock-dev/gofinexblock/finexblock/goaws"
	"github.com/finexblock-dev/gofinexblock/finexblock/trade"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
	"log"
	"time"
)

func (e *engine) Claim() {
	var xMessages []redis.XMessage
	var xPending *redis.XPending
	var err error
	var group = trade.EventGroup
	var claimer = e.Claimer(trade.EventClaimer)

	for {

		for _, stream := range streams {
			xPending, err = e.ReadPendingStream(stream, group)
			if err != nil {
				continue
			}

			if xPending.Count == 0 {
				continue
			}

			xMessages, err = e.ClaimStream(stream, group, claimer, time.Minute, []string{xPending.Lower})
			if err != nil {
				continue
			}

			for _, xMessage := range xMessages {
				go func(message redis.XMessage) {
					var event proto.Message
					var _err error

					event, _err = e.ParseMessage(stream, message)
					if _err != nil {
						// FIXME: _error handling
						return
					}

					if _err = e.Do(stream, event); _err != nil {
						log.Println("DO ERROR:", stream, _err)
						return
					}

					log.Println(stream.String(), "ACK:", e.tradeManager.AckStream(stream, group, message.ID), "\n", "IN EVENT CLAIMER", "\n", "GROUP:", group, "\n", "CLAIMER:", claimer, "\n", "MESSAGE ID:", message.ID, "\n", "MESSAGE VALUES:", message.Values, "\n", "STREAM:", stream)
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