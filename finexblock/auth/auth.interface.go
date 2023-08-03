package auth

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/admin"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/finexblock-dev/gofinexblock/finexblock/user"
	"gorm.io/gorm"
)

type Repository interface {
	types.Repository
}

type Service interface {
	types.Service
	AdminLogin(email, password string) (string, error)
	AdminToken(_admin *entity.Admin) (string, error)
	AdminRegister(email, password string) (*entity.Admin, error)
}

func NewService(db *gorm.DB) Service {
	return newAuthService(admin.NewRepository(db), user.NewRepository(db))
}