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
	announcementHandler := handler.NewAnnouncementHandler(announcementService)

	base := router.Group("/announcement", middleware.BearerTokenMiddleware(), middleware.AdminApiLogMiddleware(adminService))
	support := SupportRouter(base, adminService)
	maintainer := MaintainerRouter(base, adminService)

	support.Get("/", announcementHandler.FindAnnouncementByID())
	support.Get("/all", announcementHandler.FindAllAnnouncement())
	support.Get("/search", announcementHandler.SearchAnnouncement())
	support.Get("/category", announcementHandler.FindAllCategory())

	maintainer.Post("/", announcementHandler.CreateAnnouncement())
	maintainer.Patch("/", announcementHandler.UpdateAnnouncement())
	maintainer.Delete("/", announcementHandler.DeleteAnnouncement())
	maintainer.Post("/category", announcementHandler.CreateCategory())
	maintainer.Patch("/category", announcementHandler.UpdateCategory())
	maintainer.Delete("/category", announcementHandler.DeleteCategory())
}
