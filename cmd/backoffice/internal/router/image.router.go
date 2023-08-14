package router

import (
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/middleware"
	"github.com/finexblock-dev/gofinexblock/pkg/admin"
	"github.com/finexblock-dev/gofinexblock/pkg/image"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"os"
)

func ImageRouter(router fiber.Router, db *gorm.DB) {
	imageService := image.NewService(db, os.Getenv("AWS_BUCKET"), os.Getenv("AWS_BASE_PATH"))
	adminService := admin.NewService(db)
	imageHandler := handler.NewImageHandler(imageService)

	base := router.Group("/image", middleware.BearerTokenMiddleware(), middleware.AdminApiLogMiddleware(adminService))
	support := SupportRouter(base, adminService)
	maintainer := MaintainerRouter(base, adminService)

	support.Get("/", imageHandler.ListImage())
	maintainer.Post("/", imageHandler.UploadImage())
}
