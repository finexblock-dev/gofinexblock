package instance

import (
	context "context"
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/instance"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"gorm.io/gorm"
)

type Service interface {
	types.Service
	FindServerByName(tx *gorm.DB, name string) (*instance.FinexblockServer, error)
	InsertErrorLog(tx *gorm.DB, errorLog *instance.FinexblockErrorLog) (*instance.FinexblockErrorLog, error)
}

type instanceService struct {
	db *gorm.DB
}

func (i *instanceService) Tx(level sql.IsolationLevel) *gorm.DB {
	return i.db.Begin(&sql.TxOptions{Isolation: level})
}

func (i *instanceService) Ctx() context.Context {
	return context.Background()
}

func (i *instanceService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}

func newInstanceService(db *gorm.DB) *instanceService {
	return &instanceService{db: db}
}

func NewService(db *gorm.DB) Service {
	return newInstanceService(db)
}