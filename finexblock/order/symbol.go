package order

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/order"
	"gorm.io/gorm"
)

func (o *orderService) FindSymbolByName(tx *gorm.DB, name string) (*order.OrderSymbol, error) {
	var _symbol *order.OrderSymbol
	var err error
	if err = tx.Where("name = ?", name).First(&_symbol).Error; err != nil {
		return nil, err
	}

	return _symbol, nil
}

func (o *orderService) FindSymbolByID(tx *gorm.DB, id uint) (*order.OrderSymbol, error) {
	var _symbol *order.OrderSymbol
	var err error

	if err = tx.Where("id = ?", id).First(&_symbol).Error; err != nil {
		return nil, err
	}

	return _symbol, nil
}