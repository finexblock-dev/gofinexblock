package order

import (
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func (o *orderRepository) FindManySymbolByName(tx *gorm.DB, names []string) (result []*entity.OrderSymbol, err error) {
	var _table *entity.OrderSymbol

	if err = tx.Table(_table.TableName()).Where("name IN ?", names).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (o *orderRepository) BatchInsertOrderBook(tx *gorm.DB, orders []*entity.OrderBook) (err error) {
	var _table *entity.OrderBook

	if err = tx.Table(_table.TableName()).CreateInBatches(&orders, len(orders)).Error; err != nil {
		return err
	}
	return nil
}

func (o *orderRepository) BatchUpdateOrderBookStatus(tx *gorm.DB, orderUUIDs []string, status types.OrderStatus) (err error) {
	var _table *entity.OrderBook

	if err = tx.Table(_table.TableName()).Where("uuid IN ?", orderUUIDs).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

func (o *orderRepository) BatchInsertOrderBookDifference(tx *gorm.DB, differences []*entity.OrderBookDifference) (err error) {
	var _table *entity.OrderBookDifference

	if err = tx.Table(_table.TableName()).CreateInBatches(&differences, len(differences)).Error; err != nil {
		return err
	}
	return nil
}

func (o *orderRepository) FindRecentIntervalGroupByDuration(tx *gorm.DB) (result []*entity.OrderInterval, err error) {
	var _table *entity.OrderInterval

	if err = tx.Table(_table.Alias()).
		Joins("INNER JOIN (SELECT duration, MAX(start_time) AS max_start_time " +
			"FROM order_interval GROUP BY duration) oi2 " +
			"ON oi.duration = oi2.duration " +
			"AND oi.start_time = oi2.max_start_time").
		Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (o *orderRepository) InsertOrderInterval(tx *gorm.DB, interval *entity.OrderInterval) (result *entity.OrderInterval, err error) {
	if err = tx.Table(interval.TableName()).Create(interval).Error; err != nil {
		return nil, err
	}

	return interval, nil
}

func (o *orderRepository) FindChartByInterval(tx *gorm.DB, intervalID uint) (result []*entity.Chart, err error) {
	var _table *entity.Chart

	if err = tx.Table(_table.TableName()).Where("interval_id = ?", intervalID).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (o *orderRepository) FindChartByCond(tx *gorm.DB, intervalIDs []uint, symbolIDs []uint) (result []*entity.Chart, err error) {
	var _chart *entity.Chart
	if err = tx.Table(_chart.TableName()).
		Where("order_interval_id IN (?) AND order_symbol_id IN (?)", intervalIDs, symbolIDs).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (o *orderRepository) BatchInsertChart(tx *gorm.DB, charts []*entity.Chart) (err error) {
	var _table *entity.Chart

	if err = tx.Table(_table.TableName()).CreateInBatches(&charts, len(charts)).Error; err != nil {
		return err
	}
	return nil
}

func (o *orderRepository) BatchInsertOrderMatchingHistory(tx *gorm.DB, histories []*entity.OrderMatchingHistory) (err error) {
	var _table *entity.OrderMatchingHistory

	if err = tx.Table(_table.TableName()).CreateInBatches(&histories, len(histories)).Error; err != nil {
		return err
	}
	return nil
}

func (o *orderRepository) BatchInsertOrderMatchingEvent(tx *gorm.DB, events []*entity.OrderMatchingEvent) (err error) {
	var _table *entity.OrderMatchingEvent

	if err = tx.Table(_table.TableName()).CreateInBatches(&events, len(events)).Error; err != nil {
		return err
	}
	return nil
}

func (o *orderRepository) InsertSnapshot(tx *gorm.DB, _snapshot *entity.SnapshotOrderBook) (result *entity.SnapshotOrderBook, err error) {
	if err = tx.Table(_snapshot.TableName()).Create(_snapshot).Error; err != nil {
		return nil, err
	}

	return _snapshot, nil
}

func (o *orderRepository) FindSnapshotByOrderSymbolID(tx *gorm.DB, symbolID uint) (result *entity.SnapshotOrderBook, err error) {
	if err = tx.Table(result.TableName()).Where("order_symbol_id = ?", symbolID).Order("created_at DESC").Limit(1).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (o *orderRepository) FindSymbolByName(tx *gorm.DB, name string) (result *entity.OrderSymbol, err error) {
	if err = tx.Table(result.TableName()).Where("name = ?", name).First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (o *orderRepository) FindSymbolByID(tx *gorm.DB, id uint) (result *entity.OrderSymbol, err error) {
	if err = tx.Table(result.TableName()).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (o *orderRepository) FindManyOrderByUUID(tx *gorm.DB, uuids []string) (result []*entity.OrderBook, err error) {
	var _table = &entity.OrderBook{}

	if err = tx.Table(_table.TableName()).Where("uuid IN ?", uuids).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (o *orderRepository) FindRecentIntervalByDuration(tx *gorm.DB, duration types.Duration) (result *entity.OrderInterval, err error) {
	var _table = &entity.OrderInterval{}

	if err = tx.Table(_table.TableName()).Where("duration = ?", duration).Order("end_time desc").First(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (o *orderRepository) Tx(level sql.IsolationLevel) *gorm.DB {
	return o.db.Begin(&sql.TxOptions{Isolation: level})
}

func (o *orderRepository) Conn() *gorm.DB {
	return o.db
}

func newOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db: db}
}