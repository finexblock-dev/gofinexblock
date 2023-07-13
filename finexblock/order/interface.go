package order

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	types.Repository
	InsertSnapshot(tx *gorm.DB, symbolID uint, _snapshot *entity.SnapshotOrderBook) (result *entity.SnapshotOrderBook, err error)

	FindSymbolByName(tx *gorm.DB, name string) (result *entity.OrderSymbol, err error)
	FindSymbolByID(tx *gorm.DB, id uint) (result *entity.OrderSymbol, err error)
	FindManySymbolByName(tx *gorm.DB, names []string) (result []*entity.OrderSymbol, err error)

	FindManyOrderByUUID(tx *gorm.DB, uuids []string) (result []*entity.OrderBook, err error)
	BatchInsertOrderBook(tx *gorm.DB, orders []*entity.OrderBook) (err error)
	BatchUpdateOrderBookStatus(tx *gorm.DB, orderUUIDs []string, status types.OrderStatus) (err error)

	BatchInsertOrderBookDifference(tx *gorm.DB, differences []*entity.OrderBookDifference) (err error)

	FindRecentIntervalByDuration(tx *gorm.DB, duration types.Duration) (result *entity.OrderInterval, err error)
	FindRecentIntervalGroupByDuration(tx *gorm.DB) (result []*entity.OrderInterval, err error)
	InsertOrderInterval(tx *gorm.DB, interval *entity.OrderInterval) (result *entity.OrderInterval, err error)

	FindChartByInterval(tx *gorm.DB, intervalID uint) (result []*entity.Chart, err error)
	FindChartByCond(tx *gorm.DB, intervalIDs []uint, symbolIDs []uint) (result []*entity.Chart, err error)
	BatchInsertChart(tx *gorm.DB, charts []*entity.Chart) (err error)
	// BatchUpdateChart(tx *gorm.DB) TODO: implement this

	BatchInsertOrderMatchingHistory(tx *gorm.DB, histories []*entity.OrderMatchingHistory) (err error)

	BatchInsertOrderMatchingEvent(tx *gorm.DB, events []*entity.OrderMatchingEvent) (err error)
}

type Service interface {
	types.Service

	InsertSnapshot(symbolID uint, _snapshot *entity.SnapshotOrderBook) (result *entity.SnapshotOrderBook, err error)
	FindSymbolByName(name string) (result *entity.OrderSymbol, err error)
	FindSymbolByID(id uint) (result *entity.OrderSymbol, err error)
	FindManyOrderByUUID(uuids []string) (result []*entity.OrderBook, err error)
	FindRecentIntervalByDuration(duration types.Duration) (result *entity.OrderInterval, err error)

	OrderMatchingEventInBatch(event []*grpc_order.OrderMatching) (err error)
	ChartDraw(event []*grpc_order.OrderMatching) (err error)

	LimitOrderFulfillmentInBatch(event []*grpc_order.OrderFulfillment) (remain []*grpc_order.OrderFulfillment, err error)
	LimitOrderPartialFillInBatch(event []*grpc_order.OrderPartialFill) (remain []*grpc_order.OrderPartialFill, err error)
	LimitOrderInitializeInBatch(event []*grpc_order.OrderInitialize) (err error)
	LimitOrderCancellationInBatch(event []*grpc_order.OrderCancelled) (remain []*grpc_order.OrderCancelled, err error)

	HandleOrderInterval(name types.Duration, duration time.Duration) (err error)
}

func NewRepository(db *gorm.DB) Repository {
	return newOrderRepository(db)
}

func NewService(db *gorm.DB) Service {
	return newOrderService(db)
}
