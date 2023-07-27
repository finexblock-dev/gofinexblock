package admin

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/admin/structs"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"gorm.io/gorm"
	"time"
)

type Repository interface {
	types.Repository
	FindAdminByID(tx *gorm.DB, id uint) (result *entity.Admin, err error)
	FindManyAdminByID(tx *gorm.DB, ids []uint) (result []*entity.Admin, err error)
	FindAdminByEmail(tx *gorm.DB, email string) (result *entity.Admin, err error)
	FindAdminCredentialsByID(tx *gorm.DB, id uint) (result *entity.Admin, err error)
	FindAdminByGrade(tx *gorm.DB, grade entity.GradeType, limit, offset int) (result []*entity.Admin, err error)
	FindAllAdmin(tx *gorm.DB, limit, offset int) (result []*entity.Admin, err error)
	InsertAdmin(tx *gorm.DB, email, password string) (result *entity.Admin, err error)
	UpdateAdminByID(tx *gorm.DB, id uint, admin *entity.Admin) error
	DeleteAdminByID(tx *gorm.DB, id uint) error

	BlockAdminByID(tx *gorm.DB, id uint) error
	UnblockAdminByID(tx *gorm.DB, id uint) error

	FindAccessToken(tx *gorm.DB, limit, offset int) ([]*entity.AdminAccessToken, error)
	InsertAccessToken(tx *gorm.DB, adminID uint, expiredAt time.Time) (*entity.AdminAccessToken, error)
	DeleteAccessToken(tx *gorm.DB, id uint) error

	InsertApiLog(tx *gorm.DB, log *entity.AdminApiLog) (*entity.AdminApiLog, error)
	FindAllApiLog(tx *gorm.DB, limit, offset int) ([]*entity.AdminApiLog, error)
	FindApiLogByAdminID(tx *gorm.DB, adminID uint, limit, offset int) ([]*entity.AdminApiLog, error)
	FindApiLogByTimeCond(tx *gorm.DB, start, end time.Time, limit, offset int) ([]*entity.AdminApiLog, error)
	FindApiLogByMethodCond(tx *gorm.DB, method entity.ApiMethod, limit, offset int) ([]*entity.AdminApiLog, error)
	FindApiLogByEndpoint(tx *gorm.DB, endpoint string, limit, offset int) ([]*entity.AdminApiLog, error)
	SearchApiLog(tx *gorm.DB, query *structs.SearchApiLogInput) ([]*entity.AdminApiLog, error)

	SearchDeleteLog(tx *gorm.DB, input *structs.SearchDeleteLogInput) ([]*entity.AdminDeleteLog, error)
	FindAllDeleteLog(tx *gorm.DB, limit, offset int) ([]*entity.AdminDeleteLog, error)
	FindDeleteLogOfExecutor(tx *gorm.DB, executor uint, limit, offset int) ([]*entity.AdminDeleteLog, error)
	FindDeleteLogOfTarget(tx *gorm.DB, target uint) (*entity.AdminDeleteLog, error)
	InsertDeleteLog(tx *gorm.DB, executor, target uint) (*entity.AdminDeleteLog, error)

	InsertGradeUpdateLog(tx *gorm.DB, executor, target uint, prev, curr entity.GradeType) (*entity.AdminGradeUpdateLog, error)
	SearchGradeUpdateLog(tx *gorm.DB, input *structs.SearchGradeUpdateLogInput) ([]*entity.AdminGradeUpdateLog, error)
	FindAllGradeUpdateLog(tx *gorm.DB, limit, offset int) ([]*entity.AdminGradeUpdateLog, error)
	FindGradeUpdateLogOfExecutor(tx *gorm.DB, executor uint, limit, offset int) ([]*entity.AdminGradeUpdateLog, error)
	FindGradeUpdateLogOfTarget(tx *gorm.DB, target uint, limit, offset int) ([]*entity.AdminGradeUpdateLog, error)

	FindLoginFailedLogByAdminID(tx *gorm.DB, adminID uint, limit, offset int) ([]*entity.AdminLoginFailedLog, error)
	InsertLoginFailedLog(tx *gorm.DB, adminID uint) (*entity.AdminLoginFailedLog, error)

	FindLoginHistoryByAdminID(tx *gorm.DB, adminID uint, limit, offset int) ([]*entity.AdminLoginHistory, error)
	InsertLoginHistory(tx *gorm.DB, adminID uint) (*entity.AdminLoginHistory, error)

	SearchPasswordUpdateLog(tx *gorm.DB, input *structs.SearchPasswordUpdateLogInput) ([]*entity.AdminPasswordLog, error)
	FindAllPasswordUpdateLog(tx *gorm.DB, limit, offset int) ([]*entity.AdminPasswordLog, error)
	FindPasswordUpdateLogOfExecutor(tx *gorm.DB, executor uint, limit, offset int) ([]*entity.AdminPasswordLog, error)
	FindPasswordUpdateLogOfTarget(tx *gorm.DB, target uint, limit, offset int) ([]*entity.AdminPasswordLog, error)
	InsertPasswordUpdateLog(tx *gorm.DB, executor, target uint) (*entity.AdminPasswordLog, error)
}

