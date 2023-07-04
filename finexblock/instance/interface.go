package instance

import (
	context "context"
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

func (i *instanceService) Ctx() context.Context {
	return context.Background()
}

func (i *instanceService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}

func (i *instanceService) FindServerByName(tx *gorm.DB, name string) (*instance.FinexblockServer, error) {
	var _server *instance.FinexblockServer
	if err := tx.Table(_server.TableName()).Where("name = ?", name).First(&_server).Error; err != nil {
		return nil, err
	}
	return _server, nil
}

func (i *instanceService) InsertErrorLog(tx *gorm.DB, errorLog *instance.FinexblockErrorLog) (*instance.FinexblockErrorLog, error) {
	if err := tx.Table(errorLog.TableName()).Create(errorLog).Error; err != nil {
		return nil, err
	}
	return errorLog, nil
}

func newInstanceService(db *gorm.DB) *instanceService {
	return &instanceService{db: db}
}

func NewService(db *gorm.DB) Service {
	return newInstanceService(db)
}