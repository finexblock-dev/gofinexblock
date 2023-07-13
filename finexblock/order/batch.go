package order

import (
	"fmt"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func (o *orderService) HandleChartDraw(tx *gorm.DB, event []*grpc_order.OrderMatching) error {
	var recentInterval []uint
	var symbolNames []string
	var symbolIDs []uint
	var poleIDList []uint

	var _interval *entity.OrderInterval

	var recentPoles []*entity.Chart
	var _chart *entity.Chart

	var _symbol *entity.OrderSymbol
	var symbols []*entity.OrderSymbol

	var symbolNameMap = make(map[string]*entity.OrderSymbol)
	var symbolIDMap = make(map[uint]*entity.OrderSymbol)
	var priceMap = make(map[string]*types.PoleData)

	for _, v := range event {
		if _, exist := symbolNameMap[v.Symbol.String()]; !exist {
			symbolNames = append(symbolNames, v.Symbol.String())
			symbolNameMap[v.Symbol.String()] = new(entity.OrderSymbol)
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
