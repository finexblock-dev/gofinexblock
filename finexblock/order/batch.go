package order

import (
	"fmt"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/order"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/user"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/order/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"time"
)

func (o *orderService) HandleOrderMatchingEventInBatch(tx *gorm.DB, event []*grpc_order.OrderMatching) error {
	var symbolNames []string
	var symbols []*order.OrderSymbol
	var input []*order.OrderMatchingEvent
	var _symbol *order.OrderSymbol
	var _matchingEvent *order.OrderMatchingEvent
	var interval *order.OrderInterval

	symbolMap := make(map[string]*order.OrderSymbol, len(event))

	for _, v := range event {
		if _, ok := symbolMap[v.Symbol.String()]; !ok {
			symbolNames = append(symbolNames, v.Symbol.String())
			symbolMap[v.Symbol.String()] = new(order.OrderSymbol)
		}
	}

	if err := tx.Table(_symbol.TableName()).Where("name IN ?", symbolNames).Find(&symbols).Error; err != nil {
		return fmt.Errorf("failed to get symbol : %v", err)
	}

	if err := tx.Table(interval.TableName()).Order("start_time DESC").First(&interval).Error; err != nil {
		return fmt.Errorf("failed to get interval : %v", err)
	}

	for _, v := range symbols {
		symbolMap[v.Name] = v
	}

	for _, v := range event {
		_symbol = symbolMap[v.Symbol.String()]

		if _symbol.ID == 0 {
			continue
		}

		input = append(input, &order.OrderMatchingEvent{
			ID:              0,
			OrderSymbolID:   _symbol.ID,
			OrderIntervalID: interval.ID,
			Quantity:        decimal.NewFromFloat(v.Quantity),
			UnitPrice:       decimal.NewFromFloat(v.UnitPrice),
			OrderType:       v.OrderType.String(),
		})
	}

	if err := tx.Table(_matchingEvent.TableName()).CreateInBatches(input, len(input)).Error; err != nil {
		return fmt.Errorf("failed to create in batches into order_matching_history table : %v", err)
	}

	return nil
}

