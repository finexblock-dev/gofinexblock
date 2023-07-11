package admin

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/admin/dto"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	types.Repository
	FindAdminByID(tx *gorm.DB, id uint) (*admin.Admin, error)
	FindAdminByEmail(tx *gorm.DB, email string) (*admin.Admin, error)
	FindAdminCredentialsByID(tx *gorm.DB, id uint) (*admin.Admin, error)
	FindAdminByGrade(tx *gorm.DB, grade admin.GradeType, limit, offset int) ([]*admin.Admin, error)
	FindAllAdmin(tx *gorm.DB, limit, offset int) ([]*admin.Admin, error)
	InsertAdmin(tx *gorm.DB, email, password string) (*admin.Admin, error)
	UpdateAdminByID(tx *gorm.DB, id uint, admin *admin.Admin) error
	DeleteAdminByID(tx *gorm.DB, id uint) error

	FindAccessToken(tx *gorm.DB, limit, offset int) ([]*admin.AdminAccessToken, error)
	InsertAccessToken(tx *gorm.DB, adminID uint, expiredAt time.Time) (*admin.AdminAccessToken, error)
	DeleteAccessToken(tx *gorm.DB, id uint) error

	InsertApiLog(tx *gorm.DB, log *admin.AdminApiLog) (*admin.AdminApiLog, error)
	FindAllApiLog(tx *gorm.DB, limit, offset int) ([]*admin.AdminApiLog, error)
	FindApiLogByAdminID(tx *gorm.DB, adminID uint, limit, offset int) ([]*admin.AdminApiLog, error)
	FindApiLogByTimeCond(tx *gorm.DB, start, end time.Time, limit, offset int) ([]*admin.AdminApiLog, error)
	FindApiLogByMethodCond(tx *gorm.DB, method admin.ApiMethod, limit, offset int) ([]*admin.AdminApiLog, error)
	FindApiLogByEndpoint(tx *gorm.DB, endpoint string, limit, offset int) ([]*admin.AdminApiLog, error)
	SearchApiLog(tx *gorm.DB, query *dto.SearchApiLogInput) ([]*admin.AdminApiLog, error)

	SearchDeleteLog(tx *gorm.DB, input *dto.SearchDeleteLogInput) ([]*admin.AdminDeleteLog, error)
	FindAllDeleteLog(tx *gorm.DB, limit, offset int) ([]*admin.AdminDeleteLog, error)
	FindDeleteLogOfExecutor(tx *gorm.DB, executor uint, limit, offset int) ([]*admin.AdminDeleteLog, error)
	FindDeleteLogOfTarget(tx *gorm.DB, target uint) (*admin.AdminDeleteLog, error)
	InsertDeleteLog(tx *gorm.DB, executor, target uint) (*admin.AdminDeleteLog, error)

	InsertGradeUpdateLog(tx *gorm.DB, executor, target uint, prev, curr string) (*admin.AdminGradeUpdateLog, error)
	SearchGradeUpdateLog(tx *gorm.DB, input *dto.SearchGradeUpdateLogInput) ([]*admin.AdminGradeUpdateLog, error)
	FindAllGradeUpdateLog(tx *gorm.DB, limit, offset int) ([]*admin.AdminGradeUpdateLog, error)
	FindGradeUpdateLogOfExecutor(tx *gorm.DB, executor uint, limit, offset int) ([]*admin.AdminGradeUpdateLog, error)
	FindGradeUpdateLogOfTarget(tx *gorm.DB, target uint, limit, offset int) ([]*admin.AdminGradeUpdateLog, error)

	FindLoginFailedLogByAdminID(tx *gorm.DB, adminID uint, limit, offset int) ([]*admin.AdminLoginFailedLog, error)
	InsertLoginFailedLog(tx *gorm.DB, adminID uint) (*admin.AdminLoginFailedLog, error)

	FindLoginHistoryByAdminID(tx *gorm.DB, adminID uint, limit, offset int) ([]*admin.AdminLoginHistory, error)
	InsertLoginHistory(tx *gorm.DB, adminID uint) (*admin.AdminLoginHistory, error)

	SearchPasswordUpdateLog(tx *gorm.DB, input *dto.SearchPasswordUpdateLogInput) ([]*admin.AdminPasswordLog, error)
	FindAllPasswordUpdateLog(tx *gorm.DB, limit, offset int) ([]*admin.AdminPasswordLog, error)
	FindPasswordUpdateLogOfExecutor(tx *gorm.DB, executor uint, limit, offset int) ([]*admin.AdminPasswordLog, error)
	FindPasswordUpdateLogOfTarget(tx *gorm.DB, target uint, limit, offset int) ([]*admin.AdminPasswordLog, error)
	InsertPasswordUpdateLog(tx *gorm.DB, executor, target uint) (*admin.AdminPasswordLog, error)
}

