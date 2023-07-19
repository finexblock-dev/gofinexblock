package refund

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

func (e *engine) Claim() {
	for {
		var xMessages []redis.XMessage
		var xPending *redis.XPending
		var err error

		var claimer = e.Claimer(trade.BalanceUpdateClaimer)

		xPending, err = e.ReadPendingStream(trade.BalanceUpdateStream, trade.BalanceUpdateGroup)
		if err != nil {
			continue
		}

		if xPending.Count == 0 {
			continue
		}

		xMessages, err = e.ClaimStream(trade.BalanceUpdateStream, trade.BalanceUpdateGroup, claimer, time.Minute, []string{xPending.Lower})
		if err != nil {
			continue
		}

		for _, xMessage := range xMessages {
			go func(message redis.XMessage) {
				var _err error
				var event = new(grpc_order.BalanceUpdate)
				event, _err = e.ParseMessage(message)
				if _err != nil {
					// FIXME: _error handling
					return
				}

				if _err = e.Do(event); _err != nil {
					log.Println("DO ERROR:", trade.BalanceUpdateStream, _err)
					return
				}

				// FIXME: error handling
				log.Println(trade.BalanceUpdateStream, "ACK:", e.tradeManager.AckStream(trade.BalanceUpdateStream, trade.BalanceUpdateGroup, message.ID))
			}(xMessage)
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
	return types.Claimer(fmt.Sprintf("%s:%s", claimer.String(), privateIP))
}

func (e *engine) ReadPendingStream(stream types.Stream, group types.Group) (*redis.XPending, error) {
	return e.tradeManager.ReadPendingStream(stream, group)
}

func (e *engine) ClaimStream(stream types.Stream, group types.Group, claimer types.Claimer, minIdleTime time.Duration, ids []string) ([]redis.XMessage, error) {
	return e.tradeManager.ClaimStream(stream, group, types.Consumer(claimer), minIdleTime, ids)
}