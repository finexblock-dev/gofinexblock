package goerror

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"github.com/finexblock-dev/gofinexblock/pkg/files"
	"github.com/finexblock-dev/gofinexblock/pkg/goaws"
	"github.com/finexblock-dev/gofinexblock/pkg/instance"
	"github.com/finexblock-dev/gofinexblock/pkg/secure"
	"gorm.io/gorm"
)

type repository struct {
	db           *gorm.DB
	instanceRepo instance.Repository
	logger       files.Writer
}

func (r *repository) Log(v ...any) {
	r.logger.Println(v)
}

func (r *repository) UploadErrorLogS3(errorLog *entity.FinexblockErrorLog, bucket string) (err error) {
	var client *s3.S3
	var sess *session.Session

	sess, err = secure.GetSessionFromEnv()
	if err != nil {
		fmt.Println(err)
	}

	client = goaws.NewS3Client(sess)

	goaws.UploadErrorLog(client, errorLog, bucket)

	return nil
}

func (r *repository) InsertErrorLog(tx *gorm.DB, errorLog *entity.FinexblockErrorLog) (*entity.FinexblockErrorLog, error) {
	return r.instanceRepo.InsertErrorLog(tx, errorLog)
}

func (r *repository) Conn() *gorm.DB {
	return r.db
}

func (r *repository) Tx(level sql.IsolationLevel) *gorm.DB {
	return r.db.Begin(&sql.TxOptions{Isolation: level})
}

func (r *repository) Ctx() context.Context {
	return context.Background()
}

func (r *repository) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}

func newRepository(db *gorm.DB, instanceRepo instance.Repository, logger files.Writer) *repository {
	return &repository{db: db, instanceRepo: instanceRepo, logger: logger}
}
