package order

import (
	"context"
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/cache"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/finexblock-dev/gofinexblock/finexblock/user"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

type orderService struct {
	orderRepository Repository
	userRepository  user.Repository
}

func (o *orderService) InsertSnapshot(symbolID uint, _snapshot *entity.SnapshotOrderBook) (result *entity.SnapshotOrderBook, err error) {
	if err = o.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = o.orderRepository.InsertSnapshot(tx, symbolID, _snapshot)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}

	return result, nil
}

func (o *orderService) FindSymbolByName(name string) (result *entity.OrderSymbol, err error) {
	if err = o.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = o.orderRepository.FindSymbolByName(tx, name)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}

	return result, nil
}

func (o *orderService) FindSymbolByID(id uint) (result *entity.OrderSymbol, err error) {
	if err = o.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = o.orderRepository.FindSymbolByID(tx, id)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}

	return result, nil
}

func (o *orderService) FindManyOrderByUUID(uuids []string) (result []*entity.OrderBook, err error) {
	if err = o.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = o.orderRepository.FindManyOrderByUUID(tx, uuids)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}

	return result, nil
}

func (o *orderService) FindRecentIntervalByDuration(duration entity.Duration) (result *entity.OrderInterval, err error) {
	if err = o.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = o.orderRepository.FindRecentIntervalByDuration(tx, duration)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}

	return result, nil
}

