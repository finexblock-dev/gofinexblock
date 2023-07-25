package admin

import (
	"context"
	"database/sql"
	"errors"
	"github.com/finexblock-dev/gofinexblock/finexblock/admin/structs"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity"
	"github.com/finexblock-dev/gofinexblock/finexblock/utils"
	"gorm.io/gorm"
	"time"
)

type adminService struct {
	repo Repository
}

func (a *adminService) FindAdminByEmail(email string) (result *entity.Admin, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindAdminByEmail(tx, email)
		return err
	}, &sql.TxOptions{ReadOnly: true}); err != nil {
		return nil, err
	}
	return result, nil
}

func (a *adminService) InsertDeleteLog(executor, target uint) (result *entity.AdminDeleteLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.InsertDeleteLog(tx, executor, target)
		return err
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminService) InsertApiLog(adminID uint, method entity.ApiMethod, ip, endpoint string) (result *entity.AdminApiLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.InsertApiLog(tx, &entity.AdminApiLog{
			AdminID:  adminID,
			IP:       ip,
			Endpoint: endpoint,
			Method:   method,
		})
		return err
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminService) InsertLoginHistory(adminID uint) (result *entity.AdminLoginHistory, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.InsertLoginHistory(tx, adminID)
		return err
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminService) InsertGradeUpdateLog(executor, target uint, prev, curr entity.GradeType) (result *entity.AdminGradeUpdateLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.InsertGradeUpdateLog(tx, executor, target, prev, curr)
		return err
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminService) InsertAccessToken(adminID uint, expiredAt time.Time) (result *entity.AdminAccessToken, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.InsertAccessToken(tx, adminID, expiredAt)
		return err
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminService) InsertLoginFailedLog(adminID uint) (result *entity.AdminLoginFailedLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.InsertLoginFailedLog(tx, adminID)
		return err
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (a *adminService) UnblockAdmin(adminID uint) (err error) {
	return a.Conn().Transaction(func(tx *gorm.DB) error {
		var _admin = new(entity.Admin)

		_admin, err = a.repo.FindAdminByID(tx, adminID)
		if err != nil {
			return err
		}

		_admin.IsBlocked = false
		return a.repo.UpdateAdminByID(tx, _admin.ID, _admin)
	})
}

func (a *adminService) FindAdminByID(adminID uint) (result *entity.Admin, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindAdminByID(tx, adminID)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}
	return result, err
}

func (a *adminService) FindLoginFailedLogOfAdmin(adminID uint, limit, offset int) (result []*entity.AdminLoginFailedLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindLoginFailedLogByAdminID(tx, adminID, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindLoginHistoryOfAdmin(adminID uint, limit, offset int) (result []*entity.AdminLoginHistory, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindLoginHistoryByAdminID(tx, adminID, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) SearchApiLog(query *structs.SearchApiLogInput) (result []*entity.AdminApiLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.SearchApiLog(tx, query)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindAllApiLog(limit, offset int) (result []*entity.AdminApiLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindAllApiLog(tx, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindApiLogByAdmin(adminID uint, limit, offset int) (result []*entity.AdminApiLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindApiLogByAdminID(tx, adminID, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindApiLogByTimeCond(start, end time.Time, limit, offset int) (result []*entity.AdminApiLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindApiLogByTimeCond(tx, start, end, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindApiLogByMethodCond(method entity.ApiMethod, limit, offset int) (result []*entity.AdminApiLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindApiLogByMethodCond(tx, method, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindApiLogByEndpoint(endpoint string, limit, offset int) (result []*entity.AdminApiLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindApiLogByEndpoint(tx, endpoint, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindAllGradeUpdateLog(limit, offset int) (result []*entity.AdminGradeUpdateLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindAllGradeUpdateLog(tx, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) SearchGradeUpdateLog(query *structs.SearchGradeUpdateLogInput) (result []*entity.AdminGradeUpdateLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.SearchGradeUpdateLog(tx, query)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindGradeUpdateLogOfExecutor(executor uint, limit, offset int) (result []*entity.AdminGradeUpdateLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindGradeUpdateLogOfExecutor(tx, executor, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindGradeUpdateLogOfTarget(target uint, limit, offset int) (result []*entity.AdminGradeUpdateLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindGradeUpdateLogOfTarget(tx, target, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) SearchPasswordUpdateLog(query *structs.SearchPasswordUpdateLogInput) (result []*entity.AdminPasswordLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.SearchPasswordUpdateLog(tx, query)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindAllPasswordUpdateLog(limit, offset int) (result []*entity.AdminPasswordLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindAllPasswordUpdateLog(tx, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindPasswordUpdateLogOfExecutor(executor uint, limit, offset int) (result []*entity.AdminPasswordLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindPasswordUpdateLogOfExecutor(tx, executor, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindPasswordUpdateLogOfTarget(target uint, limit, offset int) (result []*entity.AdminPasswordLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindPasswordUpdateLogOfTarget(tx, target, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) SearchDeleteLog(query *structs.SearchDeleteLogInput) (result []*entity.AdminDeleteLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.SearchDeleteLog(tx, query)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindAllDeleteLog(limit, offset int) (result []*entity.AdminDeleteLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindAllDeleteLog(tx, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindDeleteLogOfExecutor(executor uint, limit, offset int) (result []*entity.AdminDeleteLog, err error) {
	if err = a.Conn().Transaction(func(tx *gorm.DB) error {
		result, err = a.repo.FindDeleteLogOfExecutor(tx, executor, limit, offset)
		return err
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted, ReadOnly: true}); err != nil {
		return nil, err
	}

	return result, err
}

func (a *adminService) FindDeleteLogOfTarget(target uint, limit, offset int) (result *entity.AdminDeleteLog, err error) {
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
		var _admin *entity.Admin

		_admin, err = a.repo.FindAdminByID(tx, adminID)
		if err != nil {
			return err
		}

		if _admin.Grade == entity.SUPERUSER {
			return errors.New("super admin can't be blocked")
		}

		_admin.IsBlocked = true

		return a.repo.UpdateAdminByID(tx, adminID, _admin)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (a *adminService) UpdatePassword(adminID uint, prevPassword, currentPassword string) (err error) {
	return a.Conn().Transaction(func(tx *gorm.DB) error {
		var _adminCredentials *entity.Admin

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

		return a.repo.UpdateAdminByID(tx, adminID, &entity.Admin{Password: currentPassword})
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (a *adminService) UpdateEmail(adminID uint, newEmail string) (err error) {
	return a.Conn().Transaction(func(tx *gorm.DB) error {
		var _admin = &entity.Admin{Email: newEmail}

		return a.repo.UpdateAdminByID(tx, adminID, _admin)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (a *adminService) UpdateGrade(adminID uint, grade entity.GradeType) (err error) {
	return a.Conn().Transaction(func(tx *gorm.DB) error {
		var _admin = &entity.Admin{Grade: grade}

		return a.repo.UpdateAdminByID(tx, adminID, _admin)
	}, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
}

func (a *adminService) FindAdminByGrade(grade entity.GradeType, limit, offset int) (result []*entity.Admin, err error) {
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

func (a *adminService) FindAllAdmin(limit, offset int) ([]*entity.Admin, error) {
	var result []*entity.Admin
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