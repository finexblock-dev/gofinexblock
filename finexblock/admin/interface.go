package admin

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"gorm.io/gorm"
)

type RepositoryImpl interface {
	types.Repository
	FindAdminByID(tx *gorm.DB, id uint) (*admin.Admin, error)
	FindAdminByEmail(tx *gorm.DB, email string) (*admin.Admin, error)
	FindAdminCredentialsByID(tx *gorm.DB, id uint) (*admin.Admin, error)
	FindAccessToken(tx *gorm.DB, limit, offset int) ([]*admin.AdminAccessToken, error) // find access token by admin id
	FindAdminByGrade(tx *gorm.DB, grade admin.GradeType, limit, offset int) ([]*admin.Admin, error)
	FindAllAdmin(tx *gorm.DB, limit, offset int) ([]*admin.Admin, error)

	CreateAdmin(tx *gorm.DB, email, password string) (*admin.Admin, error)

	UpdateAdminByID(tx *gorm.DB, id uint, admin *admin.Admin) error

	DeleteAdminByID(tx *gorm.DB, id uint) error
}