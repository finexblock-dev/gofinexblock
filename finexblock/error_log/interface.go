package error_log

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/instance"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"gorm.io/gorm"
)

type Repository interface {
	types.Repository
	InsertErrorLog(tx *gorm.DB, errorLog *entity.FinexblockErrorLog) (*entity.FinexblockErrorLog, error)
	UploadErrorLogS3(errorLog *entity.FinexblockErrorLog, bucket string) (err error)
}

type Service interface {
	types.Service
	Log(v ...any)
}

func NewRepository(db *gorm.DB, instanceRepo instance.Repository) Repository {
	return newRepository(db, instanceRepo)
}

func NewService(db *gorm.DB, instanceRepo instance.Repository) Service {
	return newService(NewRepository(db, instanceRepo))
}
