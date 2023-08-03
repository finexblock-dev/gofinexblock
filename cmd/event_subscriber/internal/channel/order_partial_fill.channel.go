package channel

import (
	"github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/pkg/order"
	"log"
	"time"
)

type OrderPartialFillChannel struct {
	sub chan *grpc_order.OrderPartialFill
	svc order.Service
}

func NewOrderPartialFillChannel(svc order.Service) *OrderPartialFillChannel {
	return &OrderPartialFillChannel{svc: svc, sub: make(chan *grpc_order.OrderPartialFill, 100000)}
}

func (o *OrderPartialFillChannel) Subscribe() {
	for {

		ticker := time.NewTicker(time.Millisecond * 800)
		var stack []*grpc_order.OrderPartialFill
		for {
			select {
			case <-ticker.C:
				if len(stack) == 0 {
					continue
				}
				remain, err := o.svc.LimitOrderPartialFillInBatch(stack)
				if err != nil {
					o.Send(stack, err)
					continue
				}
				if len(remain) > 0 {
					stack = remain
					continue
				}
				stack = []*grpc_order.OrderPartialFill{}
			case input := <-o.sub:
				stack = append(stack, input)
			}
		}
	}
}

func (o *OrderPartialFillChannel) Receive(event *grpc_order.OrderPartialFill) {
	o.sub <- event
}

func (o *OrderPartialFillChannel) Send(events []*grpc_order.OrderPartialFill, err error) {
	log.Println(err)
}