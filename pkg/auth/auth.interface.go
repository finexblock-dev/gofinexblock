package auth

import (
	"github.com/finexblock-dev/gofinexblock/pkg/admin"
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"github.com/finexblock-dev/gofinexblock/pkg/types"
	"github.com/finexblock-dev/gofinexblock/pkg/user"
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