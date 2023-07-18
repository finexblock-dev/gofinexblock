package refund

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/stream"
)

type Engine interface {
	stream.Consumer
	stream.Claimer
}