func (o *orderService) HandleOrderFulfillmentInBatch(tx *gorm.DB, event []*grpc_order.OrderFulfillment) ([]*grpc_order.OrderFulfillment, error) {
	var userUUIDs []string
	var symbolNames []string
	var orderUUIDs []string

	var users []*user.User
	var _user *user.User

	var symbols []*order.OrderSymbol
	var _symbol *order.OrderSymbol

	var orders []*order.OrderBook
	var _order *order.OrderBook

	var diffList []*order.OrderBookDifference
	var _diff *order.OrderBookDifference

	var matchingHistories []*order.OrderMatchingHistory
	var _matchingHistory *order.OrderMatchingHistory

	// Not initialized, or something other reason
	var remain []*grpc_order.OrderFulfillment

	// Mapping for cache user
	userMap := make(map[string]*user.User, len(event))
	// Mapping for cache symbol
	symbolMap := make(map[string]*order.OrderSymbol, len(event))
	// Mapping for cache order
	orderMap := make(map[string]*order.OrderBook, len(event))

	for _, v := range event {
		// Memoize user uuid
		if _, exists := userMap[v.UserUUID]; !exists {
			userUUIDs = append(userUUIDs, v.UserUUID)
			userMap[v.UserUUID] = &user.User{}
		}

		// Memoize symbol name
		if _, exists := symbolMap[v.Symbol.String()]; !exists {
			symbolNames = append(symbolNames, v.Symbol.String())
			symbolMap[v.Symbol.String()] = new(order.OrderSymbol)
		}

		// Memoize order uuid
		if _, exists := orderMap[v.OrderUUID]; !exists {
			orderUUIDs = append(orderUUIDs, v.OrderUUID)
			orderMap[v.OrderUUID] = new(order.OrderBook)
		}
	}

	// SELECT * FROM user WHERE uuid IN (...)
	if err := tx.Table(_user.TableName()).Where("uuid IN ?", userUUIDs).Find(&users).Error; err != nil {
		return event, fmt.Errorf("failed to get user : %v", err)
	}

	// Memoize user
	for _, v := range users {
		userMap[v.UUID] = v
	}

	// SELECT * FROM order_symbol WHERE name IN (...)
	if err := tx.Table(_symbol.TableName()).Where("name IN ?", symbolNames).Find(&symbols).Error; err != nil {
		return event, fmt.Errorf("failed to get symbol : %v", err)
	}

	// Memoize symbol
	for _, v := range symbols {
		symbolMap[v.Name] = v
	}

	// SELECT * FROM order_book WHERE order_uuid IN (...)
	if err := tx.Table(_order.TableName()).Where("order_uuid IN ?", orderUUIDs).Find(&orders).Error; err != nil {
		return event, fmt.Errorf("failed to get orders : %v", err)
	}

	// Memoize order
	for _, v := range orders {
		orderMap[v.OrderUUID] = v
	}

	// UPDATE order_book SET status = 'FULFILLED' WHERE order_uuid IN (...)
	if err := tx.Table(_order.TableName()).Where("order_uuid IN ?", orderUUIDs).UpdateColumn("status", Fulfilled).Error; err != nil {
		return event, fmt.Errorf("failed to update orders : %v", err)
	}

	for _, v := range event {
		symbolID := symbolMap[v.Symbol.String()].ID
		userID := userMap[v.UserUUID].ID
		orderData := orderMap[v.OrderUUID]

		// If event.order_uuid not exists in order_book, then append to remain
		if orderData.ID == 0 {
			remain = append(remain, v)
			continue
		}

		// order_book_difference (order_book_id, diff, reason)
		diffList = append(diffList, &order.OrderBookDifference{
			OrderBookID: orderData.ID,
			Diff:        decimal.NewFromFloat(v.FilledQuantity),
			Reason:      "FILL",
		})

		// order_matching_history (user_id, order_symbol_id, order_uuid, filled_quantity, unit_price, fee, order_type)
		matchingHistories = append(matchingHistories, &order.OrderMatchingHistory{
			UserID:         uint(userID),
			OrderSymbolID:  symbolID,
			OrderUUID:      v.OrderUUID,
			FilledQuantity: decimal.NewFromFloat(v.FilledQuantity),
			UnitPrice:      decimal.NewFromFloat(v.UnitPrice),
			Fee:            decimal.NewFromFloat(v.Fee.Amount),
			OrderType:      v.OrderType.String(),
		})
	}

	// INSERT INTO order_book_difference (order_book_id, diff, reason) VALUES (...)
	if err := tx.Table(_diff.TableName()).CreateInBatches(diffList, len(diffList)).Error; err != nil {
		return event, fmt.Errorf("failed to create in batches into order_book_difference table : %v", err)
	}

	// INSERT INTO order_matching_history (user_id, order_symbol_id, order_uuid, filled_quantity, unit_price, fee, order_type) VALUES (...)
	if err := tx.Table(_matchingHistory.TableName()).CreateInBatches(matchingHistories, len(matchingHistories)).Error; err != nil {
		return event, fmt.Errorf("failed to create in batches into order_matching_history table : %v", err)
	}

	return remain, nil
}

