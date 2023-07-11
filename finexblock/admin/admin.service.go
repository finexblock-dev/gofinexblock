package admin

import (
	"context"
	"database/sql"
	"errors"
	"github.com/finexblock-dev/gofinexblock/finexblock/admin/dto"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"
	"github.com/finexblock-dev/gofinexblock/finexblock/utils"
	"gorm.io/gorm"
	"time"
)

type adminService struct {
	repo Repository
}

func (a *adminService) FindLoginFailedLogOfAdmin(adminID uint, limit, offset int) (result []*admin.AdminLoginFailedLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindLoginFailedLogByAdminID(tx, adminID, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindLoginHistoryOfAdmin(adminID uint, limit, offset int) (result []*admin.AdminLoginHistory, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindLoginHistoryByAdminID(tx, adminID, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) SearchApiLog(query *dto.SearchApiLogInput) (result []*admin.AdminApiLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.SearchApiLog(tx, query)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindAllApiLog(limit, offset int) (result []*admin.AdminApiLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindAllApiLog(tx, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindApiLogByAdmin(adminID uint, limit, offset int) (result []*admin.AdminApiLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindApiLogByAdminID(tx, adminID, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindApiLogByTimeCond(start, end time.Time, limit, offset int) (result []*admin.AdminApiLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindApiLogByTimeCond(tx, start, end, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindApiLogByMethodCond(method admin.ApiMethod, limit, offset int) (result []*admin.AdminApiLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindApiLogByMethodCond(tx, method, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindApiLogByEndpoint(endpoint string, limit, offset int) (result []*admin.AdminApiLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindApiLogByEndpoint(tx, endpoint, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindAllGradeUpdateLog(limit, offset int) (result []*admin.AdminGradeUpdateLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindAllGradeUpdateLog(tx, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) SearchGradeUpdateLog(query *dto.SearchGradeUpdateLogInput) (result []*admin.AdminGradeUpdateLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.SearchGradeUpdateLog(tx, query)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindGradeUpdateLogOfExecutor(executor uint, limit, offset int) (result []*admin.AdminGradeUpdateLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindGradeUpdateLogOfExecutor(tx, executor, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindGradeUpdateLogOfTarget(target uint, limit, offset int) (result []*admin.AdminGradeUpdateLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindGradeUpdateLogOfTarget(tx, target, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) SearchPasswordUpdateLog(query *dto.SearchPasswordUpdateLogInput) (result []*admin.AdminPasswordLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.SearchPasswordUpdateLog(tx, query)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindAllPasswordUpdateLog(limit, offset int) (result []*admin.AdminPasswordLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindAllPasswordUpdateLog(tx, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindPasswordUpdateLogOfExecutor(executor uint, limit, offset int) (result []*admin.AdminPasswordLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindPasswordUpdateLogOfExecutor(tx, executor, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindPasswordUpdateLogOfTarget(target uint, limit, offset int) (result []*admin.AdminPasswordLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindPasswordUpdateLogOfTarget(tx, target, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) SearchDeleteLog(query *dto.SearchDeleteLogInput) (result []*admin.AdminDeleteLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.SearchDeleteLog(tx, query)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindAllDeleteLog(limit, offset int) (result []*admin.AdminDeleteLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindAllDeleteLog(tx, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindDeleteLogOfExecutor(executor uint, limit, offset int) (result []*admin.AdminDeleteLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindDeleteLogOfExecutor(tx, executor, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindDeleteLogOfTarget(target uint, limit, offset int) (result *admin.AdminDeleteLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindDeleteLogOfTarget(tx, target)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) DeleteAdmin(adminID uint) (err error) {
	return a.Conn().Transaction(func(tx *gorm.DB) error {
		return a.repo.DeleteAdminByID(tx, adminID)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (a *adminService) BlockAdmin(adminID uint) (err error) {
	return a.Conn().Transaction(func(tx *gorm.DB) error {
		var _admin *admin.Admin

		_admin, err = a.repo.FindAdminByID(tx, adminID)
		if err != nil {
			return err
		}

		if _admin.Grade == admin.SUPERUSER {
			return errors.New("super admin can't be blocked")
		}

		return a.repo.UpdateAdminByID(tx, adminID, &admin.Admin{IsBlocked: true})
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (a *adminService) UpdatePassword(adminID uint, prevPassword, currentPassword string) (err error) {
	return a.Conn().Transaction(func(tx *gorm.DB) error {
		var _adminCredentials *admin.Admin

		if utils.PasswordRegex(currentPassword) {
			return errors.New("regex error: password is not valid")
		}

		_adminCredentials, err = a.repo.FindAdminCredentialsByID(tx, adminID)
		if err != nil {
			return err
		}

		if !utils.CompareHash(_adminCredentials.Password, prevPassword) {
			return errors.New("invalid credentials: password is not valid")
		}

		return a.repo.UpdateAdminByID(tx, adminID, &admin.Admin{Password: currentPassword})
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (a *adminService) UpdateEmail(adminID uint, newEmail string) (err error) {
	return a.Conn().Transaction(func(tx *gorm.DB) error {
		var _admin = &admin.Admin{Email: newEmail}

		return a.repo.UpdateAdminByID(tx, adminID, _admin)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (a *adminService) UpdateGrade(adminID uint, grade admin.GradeType) (err error) {
	return a.Conn().Transaction(func(tx *gorm.DB) error {
		var _admin = &admin.Admin{Grade: grade}

		return a.repo.UpdateAdminByID(tx, adminID, _admin)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (a *adminService) FindAdminByGrade(grade admin.GradeType, limit, offset int) (result []*admin.Admin, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindAdminByGrade(tx, grade, limit, offset)
		if err != nil {
			return err
		}

		return nil
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminService) FindAllAdmin(limit, offset int) ([]*admin.Admin, error) {
	var result []*admin.Admin
	var err error
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindAllAdmin(tx, limit, offset)
		if err != nil {
			return err
		}

		return nil
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, nil
}

func newAdminService(repo Repository) *adminService {
	return &adminService{repo: repo}
}

func (a *adminService) Tx(level sql.IsolationLevel) *gorm.DB {
	return a.repo.Tx(level)
}

func (a *adminService) Conn() *gorm.DB {
	return a.repo.Conn()
}

func (a *adminService) Ctx() context.Context {
	return context.Background()
}

func (a *adminService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}