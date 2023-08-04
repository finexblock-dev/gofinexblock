package channel

import (
	"github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/pkg/order"
	"log"
	"time"
)

type OrderInitializeChannel struct {
	sub chan *grpc_order.OrderInitialize
	svc order.Service
}

func NewOrderInitializeChannel(svc order.Service) *OrderInitializeChannel {
	return &OrderInitializeChannel{svc: svc, sub: make(chan *grpc_order.OrderInitialize, 100000)}
}

func (o *OrderInitializeChannel) Subscribe() {
	for {
		ticker := time.NewTicker(time.Millisecond * 10)
		var stack []*grpc_order.OrderInitialize
		for {
			select {
			case <-ticker.C:
				if len(stack) == 0 {
					continue
				}
				if err := o.svc.LimitOrderInitializeInBatch(stack); err != nil {
					o.Send(stack, err)
					continue
				}
				stack = []*grpc_order.OrderInitialize{}
			case input := <-o.sub:
				stack = append(stack, input)
			}
		}
	}
}

func (o *OrderInitializeChannel) Receive(event *grpc_order.OrderInitialize) {
	o.sub <- event
}

func (o *OrderInitializeChannel) Send(events []*grpc_order.OrderInitialize, err error) {
	log.Println(err)
}