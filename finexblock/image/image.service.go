package image

import (
	"context"
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"gorm.io/gorm"
	"mime/multipart"
)

type imageService struct {
	repo Repository
}

func (i *imageService) FindAllImages(tx *gorm.DB, limit, offset int) (result []*entity.Image, err error) {
	if err = i.repo.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = i.repo.FindAllImages(tx, limit, offset)
		return err
	}); err != nil {
		return nil, err
	}
	return result, nil
}

func (i *imageService) UploadFile(f *multipart.Form) (result []*entity.Image, err error) {
	if err = i.repo.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = i.repo.UploadFiles(tx, f)
		return err
	}); err != nil {
		return nil, err
	}
	return result, nil
}

func (i *imageService) Conn() *gorm.DB {
	return i.repo.Conn()
}

func (i *imageService) Tx(level sql.IsolationLevel) *gorm.DB {
	return i.repo.Tx(level)
}

func newImageService(imageRepository Repository) *imageService {
	return &imageService{repo: imageRepository}
}

func (i *imageService) Ctx() context.Context {
	return context.Background()
}

func (i *imageService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}
