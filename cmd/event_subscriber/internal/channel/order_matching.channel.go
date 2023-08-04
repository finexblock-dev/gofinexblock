package channel

import (
	"github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/pkg/order"
	"log"
	"time"
)

type OrderMatchingChannel struct {
	sub chan *grpc_order.OrderMatching
	svc order.Service
}

func NewOrderMatchingChannel(svc order.Service) *OrderMatchingChannel {
	return &OrderMatchingChannel{svc: svc, sub: make(chan *grpc_order.OrderMatching, 100000)}
}

func (o *OrderMatchingChannel) Subscribe() {
	for {
		ticker := time.NewTicker(time.Millisecond * 800)
		var stack []*grpc_order.OrderMatching
		for {
			select {
			case <-ticker.C:
				if len(stack) == 0 {
					continue
				}
				if err := o.svc.OrderMatchingEventInBatch(stack); err != nil {
					o.Send(stack, err)
					continue
				}
				stack = []*grpc_order.OrderMatching{}
			case input := <-o.sub:
				stack = append(stack, input)
			}
		}
	}
}

func (o *OrderMatchingChannel) Receive(event *grpc_order.OrderMatching) {
	o.sub <- event
}

func (o *OrderMatchingChannel) Send(events []*grpc_order.OrderMatching, err error) {
	log.Println(err)
}