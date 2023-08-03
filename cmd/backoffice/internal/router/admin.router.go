package router

import (
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/middleware"
	"github.com/finexblock-dev/gofinexblock/pkg/admin"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AdminRouter(router fiber.Router, db *gorm.DB) {

	adminService := admin.NewService(db)

	adminRouter := router.Group("/admin", middleware.BearerTokenMiddleware(), middleware.AdminApiLogMiddleware(adminService))

	// Apis for allowed to all admin
	SupportRouter(adminRouter, adminService).Get("/", handler.FindAllAdmin(adminService))
	SupportRouter(adminRouter, adminService).Get("/online", handler.FindOnlineAdmin(adminService))
	SupportRouter(adminRouter, adminService).Get("/grade", handler.FindAdminByGrade(adminService))

	SupportRouter(adminRouter, adminService).Get("/log/failed", handler.FindLoginFailedLog(adminService))
	SupportRouter(adminRouter, adminService).Get("/log/login", handler.FindLoginHistory(adminService))

	SupportRouter(adminRouter, adminService).Get("/log/api/search", handler.SearchApiLog(adminService))
	SupportRouter(adminRouter, adminService).Get("/log/grade/search", handler.SearchGradeUpdateLog(adminService))
	SupportRouter(adminRouter, adminService).Get("/log/password/search", handler.SearchPasswordUpdateLog(adminService))
	SupportRouter(adminRouter, adminService).Get("/log/delete/search", handler.SearchDeleteLog(adminService))

	// Apis for allowed to maintainers
	MaintainerRouter(adminRouter, adminService).Patch("/block", handler.BlockAdmin(adminService))
	MaintainerRouter(adminRouter, adminService).Patch("/unblock", handler.UnblockAdmin(adminService))
	MaintainerRouter(adminRouter, adminService).Patch("/password", handler.UpdatePassword(adminService))
	MaintainerRouter(adminRouter, adminService).Patch("/email", handler.UpdateEmail(adminService))
	MaintainerRouter(adminRouter, adminService).Patch("/grade", middleware.AdminGradeUpdateLogMiddleware(adminService), handler.UpdateGrade(adminService))

	// Apis for allowed to superusers
	SuperUserRouter(adminRouter, adminService).Delete("/", middleware.AdminDeleteLogMiddleware(adminService), handler.DeleteAdmin(adminService))
}