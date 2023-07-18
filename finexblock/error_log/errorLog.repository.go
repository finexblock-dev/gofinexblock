package error_log

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/goaws"
	instance "github.com/finexblock-dev/gofinexblock/finexblock/instance"
	"github.com/finexblock-dev/gofinexblock/finexblock/secure"
	"gorm.io/gorm"
)

type repository struct {
	db           *gorm.DB
	instanceRepo instance.Repository
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

func newRepository(db *gorm.DB, instanceRepo instance.Repository) *repository {
	return &repository{db: db, instanceRepo: instanceRepo}
}
