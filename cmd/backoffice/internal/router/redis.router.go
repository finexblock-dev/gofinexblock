package router

import (
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler"
	"github.com/finexblock-dev/gofinexblock/pkg/goredis"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func RedisRouter(router fiber.Router, cluster *redis.ClusterClient) {
	svc := goredis.NewService(cluster)

	redisHandler := handler.NewRedisHandler(svc)

	redisRouter := router.Group("/redis")

	redisRouter.Get("/xrange", redisHandler.XRange())
	redisRouter.Get("/xinfostream", redisHandler.XInfoStream())
	//redisRouter.Get("/xrevrange", redisHandler.XRevRange())
	//redisRouter.Get("/xlen", redisHandler.XLen())
	//redisRouter.Get("/xadd", redisHandler.XAdd())
	//redisRouter.Get("/xdel", redisHandler.XDel())
	//redisRouter.Get("/xtrim", redisHandler.XTrim())
	//redisRouter.Get("/xgroupcreate", redisHandler.XGroupCreate())
	//redisRouter.Get("/xgroupdelconsumer", redisHandler.XGroupDelConsumer())
	//redisRouter.Get("/xgroupdestroy", redisHandler.XGroupDestroy())
	//redisRouter.Get("/xgroupsetid", redisHandler.XGroupSetID())
	//redisRouter.Get("/xgroupdel", redisHandler.XGroupDel())
	redisRouter.Get("/get", redisHandler.Get())
	redisRouter.Post("/set", redisHandler.Set())
	redisRouter.Delete("/del", redisHandler.Del())

	redisRouter.Get("/keys", redisHandler.Keys())
	redisRouter.Delete("/deleteRefreshToken", redisHandler.DeleteAllRefreshTokens())
}
