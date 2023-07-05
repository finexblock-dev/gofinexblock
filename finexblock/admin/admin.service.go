package admin

import (
	"context"
	"database/sql"
	"github.com/finexblock-dev/gofinexblock/finexblock/admin/dto"
	"github.com/finexblock-dev/gofinexblock/finexblock/entity/admin"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"time"
)

type adminService struct {
	repo Repository
}

func (a *adminService) Conn() *gorm.DB {
	return a.repo.Conn()
}

func (a *adminService) Tx(level sql.IsolationLevel) *gorm.DB {
	return a.repo.Tx(level)
}

func newAdminService(repo Repository) *adminService {
	return &adminService{repo: repo}
}

func (a *adminService) FindAllAdmin(ctx *fiber.Ctx, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) FindAdminByGrade(ctx *fiber.Ctx, grade admin.GradeType, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) FindLoginFailedLog(ctx *fiber.Ctx, adminID uint, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) FindLoginHistory(ctx *fiber.Ctx, adminID uint, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) SearchApiLog(ctx *fiber.Ctx, query *dto.SearchApiLogInput) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) FindAllApiLog(ctx *fiber.Ctx, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) FindApiLogByAdmin(ctx *fiber.Ctx, adminID uint, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) FindApiLogByTimeCond(ctx *fiber.Ctx, start, end time.Time, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) FindApiLogByMethodCond(ctx *fiber.Ctx, method admin.ApiMethod, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) FindApiLogByEndpoint(ctx *fiber.Ctx, endpoint string, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) FindAllGradeUpdateLog(ctx *fiber.Ctx, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) SearchGradeUpdateLog(ctx *fiber.Ctx, query *dto.SearchGradeUpdateLogInput) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) FindGradeUpdateLogOfExecutor(ctx *fiber.Ctx, executor uint, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) FindGradeUpdateLogOfTarget(ctx *fiber.Ctx, target uint, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) SearchPasswordUpdateLog(ctx *fiber.Ctx, query *dto.SearchPasswordUpdateLogInput) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) SearchDeleteLog(ctx *fiber.Ctx, query *dto.SearchDeleteLogInput) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) FindAllDeleteLog(ctx *fiber.Ctx, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) FindAllPasswordUpdateLog(ctx *fiber.Ctx, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) FindPasswordUpdateLogOfExecutor(ctx *fiber.Ctx, executor uint, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) FindPasswordUpdateLogOfTarget(ctx *fiber.Ctx, target uint, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) FindDeleteLogOfExecutor(ctx *fiber.Ctx, executor uint, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) FindDeleteLogOfTarget(ctx *fiber.Ctx, target uint, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) DeleteAdmin(ctx *fiber.Ctx, adminID uint) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) BlockAdmin(ctx *fiber.Ctx, adminID uint) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) UpdatePassword(ctx *fiber.Ctx, adminID uint, prevPassword, currentPassword string) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) UpdateEmail(ctx *fiber.Ctx, adminID uint, newEmail string) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) UpdateGrade(ctx *fiber.Ctx, adminID uint, grade admin.GradeType) error {
	//TODO implement me
	panic("implement me")
}

func (a *adminService) Ctx() context.Context {
	return context.Background()
}

func (a *adminService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}