type Service interface {
	types.Service
	FindOnlineAdmin(limit, offset int) (result []*entity.Admin, err error)

	FindAdminByID(adminID uint) (result *entity.Admin, err error)
	FindAllAdmin(limit, offset int) (result []*entity.Admin, err error)
	FindAdminByGrade(grade entity.GradeType, limit, offset int) (result []*entity.Admin, err error)

	FindLoginFailedLogOfAdmin(adminID uint, limit, offset int) (result []*entity.AdminLoginFailedLog, err error)

	FindLoginHistoryOfAdmin(adminID uint, limit, offset int) (result []*entity.AdminLoginHistory, err error)
	SearchApiLog(query *structs.SearchApiLogInput) (result []*entity.AdminApiLog, err error)
	FindAllApiLog(limit, offset int) (result []*entity.AdminApiLog, err error)
	FindApiLogByAdmin(adminID uint, limit, offset int) (result []*entity.AdminApiLog, err error)
	FindApiLogByTimeCond(start, end time.Time, limit, offset int) (result []*entity.AdminApiLog, err error)
	FindApiLogByMethodCond(method entity.ApiMethod, limit, offset int) (result []*entity.AdminApiLog, err error)
	FindApiLogByEndpoint(endpoint string, limit, offset int) (result []*entity.AdminApiLog, err error)

	FindAllGradeUpdateLog(limit, offset int) (result []*entity.AdminGradeUpdateLog, err error)
	SearchGradeUpdateLog(query *structs.SearchGradeUpdateLogInput) (result []*entity.AdminGradeUpdateLog, err error)
	FindGradeUpdateLogOfExecutor(executor uint, limit, offset int) (result []*entity.AdminGradeUpdateLog, err error)
	FindGradeUpdateLogOfTarget(target uint, limit, offset int) (result []*entity.AdminGradeUpdateLog, err error)

	SearchPasswordUpdateLog(query *structs.SearchPasswordUpdateLogInput) (result []*entity.AdminPasswordLog, err error)
	FindAllPasswordUpdateLog(limit, offset int) (result []*entity.AdminPasswordLog, err error)
	FindPasswordUpdateLogOfExecutor(executor uint, limit, offset int) (result []*entity.AdminPasswordLog, err error)
	FindPasswordUpdateLogOfTarget(target uint, limit, offset int) (result []*entity.AdminPasswordLog, err error)

	SearchDeleteLog(query *structs.SearchDeleteLogInput) (result []*entity.AdminDeleteLog, err error)
	FindAllDeleteLog(limit, offset int) (result []*entity.AdminDeleteLog, err error)
	FindDeleteLogOfExecutor(executor uint, limit, offset int) (result []*entity.AdminDeleteLog, err error)
	FindDeleteLogOfTarget(target uint, limit, offset int) (result *entity.AdminDeleteLog, err error)

	FindAdminByEmail(email string) (result *entity.Admin, err error)
	InsertDeleteLog(executor, target uint) (result *entity.AdminDeleteLog, err error)
	InsertApiLog(adminID uint, method entity.ApiMethod, ip, endpoint string) (result *entity.AdminApiLog, err error)
	InsertLoginHistory(adminID uint) (result *entity.AdminLoginHistory, err error)
	InsertGradeUpdateLog(executor, target uint, prev, curr entity.GradeType) (result *entity.AdminGradeUpdateLog, err error)
	InsertAccessToken(adminID uint, expiredAt time.Time) (result *entity.AdminAccessToken, err error)
	InsertLoginFailedLog(adminID uint) (result *entity.AdminLoginFailedLog, err error)
	UnblockAdmin(adminID uint) (err error)
	BlockAdmin(adminID uint) (err error)

	DeleteAdmin(adminID uint) (err error)
	UpdatePassword(adminID uint, prevPassword, currentPassword string) (err error)
	UpdateEmail(adminID uint, newEmail string) (err error)
	UpdateGrade(adminID uint, grade entity.GradeType) (err error)
}

func NewRepository(db *gorm.DB) Repository {
	return newAdminRepository(db)
}

func NewService(db *gorm.DB) Service {
	return newAdminService(NewRepository(db))
}