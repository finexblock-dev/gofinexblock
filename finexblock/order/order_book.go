package order

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"gorm.io/gorm"
)

func (o *orderService) FindManyOrderByUUID(tx *gorm.DB, uuids []string) ([]*entity.OrderBook, error) {
	var _books []*entity.OrderBook
	var err error

	if err = tx.Where("uuid IN ?", uuids).Find(&_books).Error; err != nil {
		return nil, err
	}

	return _books, nil
}
