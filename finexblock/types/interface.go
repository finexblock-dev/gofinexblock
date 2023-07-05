package types

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
)

type Model interface {
	TableName() string
	Alias() string
}

type Repository interface {
	Tx(level sql.IsolationLevel) *gorm.DB
	Conn() *gorm.DB
}

type Service interface {
	Ctx() context.Context
	CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc)
	Tx(level sql.IsolationLevel) *gorm.DB
	Conn() *gorm.DB
}