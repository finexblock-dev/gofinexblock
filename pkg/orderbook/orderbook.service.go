package orderbook

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/finexblock-dev/gofinexblock/pkg/cache"
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/pkg/goaws"
	"github.com/finexblock-dev/gofinexblock/pkg/instance"
	"github.com/finexblock-dev/gofinexblock/pkg/order"
	"github.com/finexblock-dev/gofinexblock/pkg/trade"
	"github.com/finexblock-dev/gofinexblock/pkg/utils"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"math"
)

type service struct {
	orderBookRepository            Repository
	instanceRepository             instance.Repository
	orderRepository                order.Repository
	orderCache                     *cache.DefaultKeyValueStore[grpc_order.Order]
	tradeService                   trade.Manager
	askMarketPrice, bidMarketPrice decimal.Decimal
}

func newService(cluster *redis.ClusterClient, db *gorm.DB) *service {
	return &service{
		orderBookRepository: NewRepository(),
		instanceRepository:  instance.NewRepository(db),
		orderRepository:     order.NewRepository(db),
		orderCache:          cache.NewDefaultKeyValueStore[grpc_order.Order](1000),
		tradeService:        trade.New(cluster),
		askMarketPrice:      decimal.NewFromFloat(math.MaxFloat64),
		bidMarketPrice:      decimal.Zero,
	}
}

func (s *service) Snapshot() (err error) {
	var privateIP string
	var ipModel *entity.FinexblockServerIP
	var serverModel *entity.FinexblockServer
	var symbol *entity.OrderSymbol
	var askOrderList, bidOrderList []*grpc_order.Order
	var ask, bid []byte

	privateIP, err = goaws.OwnPrivateIP()
	if err != nil {
		return status.Errorf(codes.Internal, "failed to get private ip: [%v]", err)
	}

	return s.instanceRepository.Conn().Transaction(func(tx *gorm.DB) error {
		ipModel, err = s.instanceRepository.FindServerByIP(tx, privateIP)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to find server by ip: [%v]", err)
		}

		// Find server by id
		serverModel, err = s.instanceRepository.FindServerByID(tx, ipModel.ServerID)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to find server by id: [%v]", err)
		}

		if serverModel.Name[:3] != "BTC" {
			return status.Errorf(codes.Internal, "server name is not valid: %v", serverModel.Name)
		}

		// Find order symbol
		symbol, err = s.orderRepository.FindSymbolByName(tx, serverModel.Name)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to find order symbol: [%v]", err)
		}

		// Find snapshot by server order symbol id
		askOrderList = s.orderBookRepository.AskOrder()
		bidOrderList = s.orderBookRepository.BidOrder()

		if ask, err = json.Marshal(askOrderList); err != nil {
			return status.Errorf(codes.Internal, "failed to marshal ask order list: [%v]", err)
		}

		if bid, err = json.Marshal(bidOrderList); err != nil {
			return status.Errorf(codes.Internal, "failed to marshal bid order list: [%v]", err)
		}

		if _, err = s.orderRepository.InsertSnapshot(tx, &entity.SnapshotOrderBook{
			OrderSymbolID: symbol.ID,
			AskOrderList:  string(ask),
			BidOrderList:  string(bid),
		}); err != nil {
			return status.Errorf(codes.Internal, "failed to insert snapshot: [%v]", err)
		}

		return nil
	})
}

