package image

import (
	"database/sql"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/goaws"
	"github.com/finexblock-dev/gofinexblock/finexblock/secure"
	"gorm.io/gorm"
	"mime/multipart"
)

var table = &entity.Image{}

type imageRepository struct {
	db *gorm.DB
}

func (i *imageRepository) FindAllImages(tx *gorm.DB, limit, offset int) ([]*entity.Image, error) {
	var result []*entity.Image
	var err error

	if err = tx.Table(table.TableName()).Limit(limit).Offset(offset).Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (i *imageRepository) UploadFiles(tx *gorm.DB, f *multipart.Form, bucket, basePath string) (result []*entity.Image, err error) {
	var client *s3.S3
	var uploadResult map[string]string
	var files []*multipart.FileHeader
	var sess *session.Session

	sess, err = secure.GetSessionFromEnv()
	if err != nil {
		return nil, err
	}

	client = goaws.NewS3Client(sess)

	for _, header := range f.File {
		files = append(files, header...)
	}

	uploadResult, err = goaws.UploadBatch(client, files, bucket, basePath)
	if err != nil {
		return nil, err
	}

	for filename, url := range uploadResult {
		var img = &entity.Image{Url: url, Key: filename}
		result = append(result, img)
	}

	if err = tx.Table(table.TableName()).CreateInBatches(&result, len(result)).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (i *imageRepository) Conn() *gorm.DB {
	return i.db
}

func newImageRepository(db *gorm.DB) *imageRepository {
	return &imageRepository{db: db}
}

func (i *imageRepository) Tx(level sql.IsolationLevel) *gorm.DB {
	return i.db.Begin(&sql.TxOptions{Isolation: level})
}
