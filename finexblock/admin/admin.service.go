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

type AdminService struct {
	repo Repository
}

func (a *AdminService) Tx(level sql.IsolationLevel) *gorm.DB {
	return a.repo.Tx(level)
}

func newAdminService(repo Repository) *AdminService {
	return &AdminService{repo: repo}
}

func (a *AdminService) FindAllAdmin(ctx *fiber.Ctx, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) FindAdminByGrade(ctx *fiber.Ctx, grade admin.GradeType, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) FindLoginFailedLog(ctx *fiber.Ctx, adminID uint, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) FindLoginHistory(ctx *fiber.Ctx, adminID uint, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) SearchApiLog(ctx *fiber.Ctx, query *dto.SearchApiLogInput) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) FindAllApiLog(ctx *fiber.Ctx, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) FindApiLogByAdmin(ctx *fiber.Ctx, adminID uint, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) FindApiLogByTimeCond(ctx *fiber.Ctx, start, end time.Time, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) FindApiLogByMethodCond(ctx *fiber.Ctx, method admin.ApiMethod, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) FindApiLogByEndpoint(ctx *fiber.Ctx, endpoint string, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) FindAllGradeUpdateLog(ctx *fiber.Ctx, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) SearchGradeUpdateLog(ctx *fiber.Ctx, query *dto.SearchGradeUpdateLogInput) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) FindGradeUpdateLogOfExecutor(ctx *fiber.Ctx, executor uint, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) FindGradeUpdateLogOfTarget(ctx *fiber.Ctx, target uint, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) SearchPasswordUpdateLog(ctx *fiber.Ctx, query *dto.SearchPasswordUpdateLogInput) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) SearchDeleteLog(ctx *fiber.Ctx, query *dto.SearchDeleteLogInput) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) FindAllDeleteLog(ctx *fiber.Ctx, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) FindAllPasswordUpdateLog(ctx *fiber.Ctx, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) FindPasswordUpdateLogOfExecutor(ctx *fiber.Ctx, executor uint, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) FindPasswordUpdateLogOfTarget(ctx *fiber.Ctx, target uint, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) FindDeleteLogOfExecutor(ctx *fiber.Ctx, executor uint, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) FindDeleteLogOfTarget(ctx *fiber.Ctx, target uint, limit, offset int) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) DeleteAdmin(ctx *fiber.Ctx, adminID uint) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) BlockAdmin(ctx *fiber.Ctx, adminID uint) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) UpdatePassword(ctx *fiber.Ctx, adminID uint, prevPassword, currentPassword string) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) UpdateEmail(ctx *fiber.Ctx, adminID uint, newEmail string) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) UpdateGrade(ctx *fiber.Ctx, adminID uint, grade admin.GradeType) error {
	//TODO implement me
	panic("implement me")
}

func (a *AdminService) Ctx() context.Context {
	return context.Background()
}

func (a *AdminService) CtxWithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(ctx)
}