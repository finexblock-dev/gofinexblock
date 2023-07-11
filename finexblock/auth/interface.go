package auth

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Repository interface {
	types.Repository
}

type Service interface {
	types.Service
	Login(c *fiber.Ctx, email, password string) (string, error)
	GenerateToken(c *fiber.Ctx, _admin *entity.Admin) (string, error)
	Register(c *fiber.Ctx, email, password string) (*entity.Admin, error)
}

func NewRepository(db *gorm.DB) Repository {
	return newAuthRepository(db)
}

func NewService(repo Repository) Service {
	return newAuthService(repo)
}
