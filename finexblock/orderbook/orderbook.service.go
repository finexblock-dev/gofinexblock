package orderbook

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/cache"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/shopspring/decimal"
)

type service struct {
	repository                     Repository
	orderCache                     cache.DefaultKeyValueStore[grpc_order.Order]
	askMarketPrice, bidMarketPrice decimal.Decimal
}

func (s *service) LimitAsk(ask *grpc_order.Order) error {
	unitPrice := decimal.NewFromFloat(ask.UnitPrice)

	// case if ask market price is less than ordered unit price
	if s.bidMarketPrice.GreaterThan(unitPrice) {
		s.repository.PushAsk(ask)
		return nil
	}
	s.askMarketPrice = unitPrice

}

func (s *service) LimitBid(bid *grpc_order.Order) error {
	unitPrice := decimal.NewFromFloat(bid.UnitPrice)

	// case if bid market price is less than ordered unit price
	if s.bidMarketPrice.LessThan(unitPrice) {
		s.repository.PushAsk(bid)
		return nil
	}
	s.bidMarketPrice = unitPrice
}

func (s *service) MarketAsk(ask *grpc_order.Order) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) MarketBid(bid *grpc_order.Order) error {
	//TODO implement me
	panic("implement me")
}

func (s *service) CancelOrder(uuid string) (order *grpc_order.Order, err error) {
	order, err = s.orderCache.Get(uuid)
	if err == cache.ErrKeyNotFound {
		return nil, ErrOrderCancelFailed
	}

	switch order.OrderType {
	case grpc_order.OrderType_BID:
		order = s.repository.RemoveBid(uuid)
		if order == nil {
			return nil, ErrOrderCancelFailed
		}
		return order, nil
	case grpc_order.OrderType_ASK:
		order = s.repository.RemoveAsk(uuid)
		if order == nil {
			return nil, ErrOrderCancelFailed
		}
		return order, nil
	default:
		return nil, ErrOrderTypeNotFound
	}
}

func (s *service) BidOrder() (bids []*grpc_order.Order) {
	return s.repository.BidOrder()
}

func (s *service) AskOrder() (asks []*grpc_order.Order) {
	return s.repository.AskOrder()
}
