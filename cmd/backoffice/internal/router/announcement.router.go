package router

import (
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/middleware"
	"github.com/finexblock-dev/gofinexblock/pkg/admin"
	"github.com/finexblock-dev/gofinexblock/pkg/announcement"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AnnouncementRouter(router fiber.Router, db *gorm.DB) {

	announcementService := announcement.NewService(db)
	adminService := admin.NewService(db)

	base := router.Group("/announcement", middleware.BearerTokenMiddleware(), middleware.AdminApiLogMiddleware(adminService))
	support := SupportRouter(base, adminService)
	maintainer := MaintainerRouter(base, adminService)

	support.Get("/", handler.FindAnnouncementByID(announcementService))
	support.Get("/all", handler.FindAllAnnouncement(announcementService))
	support.Get("/search", handler.SearchAnnouncement(announcementService))
	support.Get("/category", handler.FindAllCategory(announcementService))

	maintainer.Post("/", handler.CreateAnnouncement(announcementService))
	maintainer.Patch("/", handler.UpdateAnnouncement(announcementService))
	maintainer.Delete("/", handler.DeleteAnnouncement(announcementService))
	maintainer.Post("/category", handler.CreateCategory(announcementService))
	maintainer.Patch("/category", handler.UpdateCategory(announcementService))
	maintainer.Delete("/category", handler.DeleteCategory(announcementService))
}
