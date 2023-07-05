package user

import (
	context "context"
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/user"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"gorm.io/gorm"
)

type Service interface {
	types.Service
	FindUserByUUID(tx *gorm.DB, uuid string) (*user.User, error)
	FindUserByUUIDs(tx *gorm.DB, uuids []string) ([]*user.User, error)
	FindUserByID(tx *gorm.DB, id uint) (*user.User, error)

	BlockUser(tx *gorm.DB, id uint) error
	UnBlockUser(tx *gorm.DB, id uint) error
}

type userService struct {
	db *gorm.DB
}

func (u *userService) Ctx() context.Context {
	return context.Background()
}

func (u *userService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}

func (u *userService) Tx(level sql.IsolationLevel) *gorm.DB {
	return u.db.Begin(&sql.TxOptions{Isolation: level})
}

func (u *userService) Conn() *gorm.DB {
	return u.db
}

func newUserService(db *gorm.DB) *userService {
	return &userService{db: db}
}

func NewService(db *gorm.DB) Service {
	return newUserService(db)
}