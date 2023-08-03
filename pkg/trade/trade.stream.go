package trade

import (
	"context"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/pkg/types"
	"github.com/finexblock-dev/gofinexblock/pkg/utils"
	"github.com/redis/go-redis/v9"
	"time"
)

func (m *manager) SendMarketOrderMatchingStream(event *grpc_order.MarketOrderMatching) error {
	var stream []byte
	var err error

	stream, err = utils.MessagesToJson(event)
	if err != nil {
		return ErrMarshalFailed
	}

	return m.cluster.XAdd(&redis.XAddArgs{
		Stream: MarketOrderMatchingStream.String(),
		ID:     "*",
		Values: []string{"event", string(stream)},
	})
}

func (m *manager) SendMarketOrderMatchingStreamPipeline(tx redis.Pipeliner, ctx context.Context, event *grpc_order.MarketOrderMatching) error {
	var stream []byte
	var err error

	stream, err = utils.MessagesToJson(event)
	if err != nil {
		return ErrMarshalFailed
	}

	return m.cluster.XAddPipeline(tx, ctx, &redis.XAddArgs{
		Stream: MarketOrderMatchingStream.String(),
		ID:     "*",
		Values: []string{"event", string(stream)},
	})
}

func (m *manager) ReadStreams(stream []types.Stream, group types.Group, consumer types.Consumer, count int64, block time.Duration) ([]redis.XStream, error) {
	var values []string

	for _, s := range stream {
		values = append(values, s.String(), ">")
	}

	return m.cluster.XReadGroup(&redis.XReadGroupArgs{
		Group:    group.String(),
		Consumer: consumer.String(),
		Streams:  values,
		Count:    count,
		Block:    block,
	})
}

func (m *manager) AckStream(stream types.Stream, group types.Group, id string) (err error) {
	return m.cluster.XAck(stream.String(), group.String(), id)
}

func (m *manager) SendMatchStreamPipeline(tx redis.Pipeliner, ctx context.Context, matchCase types.Case, pair *grpc_order.BidAsk) error {
	var stream []byte
	var err error

	stream, err = utils.MessagesToJson(pair)
	if err != nil {
		return ErrMarshalFailed
	}

	return m.cluster.XAddPipeline(tx, ctx, &redis.XAddArgs{
		Stream: MatchStream.String(),
		ID:     "*",
		Values: []string{"pair", string(stream), "case", matchCase.String()},
	})
}

func (m *manager) SendPlacementStreamPipeline(tx redis.Pipeliner, ctx context.Context, order *grpc_order.OrderPlacement) error {
	var stream []byte
	var err error

	stream, err = utils.MessagesToJson(order)
	if err != nil {
		return ErrMarshalFailed
	}

	return m.cluster.XAddPipeline(tx, ctx, &redis.XAddArgs{
		Stream: OrderPlacementStream.String(),
		ID:     "*",
		Values: []string{"event", string(stream)},
	})
}

func (m *manager) SendErrorStreamPipeline(tx redis.Pipeliner, ctx context.Context, input *grpc_order.ErrorInput) error {
	var stream []byte
	var err error

	stream, err = utils.MessagesToJson(input)
	if err != nil {
		return ErrMarshalFailed
	}

	return m.cluster.XAddPipeline(tx, ctx, &redis.XAddArgs{
		Stream: ErrorStream.String(),
		ID:     "*",
		Values: []string{"error", string(stream)},
	})
}

func (m *manager) SendCancellationStreamPipeline(tx redis.Pipeliner, ctx context.Context, event *grpc_order.OrderCancelled) error {
	var stream []byte
	var err error

	stream, err = utils.MessagesToJson(event)
	if err != nil {
		return ErrMarshalFailed
	}

	return m.cluster.XAddPipeline(tx, ctx, &redis.XAddArgs{
		Stream: OrderCancellationStream.String(),
		ID:     "*",
		Values: []string{"event", string(stream)},
	})
}

func (m *manager) SendBalanceUpdateStreamPipeline(tx redis.Pipeliner, ctx context.Context, event *grpc_order.BalanceUpdate) error {
	var stream []byte
	var err error

	stream, err = utils.MessagesToJson(event)
	if err != nil {
		return ErrMarshalFailed
	}

	return m.cluster.XAddPipeline(tx, ctx, &redis.XAddArgs{
		Stream: BalanceUpdateStream.String(),
		ID:     "*",
		Values: []string{"event", string(stream)},
	})
}

func (m *manager) SendBalanceUpdateStream(event *grpc_order.BalanceUpdate) error {
	var stream []byte
	var err error

	stream, err = utils.MessagesToJson(event)
	if err != nil {
		return ErrMarshalFailed
	}

	return m.cluster.XAdd(&redis.XAddArgs{
		Stream: BalanceUpdateStream.String(),
		ID:     "*",
		Values: []string{"event", string(stream)},
	})
}

func (m *manager) SendOrderFulfillmentStreamPipeline(tx redis.Pipeliner, ctx context.Context, event *grpc_order.OrderFulfillment) error {
	var stream []byte
	var err error

	stream, err = utils.MessagesToJson(event)
	if err != nil {
		return ErrMarshalFailed
	}

	return m.cluster.XAddPipeline(tx, ctx, &redis.XAddArgs{
		Stream: OrderFulfillmentStream.String(),
		ID:     "*",
		Values: []string{"event", string(stream)},
	})
}

func (m *manager) SendOrderFulfillmentStream(event *grpc_order.OrderFulfillment) error {
	var stream []byte
	var err error

	stream, err = utils.MessagesToJson(event)
	if err != nil {
		return ErrMarshalFailed
	}

	return m.cluster.XAdd(&redis.XAddArgs{
		Stream: OrderFulfillmentStream.String(),
		ID:     "*",
		Values: []string{"event", string(stream)},
	})
}

