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

	support := SupportRouter(adminRouter, adminService)
	maintainer := MaintainerRouter(adminRouter, adminService)
	superUser := SuperUserRouter(adminRouter, adminService)

	// Apis for allowed to all admin
	support.Get("/", handler.FindAllAdmin(adminService))
	support.Get("/online", handler.FindOnlineAdmin(adminService))
	support.Get("/grade", handler.FindAdminByGrade(adminService))

	support.Get("/log/failed", handler.FindLoginFailedLog(adminService))
	support.Get("/log/login", handler.FindLoginHistory(adminService))

	support.Get("/log/api/search", handler.SearchApiLog(adminService))
	support.Get("/log/grade/search", handler.SearchGradeUpdateLog(adminService))
	support.Get("/log/password/search", handler.SearchPasswordUpdateLog(adminService))
	support.Get("/log/delete/search", handler.SearchDeleteLog(adminService))

	// Apis for allowed to maintainers
	maintainer.Patch("/block", handler.BlockAdmin(adminService))
	maintainer.Patch("/unblock", handler.UnblockAdmin(adminService))
	maintainer.Patch("/password", handler.UpdatePassword(adminService))
	maintainer.Patch("/email", handler.UpdateEmail(adminService))
	maintainer.Patch("/grade", middleware.AdminGradeUpdateLogMiddleware(adminService), handler.UpdateGrade(adminService))

	// Apis for allowed to superusers
	superUser.Delete("/", middleware.AdminDeleteLogMiddleware(adminService), handler.DeleteAdmin(adminService))
}
