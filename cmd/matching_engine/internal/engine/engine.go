package engine

import (
	"github.com/finexblock-dev/gofinexblock/pkg/engine/cancellation"
	"github.com/finexblock-dev/gofinexblock/pkg/engine/event"
	"github.com/finexblock-dev/gofinexblock/pkg/engine/match"
	"github.com/finexblock-dev/gofinexblock/pkg/engine/refund"
	"github.com/finexblock-dev/gofinexblock/pkg/safety"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type Engine struct {
	refund       refund.Engine
	cancellation cancellation.Engine
	match        match.Engine
	event        event.Engine
}

func NewEngine(cluster *redis.ClusterClient, eventSubscriber, chartServer *grpc.ClientConn) *Engine {
	return &Engine{
		refund:       refund.New(cluster, eventSubscriber, chartServer),
		cancellation: cancellation.New(cluster, eventSubscriber, chartServer),
		match:        match.New(cluster),
		event:        event.New(cluster, eventSubscriber, chartServer),
	}
}

func (e *Engine) Run() {
	go safety.GracefullyStopBootstrap(e.refund.Consume)
	go safety.GracefullyStopBootstrap(e.refund.Claim)

	go safety.GracefullyStopBootstrap(e.cancellation.Consume)
	go safety.GracefullyStopBootstrap(e.cancellation.Claim)

	go safety.GracefullyStopBootstrap(e.match.Consume)
	go safety.GracefullyStopBootstrap(e.match.Claim)

	go safety.GracefullyStopBootstrap(e.event.Consume)
	go safety.GracefullyStopBootstrap(e.event.Claim)
}