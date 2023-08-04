package instance

import (
	"context"
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) FindServerByIP(tx *gorm.DB, ip string) (result *entity.FinexblockServerIP, err error) {
	if err := tx.Table(result.TableName()).Where("ip = ?", ip).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) FindServerByID(tx *gorm.DB, id uint) (result *entity.FinexblockServer, err error) {
	if err := tx.Table(result.TableName()).Where("id = ?", id).First(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repository) FindServerByName(tx *gorm.DB, name string) (*entity.FinexblockServer, error) {
	var _server *entity.FinexblockServer
	if err := tx.Table(_server.TableName()).Where("name = ?", name).First(&_server).Error; err != nil {
		return nil, err
	}
	return _server, nil
}

func (r *repository) InsertErrorLog(tx *gorm.DB, errorLog *entity.FinexblockErrorLog) (*entity.FinexblockErrorLog, error) {
	if err := tx.Table(errorLog.TableName()).Create(errorLog).Error; err != nil {
		return nil, err
	}
	return errorLog, nil
}

func (r *repository) Conn() *gorm.DB {
	return r.db
}

func (r *repository) Tx(level sql.IsolationLevel) *gorm.DB {
	return r.db.Begin(&sql.TxOptions{Isolation: level})
}

func (r *repository) Ctx() context.Context {
	return context.Background()
}

func (r *repository) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}

func newRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}