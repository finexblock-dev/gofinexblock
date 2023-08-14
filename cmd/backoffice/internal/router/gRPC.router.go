package router

import (
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func GrpcRouter(router fiber.Router) {
	r := router.Group("/gRPC")

	grpcHandler := handler.NewGrpcHandler()

	r.Get("/health", grpcHandler.ProxyHealthCheck())
	r.Get("/whoami", grpcHandler.ProxyWhoAmI())
}
