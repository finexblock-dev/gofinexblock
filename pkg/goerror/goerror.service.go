package goerror

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type service struct {
	repo Repository
}

func newService(repo Repository) *service {
	return &service{repo: repo}
}

func (a *service) Tx(level sql.IsolationLevel) *gorm.DB {
	return a.repo.Tx(level)
}

func (a *service) Conn() *gorm.DB {
	return a.repo.Conn()
}

func (a *service) Ctx() context.Context {
	return context.Background()
}

func (a *service) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}
