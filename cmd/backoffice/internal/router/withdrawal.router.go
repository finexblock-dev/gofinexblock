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

	withdrawalRouter := router.Group("/withdraw", middleware.BearerTokenMiddleware(), middleware.AdminApiLogMiddleware(adminService))

	SupportRouter(withdrawalRouter, adminService).Get("/", handler.ScanWithdrawalRequestByStatus(walletService))

	MaintainerRouter(withdrawalRouter, adminService).Get("/user", handler.FindWithdrawalRequestsByUserID(walletService))

	MaintainerRouter(withdrawalRouter, adminService).Patch("/reject", handler.RejectWithdrawalRequests(walletService))

	MaintainerRouter(withdrawalRouter, adminService).Patch("/approve", handler.ApproveWithdrawalRequests(walletService))
}