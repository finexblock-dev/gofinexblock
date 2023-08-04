package error_log

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/files"
	"github.com/finexblock-dev/gofinexblock/finexblock/instance"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"gorm.io/gorm"
)

type Repository interface {
	types.Repository
	InsertErrorLog(tx *gorm.DB, errorLog *entity.FinexblockErrorLog) (*entity.FinexblockErrorLog, error)
	UploadErrorLogS3(errorLog *entity.FinexblockErrorLog, bucket string) (err error)
	Log(v ...any)
}

type Service interface {
	types.Service
}

func NewRepository(db *gorm.DB, prefix, filename string) Repository {
	return newRepository(db, instance.NewRepository(db), files.NewWriter(prefix, filename))
}

func NewService(db *gorm.DB, prefix, filename string) Service {
	return newService(NewRepository(db, prefix, filename))
}