func (s *service) LoadOrderBook() (err error) {
	var privateIP string
	var ipModel *entity.FinexblockServerIP
	var serverModel *entity.FinexblockServer
	var symbol *entity.OrderSymbol
	var snapshot *entity.SnapshotOrderBook
	var askOrderList []*grpc_order.Order
	var bidOrderList []*grpc_order.Order

	privateIP, err = goaws.OwnPrivateIP()
	if err != nil {
		return status.Errorf(codes.Internal, "failed to get private ip: [%v]", err)
	}

	if privateIP == "" {
		return status.Errorf(codes.Internal, "private ip is empty")
	}

	return s.instanceRepository.Conn().Transaction(func(tx *gorm.DB) error {
		ipModel, err = s.instanceRepository.FindServerByIP(tx, privateIP)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to find server by ip: [%v]", err)
		}

		// Find server by id
		serverModel, err = s.instanceRepository.FindServerByID(tx, ipModel.ServerID)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to find server by id: [%v]", err)
		}

		if serverModel.Name[:3] != "BTC" {
			return status.Errorf(codes.Internal, "server name is not valid: %v", serverModel.Name)
		}

		// Find order symbol
		symbol, err = s.orderRepository.FindSymbolByName(tx, serverModel.Name)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to find order symbol: [%v]", err)
		}

		// Find snapshot by server order symbol id
		snapshot, err = s.orderRepository.FindSnapshotByOrderSymbolID(tx, symbol.ID)
		if err != nil && err != gorm.ErrRecordNotFound {
			return status.Errorf(codes.Internal, "failed to find snapshot: [%v]", err)
		}

		if err == gorm.ErrRecordNotFound {
			return nil
		}

		if err = json.Unmarshal([]byte(snapshot.AskOrderList), &askOrderList); err != nil {
			return status.Errorf(codes.Internal, "failed to unmarshal ask order list: [%v]", err)
		}

		if err = json.Unmarshal([]byte(snapshot.BidOrderList), &bidOrderList); err != nil {
			return status.Errorf(codes.Internal, "failed to unmarshal bid order list: [%v]", err)
		}

		//for _, ask := range askOrderList {
		//	if err = s.orderCache.Set(ask.OrderUUID, ask); err == cache.ErrCacheFull {
		//		_ = s.orderCache.Resize(s.orderCache.CurrentSize() * 2)
		//		_ = s.orderCache.Set(ask.OrderUUID, ask)
		//	}
		//}
		//
		//for _, bid := range bidOrderList {
		//	if err = s.orderCache.Set(bid.OrderUUID, bid); err == cache.ErrCacheFull {
		//		_ = s.orderCache.Resize(s.orderCache.CurrentSize() * 2)
		//		_ = s.orderCache.Set(bid.OrderUUID, bid)
		//	}
		//}

		// Load order book
		if err = s.orderBookRepository.LoadOrderBook(bidOrderList, askOrderList); err != nil {
			return status.Errorf(codes.Internal, "failed to load order book: [%v]", err)
		}

		s.askMarketPrice = s.orderBookRepository.AskMarketPrice()
		s.bidMarketPrice = s.orderBookRepository.BidMarketPrice()

		return nil
	})
}

