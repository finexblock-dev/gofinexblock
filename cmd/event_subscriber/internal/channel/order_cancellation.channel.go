package channel

import (
	"github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/pkg/order"
	"log"
	"time"
)

type OrderCancellationChannel struct {
	sub chan *grpc_order.OrderCancelled
	svc order.Service
}

func NewOrderCancellationChannel(svc order.Service) *OrderCancellationChannel {
	return &OrderCancellationChannel{svc: svc, sub: make(chan *grpc_order.OrderCancelled, 100000)}
}

// Subscribe : Subscribe events and insert events to database.
func (o *OrderCancellationChannel) Subscribe() {
	for {
		ticker := time.NewTicker(time.Millisecond * 500)
		var stack []*grpc_order.OrderCancelled
		for {
			select {
			case <-ticker.C:
				if len(stack) == 0 {
					continue
				}
				remain, err := o.svc.LimitOrderCancellationInBatch(stack)
				if err != nil {
					o.Send(stack, err)
					continue
				}

				if len(remain) > 0 {
					stack = remain
					continue
				}
				stack = []*grpc_order.OrderCancelled{}
			case input := <-o.sub:
				stack = append(stack, input)
			}
		}
	}
}

// Receive : Receive event and send to its own channel.
func (o *OrderCancellationChannel) Receive(event *grpc_order.OrderCancelled) {
	o.sub <- event
}

// Send : Send error message to slack.
func (o *OrderCancellationChannel) Send(events []*grpc_order.OrderCancelled, err error) {
	log.Println(err)
}