func (o *orderService) OrderMatchingEventInBatch(event []*grpc_order.OrderMatching) (err error) {
	return o.Conn().Transaction(func(tx *gorm.DB) error {
		// symbol cache for memoization
		var symbolCache = cache.NewDefaultKeyValueStore[entity.OrderSymbol](len(event))
		// symbols
		var symbols []*entity.OrderSymbol
		// symbol
		var _symbol *entity.OrderSymbol
		// symbol names for select query
		var symbolNames []string
		// recent one minute interval
		var recentInterval *entity.OrderInterval
		// order matching events for batch insert operation
		var orderMatchingEvents []*entity.OrderMatchingEvent

		for _, v := range event {
			if _, err = symbolCache.Get(v.Symbol.String()); err == cache.ErrKeyNotFound {
				symbolNames = append(symbolNames, v.Symbol.String())
				// absolute ignore error
				_ = symbolCache.Set(v.Symbol.String(), &entity.OrderSymbol{Name: v.Symbol.String()})
			}
		}

		// find all symbols
		symbols, err = o.orderRepository.FindManySymbolByName(tx, symbolNames)
		if err != nil {
			return err
		}

		for _, symbol := range symbols {
			// absolute ignore error
			_ = symbolCache.Set(symbol.Name, symbol)
		}

		recentInterval, err = o.orderRepository.FindRecentIntervalByDuration(tx, entity.OneMinute)
		if err != nil {
			return err
		}

		for _, v := range event {
			_symbol, err = symbolCache.Get(v.Symbol.String())
			if err == cache.ErrKeyNotFound {
				continue
			}

			orderMatchingEvents = append(orderMatchingEvents, &entity.OrderMatchingEvent{
				OrderSymbolID:   _symbol.ID,
				OrderIntervalID: recentInterval.ID,
				Quantity:        decimal.NewFromFloat(v.Quantity),
				UnitPrice:       decimal.NewFromFloat(v.UnitPrice),
				OrderType:       v.OrderType.String(),
			})
		}

		return o.orderRepository.BatchInsertOrderMatchingEvent(tx, orderMatchingEvents)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (o *orderService) LimitOrderFulfillmentInBatch(event []*grpc_order.OrderFulfillment) (remain []*grpc_order.OrderFulfillment, err error) {
	// case of failed to handle order fulfillment event
	if err = o.orderRepository.Conn().Transaction(func(tx *gorm.DB) error {
		var userCache = cache.NewDefaultKeyValueStore[entity.User](len(event))
		var symbolCache = cache.NewDefaultKeyValueStore[entity.OrderSymbol](len(event))
		var orderCache = cache.NewDefaultKeyValueStore[entity.OrderBook](len(event))

		var orderBookDifferences []*entity.OrderBookDifference

		var orderMatchingHistories []*entity.OrderMatchingHistory

		var users []*entity.User
		var _user *entity.User
		var userUUIDs []string

		var symbols []*entity.OrderSymbol
		var _symbol *entity.OrderSymbol
		var symbolNames []string

		var orders []*entity.OrderBook
		var _order *entity.OrderBook
		var orderUUIDs []string

		// Cache all data for select and batch insert
		for _, v := range event {
			if _, err = userCache.Get(v.UserUUID); err == cache.ErrKeyNotFound {
				userUUIDs = append(userUUIDs, v.UserUUID)
				_ = userCache.Set(v.UserUUID, &entity.User{UUID: v.UserUUID})
			}

			if _, err = symbolCache.Get(v.Symbol.String()); err == cache.ErrKeyNotFound {
				symbolNames = append(symbolNames, v.Symbol.String())
				_ = symbolCache.Set(v.Symbol.String(), &entity.OrderSymbol{Name: v.Symbol.String()})
			}

			if _, err = orderCache.Get(v.OrderUUID); err == cache.ErrKeyNotFound {
				orderUUIDs = append(orderUUIDs, v.OrderUUID)
				_ = orderCache.Set(v.OrderUUID, &entity.OrderBook{OrderUUID: v.OrderUUID})
			}
		}

		// find all users
		users, err = o.userRepository.FindManyUserByUUID(tx, userUUIDs)
		if err != nil {
			return err
		}

		for _, _user := range users {
			_ = userCache.Set(_user.UUID, _user)
		}

		// find all symbols
		symbols, err = o.orderRepository.FindManySymbolByName(tx, symbolNames)
		if err != nil {
			return err
		}

		for _, _symbol := range symbols {
			_ = symbolCache.Set(_symbol.Name, _symbol)
		}

		// find all orders
		orders, err = o.orderRepository.FindManyOrderByUUID(tx, orderUUIDs)
		if err != nil {
			return err
		}

		// update all orders
		if err = o.orderRepository.BatchUpdateOrderBookStatus(tx, orderUUIDs, entity.Fulfilled); err != nil {
			return err
		}

		for _, _order := range orders {
			_ = orderCache.Set(_order.OrderUUID, _order)
		}

		for _, v := range event {
			_symbol, err = symbolCache.Get(v.Symbol.String())
			if err == cache.ErrKeyNotFound {
				remain = append(remain, v)
				continue
			}

			_user, err = userCache.Get(v.UserUUID)
			if err == cache.ErrKeyNotFound {
				remain = append(remain, v)
				continue
			}

			_order, err = orderCache.Get(v.OrderUUID)
			if err == cache.ErrKeyNotFound {
				remain = append(remain, v)
				continue
			}

			orderBookDifferences = append(orderBookDifferences, &entity.OrderBookDifference{
				OrderBookID: _order.ID,
				Reason:      types.Fill,
				Diff:        decimal.NewFromFloat(v.FilledQuantity),
			})

			orderMatchingHistories = append(orderMatchingHistories, &entity.OrderMatchingHistory{
				UserID:         _user.ID,
				OrderUUID:      _order.OrderUUID,
				OrderSymbolID:  _symbol.ID,
				FilledQuantity: decimal.NewFromFloat(v.FilledQuantity),
				UnitPrice:      decimal.NewFromFloat(v.UnitPrice),
				Fee:            decimal.NewFromFloat(v.Fee.Amount),
				OrderType:      _order.OrderType,
			})
		}

		if err = o.orderRepository.BatchInsertOrderBookDifference(tx, orderBookDifferences); err != nil {
			return err
		}

		return o.orderRepository.BatchInsertOrderMatchingHistory(tx, orderMatchingHistories)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return event, err
	}

	return remain, err
}

func (o *orderService) LimitOrderPartialFillInBatch(event []*grpc_order.OrderPartialFill) (remain []*grpc_order.OrderPartialFill, err error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderService) LimitOrderInitializeInBatch(event []*grpc_order.OrderInitialize) (err error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderService) ChartDraw(event []*grpc_order.OrderMatching) (err error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderService) LimitOrderCancellationInBatch(event []*grpc_order.OrderCancelled) (result []*grpc_order.OrderCancelled, err error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderService) HandleOrderInterval(name string, duration time.Duration) (err error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderService) Conn() *gorm.DB {
	return o.orderRepository.Conn()
}

func (o *orderService) Tx(level sql.IsolationLevel) *gorm.DB {
	return o.orderRepository.Tx(level)
}

func (o *orderService) Ctx() context.Context {
	return context.Background()
}

func (o *orderService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}

func newOrderService(db *gorm.DB) *orderService {
	return &orderService{orderRepository: newOrderRepository(db)}
}
