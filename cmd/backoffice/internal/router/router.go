package router

import (
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/middleware"
	"github.com/finexblock-dev/gofinexblock/pkg/admin"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Router(db *gorm.DB, cluster *redis.ClusterClient) *fiber.App {
	app := fiber.New()

	app.Use(recover.New()).Use(cors.New()).Use(logger.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello world!")
	})

	Swagger(app)

	AdminRouter(app, db)

	AuthRouter(app, db)

	AnnouncementRouter(app, db)

	ImageRouter(app, db)

	UserRouter(app, db, cluster)

	AssetRouter(app, db, cluster)

	WithdrawalRouter(app, db, cluster)

	RedisRouter(app, cluster)

	GrpcRouter(app)
	return app
}

func SupportRouter(router fiber.Router, service admin.Service) fiber.Router {
	return router.Use(middleware.SupportGuard(service))
}

func MaintainerRouter(router fiber.Router, service admin.Service) fiber.Router {
	return router.Use(middleware.MaintainerGuard(service))
}

func SuperUserRouter(router fiber.Router, service admin.Service) fiber.Router {
	return router.Use(middleware.SuperUserGuard(service))
}