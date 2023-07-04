package instance

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/instance"
	"gorm.io/gorm"
)

func (i *instanceService) FindServerByName(tx *gorm.DB, name string) (*instance.FinexblockServer, error) {
	var _server *instance.FinexblockServer
	if err := tx.Table(_server.TableName()).Where("name = ?", name).First(&_server).Error; err != nil {
		return nil, err
	}
	return _server, nil
}