func (m *manager) SendOrderPartialFillStreamPipeline(tx redis.Pipeliner, ctx context.Context, event *grpc_order.OrderPartialFill) error {
	var stream []byte
	var err error

	stream, err = utils.MessagesToJson(event)
	if err != nil {
		return ErrMarshalFailed
	}

	return m.cluster.XAddPipeline(tx, ctx, &redis.XAddArgs{
		Stream: OrderPartialFillStream.String(),
		ID:     "*",
		Values: []string{"event", string(stream)},
	})
}

func (m *manager) SendOrderPartialFillStream(event *grpc_order.OrderPartialFill) error {
	var stream []byte
	var err error

	stream, err = utils.MessagesToJson(event)
	if err != nil {
		return ErrMarshalFailed
	}

	return m.cluster.XAdd(&redis.XAddArgs{
		Stream: OrderPartialFillStream.String(),
		ID:     "*",
		Values: []string{"event", string(stream)},
	})
}

func (m *manager) SendOrderMatchingStreamPipeline(tx redis.Pipeliner, ctx context.Context, event *grpc_order.OrderMatching) error {
	var stream []byte
	var err error

	stream, err = utils.MessagesToJson(event)
	if err != nil {
		return ErrMarshalFailed
	}

	return m.cluster.XAddPipeline(tx, ctx, &redis.XAddArgs{
		Stream: OrderMatchingStream.String(),
		ID:     "*",
		Values: []string{"event", string(stream)},
	})
}

func (m *manager) SendOrderMatchingStream(event *grpc_order.OrderMatching) error {
	var stream []byte
	var err error

	stream, err = utils.MessagesToJson(event)
	if err != nil {
		return ErrMarshalFailed
	}

	return m.cluster.XAdd(&redis.XAddArgs{
		Stream: OrderMatchingStream.String(),
		ID:     "*",
		Values: []string{"event", string(stream)},
	})
}

func (m *manager) SendInitializeStreamPipeline(tx redis.Pipeliner, ctx context.Context, event *grpc_order.OrderInitialize) error {
	var stream []byte
	var err error

	stream, err = utils.MessagesToJson(event)
	if err != nil {
		return ErrMarshalFailed
	}

	return m.cluster.XAddPipeline(tx, ctx, &redis.XAddArgs{
		Stream: OrderInitializeStream.String(),
		ID:     "*",
		Values: []string{"event", string(stream)},
	})
}

func (m *manager) SendInitializeStream(event *grpc_order.OrderInitialize) error {
	var stream []byte
	var err error

	stream, err = utils.MessagesToJson(event)
	if err != nil {
		return ErrMarshalFailed
	}

	return m.cluster.XAdd(&redis.XAddArgs{
		Stream: OrderInitializeStream.String(),
		ID:     "*",
		Values: []string{"event", string(stream)},
	})
}

func (m *manager) ReadStreamInfo(stream types.Stream) (*redis.XInfoStream, error) {
	return m.cluster.XInfoStream(stream.String())
}

func (m *manager) ClaimStream(stream types.Stream, group types.Group, consumer types.Consumer, minIdleTime time.Duration, ids []string) ([]redis.XMessage, error) {
	return m.cluster.XClaim(&redis.XClaimArgs{
		Stream:   stream.String(),
		Group:    group.String(),
		Consumer: consumer.String(),
		MinIdle:  minIdleTime,
		Messages: ids,
	})
}

func (m *manager) ReadPendingStream(stream types.Stream, group types.Group) (*redis.XPending, error) {
	return m.cluster.XPending(stream.String(), group.String())
}

func (m *manager) ReadStream(stream types.Stream, group types.Group, consumer types.Consumer, count int64, block time.Duration) ([]redis.XStream, error) {
	return m.cluster.XReadGroup(&redis.XReadGroupArgs{
		Group:    group.String(),
		Consumer: consumer.String(),
		Streams:  []string{stream.String(), ">"},
		Count:    count,
		Block:    block,
	})
}

func (m *manager) SendCancellationStream(event *grpc_order.OrderCancelled) error {
	var stream []byte
	var err error

	stream, err = utils.MessagesToJson(event)
	if err != nil {
		return ErrMarshalFailed
	}

	return m.cluster.XAdd(&redis.XAddArgs{
		Stream: OrderCancellationStream.String(),
		ID:     "*",
		Values: []string{"event", string(stream)},
	})
}

func (m *manager) SendMatchStream(matchCase types.Case, pair *grpc_order.BidAsk) error {
	var stream []byte
	var err error

	stream, err = utils.MessagesToJson(pair)
	if err != nil {
		return ErrMarshalFailed
	}

	return m.cluster.XAdd(&redis.XAddArgs{
		Stream: MatchStream.String(),
		ID:     "*",
		Values: []string{"pair", string(stream), "case", matchCase.String()},
	})
}

func (m *manager) SendPlacementStream(event *grpc_order.OrderPlacement) error {
	var stream []byte
	var err error

	stream, err = utils.MessagesToJson(event)
	if err != nil {
		return ErrMarshalFailed
	}

	return m.cluster.XAdd(&redis.XAddArgs{
		Stream: OrderPlacementStream.String(),
		ID:     "*",
		Values: []string{"event", string(stream)},
	})
}

func (m *manager) SendErrorStream(input *grpc_order.ErrorInput) error {
	var stream []byte
	var err error

	stream, err = utils.MessagesToJson(input)
	if err != nil {
		return ErrMarshalFailed
	}

	return m.cluster.XAdd(&redis.XAddArgs{
		Stream: ErrorStream.String(),
		ID:     "*",
		Values: []string{"error", string(stream)},
	})
}