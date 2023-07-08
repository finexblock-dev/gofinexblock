package auth

import (
	"context"
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type authService struct {
	authRepository Repository
}

func (a *authService) Conn() *gorm.DB {
	return a.authRepository.Conn()
}

func (a *authService) Login(c *fiber.Ctx, email, password string) (string, error) {
	panic("implement me")
}

func (a *authService) GenerateToken(c *fiber.Ctx, _admin *admin.Admin) (string, error) {
	panic("implement me")
}

func (a *authService) Register(c *fiber.Ctx, email, password string) (*admin.Admin, error) {
	panic("implement me")

}

func (a *authService) Tx(level sql.IsolationLevel) *gorm.DB {
	return a.authRepository.Tx(level)
}

func newAuthService(authRepository Repository) *authService {
	return &authService{authRepository: authRepository}
}

func (a *authService) Ctx() context.Context {
	return context.Background()
}

func (a *authService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}
