package router

import (
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func GrpcRouter(router fiber.Router) {
	r := router.Group("/gRPC")
	r.Get("/health", handler.ProxyHealthCheck())
}