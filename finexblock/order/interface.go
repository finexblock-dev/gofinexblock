package order

import (
	"context"
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/order"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"gorm.io/gorm"
	"time"
)

type Service interface {
	types.Service
	Snapshot(tx *gorm.DB, symbolID uint, bids []*grpc_order.Order, asks []*grpc_order.Order) error
	FindSymbolByName(tx *gorm.DB, name string) (*order.OrderSymbol, error)
	FindSymbolByID(tx *gorm.DB, id uint) (*order.OrderSymbol, error)
	FindManyOrderByUUID(tx *gorm.DB, uuids []string) ([]*order.OrderBook, error)
	FindRecentIntervalByName(tx *gorm.DB, name string) (*order.OrderInterval, error)

	HandleOrderMatchingEventInBatch(tx *gorm.DB, event []*grpc_order.OrderMatching) error
	HandleOrderFulfillmentInBatch(tx *gorm.DB, event []*grpc_order.OrderFulfillment) ([]*grpc_order.OrderFulfillment, error)
	HandleOrderPartialFillInBatch(tx *gorm.DB, event []*grpc_order.OrderPartialFill) ([]*grpc_order.OrderPartialFill, error)
	HandleOrderInitializeInBatch(tx *gorm.DB, event []*grpc_order.OrderInitialize) error
	HandleOrderInterval(tx *gorm.DB, name string, duration time.Duration) error
	HandleChartDraw(tx *gorm.DB, event []*grpc_order.OrderMatching) error
	HandleOrderCancellationInBatch(tx *gorm.DB, event []*grpc_order.OrderCancelled) ([]*grpc_order.OrderCancelled, error)
}

type orderService struct {
	db *gorm.DB
}

func (o *orderService) Conn() *gorm.DB {
	return o.db
}

func (o *orderService) Tx(level sql.IsolationLevel) *gorm.DB {
	return o.db.Begin(&sql.TxOptions{Isolation: level})
}

func (o *orderService) Ctx() context.Context {
	return context.Background()
}

func (o *orderService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}

func newOrderService(db *gorm.DB) *orderService {
	return &orderService{db: db}
}

func NewService(db *gorm.DB) Service {
	return newOrderService(db)
}