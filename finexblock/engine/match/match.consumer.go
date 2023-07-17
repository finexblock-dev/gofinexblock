package match

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/trade"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/redis/go-redis/v9"
	"time"
)

func (e *engine) Consume() {
	for {
		var _case types.Case
		var xStreams []redis.XStream
		var err error

		var pair = new(grpc_order.BidAsk)

		xStreams, err = e.ReadStream(trade.MatchStream, trade.MatchGroup, trade.MatchConsumer, 1, 0)
		if err != nil {
			// FIXME: error handling
			continue
		}

		for _, xStream := range xStreams {
			for _, message := range xStream.Messages {
				go func(message redis.XMessage) {
					_case, pair, err = e.ParseMessage(message)
					if err != nil {
						// FIXME: error handling
						return
					}

					if err = e.Do(_case, pair); err != nil {
						// FIXME: error handling
						e.tradeManager.AckStream(trade.MatchStream, trade.MatchGroup, message.ID)
					}
				}(message)
			}
		}
	}
}

func (e *engine) ReadStream(stream types.Stream, group types.Group, consumer types.Consumer, count int64, block time.Duration) ([]redis.XStream, error) {
	return e.tradeManager.ReadStream(trade.MatchStream, trade.MatchGroup, trade.MatchConsumer, 1, 0)
}