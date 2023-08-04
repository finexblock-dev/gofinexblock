package router

import (
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/middleware"
	"github.com/finexblock-dev/gofinexblock/pkg/admin"
	"github.com/finexblock-dev/gofinexblock/pkg/auth"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRouter(router fiber.Router, db *gorm.DB) {
	authService := auth.NewService(db)
	adminService := admin.NewService(db)

	authRouter := router.Group("/auth")

	authRouter.Post("/login", middleware.LoginMiddleware(adminService), handler.Login(authService))

	//SuperUserRouter(authRouter, adminService).Post("/register", middleware.AdminApiLogMiddleware(adminService), handler.Register(authService))
	authRouter.Post("/register", handler.Register(authService))
}