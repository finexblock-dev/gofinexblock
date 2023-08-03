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

	announcementRouter := router.Group("/announcement", middleware.BearerTokenMiddleware(), middleware.AdminApiLogMiddleware(adminService))

	SupportRouter(announcementRouter, adminService).Get("/", handler.FindAnnouncementByID(announcementService))
	SupportRouter(announcementRouter, adminService).Get("/all", handler.FindAllAnnouncement(announcementService))
	SupportRouter(announcementRouter, adminService).Get("/search", handler.SearchAnnouncement(announcementService))
	SupportRouter(announcementRouter, adminService).Get("/category", handler.FindAllCategory(announcementService))

	MaintainerRouter(announcementRouter, adminService).Post("/", handler.CreateAnnouncement(announcementService))

	MaintainerRouter(announcementRouter, adminService).Patch("/", handler.UpdateAnnouncement(announcementService))

	MaintainerRouter(announcementRouter, adminService).Delete("/", handler.DeleteAnnouncement(announcementService))

	MaintainerRouter(announcementRouter, adminService).Post("/category", handler.CreateCategory(announcementService))
	MaintainerRouter(announcementRouter, adminService).Patch("/category", handler.UpdateCategory(announcementService))
	MaintainerRouter(announcementRouter, adminService).Delete("/category", handler.DeleteCategory(announcementService))
}