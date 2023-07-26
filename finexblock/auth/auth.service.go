package auth

import (
	"context"
	"database/sql"
	"errors"
	"github.com/finexblock-dev/gofinexblock/finexblock/admin"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/user"
	"github.com/finexblock-dev/gofinexblock/finexblock/utils"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"os"
	"time"
)

type authService struct {
	userRepository  user.Repository
	adminRepository admin.Repository
}

func (a *authService) AdminLogin(email, password string) (result string, err error) {
	var _admin *entity.Admin
	var _credentials *entity.Admin
	var _token string
	if err = a.adminRepository.Conn().Transaction(func(tx *gorm.DB) error {

		_admin, err = a.adminRepository.FindAdminByEmail(tx, email)
		if err != nil {
			return err
		}

		if _admin.IsBlocked {
			return errors.New("you've been blocked")
		}

		_credentials, err = a.adminRepository.FindAdminCredentialsByID(tx, _admin.ID)
		if err != nil {
			return err
		}

		if !utils.CompareHash(_credentials.Password, password) {
			return errors.New("invalid credentials")
		}

		_token, err = a.AdminToken(_admin)
		if err != nil {
			return err
		}

		return nil
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted}); err != nil {
		return "", err
	}

	return _token, nil
}

func (a *authService) AdminToken(_admin *entity.Admin) (string, error) {
	claims := jwt.MapClaims{
		"adminId": _admin.ID,
		"grade":   _admin.Grade,
		"email":   _admin.Email,
		"exp":     time.Now().Add(time.Hour * 8).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (a *authService) AdminRegister(email, password string) (result *entity.Admin, err error) {
	if err = a.adminRepository.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.adminRepository.InsertAdmin(tx, email, password)
		return err
	}); err != nil {
		return nil, err
	}
	return result, nil
}

func (a *authService) Tx(level sql.IsolationLevel) *gorm.DB {
	return a.adminRepository.Tx(level)
}

func (a *authService) Ctx() context.Context {
	return context.Background()
}

func (a *authService) Conn() *gorm.DB {
	return a.adminRepository.Conn()
}

func (a *authService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}

func newAuthService(adminRepository admin.Repository, userRepository user.Repository) *authService {
	return &authService{adminRepository: adminRepository, userRepository: userRepository}
}