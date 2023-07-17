package cancellation

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/event"
	"github.com/finexblock-dev/gofinexblock/finexblock/safety"
	"github.com/finexblock-dev/gofinexblock/finexblock/stream"
)

type Engine interface {
	safety.Subscriber
	stream.Consumer
	stream.Claimer
	event.Hook
}