package user

import (
	"context"
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/user/structs"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type userService struct {
	repo Repository
}

func newUserService(db *gorm.DB, cluster *redis.ClusterClient) *userService {
	return &userService{repo: NewRepository(db, cluster)}
}

func (u *userService) Ctx() context.Context {
	return context.Background()
}

func (u *userService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}

func (u *userService) Tx(level sql.IsolationLevel) *gorm.DB {
	return u.repo.Tx(level)
}

func (u *userService) Conn() *gorm.DB {
	return u.repo.Conn()
}

func (u *userService) FindUserByUUID(uuid string) (result *entity.User, err error) {
	if err = u.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = u.repo.FindUserByUUID(tx, uuid)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userService) FindUserByUUIDs(uuids []string) (result []*entity.User, err error) {
	if err = u.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = u.repo.FindManyUserByUUID(tx, uuids)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userService) FindUserByID(id uint) (result *entity.User, err error) {
	if err = u.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = u.repo.FindUserByID(tx, id)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userService) FindUserMetadata(id uint) (result *entity.UserMetadata, err error) {
	if err = u.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = u.repo.FindUserMetadata(tx, id)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userService) SearchUser(input *structs.SearchUserInput) (result []*entity.UserMetadata, err error) {
	if err = u.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = u.repo.SearchUser(tx, input)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userService) BlockUser(id uint) (err error) {
	return u.Conn().Transaction(func(tx *gorm.DB) error {
		return u.repo.BlockUser(tx, id)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (u *userService) UnBlockUser(id uint) (err error) {
	return u.Conn().Transaction(func(tx *gorm.DB) error {
		return u.repo.UnBlockUser(tx, id)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (u *userService) CreateMemo(id uint, desc string) (err error) {
	return u.Conn().Transaction(func(tx *gorm.DB) error {
		return u.repo.CreateMemo(tx, id, desc)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}
