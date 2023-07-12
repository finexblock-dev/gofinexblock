package order

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"gorm.io/gorm"
)

func (o *orderService) FindRecentIntervalByName(tx *gorm.DB, name string) (*entity.OrderInterval, error) {
	var _interval *entity.OrderInterval
	var err error

	if err = tx.Table(_interval.TableName()).Where("duration = ?", name).Order("end_time desc").First(&_interval).Error; err != nil {
		return nil, err
	}

	return _interval, nil
}
