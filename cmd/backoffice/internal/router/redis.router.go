package router

import (
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler"
	"github.com/finexblock-dev/gofinexblock/pkg/goredis"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func RedisRouter(router fiber.Router, cluster *redis.ClusterClient) {
	svc := goredis.NewService(cluster)

	redisRouter := router.Group("/redis")

	redisRouter.Get("/xrange", handler.XRange(svc))
	redisRouter.Get("/xinfostream", handler.XInfoStream(svc))
	//redisRouter.Get("/xrevrange", handler.XRevRange(svc))
	//redisRouter.Get("/xlen", handler.XLen(svc))
	//redisRouter.Get("/xadd", handler.XAdd(svc))
	//redisRouter.Get("/xdel", handler.XDel(svc))
	//redisRouter.Get("/xtrim", handler.XTrim(svc))
	//redisRouter.Get("/xgroupcreate", handler.XGroupCreate(svc))
	//redisRouter.Get("/xgroupdelconsumer", handler.XGroupDelConsumer(svc))
	//redisRouter.Get("/xgroupdestroy", handler.XGroupDestroy(svc))
	//redisRouter.Get("/xgroupsetid", handler.XGroupSetID(svc))
	//redisRouter.Get("/xgroupdel", handler.XGroupDel(svc))
	redisRouter.Get("/get", handler.Get(svc))
	redisRouter.Post("/set", handler.Set(svc))
	redisRouter.Delete("/del", handler.Del(svc))
}