package orderbook

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
)

type queue struct {
	service Service

	limitAsk  chan *types.ErrReceiveContext[*grpc_order.Order] // limitAsk is channel for limit ask order
	marketAsk chan *types.ErrReceiveContext[*grpc_order.Order] // marketAsk is channel for market ask order
	limitBid  chan *types.ErrReceiveContext[*grpc_order.Order] // limitBid is channel for limit bid order
	marketBid chan *types.ErrReceiveContext[*grpc_order.Order] // marketBid is channel for market bid order

	askRemove chan *types.ResultReceiveContext[string, *grpc_order.Order] // askRemove is channel for cancel ask order
	bidRemove chan *types.ResultReceiveContext[string, *grpc_order.Order] // bidRemove is channel for cancel bid order

	// bidMarketPrice
	// askMarketPrice
}

func (q *queue) Subscribe() {

	for {
		select {
		case ctx := <-q.limitAsk:
			ctx.Tunnel <- q.service.LimitAsk(ctx.Value)
		case ctx := <-q.limitBid:
			ctx.Tunnel <- q.service.LimitBid(ctx.Value)
		case ctx := <-q.marketAsk:
			ctx.Tunnel <- q.service.MarketAsk(ctx.Value)
		case ctx := <-q.marketBid:
			ctx.Tunnel <- q.service.MarketBid(ctx.Value)
		case ctx := <-q.askRemove:
			order, _ := q.service.CancelOrder(ctx.Value)
			ctx.Tunnel <- order
		case ctx := <-q.bidRemove:
			order, _ := q.service.CancelOrder(ctx.Value)
			ctx.Tunnel <- order
		}
	}
}

func (q *queue) LimitAskInsert(ask *grpc_order.Order) (order *grpc_order.Order, err error) {
	ctx := &types.ErrReceiveContext[*grpc_order.Order]{
		Tunnel: make(chan error),
		Value:  ask,
	}
	q.limitAsk <- ctx
	if <-ctx.Tunnel != nil {
		return nil, err
	}
	return ask, nil
}

func (q *queue) LimitBidInsert(bid *grpc_order.Order) (order *grpc_order.Order, err error) {
	ctx := &types.ErrReceiveContext[*grpc_order.Order]{
		Tunnel: make(chan error),
		Value:  bid,
	}
	q.limitBid <- ctx
	if <-ctx.Tunnel != nil {
		return nil, err
	}
	return bid, nil
}

func (q *queue) MarketAskInsert(ask *grpc_order.Order) (order *grpc_order.Order, err error) {
	ctx := &types.ErrReceiveContext[*grpc_order.Order]{
		Tunnel: make(chan error),
		Value:  ask,
	}
	q.marketAsk <- ctx
	if <-ctx.Tunnel != nil {
		return nil, err
	}
	return ask, nil
}

func (q *queue) MarketBidInsert(bid *grpc_order.Order) (order *grpc_order.Order, err error) {
	ctx := &types.ErrReceiveContext[*grpc_order.Order]{
		Tunnel: make(chan error),
		Value:  bid,
	}
	q.marketBid <- ctx
	if <-ctx.Tunnel != nil {
		return nil, err
	}
	return bid, nil
}

func (q *queue) AskRemove(uuid string) (order *grpc_order.Order, err error) {
	ctx := &types.ResultReceiveContext[string, *grpc_order.Order]{
		Tunnel: make(chan *grpc_order.Order),
		Value:  uuid,
	}
	q.askRemove <- ctx
	order = <-ctx.Tunnel
	if order == nil {
		return nil, ErrOrderCancelFailed
	}
	return order, nil
}

func (q *queue) BidRemove(uuid string) (order *grpc_order.Order, err error) {
	ctx := &types.ResultReceiveContext[string, *grpc_order.Order]{
		Tunnel: make(chan *grpc_order.Order),
		Value:  uuid,
	}
	q.bidRemove <- ctx
	order = <-ctx.Tunnel
	if order == nil {
		return nil, ErrOrderCancelFailed
	}
	return order, nil
}

func (q *queue) BidOrder() (bids []*grpc_order.Order, err error) {
	return q.service.BidOrder()
}

func (q *queue) AskOrder() (asks []*grpc_order.Order, err error) {
	return q.service.AskOrder()
}