func (s *service) LimitAsk(ask *grpc_order.Order) (err error) {
	var group *errgroup.Group

	defer func() {
		s.askMarketPrice = s.orderBookRepository.AskMarketPrice()
		s.bidMarketPrice = s.orderBookRepository.BidMarketPrice()

	}()

	// Cache order information
	// FIXME: expiration time is 4 week now.
	if err = s.tradeService.SetOrder(ask.OrderUUID, ask.OrderType.String()); err != nil {
		return status.Errorf(codes.Internal, "failed to set order: [%v]", err)
	}

	//if err = s.orderCache.SetEX(ask.OrderUUID, ask, time.Hour*24*7*4); err == cache.ErrCacheFull {
	//	_ = s.orderCache.Resize(s.orderCache.CurrentSize() * 2)
	//	_ = s.orderCache.SetEX(ask.OrderUUID, ask, time.Hour*24*7*4)
	//}

	askUnitPrice := decimal.NewFromFloat(ask.UnitPrice)

	// Set quantity to decimal for safe math
	quantity := decimal.NewFromFloat(ask.Quantity)

	// case if ask market price is less than ordered unit price
	if askUnitPrice.GreaterThan(s.bidMarketPrice) {
		s.orderBookRepository.PushAsk(ask)
		placement := utils.NewOrderPlacement(ask.UserUUID, ask.OrderUUID, quantity, askUnitPrice, ask.OrderType, ask.Symbol)
		if err = s.tradeService.SendPlacementStream(placement); err != nil {
			return status.Errorf(codes.Internal, "failed to send placement stream: [%v]", err)
		}
		return nil
	}

	group, _ = errgroup.WithContext(context.Background())

	// Loop while quantity is greater than zero
	// Break if there is no ask order to match or ask order has fulfilled
	for quantity.GreaterThan(decimal.Zero) {
		ask.Quantity = quantity.InexactFloat64()

		// Get bid order to match
		bid := s.orderBookRepository.PopBid()

		// If there is no bid order, just place order
		if bid == nil {
			group.Go(func() error {
				// Place order (Send Redis Stream)
				placement := utils.NewOrderPlacement(ask.UserUUID, ask.OrderUUID, quantity, askUnitPrice, ask.OrderType, ask.Symbol)
				if err = s.tradeService.SendPlacementStream(placement); err != nil {
					return status.Errorf(codes.Internal, "failed to send placement stream: [%v]", err)
				}

				return nil
			})

			// Push ask order
			s.orderBookRepository.PushAsk(ask)

			return group.Wait()
		}

		bidUnitPrice := decimal.NewFromFloat(bid.UnitPrice)
		// When case of ask unit price is greater than bid unit price
		if askUnitPrice.GreaterThan(bidUnitPrice) {
			group.Go(func() error {
				// Place order (Send Redis Stream)
				placement := utils.NewOrderPlacement(ask.UserUUID, ask.OrderUUID, quantity, askUnitPrice, ask.OrderType, ask.Symbol)
				if err = s.tradeService.SendPlacementStream(placement); err != nil {
					return status.Errorf(codes.Internal, "failed to send placement stream: [%v]", err)
				}
				return nil
			})

			// Push ask order
			// Update market price or not
			s.orderBookRepository.PushAsk(ask)

			// Push bid order
			// Update market price or not
			s.orderBookRepository.PushBid(bid)

			return group.Wait()
		}

		// Quantity of opponent bid order
		bidQuantity := decimal.NewFromFloat(bid.Quantity)

		switch {
		// Case of bid order quantity is greater than ask order quantity.
		// Ask order : Fulfilled, Bid order : Partial filled
		case bidQuantity.GreaterThan(quantity):
			// Place order (Send Redis Stream)
			group.Go(func() error {
				if err = s.tradeService.SendMatchStream(trade.CaseLimitAskBigger, &grpc_order.BidAsk{Ask: ask, Bid: bid}); err != nil {
					return status.Errorf(codes.Internal, "failed to send match stream: [%v]", err)
				}
				return nil
			})

			// Minus bid order quantity
			bid.Quantity = bidQuantity.Sub(quantity).InexactFloat64()
			s.orderBookRepository.PushBid(bid)

			//s.orderCache.Del(ask.OrderUUID)
			return group.Wait()
		// Case of bid order quantity is equal to ask order quantity.
		// Ask order : Fulfilled, Bid order : Fulfilled
		case bidQuantity.Equal(quantity):
			group.Go(func() error {
				if err = s.tradeService.SendMatchStream(trade.CaseLimitAskEqual, &grpc_order.BidAsk{Ask: ask, Bid: bid}); err != nil {
					return status.Errorf(codes.Internal, "failed to send match stream: [%v]", err)
				}
				return nil
			})
			//s.orderCache.Del(ask.OrderUUID)
			//s.orderCache.Del(bid.OrderUUID)
			return group.Wait()
		// Case of bid order quantity is less than ask order quantity.
		// Ask order : Partial filled, Bid order : Fulfilled
		case bidQuantity.LessThan(quantity):
			group.Go(func() error {
				if err = s.tradeService.SendMatchStream(trade.CaseLimitAskSmaller, &grpc_order.BidAsk{Ask: ask, Bid: bid}); err != nil {
					return status.Errorf(codes.Internal, "failed to send match stream: [%v]", err)
				}
				return nil
			})

			// Minus quantity and continue process...
			quantity = quantity.Sub(bidQuantity)

			//s.orderCache.Del(bid.OrderUUID)
		}
	}

	return group.Wait()
}

