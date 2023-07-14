package order

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/finexblock-dev/gofinexblock/finexblock/cache"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/finexblock-dev/gofinexblock/finexblock/user"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"math"
	"strconv"
	"time"
)

type orderService struct {
	orderRepository Repository
	userRepository  user.Repository
	userCache       *cache.DefaultKeyValueStore[entity.User]
}

func (o *orderService) InsertSnapshot(symbolID uint, bid, ask []*grpc_order.Order) (result *entity.SnapshotOrderBook, err error) {
	if err = o.Conn().Transaction(func(tx *gorm.DB) error {
		var bidMarshal, askMarshal []byte

		bidMarshal, err = json.Marshal(bid)
		if err != nil {
			return err
		}

		askMarshal, err = json.Marshal(ask)
		if err != nil {
			return err
		}

		result, err = o.orderRepository.InsertSnapshot(tx, &entity.SnapshotOrderBook{
			OrderSymbolID: symbolID,
			BidOrderList:  string(bidMarshal),
			AskOrderList:  string(askMarshal),
		})
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

func (o *orderService) FindRecentIntervalByDuration(duration types.Duration) (result *entity.OrderInterval, err error) {
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

		recentInterval, err = o.orderRepository.FindRecentIntervalByDuration(tx, types.OneMinute)
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
		if err = o.orderRepository.BatchUpdateOrderBookStatus(tx, orderUUIDs, types.PartialFilled); err != nil {
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
	}, &sql.TxOptions{Isolation: sql.LevelRepeatableRead}); err != nil {
		return event, err
	}

	return remain, err
}

func (o *orderService) LimitOrderPartialFillInBatch(event []*grpc_order.OrderPartialFill) (remain []*grpc_order.OrderPartialFill, err error) {
	// case of failed to handle order partial fill event
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
		// checkpoint: do not update order status
		//if err = o.orderRepository.BatchUpdateOrderBookStatus(tx, orderUUIDs, types.Fulfilled); err != nil {
		//	return err
		//}

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
	}, &sql.TxOptions{Isolation: sql.LevelRepeatableRead}); err != nil {
		return event, err
	}

	return remain, err
}

