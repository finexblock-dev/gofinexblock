package event

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/stream"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type Engine interface {
	stream.MultiStreamConsumer
	stream.Claimer
}

func New(cluster *redis.ClusterClient, conn *grpc.ClientConn) Engine {
	return newEngine(cluster, conn)
}