func (s *service) LimitBid(bid *grpc_order.Order) (err error) {
	defer func() {
		s.askMarketPrice = s.orderBookRepository.AskMarketPrice()
		s.bidMarketPrice = s.orderBookRepository.BidMarketPrice()
	}()

	// Cache order information
	// FIXME: expiration time is 4 week now.
	if err = s.tradeService.SetOrder(bid.OrderUUID, bid.OrderType.String()); err != nil {
		return status.Errorf(codes.Internal, "failed to set order: [%v]", err)
	}

	//if err = s.orderCache.SetEX(bid.OrderUUID, bid, time.Hour*24*7*4); err == cache.ErrCacheFull {
	//	_ = s.orderCache.Resize(s.orderCache.CurrentSize() * 2)
	//	_ = s.orderCache.SetEX(bid.OrderUUID, bid, time.Hour*24*7*4)
	//}

	bidUnitPrice := decimal.NewFromFloat(bid.UnitPrice)

	// Set quantity to decimal for safe math
	quantity := decimal.NewFromFloat(bid.Quantity)

	// If unit price is less than ask market price, place order.
	if bidUnitPrice.LessThan(s.askMarketPrice) {
		// Push bid order to heap
		// Update market price or not
		s.orderBookRepository.PushBid(bid)

		// Send placement event
		placement := utils.NewOrderPlacement(bid.UserUUID, bid.OrderUUID, quantity, bidUnitPrice, bid.OrderType, bid.Symbol)
		if err = s.tradeService.SendPlacementStream(placement); err != nil {
			return status.Errorf(codes.Internal, "failed to send placement stream: [%v]", err)
		}

		return nil
	}

	var group *errgroup.Group

	group, _ = errgroup.WithContext(context.Background())

	// Loop while quantity is greater than zero
	// Break if there is no ask order to match or bid order has fulfilled
	for quantity.GreaterThan(decimal.Zero) {
		bid.Quantity = quantity.InexactFloat64()

		// Get ask order to match
		ask := s.orderBookRepository.PopAsk()

		// If there is no ask order, just place order
		if ask == nil {
			// Place order
			s.orderBookRepository.PushBid(bid)

			group.Go(func() error {
				placement := utils.NewOrderPlacement(bid.UserUUID, bid.OrderUUID, quantity, bidUnitPrice, bid.OrderType, bid.Symbol)
				if err = s.tradeService.SendPlacementStream(placement); err != nil {
					return status.Errorf(codes.Internal, "failed to send placement stream: [%v]", err)
				}
				return nil
			})

			return group.Wait()
		}

		askUnitPrice := decimal.NewFromFloat(ask.UnitPrice)

		// When case of ask unit price is greater than bid unit price
		if askUnitPrice.GreaterThan(bidUnitPrice) {
			// Push ask order
			// Update market price or not
			s.orderBookRepository.PushAsk(ask)

			// Push bid order
			// Update market price or not
			s.orderBookRepository.PushBid(bid)

			// Send placement event
			group.Go(func() error {
				placement := utils.NewOrderPlacement(bid.UserUUID, bid.OrderUUID, quantity, bidUnitPrice, bid.OrderType, bid.Symbol)

				if err = s.tradeService.SendPlacementStream(placement); err != nil {
					return status.Errorf(codes.Internal, "failed to send placement stream: [%v]", err)
				}

				return nil
			})

			return group.Wait()
		}

		// Quantity of opponent ask order
		opQuantity := decimal.NewFromFloat(ask.Quantity)
		switch {
		// Case of ask order quantity is greater than bid order quantity.
		// Bid order : Fulfilled, Ask order : Partial filled
		case opQuantity.GreaterThan(quantity):
			group.Go(func() error {
				if err = s.tradeService.SendMatchStream(trade.CaseLimitBidBigger, &grpc_order.BidAsk{Bid: bid, Ask: ask}); err != nil {
					return status.Errorf(codes.Internal, "failed to send match stream: [%v]", err)
				}

				return nil
			})

			// Minus ask order quantity
			ask.Quantity = opQuantity.Sub(quantity).InexactFloat64()

			// Push ask order
			// Update market price or not
			s.orderBookRepository.PushAsk(ask)

			//s.orderCache.Del(bid.OrderUUID)

			return group.Wait()
		// Case of ask order quantity is equal to bid order quantity.
		// Bid order : Fulfilled, Ask order : Fulfilled
		case opQuantity.Equal(quantity):
			group.Go(func() error {
				if err = s.tradeService.SendMatchStream(trade.CaseLimitBidEqual, &grpc_order.BidAsk{Bid: bid, Ask: ask}); err != nil {
					return status.Errorf(codes.Internal, "failed to send match stream: [%v]", err)
				}

				return nil
			})

			//s.orderCache.Del(ask.OrderUUID)
			//s.orderCache.Del(bid.OrderUUID)
			return group.Wait()
		// Case of ask order quantity is less than bid order quantity.
		// Bid order : Partial filled, Ask order : Fulfilled
		case opQuantity.LessThan(quantity):
			group.Go(func() error {
				if err = s.tradeService.SendMatchStream(trade.CaseLimitBidSmaller, &grpc_order.BidAsk{Bid: bid, Ask: ask}); err != nil {
					return status.Errorf(codes.Internal, "failed to send match stream: [%v]", err)
				}

				return nil
			})

			// Minus quantity and continue process...
			quantity = quantity.Sub(opQuantity)

			//s.orderCache.Del(ask.OrderUUID)
		}
	}

	return group.Wait()
}

