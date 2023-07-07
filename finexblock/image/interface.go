package image

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/image"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"mime/multipart"
)

type Repository interface {
	types.Repository
	FindAllImages(tx *gorm.DB, limit, offset int) ([]*image.Image, error)
	UploadFile(tx *gorm.DB, f *multipart.Form) ([]*image.Image, error)
}

type Service interface {
	types.Service
	UploadFile(c *fiber.Ctx, f *multipart.Form) ([]*image.Image, error)
}

func NewRepository(db *gorm.DB) Repository {
	return newImageRepository(db)
}

func NewService(repo Repository) Service {
	return newImageService(repo)
}