func (o *orderService) HandleOrderPartialFillInBatch(tx *gorm.DB, event []*grpc_order.OrderPartialFill) ([]*grpc_order.OrderPartialFill, error) {
	var userUUIDs []string
	var symbolNames []string
	var orderUUIDs []string

	var users []*user.User
	var _user *user.User

	var symbols []*order.OrderSymbol
	var _symbol *order.OrderSymbol

	var orders []*order.OrderBook
	var _book *order.OrderBook

	var diffList []*order.OrderBookDifference
	var _diff *order.OrderBookDifference

	var matchingHistories []*order.OrderMatchingHistory
	var _matchingHistory *order.OrderMatchingHistory

	// Not initialized, or something other reason
	var remain []*grpc_order.OrderPartialFill

	// Mapping for user, symbol, order
	userMap := make(map[string]*user.User, len(event))
	symbolMap := make(map[string]*order.OrderSymbol, len(event))
	orderMap := make(map[string]*order.OrderBook, len(event))

	// Memoize user uuid, symbol name, order uuid
	for _, v := range event {
		if _, exists := userMap[v.UserUUID]; !exists {
			userUUIDs = append(userUUIDs, v.UserUUID)
			userMap[v.UserUUID] = new(user.User)
		}
		if _, exists := symbolMap[v.Symbol.String()]; !exists {
			symbolNames = append(symbolNames, v.Symbol.String())
			symbolMap[v.Symbol.String()] = new(order.OrderSymbol)
		}
		if _, exists := orderMap[v.OrderUUID]; !exists {
			orderUUIDs = append(orderUUIDs, v.OrderUUID)
			orderMap[v.OrderUUID] = new(order.OrderBook)
		}
	}

	// SELECT * FROM user WHERE uuid IN (...)
	if err := tx.Table(_user.TableName()).Where("uuid IN ?", userUUIDs).Find(&users).Error; err != nil {
		return event, fmt.Errorf("failed to get user : %v", err)
	}

	// Memoize user
	for _, v := range users {
		userMap[v.UUID] = v
	}

	// SELECT * FROM order_symbol WHERE name IN (...)
	if err := tx.Table(_symbol.TableName()).Where("name IN ?", symbolNames).Find(&symbols).Error; err != nil {
		return event, fmt.Errorf("failed to get symbol : %v", err)
		return event, fmt.Errorf("failed to get symbol : %v", err)
	}

	// Memoize symbol
	for _, v := range symbols {
		symbolMap[v.Name] = v
	}

	// SELECT * FROM order_book WHERE order_uuid IN (...)
	if err := tx.Table(_book.TableName()).Where("order_uuid IN ?", orderUUIDs).Find(&orders).Error; err != nil {
		return event, fmt.Errorf("failed to get orders : %v", err)
	}

	// Memoize order
	for _, v := range orders {
		orderMap[v.OrderUUID] = v
	}

	// Create diff, matching history
	for _, v := range event {
		_s := symbolMap[v.Symbol.String()]
		u := userMap[v.UserUUID]
		orderData := orderMap[v.OrderUUID]

		// If event.order_uuid not exists in order_book, then append to remain
		if orderData.ID == 0 {
			remain = append(remain, v)
			continue
		}

		// order_book_difference (order_book_id, diff, reason)
		diffList = append(diffList, &order.OrderBookDifference{
			OrderBookID: orderData.ID,
			Diff:        decimal.NewFromFloat(v.FilledQuantity),
			Reason:      "FILL",
		})

		// order_matching_history (user_id, order_symbol_id, order_uuid, filled_quantity, unit_price, fee, order_type)
		matchingHistories = append(matchingHistories, &order.OrderMatchingHistory{
			UserID:         uint(u.ID),
			OrderSymbolID:  _s.ID,
			OrderUUID:      v.OrderUUID,
			FilledQuantity: decimal.NewFromFloat(v.FilledQuantity),
			UnitPrice:      decimal.NewFromFloat(v.UnitPrice),
			Fee:            decimal.NewFromFloat(v.Fee.Amount),
			OrderType:      v.OrderType.String(),
		})
	}

	// INSERT INTO order_book_difference (order_book_id, diff, reason) VALUES (...) ON CONFLICT DO NOTHING
	if err := tx.Table(_diff.TableName()).CreateInBatches(diffList, len(diffList)).Error; err != nil {
		return event, fmt.Errorf("failed to create in batches into order_book_difference table : %v", err)
	}

	// INSERT INTO order_matching_history (user_id, order_symbol_id, order_uuid, filled_quantity, unit_price, fee, order_type) VALUES (...) ON CONFLICT DO NOTHING
	if err := tx.Table(_matchingHistory.TableName()).CreateInBatches(matchingHistories, len(matchingHistories)).Error; err != nil {
		return event, fmt.Errorf("failed to create in batches into order_matching_history table : %v", err)
	}

	return remain, nil
}