func (s *service) MarketAsk(ask *grpc_order.Order) (err error) {
	var group *errgroup.Group

	defer func() {
		s.askMarketPrice = s.orderBookRepository.AskMarketPrice()
		s.bidMarketPrice = s.orderBookRepository.BidMarketPrice()
	}()
	// Set quantity to decimal for safe math
	quantity := decimal.NewFromFloat(ask.Quantity)

	group, _ = errgroup.WithContext(context.Background())

	// Loop while quantity is greater than zero
	// Break if there is no ask order to match
	for quantity.GreaterThan(decimal.Zero) {
		ask.Quantity = quantity.InexactFloat64()

		// Get ask order to match
		bid := s.orderBookRepository.PopBid()

		// If there is no bid order, refund market ask order
		if bid == nil {
			group.Go(func() error {
				event := utils.NewBalanceUpdate(ask.UserUUID, quantity, utils.OpponentCurrency(ask.Symbol), grpc_order.Reason_REFUND)
				if err = s.tradeService.SendBalanceUpdateStream(event); err != nil {
					return status.Errorf(codes.Internal, "failed to send refund stream: [%v]", err)
				}

				return nil
			})

			return group.Wait()
		}

		mul := utils.CoinDecimal(utils.OpponentCurrency(ask.Symbol))
		bidQuantity := decimal.NewFromFloat(bid.Quantity)
		actualAskQuantity := quantity.Div(mul)

		switch {
		// Case of bid order quantity is greater than ask order quantity.
		// Ask order : Fulfilled, Bid order : Partial filled
		case bidQuantity.GreaterThan(actualAskQuantity):
			group.Go(func() error {
				// End loop
				if err = s.tradeService.SendMatchStream(trade.CaseMarketAskBigger, &grpc_order.BidAsk{Bid: bid, Ask: ask}); err != nil {
					return status.Errorf(codes.Internal, "failed to send match stream: [%v]", err)
				}
				return nil
			})

			// Minus bid quantity,
			// bid.Quantity represents coin quantity (decimal)
			// actualAskQuantity = filled quantity
			bid.Quantity = bidQuantity.Sub(actualAskQuantity).InexactFloat64()

			// Push bid order
			s.orderBookRepository.PushBid(bid)

			return group.Wait()
		// Case of bid order quantity is equal to ask order quantity.
		// Ask order : Fulfilled, Bid order : Fulfilled
		case bidQuantity.Equal(actualAskQuantity):

			group.Go(func() error {
				// End loop
				if err = s.tradeService.SendMatchStream(trade.CaseMarketAskEqual, &grpc_order.BidAsk{Bid: bid, Ask: ask}); err != nil {
					return status.Errorf(codes.Internal, "failed to send match stream: [%v]", err)
				}

				return nil
			})

			//s.orderCache.Del(bid.OrderUUID)

			return group.Wait()
		// Case of bid order quantity is less than ask order quantity.
		// Ask order : Partial filled, Bid order : Fulfilled
		case bidQuantity.LessThan(actualAskQuantity):
			group.Go(func() error {
				if err = s.tradeService.SendMatchStream(trade.CaseMarketAskSmaller, &grpc_order.BidAsk{Bid: bid, Ask: ask}); err != nil {
					return status.Errorf(codes.Internal, "failed to send match stream: [%v]", err)
				}
				return nil
			})

			// Minus quantity and continue process...
			quantity = quantity.Sub(bidQuantity.Mul(mul))

			//s.orderCache.Del(bid.OrderUUID)
		}
	}

	return group.Wait()
}

