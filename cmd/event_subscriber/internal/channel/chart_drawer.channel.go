package channel

import (
	"github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/pkg/order"
	"log"
	"time"
)

type ChartDrawerChannel struct {
	sub chan *grpc_order.OrderMatching
	svc order.Service
}

func NewChartDrawerChannel(order order.Service) *ChartDrawerChannel {
	return &ChartDrawerChannel{
		sub: make(chan *grpc_order.OrderMatching, 100000),
		svc: order,
	}
}

func (c *ChartDrawerChannel) Subscribe() {
	for {
		ticker := time.NewTicker(time.Millisecond * 500)
		var stack []*grpc_order.OrderMatching
		for {
			select {
			case <-ticker.C:
				if len(stack) == 0 {
					continue
				}
				if err := c.svc.ChartDraw(stack); err != nil {
					c.Send(stack, err)
					continue
				}
				stack = []*grpc_order.OrderMatching{}
			case input := <-c.sub:
				stack = append(stack, input)
			}
		}
	}
}

func (c *ChartDrawerChannel) Receive(event *grpc_order.OrderMatching) {
	c.sub <- event
}

func (c *ChartDrawerChannel) Send(events []*grpc_order.OrderMatching, err error) {
	log.Println(err)
}