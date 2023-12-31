package order

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/finexblock-dev/gofinexblock/pkg/cache"
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/pkg/order/structs"
	"github.com/finexblock-dev/gofinexblock/pkg/types"
	"github.com/finexblock-dev/gofinexblock/pkg/user"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type orderService struct {
	orderRepository Repository
	userRepository  user.Repository
}

func (o *orderService) SearchOrderMatchingHistory(input *structs.SearchOrderMatchingHistoryInput) (result []*entity.OrderMatchingHistory, err error) {
	if err = o.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = o.orderRepository.SearchOrderMatchingHistory(tx, input)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		var msg = errors.New("failed to get order matching history")
		return nil, errors.Join(msg, err)
	}

	return result, nil
}

func (o *orderService) MarketOrderMatchingInBatch(event []*grpc_order.MarketOrderMatching) (err error) {
	return o.Conn().Transaction(func(tx *gorm.DB) error {
		var userCache = cache.NewDefaultKeyValueStore[entity.User](len(event))
		var symbolCache = cache.NewDefaultKeyValueStore[entity.OrderSymbol](len(event))

		var userUUIDs []string
		var symbolNames []string

		var users []*entity.User
		var _user *entity.User

		var symbols []*entity.OrderSymbol
		var _symbol *entity.OrderSymbol

		var matchingHistories []*entity.OrderMatchingHistory

		for _, v := range event {
			// Memoize user uuid
			if _user, err = userCache.Get(v.UserUUID); err != nil {
				userUUIDs = append(userUUIDs, v.UserUUID)
				//_ = userCache.Set(v.UserUUID, new(entity.User))
			}

			// Memoize symbol name
			if _symbol, err = symbolCache.Get(v.Symbol.String()); err != nil {
				symbolNames = append(symbolNames, v.Symbol.String())
				//_ = symbolCache.Set(v.Symbol.String(), new(entity.OrderSymbol))
			}
		}

		// SELECT * FROM user WHERE uuid IN (...)
		users, err = o.userRepository.FindManyUserByUUID(tx, userUUIDs)
		if err != nil {
			return fmt.Errorf("failed to get user : %v", err)
		}

		// Memoize user
		for _, v := range users {
			_ = userCache.Set(v.UUID, v)
		}

		// SELECT * FROM order_symbol WHERE name IN (...)
		symbols, err = o.orderRepository.FindManySymbolByName(tx, symbolNames)
		if err != nil {
			return fmt.Errorf("failed to get symbol : %v", err)
		}

		// Memoize symbol
		for _, v := range symbols {
			_ = symbolCache.Set(v.Name, v)
		}

		for _, v := range event {
			_symbol, err = symbolCache.Get(v.GetSymbol().String())
			if err == cache.ErrKeyNotFound {
				continue
			}

			_user, err = userCache.Get(v.GetUserUUID())
			if err == cache.ErrKeyNotFound {
				continue
			}

			if _symbol.ID == 0 || _user.ID == 0 {
				continue
			}

			// order_matching_history (user_id, order_symbol_id, order_uuid, filled_quantity, unit_price, fee, order_type)
			matchingHistories = append(matchingHistories, &entity.OrderMatchingHistory{
				UserID:         _user.ID,
				OrderSymbolID:  _symbol.ID,
				OrderUUID:      v.OrderUUID,
				FilledQuantity: decimal.NewFromFloat(v.Quantity),
				UnitPrice:      decimal.NewFromFloat(v.UnitPrice),
				Fee:            decimal.NewFromFloat(v.Fee.Amount),
				OrderType:      types.OrderType(v.OrderType.String()),
			})
		}

		return o.orderRepository.BatchInsertOrderMatchingHistory(tx, matchingHistories)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (o *orderService) FindSnapshotByOrderSymbolID(symbolID uint) (result *entity.SnapshotOrderBook, err error) {
	if err = o.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = o.orderRepository.FindSnapshotByOrderSymbolID(tx, symbolID)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}

	return result, nil
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

		var symbolCache = cache.NewDefaultKeyValueStore[entity.OrderSymbol](len(event))

		for _, v := range event {
			symbolNames = append(symbolNames, v.Symbol.String())
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

		var userCache = cache.NewDefaultKeyValueStore[entity.User](len(event))
		var symbolCache = cache.NewDefaultKeyValueStore[entity.OrderSymbol](len(event))
		var orderCache = cache.NewDefaultKeyValueStore[entity.OrderBook](len(event))

		// Cache all data for select and batch insert
		for _, v := range event {
			userUUIDs = append(userUUIDs, v.UserUUID)
			symbolNames = append(symbolNames, v.Symbol.String())
			orderUUIDs = append(orderUUIDs, v.OrderUUID)
		}

		// find all users
		users, err = o.userRepository.FindManyUserByUUID(tx, userUUIDs)
		if err != nil {
			return err
		}

		for _, _user = range users {
			_ = userCache.Set(_user.UUID, _user)
		}

		// find all symbols
		symbols, err = o.orderRepository.FindManySymbolByName(tx, symbolNames)
		if err != nil {
			return err
		}

		for _, _symbol = range symbols {
			_ = symbolCache.Set(_symbol.Name, _symbol)
		}

		// find all orders
		orders, err = o.orderRepository.FindManyOrderByUUID(tx, orderUUIDs)
		if err != nil {
			return err
		}

		// update all orders
		if err = o.orderRepository.BatchUpdateOrderBookStatus(tx, orderUUIDs, types.Fulfilled); err != nil {
			return err
		}

		for _, _order = range orders {
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
	// case of failed to handle order partial fill event
	if err = o.orderRepository.Conn().Transaction(func(tx *gorm.DB) error {
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

		var userCache = cache.NewDefaultKeyValueStore[entity.User](len(event))
		var symbolCache = cache.NewDefaultKeyValueStore[entity.OrderSymbol](len(event))
		var orderCache = cache.NewDefaultKeyValueStore[entity.OrderBook](len(event))

		// Cache all data for select and batch insert
		for _, v := range event {
			userUUIDs = append(userUUIDs, v.UserUUID)
			symbolNames = append(symbolNames, v.Symbol.String())
			orderUUIDs = append(orderUUIDs, v.OrderUUID)
		}

		// find all users
		users, err = o.userRepository.FindManyUserByUUID(tx, userUUIDs)
		if err != nil {
			return err
		}

		for _, _user = range users {
			_ = userCache.Set(_user.UUID, _user)
		}

		// find all symbols
		symbols, err = o.orderRepository.FindManySymbolByName(tx, symbolNames)
		if err != nil {
			return err
		}

		for _, _symbol = range symbols {
			_ = symbolCache.Set(_symbol.Name, _symbol)
		}

		// find all orders
		orders, err = o.orderRepository.FindManyOrderByUUID(tx, orderUUIDs)
		if err != nil {
			return err
		}

		for _, _order = range orders {
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
			userUUIDs = append(userUUIDs, v.UserUUID)
			symbolNames = append(symbolNames, v.Symbol.String())
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
				OrderType:     types.OrderType(v.OrderType.String()),
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

			if data.LowPrice.GreaterThan(unitPrice) {
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

		for _, v := range event {
			orderUUIDs = append(orderUUIDs, v.OrderUUID)
		}

		var orderCache = cache.NewDefaultKeyValueStore[entity.OrderBook](len(event))

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

		for _, uuid := range orderUUIDs {
			orderCache.ConcurrentDelete(uuid)
		}

		return o.orderRepository.BatchUpdateOrderBookStatus(tx, orderUUIDs, types.Cancelled)

	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return event, err
	}

	return remain, err
}

func (o *orderService) InsertOrderInterval(name types.Duration, duration time.Duration) (err error) {
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

	return &orderService{
		orderRepository: newOrderRepository(db),
		userRepository:  user.NewRepository(db),
	}
}