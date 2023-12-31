package orderbook

import (
	"github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/pkg/types"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log"
	"time"
)

type manager struct {
	service Service

	limitAsk  chan *types.ErrReceiveContext[*grpc_order.Order] // limitAsk is channel for limit ask order
	marketAsk chan *types.ErrReceiveContext[*grpc_order.Order] // marketAsk is channel for market ask order
	limitBid  chan *types.ErrReceiveContext[*grpc_order.Order] // limitBid is channel for limit bid order
	marketBid chan *types.ErrReceiveContext[*grpc_order.Order] // marketBid is channel for market bid order

	cancel chan *types.ResultReceiveContext[string, *grpc_order.Order] // cancel is channel for cancel order
}

func newManager(cluster *redis.ClusterClient, db *gorm.DB) *manager {
	return &manager{
		service:   NewService(cluster, db),
		limitAsk:  make(chan *types.ErrReceiveContext[*grpc_order.Order], 1000000),
		marketAsk: make(chan *types.ErrReceiveContext[*grpc_order.Order], 1000000),
		limitBid:  make(chan *types.ErrReceiveContext[*grpc_order.Order], 1000000),
		marketBid: make(chan *types.ErrReceiveContext[*grpc_order.Order], 1000000),
		cancel:    make(chan *types.ResultReceiveContext[string, *grpc_order.Order], 1000000),
	}
}
func (m *manager) Subscribe() {
	for {
		select {
		case ctx := <-m.limitAsk:
			ctx.Tunnel <- m.service.LimitAsk(ctx.Value)
		case ctx := <-m.limitBid:
			ctx.Tunnel <- m.service.LimitBid(ctx.Value)
		case ctx := <-m.marketAsk:
			ctx.Tunnel <- m.service.MarketAsk(ctx.Value)
		case ctx := <-m.marketBid:
			ctx.Tunnel <- m.service.MarketBid(ctx.Value)
		case ctx := <-m.cancel:
			order, err := m.service.CancelOrder(ctx.Value)
			ctx.Result <- order
			ctx.Err <- err
		}
	}
}

func (m *manager) SnapshotCron(duration time.Duration) {
	var err error
	var tick <-chan time.Time

	tick = time.Tick(duration)

	for range tick {
		if err = m.service.Snapshot(); err != nil {
			// FIXME: error handling
			log.Println("snapshot error: ", err)
		}
	}
}

func (m *manager) LoadOrderBook() (err error) {
	return m.service.LoadOrderBook()
}

func (m *manager) LimitAskInsert(ask *grpc_order.Order) (order *grpc_order.Order, err error) {
	ctx := &types.ErrReceiveContext[*grpc_order.Order]{
		Tunnel: make(chan error),
		Value:  ask,
	}
	m.limitAsk <- ctx
	if <-ctx.Tunnel != nil {
		return nil, err
	}
	return ask, nil
}

func (m *manager) LimitBidInsert(bid *grpc_order.Order) (order *grpc_order.Order, err error) {
	ctx := &types.ErrReceiveContext[*grpc_order.Order]{
		Tunnel: make(chan error),
		Value:  bid,
	}
	m.limitBid <- ctx
	if <-ctx.Tunnel != nil {
		return nil, err
	}
	return bid, nil
}

func (m *manager) MarketAskInsert(ask *grpc_order.Order) (order *grpc_order.Order, err error) {
	ctx := &types.ErrReceiveContext[*grpc_order.Order]{
		Tunnel: make(chan error),
		Value:  ask,
	}
	m.marketAsk <- ctx
	if <-ctx.Tunnel != nil {
		return nil, err
	}
	return ask, nil
}

func (m *manager) MarketBidInsert(bid *grpc_order.Order) (order *grpc_order.Order, err error) {
	ctx := &types.ErrReceiveContext[*grpc_order.Order]{
		Tunnel: make(chan error),
		Value:  bid,
	}
	m.marketBid <- ctx
	if <-ctx.Tunnel != nil {
		return nil, err
	}
	return bid, nil
}

func (m *manager) CancelOrder(uuid string) (order *grpc_order.Order, err error) {
	ctx := &types.ResultReceiveContext[string, *grpc_order.Order]{
		Result: make(chan *grpc_order.Order),
		Err:    make(chan error),
		Value:  uuid,
	}
	m.cancel <- ctx
	order = <-ctx.Result
	err = <-ctx.Err

	if err != nil {
		return nil, status.Errorf(codes.Canceled, "failed to cancel order: [%v]", err)
	}

	if order == nil {
		return nil, status.Error(codes.Unknown, ErrOrderCancelFailed.Error())
	}
	return order, nil
}

func (m *manager) BidOrder() (bids []*grpc_order.Order, err error) {
	return m.service.BidOrder()
}

func (m *manager) AskOrder() (asks []*grpc_order.Order, err error) {
	return m.service.AskOrder()
}