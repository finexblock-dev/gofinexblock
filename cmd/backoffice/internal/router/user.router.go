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
	userHandler := handler.NewUserHandler(userService)

	base := router.Group("/user", middleware.BearerTokenMiddleware(), middleware.AdminApiLogMiddleware(adminService))
	support := SupportRouter(base, adminService)
	maintainer := MaintainerRouter(base, adminService)

	support.Get("/", userHandler.FindUserByID())
	support.Get("/search", userHandler.SearchUser())

	maintainer.Patch("/block", userHandler.BlockUser())
	maintainer.Patch("/unblock", userHandler.UnblockUser())
	maintainer.Post("/memo", userHandler.CreateMemo())
}
