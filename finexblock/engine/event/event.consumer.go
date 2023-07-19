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

var (
	streams = []types.Stream{
		trade.OrderInitializeStream,
		trade.OrderPlacementStream,
		trade.OrderFulfillmentStream,
		trade.OrderPartialFillStream,
		trade.OrderMatchingStream,
		trade.MarketOrderMatchingStream,
	}
)

func (e *engine) Consume() {
	var event proto.Message
	var xStreams []redis.XStream
	var err error

	var group = trade.EventGroup
	var consumer = e.Consumer(trade.EventConsumer)

	for {
		for _, s := range streams {
			xStreams, err = e.ReadStreams([]types.Stream{s}, group, consumer, 1000, -1)
			if err != nil {
				continue
			}

			for _, stream := range xStreams {
				for _, message := range stream.Messages {
					go func(message redis.XMessage) {
						event, err = e.ParseMessage(types.Stream(stream.Stream), message)
						if err != nil {
							// FIXME: error handling
							return
						}
						if err = e.Do(types.Stream(stream.Stream), event); err != nil {
							log.Println("DO ERROR:", stream.Stream, err)
							return
						}

						log.Println("ACK:", e.tradeManager.AckStream(types.Stream(stream.Stream), group, message.ID))
					}(message)
				}
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
	return types.Consumer(fmt.Sprintf("%s:%s", consumer, privateIP))
}

func (e *engine) ReadStreams(streams []types.Stream, group types.Group, consumer types.Consumer, count int64, block time.Duration) ([]redis.XStream, error) {
	return e.tradeManager.ReadStreams(streams, group, consumer, count, block)
}