package instance

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"gorm.io/gorm"
)

type Repository interface {
	types.Repository
	FindServerByIP(tx *gorm.DB, ip string) (result *entity.FinexblockServerIP, err error)
	FindServerByName(tx *gorm.DB, name string) (*entity.FinexblockServer, error)
	InsertErrorLog(tx *gorm.DB, errorLog *entity.FinexblockErrorLog) (*entity.FinexblockErrorLog, error)
}

func NewRepository(db *gorm.DB) Repository {
	return newRepository(db)
}
