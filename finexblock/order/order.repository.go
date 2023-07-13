package order

import (
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"gorm.io/gorm"
)

type orderRepository struct {
	db *gorm.DB
}

func (o *orderRepository) InsertSnapshot(tx *gorm.DB, symbolID uint, _snapshot *entity.SnapshotOrderBook) (result *entity.SnapshotOrderBook, err error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderRepository) FindSymbolByName(tx *gorm.DB, name string) (result *entity.OrderSymbol, err error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderRepository) FindSymbolByID(tx *gorm.DB, id uint) (result *entity.OrderSymbol, err error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderRepository) FindManyOrderByUUID(tx *gorm.DB, uuids []string) (result []*entity.OrderBook, err error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderRepository) FindRecentIntervalByName(tx *gorm.DB, name string) (result *entity.OrderInterval, err error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderRepository) Tx(level sql.IsolationLevel) *gorm.DB {
	//TODO implement me
	panic("implement me")
}

func (o *orderRepository) Conn() *gorm.DB {
	//TODO implement me
	panic("implement me")
}

func newOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db: db}
}