func (o *orderService) HandleOrderInitializeInBatch(tx *gorm.DB, event []*grpc_order.OrderInitialize) error {
	var userUUIDs []string
	var symbolNames []string
	var orderBooks []*order.OrderBook
	var _orderBook *order.OrderBook

	var users []*user.User
	var _user *user.User

	var symbols []*order.OrderSymbol
	var _symbol *order.OrderSymbol

	// Create user, symbol map
	symbolMap := make(map[string]*order.OrderSymbol, len(event))
	userMap := make(map[string]*user.User, len(event))

	for _, v := range event {
		if _, exists := userMap[v.UserUUID]; !exists {
			userUUIDs = append(userUUIDs, v.UserUUID)
			userMap[v.UserUUID] = new(user.User)
		}
		if _, exists := symbolMap[v.Symbol.String()]; !exists {
			symbolNames = append(symbolNames, v.Symbol.String())
			symbolMap[v.Symbol.String()] = new(order.OrderSymbol)
		}
	}

	// SELECT * FROM user WHERE uuid IN (...)
	if err := tx.Table(_user.TableName()).Where("uuid IN ?", userUUIDs).Find(&users).Error; err != nil {
		return fmt.Errorf("failed to get user : %v", err)
	}

	// Memoize user
	for _, v := range users {
		userMap[v.UUID] = v
	}

	// SELECT * FROM order_symbol WHERE name IN (...)
	if err := tx.Table(_symbol.TableName()).Where("name IN ?", symbolNames).Find(&symbols).Error; err != nil {
		return fmt.Errorf("failed to get symbol : %v", err)
	}

	// Memoize symbol
	for _, v := range symbols {
		symbolMap[v.Name] = v
	}

	for _, v := range event {
		_s := symbolMap[v.Symbol.String()]
		u := userMap[v.UserUUID]

		// If symbol or user not exists, then skip
		if _s.ID == 0 || u.ID == 0 {
			continue
		}

		// order_book (order_symbol_id, user_id, order_uuid, unit_price, quantity, order_type, status)
		orderBooks = append(orderBooks, &order.OrderBook{
			OrderSymbolID: _s.ID,
			UserID:        uint(u.ID),
			OrderUUID:     v.OrderUUID,
			UnitPrice:     decimal.NewFromFloat(v.UnitPrice),
			Quantity:      decimal.NewFromFloat(v.Quantity),
			OrderType:     v.OrderType.String(),
			Status:        "PLACED",
		})
	}

	// INSERT INTO order_book (order_symbol_id, user_id, order_uuid, unit_price, quantity, order_type, status) VALUES (...) ON CONFLICT DO NOTHING
	if err := tx.Table(_orderBook.TableName()).CreateInBatches(orderBooks, len(orderBooks)).Error; err != nil {
		return fmt.Errorf("failed to create in batches into order_book table : %v", err)
	}

	return nil
}

