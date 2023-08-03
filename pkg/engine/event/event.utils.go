package event

import (
	"context"
	"encoding/json"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/pkg/trade"
	"github.com/finexblock-dev/gofinexblock/pkg/types"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func (e *engine) ParseMessage(stream types.Stream, message redis.XMessage) (event proto.Message, err error) {
	var bytes []byte
	var data = make(map[string]string)

	bytes, err = json.Marshal(message.Values)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	switch stream {
	case trade.OrderInitializeStream:
		event = new(grpc_order.OrderInitialize)
	case trade.OrderPlacementStream:
		event = new(grpc_order.OrderPlacement)
	case trade.OrderFulfillmentStream:
		event = new(grpc_order.OrderFulfillment)
	case trade.OrderPartialFillStream:
		event = new(grpc_order.OrderPartialFill)
	case trade.OrderMatchingStream:
		event = new(grpc_order.OrderMatching)
	case trade.MarketOrderMatchingStream:
		event = new(grpc_order.MarketOrderMatching)
	default:
		return nil, status.Errorf(codes.InvalidArgument, "invalid stream: %s", stream)
	}

	if err = protojson.Unmarshal([]byte(data["event"]), event); err != nil {
		return nil, err
	}

	return event, nil
}

func (e *engine) Do(stream types.Stream, message proto.Message) (err error) {
	var ctx = context.Background()
	switch stream {
	case trade.OrderInitializeStream:
		if event, ok := message.(*grpc_order.OrderInitialize); ok {
			_, err = e.eventSubscriber.OrderInitializeEvent(ctx, event)
			return err
		}
		return status.Errorf(codes.InvalidArgument, "invalid message: %s, %v", stream)
	case trade.OrderPlacementStream:
		if event, ok := message.(*grpc_order.OrderPlacement); ok {
			_, err = e.chartServer.OrderPlacementEvent(ctx, event)
			_, err = e.eventSubscriber.OrderPlacementEvent(ctx, event)
			return err
		}
		return status.Errorf(codes.InvalidArgument, "invalid message: %s, %v", stream)
	case trade.OrderFulfillmentStream:
		if event, ok := message.(*grpc_order.OrderFulfillment); ok {
			_, err = e.chartServer.OrderFulfillmentEvent(ctx, event)
			_, err = e.eventSubscriber.OrderFulfillmentEvent(ctx, event)
			return err
		}
		return status.Errorf(codes.InvalidArgument, "invalid message: %s, %v", stream)
	case trade.OrderPartialFillStream:
		if event, ok := message.(*grpc_order.OrderPartialFill); ok {
			_, err = e.chartServer.OrderPartialFillEvent(ctx, event)
			_, err = e.eventSubscriber.OrderPartialFillEvent(ctx, event)
			return err
		}
		return status.Errorf(codes.InvalidArgument, "invalid message: %s, %v", stream)
	case trade.OrderMatchingStream:
		if event, ok := message.(*grpc_order.OrderMatching); ok {
			_, err = e.chartServer.OrderMatchingEvent(ctx, event)
			_, err = e.eventSubscriber.OrderMatchingEvent(ctx, event)
			return err
		}
		return status.Errorf(codes.InvalidArgument, "invalid message: %s, %v", stream)
	case trade.MarketOrderMatchingStream:
		if event, ok := message.(*grpc_order.MarketOrderMatching); ok {
			_, err = e.eventSubscriber.MarketOrderMatchingEvent(ctx, event)
			return err
		}
		return status.Errorf(codes.InvalidArgument, "invalid message: %s, %v", stream)
	default:
		return status.Errorf(codes.InvalidArgument, "invalid stream: %s", stream)
	}
}