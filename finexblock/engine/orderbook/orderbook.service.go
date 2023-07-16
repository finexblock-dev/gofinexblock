package orderbook

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/cache"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/trade"
	"github.com/finexblock-dev/gofinexblock/finexblock/utils"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
	"math"
	"time"
)

type service struct {
	repository                     Repository
	orderCache                     *cache.DefaultKeyValueStore[grpc_order.Order]
	tradeService                   trade.Service
	askMarketPrice, bidMarketPrice decimal.Decimal
}

func newService(cluster *redis.ClusterClient) *service {
	return &service{
		repository:     NewRepository(),
		orderCache:     cache.NewDefaultKeyValueStore[grpc_order.Order](1000),
		tradeService:   trade.NewService(cluster),
		askMarketPrice: decimal.NewFromFloat(math.MaxFloat64),
		bidMarketPrice: decimal.Zero,
	}
}

func (s *service) LimitAsk(ask *grpc_order.Order) (err error) {
	defer func() {
		s.askMarketPrice = s.repository.AskMarketPrice()
		s.bidMarketPrice = s.repository.BidMarketPrice()

		// Cache order information
		// FIXME: expiration time is 4 week now.
		if err = s.orderCache.SetEX(ask.OrderUUID, ask, time.Hour*24*7*4); err == cache.ErrCacheFull {
			_ = s.orderCache.Resize(s.orderCache.CurrentSize() * 2)
			_ = s.orderCache.SetEX(ask.OrderUUID, ask, time.Hour*24*7*4)
		}
	}()

	askUnitPrice := decimal.NewFromFloat(ask.UnitPrice)

	// case if ask market price is less than ordered unit price
	if askUnitPrice.GreaterThan(s.bidMarketPrice) {
		s.repository.PushAsk(ask)
		return s.tradeService.SendPlacementStream(ask)
	}

	// Set quantity to decimal for safe math
	quantity := decimal.NewFromFloat(ask.Quantity)

	// Loop while quantity is greater than zero
	// Break if there is no ask order to match or ask order has fulfilled
	for quantity.GreaterThan(decimal.Zero) {
		ask.Quantity = quantity.InexactFloat64()

		// Get bid order to match
		bid := s.repository.PopBid()

		// If there is no bid order, just place order
		if bid == nil {
			// Push ask order
			s.repository.PushAsk(ask)
			// Place order (Send Redis Stream)
			return s.tradeService.SendPlacementStream(ask)
		}

		bidUnitPrice := decimal.NewFromFloat(bid.UnitPrice)
		// When case of ask unit price is greater than bid unit price
		if askUnitPrice.GreaterThan(bidUnitPrice) {
			// Push ask order
			// Update market price or not
			s.repository.PushAsk(ask)

			// Push bid order
			// Update market price or not
			s.repository.PushBid(bid)

			// Place order (Send Redis Stream)
			return s.tradeService.SendPlacementStream(ask)
		}

		// Quantity of opponent bid order
		bidQuantity := decimal.NewFromFloat(bid.Quantity)

		switch {
		// Case of bid order quantity is greater than ask order quantity.
		// Ask order : Fulfilled, Bid order : Partial filled
		case bidQuantity.GreaterThan(quantity):
			// Minus bid order quantity
			bid.Quantity = bidQuantity.Sub(quantity).InexactFloat64()
			s.repository.PushBid(bid)
			// Place order (Send Redis Stream)
			return s.tradeService.SendMatchStream(trade.CaseLimitAskBigger, &grpc_order.BidAsk{Ask: ask, Bid: bid})
		// Case of bid order quantity is equal to ask order quantity.
		// Ask order : Fulfilled, Bid order : Fulfilled
		case bidQuantity.Equal(quantity):
			return s.tradeService.SendMatchStream(trade.CaseLimitAskEqual, &grpc_order.BidAsk{Ask: ask, Bid: bid})
		// Case of bid order quantity is less than ask order quantity.
		// Ask order : Partial filled, Bid order : Fulfilled
		case bidQuantity.LessThan(quantity):
			if err = s.tradeService.SendMatchStream(trade.CaseLimitAskSmaller, &grpc_order.BidAsk{Ask: ask, Bid: bid}); err != nil {
				return err
			}
			// Minus quantity and continue process...
			quantity = quantity.Sub(bidQuantity)
		}
	}

	return nil
}

