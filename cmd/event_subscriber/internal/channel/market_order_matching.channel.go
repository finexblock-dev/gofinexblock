package channel

import (
	"github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/pkg/order"
	"log"
	"time"
)

type MarketOrderMatchingChannel struct {
	sub chan *grpc_order.MarketOrderMatching
	svc order.Service
}

func NewMarketOrderMatchingChannel(svc order.Service) *MarketOrderMatchingChannel {
	return &MarketOrderMatchingChannel{
		sub: make(chan *grpc_order.MarketOrderMatching, 100000),
		svc: svc,
	}
}

func (m *MarketOrderMatchingChannel) Subscribe() {
	ticker := time.NewTicker(time.Millisecond * 500)
	var stack []*grpc_order.MarketOrderMatching
	for {
		select {
		case <-ticker.C:
			if len(stack) == 0 {
				continue
			}
			if err := m.svc.MarketOrderMatchingInBatch(stack); err != nil {
				m.Send(stack, err)
				continue
			}
			stack = []*grpc_order.MarketOrderMatching{}
		case input := <-m.sub:
			stack = append(stack, input)
		}
	}
}

func (m *MarketOrderMatchingChannel) Send(events []*grpc_order.MarketOrderMatching, err error) {
	log.Println(err)
}

func (m *MarketOrderMatchingChannel) Receive(event *grpc_order.MarketOrderMatching) {
	m.sub <- event
}