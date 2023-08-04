package image

import (
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"github.com/finexblock-dev/gofinexblock/pkg/types"
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
	FindAllImages(limit, offset int) ([]*entity.Image, error)
}

func NewRepository(db *gorm.DB) Repository {
	return newImageRepository(db)
}

func NewService(db *gorm.DB, bucket, basePath string) Service {
	return newImageService(NewRepository(db), bucket, basePath)
}