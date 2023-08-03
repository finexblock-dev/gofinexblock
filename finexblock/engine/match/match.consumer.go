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
	var consumer = e.Consumer(trade.MatchConsumer)
	for {
		var xStreams []redis.XStream
		var err error

		xStreams, err = e.ReadStream(trade.MatchStream, trade.MatchGroup, consumer, 1000, 0)
		if err != nil {
			log.Printf("failed to read stream: %v", err)
			// FIXME: error handling
			continue
		}

		for _, xStream := range xStreams {
			for _, message := range xStream.Messages {
				go func(message redis.XMessage) {
					var _case types.Case
					var pair = new(grpc_order.BidAsk)
					var _err error

					_case, pair, _err = e.ParseMessage(message)
					if _err != nil {
						log.Printf("failed to parse message: %v", _err)
						return
					}

					if _err = e.Do(_case, pair); _err != nil {
						log.Printf("failed to do: %v", _err)
						return
					}

					log.Println(xStream.Stream, "ACK:", e.tradeManager.AckStream(trade.MatchStream, trade.MatchGroup, message.ID))
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