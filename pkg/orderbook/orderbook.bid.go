package orderbook

import (
	"github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
)

type bidOrderBook struct {
	IndexMap map[string]int
	Requests []*grpc_order.Order
}

func newBidOrderBook() *bidOrderBook {
	return &bidOrderBook{IndexMap: make(map[string]int, 1000000), Requests: []*grpc_order.Order{}}
}

func (b *bidOrderBook) Len() int {
	return len(b.Requests)
}

func (b *bidOrderBook) Less(i, j int) bool {
	if b.Requests[i].UnitPrice == b.Requests[j].UnitPrice {
		return b.Requests[i].Quantity > b.Requests[j].Quantity
	}
	return b.Requests[i].UnitPrice > b.Requests[j].UnitPrice
}

func (b *bidOrderBook) Swap(i, j int) {
	if len(b.Requests) > i && len(b.Requests) > j {
		b.Requests[i], b.Requests[j] = b.Requests[j], b.Requests[i]
		b.IndexMap[b.Requests[i].OrderUUID] = i
		b.IndexMap[b.Requests[j].OrderUUID] = j
	}
}

func (b *bidOrderBook) Pop() interface{} {
	if len(b.Requests) > 0 {
		old := b.Requests
		n := len(old)
		item := old[n-1]
		b.Requests = old[0 : n-1]
		delete(b.IndexMap, item.OrderUUID)
		return item
	}
	return nil
}

func (b *bidOrderBook) Push(x interface{}) {
	item := x.(*grpc_order.Order)
	b.IndexMap[item.OrderUUID] = len(b.Requests)
	b.Requests = append(b.Requests, item)
}