func (o *orderService) HandleOrderInterval(tx *gorm.DB, name string, duration time.Duration) error {
	var _interval *order.OrderInterval
	var newInterval *order.OrderInterval

	var recentPoles []*order.Chart
	var _chart *order.Chart

	var recentInterval *order.OrderInterval
	var newPoles []*order.Chart

	if err := tx.Table(_interval.TableName()).Where("duration = ?", name).Order("start_time desc").First(&recentInterval).Error; err != nil {
		return fmt.Errorf("failed to get most recent interval which is corresponding to %v : %v", name, err)
	}

	if err := tx.Table(_chart.TableName()).Where("order_interval_id = ?", recentInterval.ID).Find(&recentPoles).Error; err != nil {
		return fmt.Errorf("failed to get recent poles corresponding to %v : %v", name, err)
	}

	newInterval = &order.OrderInterval{
		Duration:  name,
		StartTime: recentInterval.StartTime.Add(duration),
		EndTime:   recentInterval.EndTime.Add(duration),
	}

	if err := tx.Table(_interval.TableName()).Create(newInterval).Error; err != nil {
		return fmt.Errorf("failed to create new interval : %v", err)
	}

	for _, v := range recentPoles {
		newPoles = append(newPoles, &order.Chart{
			OrderSymbolID:   v.OrderSymbolID,
			OrderIntervalID: newInterval.ID,
			OpenPrice:       v.ClosePrice,
			LowPrice:        v.ClosePrice,
			HighPrice:       v.ClosePrice,
			ClosePrice:      v.ClosePrice,
			Volume:          decimal.NewFromFloat(0),
		})
	}

	if err := tx.Table(_chart.TableName()).CreateInBatches(newPoles, len(newPoles)).Error; err != nil {
		return fmt.Errorf("failed to create in batches into chart table : %v", err)
	}

	return nil
}

func (o *orderService) HandleChartDraw(tx *gorm.DB, event []*grpc_order.OrderMatching) error {
	var recentInterval []uint
	var symbolNames []string
	var symbolIDs []uint
	var poleIDList []uint

	var _interval *order.OrderInterval

	var recentPoles []*order.Chart
	var _chart *order.Chart

	var _symbol *order.OrderSymbol
	var symbols []*order.OrderSymbol

	var symbolNameMap = make(map[string]*order.OrderSymbol)
	var symbolIDMap = make(map[uint]*order.OrderSymbol)
	var priceMap = make(map[string]*types.PoleData)

	for _, v := range event {
		if _, exist := symbolNameMap[v.Symbol.String()]; !exist {
			symbolNames = append(symbolNames, v.Symbol.String())
			symbolNameMap[v.Symbol.String()] = new(order.OrderSymbol)
		}
		if _, exist := priceMap[v.Symbol.String()]; !exist {
			priceMap[v.Symbol.String()] = types.NewPriceSet()
		}
		if priceMap[v.Symbol.String()].HighPrice < v.UnitPrice {
			priceMap[v.Symbol.String()].HighPrice = v.UnitPrice
		}
		if priceMap[v.Symbol.String()].LowPrice > v.UnitPrice {
			priceMap[v.Symbol.String()].LowPrice = v.UnitPrice
		}
		priceMap[v.Symbol.String()].ClosePrice = v.UnitPrice
		priceMap[v.Symbol.String()].TradingValue += v.UnitPrice * v.Quantity
		priceMap[v.Symbol.String()].Volume += v.Quantity
	}

	if err := tx.Table(_symbol.TableName()).Where("name IN ?", symbolNames).Find(&symbols).Error; err != nil {
		return fmt.Errorf("failed to get symbol : %v", err)
	}

	for _, v := range symbols {
		symbolIDs = append(symbolIDs, v.ID)
		symbolNameMap[v.Name] = v
		symbolIDMap[v.ID] = v
	}

	// 현재 기준, duration 별로 가장 최근 column을 order_interval에서 가져온다.
	if err := tx.Table(fmt.Sprintf("%v as o1", _interval.TableName())).
		Select("o1.id").
		Joins("INNER JOIN (SELECT duration, MAX(start_time) AS max_start_time " +
			"FROM order_interval GROUP BY duration) o2 " +
			"ON o1.duration = o2.duration " +
			"AND o1.start_time = o2.max_start_time").
		Scan(&recentInterval).Error; err != nil {
		return fmt.Errorf("failed to get recent interval : %v", err)
	}

	// 현재 기준, order_interval 별로 각 symbol 별 가장 최근 row 가져온다.
	if err := tx.Table(fmt.Sprintf("%v as c", _chart.TableName())).
		Select("c.*").
		Where("c.order_interval_id IN (?)", recentInterval).
		Where("c.order_symbol_id IN (?)", symbolIDs).
		Scan(&recentPoles).Error; err != nil {
		return fmt.Errorf("failed to get recent poles : %v", err)
	}

	for _, pole := range recentPoles {
		poleIDList = append(poleIDList, pole.ID)
		high, _ := pole.HighPrice.Float64()
		low, _ := pole.LowPrice.Float64()
		eventLow := priceMap[symbolIDMap[pole.OrderSymbolID].Name].LowPrice
		eventHigh := priceMap[symbolIDMap[pole.OrderSymbolID].Name].HighPrice
		pole.ClosePrice = decimal.NewFromFloat(priceMap[symbolIDMap[pole.OrderSymbolID].Name].ClosePrice)
		pole.TradingValue = pole.TradingValue.Add(decimal.NewFromFloat(priceMap[symbolIDMap[pole.OrderSymbolID].Name].TradingValue))
		pole.Volume = pole.Volume.Add(decimal.NewFromFloat(priceMap[symbolIDMap[pole.OrderSymbolID].Name].Volume))
		if high < eventHigh {
			pole.HighPrice = decimal.NewFromFloat(eventHigh)
		}
		if low > eventLow {
			pole.LowPrice = decimal.NewFromFloat(eventLow)
		}
	}

	updateQuery := "UPDATE " + _chart.TableName() + " SET high_price = CASE id "
	for _, pole := range recentPoles {
		updateQuery += fmt.Sprintf("WHEN %v THEN %v ", pole.ID, pole.HighPrice)
	}
	updateQuery += "END, low_price = CASE id "
	for _, pole := range recentPoles {
		updateQuery += fmt.Sprintf("WHEN %v THEN %v ", pole.ID, pole.LowPrice)
	}
	updateQuery += "END, close_price = CASE id "
	for _, pole := range recentPoles {
		updateQuery += fmt.Sprintf("WHEN %v THEN %v ", pole.ID, pole.ClosePrice)
	}
	updateQuery += "END, volume = CASE id "
	for _, pole := range recentPoles {
		updateQuery += fmt.Sprintf("WHEN %v THEN %v ", pole.ID, pole.Volume)
	}
	updateQuery += "END, trading_value = CASE id "
	for _, pole := range recentPoles {
		updateQuery += fmt.Sprintf("WHEN %v THEN %v ", pole.ID, pole.TradingValue)
	}
	updateQuery += fmt.Sprintf("END WHERE id IN (?)")
	if err := tx.Exec(updateQuery, poleIDList).Error; err != nil {
		return fmt.Errorf("failed to batch update chart table : %v", err)
	}

	return nil
}

