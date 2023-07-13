package image

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"gorm.io/gorm"
	"mime/multipart"
)

type Repository interface {
	types.Repository
	FindAllImages(tx *gorm.DB, limit, offset int) ([]*entity.Image, error)
	UploadFiles(tx *gorm.DB, f *multipart.Form, bucket, basePath string) ([]*entity.Image, error)
}

type Service interface {
	types.Service
	UploadFile(f *multipart.Form) ([]*entity.Image, error)
	FindAllImages(tx *gorm.DB, limit, offset int) ([]*entity.Image, error)
}

func NewRepository(db *gorm.DB) Repository {
	return newImageRepository(db)
}

func NewService(db *gorm.DB) Service {
	return newImageService(NewRepository(db))
}
