package router

import (
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/middleware"
	"github.com/finexblock-dev/gofinexblock/pkg/admin"
	"github.com/finexblock-dev/gofinexblock/pkg/wallet"
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func WithdrawalRouter(router fiber.Router, db *gorm.DB, cluster *redis.ClusterClient) {
	walletService := wallet.NewService(db, cluster)
	adminService := admin.NewService(db)
	withdrawalHandler := handler.NewWithdrawalHandler(walletService)

	base := router.Group("/withdraw", middleware.BearerTokenMiddleware(), middleware.AdminApiLogMiddleware(adminService))
	support := SupportRouter(base, adminService)
	maintainer := MaintainerRouter(base, adminService)

	support.Get("/", withdrawalHandler.ScanWithdrawalRequestByStatus())

	maintainer.Get("/user", withdrawalHandler.FindWithdrawalRequestsByUserID())

	maintainer.Patch("/reject", withdrawalHandler.RejectWithdrawalRequests())

	maintainer.Patch("/approve", withdrawalHandler.ApproveWithdrawalRequests())
}
