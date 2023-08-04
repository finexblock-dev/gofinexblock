package channel

import (
	"github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/pkg/order"
	"log"
	"time"
)

type OrderFulfillmentChannel struct {
	sub chan *grpc_order.OrderFulfillment
	svc order.Service
}

func NewOrderFulfillmentChannel(svc order.Service) *OrderFulfillmentChannel {
	return &OrderFulfillmentChannel{svc: svc, sub: make(chan *grpc_order.OrderFulfillment, 100000)}
}

func (o *OrderFulfillmentChannel) Subscribe() {
	for {
		ticker := time.NewTicker(time.Millisecond * 800)
		var stack []*grpc_order.OrderFulfillment
		for {
			select {
			case <-ticker.C:
				if len(stack) == 0 {
					continue
				}
				remain, err := o.svc.LimitOrderFulfillmentInBatch(stack)
				if err != nil {
					o.Send(stack, err)
					continue
				}
				if len(remain) > 0 {
					stack = remain
					continue
				}
				stack = []*grpc_order.OrderFulfillment{}
			case input := <-o.sub:
				stack = append(stack, input)
			}
		}

	}
}

func (o *OrderFulfillmentChannel) Receive(event *grpc_order.OrderFulfillment) {
	o.sub <- event
}

func (o *OrderFulfillmentChannel) Send(events []*grpc_order.OrderFulfillment, err error) {
	log.Println(err)
}