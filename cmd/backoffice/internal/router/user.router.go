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

	base := router.Group("/user", middleware.BearerTokenMiddleware(), middleware.AdminApiLogMiddleware(adminService))
	support := SupportRouter(base, adminService)
	maintainer := MaintainerRouter(base, adminService)

	support.Get("/", handler.FindUserByID(userService))
	support.Get("/search", handler.SearchUser(userService))

	maintainer.Patch("/block", handler.BlockUser(userService))
	maintainer.Patch("/unblock", handler.UnblockUser(userService))
	maintainer.Post("/memo", handler.CreateMemo(userService))
}
