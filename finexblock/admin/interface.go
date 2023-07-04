package admin

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/admin/dto"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
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

type Service interface {
	types.Service
	FindAllAdmin(ctx *fiber.Ctx, limit, offset int) error
	FindAdminByGrade(ctx *fiber.Ctx, grade admin.GradeType, limit, offset int) error
	FindLoginFailedLog(ctx *fiber.Ctx, adminID uint, limit, offset int) error
	//FindOnlineAdmin(ctx *fiber.Ctx) error
	FindLoginHistory(ctx *fiber.Ctx, adminID uint, limit, offset int) error
	SearchApiLog(ctx *fiber.Ctx, query *dto.SearchApiLogInput) error
	FindAllApiLog(ctx *fiber.Ctx, limit, offset int) error
	FindApiLogByAdmin(ctx *fiber.Ctx, adminID uint, limit, offset int) error
	FindApiLogByTimeCond(ctx *fiber.Ctx, start, end time.Time, limit, offset int) error
	FindApiLogByMethodCond(ctx *fiber.Ctx, method admin.ApiMethod, limit, offset int) error
	FindApiLogByEndpoint(ctx *fiber.Ctx, endpoint string, limit, offset int) error
	FindAllGradeUpdateLog(ctx *fiber.Ctx, limit, offset int) error
	SearchGradeUpdateLog(ctx *fiber.Ctx, query *dto.SearchGradeUpdateLogInput) error
	FindGradeUpdateLogOfExecutor(ctx *fiber.Ctx, executor uint, limit, offset int) error
	FindGradeUpdateLogOfTarget(ctx *fiber.Ctx, target uint, limit, offset int) error
	SearchPasswordUpdateLog(ctx *fiber.Ctx, query *dto.SearchPasswordUpdateLogInput) error
	SearchDeleteLog(ctx *fiber.Ctx, query *dto.SearchDeleteLogInput) error
	FindAllDeleteLog(ctx *fiber.Ctx, limit, offset int) error
	FindAllPasswordUpdateLog(ctx *fiber.Ctx, limit, offset int) error
	FindPasswordUpdateLogOfExecutor(ctx *fiber.Ctx, executor uint, limit, offset int) error
	FindPasswordUpdateLogOfTarget(ctx *fiber.Ctx, target uint, limit, offset int) error
	FindDeleteLogOfExecutor(ctx *fiber.Ctx, executor uint, limit, offset int) error
	FindDeleteLogOfTarget(ctx *fiber.Ctx, target uint, limit, offset int) error
	DeleteAdmin(ctx *fiber.Ctx, adminID uint) error
	BlockAdmin(ctx *fiber.Ctx, adminID uint) error
	UpdatePassword(ctx *fiber.Ctx, adminID uint, prevPassword, currentPassword string) error
	UpdateEmail(ctx *fiber.Ctx, adminID uint, newEmail string) error
	UpdateGrade(ctx *fiber.Ctx, adminID uint, grade admin.GradeType) error
}

func NewRepository(db *gorm.DB) Repository {
	return newAdminRepository(db)
}

func NewService(repo Repository) Service {
	return newAdminService(repo)
}