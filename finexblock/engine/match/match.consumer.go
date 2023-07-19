package match

import (
	"fmt"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/goaws"
	"github.com/finexblock-dev/gofinexblock/finexblock/trade"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

func (e *engine) Consume() {
	for {
		var _case types.Case
		var xStreams []redis.XStream
		var err error

		var pair = new(grpc_order.BidAsk)

		xStreams, err = e.ReadStream(trade.MatchStream, trade.MatchGroup, e.Consumer(trade.MatchConsumer), 1000, 0)
		if err != nil {
			log.Printf("failed to read stream: %v", err)
			// FIXME: error handling
			continue
		}

		for _, xStream := range xStreams {
			for _, message := range xStream.Messages {
				go func(message redis.XMessage) {
					_case, pair, err = e.ParseMessage(message)
					if err != nil {
						log.Printf("failed to parse message: %v", err)
						return
					}

					if err = e.Do(_case, pair); err != nil {
						log.Printf("failed to do: %v", err)
						return
					}

					log.Println("ACK:", e.tradeManager.AckStream(trade.MatchStream, trade.MatchGroup, message.ID))
				}(message)
			}
		}
	}
}

func (e *engine) Consumer(consumer types.Consumer) types.Consumer {
	var privateIP string
	var err error
	privateIP, err = goaws.OwnPrivateIP()
	if err != nil {
		panic(err)
	}
	return types.Consumer(fmt.Sprintf("%s:%s", consumer.String(), privateIP))
}

func (e *engine) ReadStream(stream types.Stream, group types.Group, consumer types.Consumer, count int64, block time.Duration) ([]redis.XStream, error) {
	return e.tradeManager.ReadStream(stream, group, consumer, count, block)
}