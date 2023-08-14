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
	adminHandler := handler.NewAdminHandler(adminService)

	adminRouter := router.Group("/admin", middleware.BearerTokenMiddleware(), middleware.AdminApiLogMiddleware(adminService))

	support := SupportRouter(adminRouter, adminService)
	maintainer := MaintainerRouter(adminRouter, adminService)
	superUser := SuperUserRouter(adminRouter, adminService)

	// Apis for allowed to all admin
	support.Get("/", adminHandler.FindAllAdmin())
	support.Get("/online", adminHandler.FindOnlineAdmin())
	support.Get("/grade", adminHandler.FindAdminByGrade())

	support.Get("/log/failed", adminHandler.FindLoginFailedLog())
	support.Get("/log/login", adminHandler.FindLoginHistory())

	support.Get("/log/api/search", adminHandler.SearchApiLog())
	support.Get("/log/grade/search", adminHandler.SearchGradeUpdateLog())
	support.Get("/log/password/search", adminHandler.SearchPasswordUpdateLog())
	support.Get("/log/delete/search", adminHandler.SearchDeleteLog())

	// Apis for allowed to maintainers
	maintainer.Patch("/block", adminHandler.BlockAdmin())
	maintainer.Patch("/unblock", adminHandler.UnblockAdmin())
	maintainer.Patch("/password", adminHandler.UpdatePassword())
	maintainer.Patch("/email", adminHandler.UpdateEmail())
	maintainer.Patch("/grade", middleware.AdminGradeUpdateLogMiddleware(adminService), adminHandler.UpdateGrade())

	// Apis for allowed to superusers
	superUser.Delete("/", middleware.AdminDeleteLogMiddleware(adminService), adminHandler.DeleteAdmin())
}