func (o *orderService) LimitOrderInitializeInBatch(event []*grpc_order.OrderInitialize) (err error) {
	return o.Conn().Transaction(func(tx *gorm.DB) error {
		var userUUIDs []string
		var symbolNames []string

		var users []*entity.User
		var _user *entity.User
		var symbols []*entity.OrderSymbol
		var _symbol *entity.OrderSymbol

		var orders []*entity.OrderBook

		var userCache = cache.NewDefaultKeyValueStore[entity.User](len(event))
		var symbolCache = cache.NewDefaultKeyValueStore[entity.OrderSymbol](len(event))

		for _, v := range event {
			if _, err = userCache.Get(v.UserUUID); err == cache.ErrKeyNotFound {
				userUUIDs = append(userUUIDs, v.UserUUID)
				_ = userCache.Set(v.UserUUID, &entity.User{UUID: v.UserUUID})
			}

			if _, err = symbolCache.Get(v.Symbol.String()); err == cache.ErrKeyNotFound {
				symbolNames = append(symbolNames, v.Symbol.String())
				_ = symbolCache.Set(v.Symbol.String(), &entity.OrderSymbol{Name: v.Symbol.String()})
			}
		}

		// find all users
		users, err = o.userRepository.FindManyUserByUUID(tx, userUUIDs)
		if err != nil {
			return err
		}

		for _, _user = range users {
			_ = userCache.Set(_user.UUID, _user)
		}

		symbols, err = o.orderRepository.FindManySymbolByName(tx, symbolNames)
		if err != nil {
			return err
		}

		for _, _symbol = range symbols {
			_ = symbolCache.Set(_symbol.Name, _symbol)
		}

		for _, v := range event {
			_symbol, err = symbolCache.Get(v.Symbol.String())
			if err == cache.ErrKeyNotFound {
				continue
			}

			_user, err = userCache.Get(v.UserUUID)
			if err == cache.ErrKeyNotFound {
				continue
			}

			orders = append(orders, &entity.OrderBook{
				UserID:        _user.ID,
				OrderUUID:     v.OrderUUID,
				OrderSymbolID: _symbol.ID,
				UnitPrice:     decimal.NewFromFloat(v.UnitPrice),
				Quantity:      decimal.NewFromFloat(v.Quantity),
				OrderType:     types.OrderType(v.OrderType),
				Status:        types.Placed,
			})
		}

		return o.orderRepository.BatchInsertOrderBook(tx, orders)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (o *orderService) ChartDraw(event []*grpc_order.OrderMatching) (err error) {
	return o.Conn().Transaction(func(tx *gorm.DB) error {
		var data *types.PoleData

		var symbols []*entity.OrderSymbol
		var _symbol *entity.OrderSymbol
		var symbolNames []string
		var symbolIDs []uint

		var recentIntervals []*entity.OrderInterval
		var _interval *entity.OrderInterval
		var recentIntervalIDs []uint

		var recentPoles []*entity.Chart
		var poleIDs []uint
		var pole *entity.Chart

		// cache symbol
		var symbolCache = cache.NewDefaultKeyValueStore[entity.OrderSymbol](len(event))

		// cache price for each symbol
		var poleDataCache = cache.NewDefaultKeyValueStore[types.PoleData](len(event))

		for _, v := range event {

			name := v.Symbol.String()
			symbolNames = append(symbolNames, name)

			unitPrice := decimal.NewFromFloat(v.UnitPrice)
			quantity := decimal.NewFromFloat(v.Quantity)
			if _, err = poleDataCache.Get(name); err == cache.ErrKeyNotFound {
				_ = poleDataCache.Set(name, &types.PoleData{
					LowPrice:     unitPrice,
					HighPrice:    unitPrice,
					ClosePrice:   unitPrice,
					Volume:       quantity,
					TradingValue: unitPrice.Mul(quantity),
				})
				continue
			}

			data, _ = poleDataCache.Get(name)

			if data.HighPrice.LessThan(unitPrice) {
				data.HighPrice = unitPrice
			}

			if data, _ = poleDataCache.Get(name); data.LowPrice.GreaterThan(unitPrice) {
				data.LowPrice = unitPrice
			}

			data.ClosePrice = unitPrice
			data.TradingValue.Add(unitPrice.Mul(quantity))
			data.Volume.Add(quantity)

			_ = poleDataCache.Set(name, data)
		}

		// find all symbols
		symbols, err = o.orderRepository.FindManySymbolByName(tx, symbolNames)
		if err != nil {
			return err
		}

		for _, v := range symbols {
			_ = symbolCache.Set(strconv.Itoa(int(v.ID)), v)
			symbolIDs = append(symbolIDs, v.ID)
		}

		recentIntervals, err = o.orderRepository.FindRecentIntervalGroupByDuration(tx)
		if err != nil {
			return err
		}

		for _, _interval = range recentIntervals {
			recentIntervalIDs = append(recentIntervalIDs, _interval.ID)
		}

		recentPoles, err = o.orderRepository.FindChartByCond(tx, recentIntervalIDs, symbolIDs)
		if err != nil {
			return err
		}

		for _, pole = range recentPoles {
			poleIDs = append(poleIDs, pole.ID)
			_symbol, err = symbolCache.Get(strconv.Itoa(int(pole.OrderSymbolID)))
			if err == cache.ErrKeyNotFound {
				continue
			}

			data, err = poleDataCache.Get(_symbol.Name)
			if err == cache.ErrKeyNotFound {
				continue
			}

			if pole.LowPrice.GreaterThan(data.LowPrice) {
				pole.LowPrice = data.LowPrice
			}

			if pole.HighPrice.LessThan(data.HighPrice) {
				pole.HighPrice = data.HighPrice
			}

			pole.Volume = data.Volume.Add(data.Volume)
			pole.TradingValue = data.TradingValue.Add(data.TradingValue)
			pole.ClosePrice = data.ClosePrice

		}

		updateQuery := "UPDATE " + pole.TableName() + " SET high_price = CASE id "
		for _, p := range recentPoles {
			updateQuery += fmt.Sprintf("WHEN %v THEN %v ", p.ID, p.HighPrice)
		}
		updateQuery += "END, low_price = CASE id "
		for _, p := range recentPoles {
			updateQuery += fmt.Sprintf("WHEN %v THEN %v ", p.ID, p.LowPrice)
		}
		updateQuery += "END, close_price = CASE id "
		for _, p := range recentPoles {
			updateQuery += fmt.Sprintf("WHEN %v THEN %v ", p.ID, p.ClosePrice)
		}
		updateQuery += "END, volume = CASE id "
		for _, p := range recentPoles {
			updateQuery += fmt.Sprintf("WHEN %v THEN %v ", p.ID, p.Volume)
		}
		updateQuery += "END, trading_value = CASE id "
		for _, p := range recentPoles {
			updateQuery += fmt.Sprintf("WHEN %v THEN %v ", p.ID, p.TradingValue)
		}

		updateQuery += fmt.Sprintf("END WHERE id IN (?)")
		return tx.Exec(updateQuery, poleIDs).Error
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (o *orderService) LimitOrderCancellationInBatch(event []*grpc_order.OrderCancelled) (remain []*grpc_order.OrderCancelled, err error) {
	if err = o.Conn().Transaction(func(tx *gorm.DB) error {
		var orderUUIDs []string
		var orders []*entity.OrderBook
		var orderCache = cache.NewDefaultKeyValueStore[entity.OrderBook](len(event))

		for _, v := range event {
			orderUUIDs = append(orderUUIDs, v.OrderUUID)
		}

		orders, err = o.orderRepository.FindManyOrderByUUID(tx, orderUUIDs)
		if err != nil {
			return err
		}

		for _, v := range orders {
			_ = orderCache.Set(v.OrderUUID, v)
		}

		for _, cancellation := range event {
			if _, err = orderCache.Get(cancellation.OrderUUID); err == cache.ErrKeyNotFound {
				remain = append(remain, cancellation)
			}
		}

		return o.orderRepository.BatchUpdateOrderBookStatus(tx, orderUUIDs, types.Cancelled)

	}, &sql.TxOptions{Isolation: sql.LevelRepeatableRead}); err != nil {
		return event, err
	}

	return remain, err
}

func (o *orderService) HandleOrderInterval(name types.Duration, duration time.Duration) (err error) {
	return o.Conn().Transaction(func(tx *gorm.DB) error {
		var recentInterval *entity.OrderInterval
		var recentPoles []*entity.Chart
		var newInterval *entity.OrderInterval
		var newPoles []*entity.Chart

		recentInterval, err = o.orderRepository.FindRecentIntervalByDuration(tx, name)
		if err != nil {
			return err
		}

		recentPoles, err = o.orderRepository.FindChartByInterval(tx, recentInterval.ID)
		if err != nil {
			return err
		}

		newInterval, err = o.orderRepository.InsertOrderInterval(tx, &entity.OrderInterval{
			Duration:  name,
			StartTime: recentInterval.StartTime.Add(duration),
			EndTime:   recentInterval.EndTime.Add(duration),
		})

		if err != nil {
			return err
		}

		for _, pole := range recentPoles {
			newPoles = append(newPoles, &entity.Chart{
				OrderIntervalID: newInterval.ID,
				OrderSymbolID:   pole.OrderSymbolID,
				OpenPrice:       pole.ClosePrice,
				LowPrice:        pole.ClosePrice,
				HighPrice:       pole.ClosePrice,
				ClosePrice:      pole.ClosePrice,
				Volume:          decimal.Zero,
			})
		}

		return o.orderRepository.BatchInsertChart(tx, newPoles)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
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
	return &orderService{orderRepository: newOrderRepository(db), userRepository: user.NewRepository(db), userCache: cache.NewDefaultKeyValueStore[entity.User](math.MaxInt)}
}