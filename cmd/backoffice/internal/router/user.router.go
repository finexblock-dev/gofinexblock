package router

import (
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/middleware"
	"github.com/finexblock-dev/gofinexblock/pkg/admin"
	"github.com/finexblock-dev/gofinexblock/pkg/user"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func UserRouter(router fiber.Router, db *gorm.DB, cluster *redis.ClusterClient) {
	userService := user.NewService(db, cluster)
	adminService := admin.NewService(db)

	userRouter := router.Group("/user", middleware.BearerTokenMiddleware(), middleware.AdminApiLogMiddleware(adminService))

	SupportRouter(userRouter, adminService).Get("/", handler.FindUserByID(userService))
	SupportRouter(userRouter, adminService).Get("/search", handler.SearchUser(userService))

	MaintainerRouter(userRouter, adminService).Patch("/block", handler.BlockUser(userService))
	MaintainerRouter(userRouter, adminService).Patch("/unblock", handler.UnblockUser(userService))
	MaintainerRouter(userRouter, adminService).Post("/memo", handler.CreateMemo(userService))
}