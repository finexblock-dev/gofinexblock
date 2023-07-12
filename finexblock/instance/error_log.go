package instance

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"gorm.io/gorm"
)

func (i *instanceService) InsertErrorLog(tx *gorm.DB, errorLog *entity.FinexblockErrorLog) (*entity.FinexblockErrorLog, error) {
	if err := tx.Table(errorLog.TableName()).Create(errorLog).Error; err != nil {
		return nil, err
	}
	return errorLog, nil
}
