package orderbook

import "github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"

type askOrderBook struct {
	IndexMap map[string]int
	Requests []*grpc_order.Order
}

func newAskOrderBook() *askOrderBook {
	return &askOrderBook{IndexMap: make(map[string]int, 1000000), Requests: []*grpc_order.Order{}}
}

func (a *askOrderBook) Len() int {
	return len(a.Requests)
}

func (a *askOrderBook) Less(i, j int) bool {
	if a.Requests[i].UnitPrice == a.Requests[j].UnitPrice {
		return a.Requests[i].Quantity > a.Requests[j].Quantity
	}
	return a.Requests[i].UnitPrice < a.Requests[j].UnitPrice
}

func (a *askOrderBook) Swap(i, j int) {
	if len(a.Requests) > i && len(a.Requests) > j {
		a.Requests[i], a.Requests[j] = a.Requests[j], a.Requests[i]
		a.IndexMap[a.Requests[i].OrderUUID] = i
		a.IndexMap[a.Requests[j].OrderUUID] = j
	}
}

func (a *askOrderBook) Pop() interface{} {
	if len(a.Requests) > 0 {
		old := a.Requests
		n := len(old)
		item := old[n-1]
		a.Requests = old[0 : n-1]
		delete(a.IndexMap, item.OrderUUID)
		return item
	}
	return nil
}

func (a *askOrderBook) Push(x interface{}) {
	item := x.(*grpc_order.Order)
	a.IndexMap[item.OrderUUID] = len(a.Requests)
	a.Requests = append(a.Requests, item)
}