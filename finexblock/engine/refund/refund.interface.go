package refund

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/event"
	"github.com/finexblock-dev/gofinexblock/finexblock/stream"
)

type Engine interface {
	stream.Consumer
	stream.Claimer
	event.Hook
}