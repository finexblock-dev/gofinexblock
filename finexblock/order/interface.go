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
	FindManyOrderByUUID(tx *gorm.DB, uuids []string) (result []*entity.OrderBook, err error)
	FindRecentIntervalByName(tx *gorm.DB, name string) (result *entity.OrderInterval, err error)
}

type Service interface {
	types.Service
	Snapshot(symbolID uint, bids []*grpc_order.Order, asks []*grpc_order.Order) error
	FindSymbolByName(name string) (result *entity.OrderSymbol, err error)
	FindSymbolByID(id uint) (result *entity.OrderSymbol, err error)
	FindManyOrderByUUID(uuids []string) (result []*entity.OrderBook, err error)
	FindRecentIntervalByName(name string) (result *entity.OrderInterval, err error)

	HandleOrderMatchingEventInBatch(event []*grpc_order.OrderMatching) (err error)
	HandleOrderFulfillmentInBatch(event []*grpc_order.OrderFulfillment) (remain []*grpc_order.OrderFulfillment, err error)
	HandleOrderPartialFillInBatch(event []*grpc_order.OrderPartialFill) (remain []*grpc_order.OrderPartialFill, err error)
	HandleOrderInitializeInBatch(event []*grpc_order.OrderInitialize) (err error)
	HandleOrderInterval(name string, duration time.Duration) (err error)
	HandleChartDraw(event []*grpc_order.OrderMatching) (err error)
	HandleOrderCancellationInBatch(event []*grpc_order.OrderCancelled) (result []*grpc_order.OrderCancelled, err error)
}

func NewService(db *gorm.DB) Service {
	return newOrderService(db)
}