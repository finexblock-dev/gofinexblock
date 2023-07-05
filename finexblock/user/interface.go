package user

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/user"
	"gorm.io/gorm"
)

type Service interface {
	FindUserByUUID(tx *gorm.DB, uuid string) (*user.User, error)
	FindUserByUUIDs(tx *gorm.DB, uuids []string) ([]*user.User, error)
	FindUserByID(tx *gorm.DB, id uint) (*user.User, error)

	BlockUser(tx *gorm.DB, id uint) error
	UnBlockUser(tx *gorm.DB, id uint) error
}

type userService struct {
	db *gorm.DB
}

func newUserService(db *gorm.DB) *userService {
	return &userService{db: db}
}

func NewService(db *gorm.DB) Service {
	return newUserService(db)
}