func (s *service) LimitBid(bid *grpc_order.Order) (err error) {
	defer func() {
		s.askMarketPrice = s.repository.AskMarketPrice()
		s.bidMarketPrice = s.repository.BidMarketPrice()

		// Cache order information
		// FIXME: expiration time is 4 week now.
		if err = s.orderCache.SetEX(bid.OrderUUID, bid, time.Hour*24*7*4); err == cache.ErrCacheFull {
			_ = s.orderCache.Resize(s.orderCache.CurrentSize() * 2)
			_ = s.orderCache.SetEX(bid.OrderUUID, bid, time.Hour*24*7*4)
		}
	}()

	bidUnitPrice := decimal.NewFromFloat(bid.UnitPrice)
	// If unit price is less than ask market price, place order.
	if bidUnitPrice.LessThan(s.askMarketPrice) {
		// Push bid order to heap
		// Update market price or not
		s.repository.PushBid(bid)

		// Send placement event
		return s.tradeService.SendPlacementStream(bid)
	}

	// Set quantity to decimal for safe math
	quantity := decimal.NewFromFloat(bid.Quantity)

	// Loop while quantity is greater than zero
	// Break if there is no ask order to match or bid order has fulfilled
	for quantity.GreaterThan(decimal.Zero) {
		bid.Quantity = quantity.InexactFloat64()

		// Get ask order to match
		ask := s.repository.PopAsk()

		// If there is no ask order, just place order
		if ask == nil {
			// Place order
			s.repository.PushBid(bid)
			return s.tradeService.SendPlacementStream(bid)
		}

		askUnitPrice := decimal.NewFromFloat(ask.UnitPrice)
		// When case of ask unit price is greater than bid unit price
		if askUnitPrice.GreaterThan(bidUnitPrice) {
			// Push ask order
			// Update market price or not
			s.repository.PushAsk(ask)

			// Push bid order
			// Update market price or not
			s.repository.PushBid(bid)

			// Send placement event
			return s.tradeService.SendPlacementStream(bid)
		}

		// Quantity of opponent ask order
		opQuantity := decimal.NewFromFloat(ask.Quantity)
		switch {
		// Case of ask order quantity is greater than bid order quantity.
		// Bid order : Fulfilled, Ask order : Partial filled
		case opQuantity.GreaterThan(quantity):
			// Minus ask order quantity
			ask.Quantity = opQuantity.Sub(quantity).InexactFloat64()

			// Push ask order
			// Update market price or not
			s.repository.PushAsk(ask)

			return s.tradeService.SendMatchStream(trade.CaseLimitBidBigger, &grpc_order.BidAsk{Bid: bid, Ask: ask})
		// Case of ask order quantity is equal to bid order quantity.
		// Bid order : Fulfilled, Ask order : Fulfilled
		case opQuantity.Equal(quantity):
			return s.tradeService.SendMatchStream(trade.CaseLimitBidEqual, &grpc_order.BidAsk{Bid: bid, Ask: ask})
		// Case of ask order quantity is less than bid order quantity.
		// Bid order : Partial filled, Ask order : Fulfilled
		case opQuantity.LessThan(quantity):
			// Minus quantity and continue process...
			quantity = quantity.Sub(opQuantity)

			if err = s.tradeService.SendMatchStream(trade.CaseLimitBidSmaller, &grpc_order.BidAsk{Bid: bid, Ask: ask}); err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *service) MarketAsk(ask *grpc_order.Order) (err error) {
	defer func() {
		s.askMarketPrice = s.repository.AskMarketPrice()
		s.bidMarketPrice = s.repository.BidMarketPrice()
	}()
	// Set quantity to decimal for safe math
	quantity := decimal.NewFromFloat(ask.Quantity)

	// Loop while quantity is greater than zero
	// Break if there is no ask order to match
	for quantity.GreaterThan(decimal.Zero) {
		ask.Quantity = quantity.InexactFloat64()

		// Get ask order to match
		bid := s.repository.PopBid()

		// If there is no bid order, refund market ask order
		if bid == nil {
			return s.tradeService.SendRefundStream(ask)
		}

		mul := utils.CoinDecimal(utils.OpponentCurrency(ask.Symbol))
		bidQuantity := decimal.NewFromFloat(bid.Quantity)
		actualAskQuantity := quantity.Div(mul)

		switch {
		// Case of bid order quantity is greater than ask order quantity.
		// Ask order : Fulfilled, Bid order : Partial filled
		case bidQuantity.GreaterThan(actualAskQuantity):
			// Minus bid quantity,
			// bid.Quantity represents coin quantity (decimal)
			// actualAskQuantity = filled quantity
			bid.Quantity = bidQuantity.Sub(actualAskQuantity).InexactFloat64()

			// Push bid order
			s.repository.PushBid(bid)

			// End loop
			return s.tradeService.SendMatchStream(trade.CaseMarketAskBigger, &grpc_order.BidAsk{Bid: bid, Ask: ask})
		// Case of bid order quantity is equal to ask order quantity.
		// Ask order : Fulfilled, Bid order : Fulfilled
		case bidQuantity.Equal(actualAskQuantity):
			// End loop
			return s.tradeService.SendMatchStream(trade.CaseMarketAskEqual, &grpc_order.BidAsk{Bid: bid, Ask: ask})
		// Case of bid order quantity is less than ask order quantity.
		// Ask order : Partial filled, Bid order : Fulfilled
		case bidQuantity.LessThan(actualAskQuantity):

			if err = s.tradeService.SendMatchStream(trade.CaseMarketAskSmaller, &grpc_order.BidAsk{Bid: bid, Ask: ask}); err != nil {
				return err
			}
			// Minus quantity and continue process...
			quantity = quantity.Sub(bidQuantity.Mul(mul))
		}
	}

	return nil
}

func (s *service) MarketBid(bid *grpc_order.Order) (err error) {
	defer func() {
		s.askMarketPrice = s.repository.AskMarketPrice()
		s.bidMarketPrice = s.repository.BidMarketPrice()
	}()
	// Set quantity to decimal for safe math
	quantity := decimal.NewFromFloat(bid.Quantity)

	// Loop while quantity is greater than zero
	// Break if there is no ask order to match
	for quantity.GreaterThan(decimal.Zero) {
		bid.Quantity = quantity.InexactFloat64()

		// Get ask order to match
		ask := s.repository.PopAsk()

		// If there is no ask order, refund order
		if ask == nil {
			// Refund order
			return s.tradeService.SendRefundStream(bid)
		}

		// convert ask order quantity to satoshi(1 => 1e8)
		// ask.Quantity : amount of order
		// bid.Quantity : satoshi
		askUnitPrice := decimal.NewFromFloat(ask.UnitPrice)
		askQuantity := decimal.NewFromFloat(ask.Quantity)
		btcQuantity := askQuantity.Mul(askUnitPrice)

		switch {
		// Case of ask order quantity is greater than bid order quantity.
		// Bid order : Fulfilled, Ask order : Partial filled
		case btcQuantity.GreaterThan(quantity):
			// Minus ask quantity,
			// bid.Quantity represents coin quantity (decimal)
			// ask.UnitPrice represents price meets the deal (decimal)
			// quantity.Div(askUnitPrice).InexactFloat64() = filled quantity
			ask.Quantity = askQuantity.Sub(quantity.Div(askUnitPrice)).InexactFloat64()

			// Push ask order
			s.repository.PushAsk(ask)

			// End loop
			return s.tradeService.SendMatchStream(trade.CaseMarketBidBigger, &grpc_order.BidAsk{Bid: bid, Ask: ask})
		// Case of ask order quantity is equal to bid order quantity.
		// Bid order : Fulfilled, Ask order : Fulfilled
		case btcQuantity.Equal(quantity):
			// End loop
			return s.tradeService.SendMatchStream(trade.CaseMarketBidEqual, &grpc_order.BidAsk{Bid: bid, Ask: ask})
		// Case of ask order quantity is less than bid order quantity.
		// Bid order : Partial filled, Ask order : Fulfilled
		case btcQuantity.LessThan(quantity):
			if err = s.tradeService.SendMatchStream(trade.CaseMarketBidSmaller, &grpc_order.BidAsk{Bid: bid, Ask: ask}); err != nil {
				return err
			}
			// Minus quantity and continue process...
			quantity = quantity.Sub(askQuantity.Mul(askUnitPrice))
		}
	}

	return nil
}

func (s *service) CancelOrder(uuid string) (order *grpc_order.Order, err error) {
	defer func() {
		s.askMarketPrice = s.repository.AskMarketPrice()
		s.bidMarketPrice = s.repository.BidMarketPrice()
	}()
	order, err = s.orderCache.Get(uuid)
	if err == cache.ErrKeyNotFound {
		return nil, ErrOrderNotFound
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

func (s *service) BidOrder() (bids []*grpc_order.Order, err error) {
	if len(s.repository.BidOrder()) == 0 {
		return nil, ErrOrderBookEmpty
	}
	return s.repository.BidOrder(), nil
}

func (s *service) AskOrder() (asks []*grpc_order.Order, err error) {
	if len(s.repository.AskOrder()) == 0 {
		return nil, ErrOrderBookEmpty
	}
	return s.repository.AskOrder(), nil
}
