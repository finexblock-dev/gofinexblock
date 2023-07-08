package image

import (
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/image"
	"gorm.io/gorm"
	"mime/multipart"
)

var table = &image.Image{}

type imageRepository struct {
	db *gorm.DB
}

func (i *imageRepository) Conn() *gorm.DB {
	return i.db
}

func (i *imageRepository) FindAllImages(tx *gorm.DB, limit, offset int) ([]*image.Image, error) {
	var result []*image.Image
	var err error

	if err = tx.Table(table.TableName()).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (i *imageRepository) UploadFile(tx *gorm.DB, f *multipart.Form) ([]*image.Image, error) {
	panic("implement me")
}

func newImageRepository(db *gorm.DB) *imageRepository {
	return &imageRepository{db: db}
}

func (i *imageRepository) Tx(level sql.IsolationLevel) *gorm.DB {
	return i.db.Begin(&sql.TxOptions{Isolation: level})
}