type Service interface {
	types.Service
	FindAllAdmin(limit, offset int) (result []*admin.Admin, err error)
	FindAdminByGrade(grade admin.GradeType, limit, offset int) (result []*admin.Admin, err error)

	FindLoginFailedLogOfAdmin(adminID uint, limit, offset int) (result []*admin.AdminLoginFailedLog, err error)

	FindLoginHistoryOfAdmin(adminID uint, limit, offset int) (result []*admin.AdminLoginHistory, err error)

	SearchApiLog(query *dto.SearchApiLogInput) (result []*admin.AdminApiLog, err error)
	FindAllApiLog(limit, offset int) (result []*admin.AdminApiLog, err error)
	FindApiLogByAdmin(adminID uint, limit, offset int) (result []*admin.AdminApiLog, err error)
	FindApiLogByTimeCond(start, end time.Time, limit, offset int) (result []*admin.AdminApiLog, err error)
	FindApiLogByMethodCond(method admin.ApiMethod, limit, offset int) (result []*admin.AdminApiLog, err error)
	FindApiLogByEndpoint(endpoint string, limit, offset int) (result []*admin.AdminApiLog, err error)

	FindAllGradeUpdateLog(limit, offset int) (result []*admin.AdminGradeUpdateLog, err error)
	SearchGradeUpdateLog(query *dto.SearchGradeUpdateLogInput) (result []*admin.AdminGradeUpdateLog, err error)
	FindGradeUpdateLogOfExecutor(executor uint, limit, offset int) (result []*admin.AdminGradeUpdateLog, err error)
	FindGradeUpdateLogOfTarget(target uint, limit, offset int) (result []*admin.AdminGradeUpdateLog, err error)

	SearchPasswordUpdateLog(query *dto.SearchPasswordUpdateLogInput) (result []*admin.AdminPasswordLog, err error)
	FindAllPasswordUpdateLog(limit, offset int) (result []*admin.AdminPasswordLog, err error)
	FindPasswordUpdateLogOfExecutor(executor uint, limit, offset int) (result []*admin.AdminPasswordLog, err error)
	FindPasswordUpdateLogOfTarget(target uint, limit, offset int) (result []*admin.AdminPasswordLog, err error)

	SearchDeleteLog(query *dto.SearchDeleteLogInput) (result []*admin.AdminDeleteLog, err error)
	FindAllDeleteLog(limit, offset int) (result []*admin.AdminDeleteLog, err error)
	FindDeleteLogOfExecutor(executor uint, limit, offset int) (result []*admin.AdminDeleteLog, err error)
	FindDeleteLogOfTarget(target uint, limit, offset int) (result *admin.AdminDeleteLog, err error)

	DeleteAdmin(adminID uint) (err error)
	BlockAdmin(adminID uint) (err error)
	UpdatePassword(adminID uint, prevPassword, currentPassword string) (err error)
	UpdateEmail(adminID uint, newEmail string) (err error)
	UpdateGrade(adminID uint, grade admin.GradeType) (err error)
}

func NewRepository(db *gorm.DB) Repository {
	return newAdminRepository(db)
}

func NewService(db *gorm.DB) Service {
	return newAdminService(NewRepository(db))
}
