package router

import (
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/middleware"
	"github.com/finexblock-dev/gofinexblock/pkg/admin"
	"github.com/finexblock-dev/gofinexblock/pkg/order"
	"github.com/finexblock-dev/gofinexblock/pkg/wallet"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func TradeRouter(router fiber.Router, db *gorm.DB, cluster *redis.ClusterClient) {
	orderService := order.NewService(db)
	walletService := wallet.NewService(db, cluster)
	adminService := admin.NewService(db)
	tradeHandler := handler.NewTradeHandler(walletService, orderService)

	base := router.Group("/trade", middleware.BearerTokenMiddleware(), middleware.AdminApiLogMiddleware(adminService))

	support := SupportRouter(base, adminService)

	support.Get("/search", tradeHandler.SearchTradeHistory())
}