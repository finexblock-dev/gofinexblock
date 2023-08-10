package router

import (
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/middleware"
	"github.com/finexblock-dev/gofinexblock/pkg/admin"
	"github.com/finexblock-dev/gofinexblock/pkg/wallet"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func AssetRouter(router fiber.Router, db *gorm.DB, cluster *redis.ClusterClient) {
	walletService := wallet.NewService(db, cluster)
	adminService := admin.NewService(db)

	assetRouter := router.Group("/asset", middleware.BearerTokenMiddleware(), middleware.AdminApiLogMiddleware(adminService))

	SupportRouter(assetRouter, adminService).Get("/", handler.FindUserAssets(walletService))
	SupportRouter(assetRouter, adminService).Get("/balance/log", handler.FindUserBalanceUpdateLog(walletService))
	SupportRouter(assetRouter, adminService).Get("/search", handler.FindUserAssetsByCond(walletService))
}