func (o *orderService) HandleOrderCancellationInBatch(tx *gorm.DB, event []*grpc_order.OrderCancelled) ([]*grpc_order.OrderCancelled, error) {
	var orderUUIDs []string
	var _orderBook *order.OrderBook

	var orderBookList []*order.OrderBook
	var remain []*grpc_order.OrderCancelled

	for _, cancellation := range event {
		orderUUIDs = append(orderUUIDs, cancellation.OrderUUID)
	}

	// order book table에서 order_uuid가 event에 있는지 확인한다.
	if err := tx.Table(_orderBook.TableName()).Where("order_uuid IN ?", orderUUIDs).Find(&orderBookList).Error; err != nil {
		return nil, fmt.Errorf("failed to batch update order book table : %v", err)
	}

	for _, e := range event {
		flag := func(e *grpc_order.OrderCancelled) bool {
			for _, orderBook := range orderBookList {
				if e.OrderUUID == orderBook.OrderUUID {
					return true
				}
			}
			return false
		}(e)

		// order book table에 없으면 remain에 append한다.
		if !flag {
			remain = append(remain, e)
		}
	}

	// order book table에 없는 event는 무시한다.
	if err := tx.Table(_orderBook.TableName()).Where("order_uuid IN ?", orderUUIDs).UpdateColumn("status", Cancelled).Error; err != nil {
		return nil, fmt.Errorf("failed to batch update order book table : %v", err)
	}

	return remain, nil
}