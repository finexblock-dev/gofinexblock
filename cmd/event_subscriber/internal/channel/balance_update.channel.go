package channel

import (
	"github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/pkg/wallet"
	"log"
	"time"
)

type BalanceUpdateChannel struct {
	sub chan *grpc_order.BalanceUpdate
	svc wallet.Service
}

func NewBalanceUpdateChannel(svc wallet.Service) *BalanceUpdateChannel {
	return &BalanceUpdateChannel{svc: svc, sub: make(chan *grpc_order.BalanceUpdate, 100000)}
}

// Subscribe : Subscribe events and insert events to database.
func (b *BalanceUpdateChannel) Subscribe() {
	for {
		ticker := time.NewTicker(time.Millisecond * 500)
		var stack []*grpc_order.BalanceUpdate
		for {
			select {
			case <-ticker.C:
				if len(stack) == 0 {
					continue
				}
				if err := b.svc.BalanceUpdateInBatch(stack); err != nil {
					b.Send(stack, err)
					continue
				}
				stack = []*grpc_order.BalanceUpdate{}
			case input := <-b.sub:
				stack = append(stack, input)
			}
		}
	}
}

// Receive : Receive event and send to its own channel.
func (b *BalanceUpdateChannel) Receive(event *grpc_order.BalanceUpdate) {
	b.sub <- event
}

// Send : Send error message to slack.
func (b *BalanceUpdateChannel) Send(event []*grpc_order.BalanceUpdate, err error) {
	log.Println(err)
}