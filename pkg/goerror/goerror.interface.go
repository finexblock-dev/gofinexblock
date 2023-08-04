package goerror

import (
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"github.com/finexblock-dev/gofinexblock/pkg/files"
	"github.com/finexblock-dev/gofinexblock/pkg/instance"
	"github.com/finexblock-dev/gofinexblock/pkg/types"
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
