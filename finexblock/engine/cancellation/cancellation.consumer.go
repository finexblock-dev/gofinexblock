package cancellation

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
		var xStreams []redis.XStream
		var err error

		var consumer = e.Consumer(trade.OrderCancellationConsumer)

		xStreams, err = e.ReadStream(trade.OrderCancellationStream, trade.OrderCancellationGroup, consumer, 1000, 0)
		if err != nil {
			// FIXME: error handling
			continue
		}

		for _, xStream := range xStreams {
			for _, message := range xStream.Messages {
				go func(message redis.XMessage) {
					var _err error
					var event = new(grpc_order.OrderCancelled)
					event, _err = e.ParseMessage(message)
					if _err != nil {
						// FIXME: error handling
						return
					}

					if _err = e.Do(event); _err != nil {
						log.Println("DO ERROR:", trade.OrderCancellationStream, _err)
						return
					}

					log.Println(xStream.Stream, "ACK:", e.tradeManager.AckStream(trade.OrderCancellationStream, trade.OrderCancellationGroup, message.ID))
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