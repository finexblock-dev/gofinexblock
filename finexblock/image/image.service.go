package image

import (
	"context"
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/image"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"mime/multipart"
)

type imageService struct {
	imageRepository Repository
}

func (i *imageService) Conn() *gorm.DB {
	return i.imageRepository.Conn()
}

func (i *imageService) Tx(level sql.IsolationLevel) *gorm.DB {
	return i.imageRepository.Tx(level)
}

func newImageService(imageRepository Repository) *imageService {
	return &imageService{imageRepository: imageRepository}
}

func (i *imageService) UploadFile(c *fiber.Ctx, f *multipart.Form) ([]*image.Image, error) {
	panic("implement me")
}

func (i *imageService) Ctx() context.Context {
	return context.Background()
}

func (i *imageService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}
