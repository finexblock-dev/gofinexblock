package order

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type orderService struct {
	repo Repository
}

func (o *orderService) Conn() *gorm.DB {
	return o.repo.Conn()
}

func (o *orderService) Tx(level sql.IsolationLevel) *gorm.DB {
	return o.repo.Tx(level)
}

func (o *orderService) Ctx() context.Context {
	return context.Background()
}

func (o *orderService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}

func newOrderService(db *gorm.DB) *orderService {
	return &orderService{repo: newOrderRepository(db)}
}