func (s *service) MarketBid(bid *grpc_order.Order) (err error) {
	var group *errgroup.Group

	defer func() {
		s.askMarketPrice = s.orderBookRepository.AskMarketPrice()
		s.bidMarketPrice = s.orderBookRepository.BidMarketPrice()
	}()
	// Set quantity to decimal for safe math
	quantity := decimal.NewFromFloat(bid.Quantity)

	group, _ = errgroup.WithContext(context.Background())

	// Loop while quantity is greater than zero
	// Break if there is no ask order to match
	for quantity.GreaterThan(decimal.Zero) {
		bid.Quantity = quantity.InexactFloat64()

		// Get ask order to match
		ask := s.orderBookRepository.PopAsk()

		// If there is no ask order, refund order
		if ask == nil {
			group.Go(func() error {
				// Refund order
				event := utils.NewBalanceUpdate(bid.UserUUID, quantity, grpc_order.Currency_BTC, grpc_order.Reason_REFUND)
				if err = s.tradeService.SendBalanceUpdateStream(event); err != nil {
					return status.Errorf(codes.Internal, "failed to send refund stream: [%v]", err)
				}

				return nil
			})

			return group.Wait()
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
			group.Go(func() error {
				// End loop
				if err = s.tradeService.SendMatchStream(trade.CaseMarketBidBigger, &grpc_order.BidAsk{Bid: bid, Ask: ask}); err != nil {
					return status.Errorf(codes.Internal, "failed to send match stream: [%v]", err)
				}
				return nil
			})

			// Minus ask quantity,
			// bid.Quantity represents coin quantity (decimal)
			// ask.UnitPrice represents price meets the deal (decimal)
			// quantity.Div(askUnitPrice).InexactFloat64() = filled quantity
			ask.Quantity = askQuantity.Sub(quantity.Div(askUnitPrice)).InexactFloat64()

			// Push ask order
			s.orderBookRepository.PushAsk(ask)

			return group.Wait()
		// Case of ask order quantity is equal to bid order quantity.
		// Bid order : Fulfilled, Ask order : Fulfilled
		case btcQuantity.Equal(quantity):
			group.Go(func() error {
				// End loop
				if err = s.tradeService.SendMatchStream(trade.CaseMarketBidEqual, &grpc_order.BidAsk{Bid: bid, Ask: ask}); err != nil {
					return status.Errorf(codes.Internal, "failed to send match stream: [%v]", err)
				}
				return nil
			})

			//s.orderCache.Del(ask.OrderUUID)
			return group.Wait()
		// Case of ask order quantity is less than bid order quantity.
		// Bid order : Partial filled, Ask order : Fulfilled
		case btcQuantity.LessThan(quantity):
			group.Go(func() error {
				if err = s.tradeService.SendMatchStream(trade.CaseMarketBidSmaller, &grpc_order.BidAsk{Bid: bid, Ask: ask}); err != nil {
					return status.Errorf(codes.Internal, "failed to send match stream: [%v]", err)
				}

				return nil
			})
			// Minus quantity and continue process...
			quantity = quantity.Sub(askQuantity.Mul(askUnitPrice))
			//s.orderCache.Del(ask.OrderUUID)
		}
	}

	return group.Wait()
}

func (s *service) CancelOrder(uuid string) (order *grpc_order.Order, err error) {
	defer func() {
		s.askMarketPrice = s.orderBookRepository.AskMarketPrice()
		s.bidMarketPrice = s.orderBookRepository.BidMarketPrice()
		if err == nil {
			//s.orderCache.Del(uuid)
			_ = s.tradeService.DeleteOrder(uuid)
		}
	}()

	var pushFunc func(*grpc_order.Order)
	var value string

	value, err = s.tradeService.GetOrder(uuid)
	if err != nil {
		return nil, err
	}

	//order, err = s.orderCache.Get(uuid)
	//if err == cache.ErrKeyNotFound {
	//	return nil, errors.Join(ErrOrderNotFound, cache.ErrKeyNotFound)
	//}

	switch value {
	case grpc_order.OrderType_BID.String():
		order = s.orderBookRepository.RemoveBid(uuid)
		if order == nil {
			return nil, errors.Join(ErrOrderCancelFailed, ErrOrderNotFound)
		}
		pushFunc = s.orderBookRepository.PushBid

	case grpc_order.OrderType_ASK.String():
		order = s.orderBookRepository.RemoveAsk(uuid)
		if order == nil {
			return nil, errors.Join(ErrOrderCancelFailed, ErrOrderNotFound)
		}
		pushFunc = s.orderBookRepository.PushAsk
	default:
		return nil, ErrOrderTypeNotFound
	}

	// Send Cancellation event
	quantity := decimal.NewFromFloat(order.Quantity)
	unitPrice := decimal.NewFromFloat(order.UnitPrice)
	cancelled := utils.NewOrderCancelled(order.UserUUID, order.OrderUUID, quantity, unitPrice, order.OrderType, order.Symbol)
	if err = s.tradeService.SendCancellationStream(cancelled); err != nil {
		pushFunc(order)
		return nil, errors.Join(ErrOrderCancelFailed, err)
	}
	return order, nil
}

func (s *service) BidOrder() (bids []*grpc_order.Order, err error) {
	if len(s.orderBookRepository.BidOrder()) == 0 {
		return nil, ErrOrderBookEmpty
	}
	return s.orderBookRepository.BidOrder(), nil
}

func (s *service) AskOrder() (asks []*grpc_order.Order, err error) {
	if len(s.orderBookRepository.AskOrder()) == 0 {
		return nil, ErrOrderBookEmpty
	}
	return s.orderBookRepository.AskOrder(), nil
}