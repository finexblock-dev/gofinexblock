package order

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/order"
	"gorm.io/gorm"
)

func (o *orderService) FindRecentIntervalByName(tx *gorm.DB, name string) (*order.OrderInterval, error) {
	var _interval *order.OrderInterval
	var err error

	if err = tx.Table(_interval.TableName()).Where("duration = ?", name).Order("end_time desc").First(&_interval).Error; err != nil {
		return nil, err
	}

	return _interval, nil
}