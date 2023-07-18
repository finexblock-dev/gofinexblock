package trade

import (
	"context"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/goredis"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/finexblock-dev/gofinexblock/finexblock/utils"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"
	"math"
	"time"
)

type manager struct {
	cluster goredis.Service
}

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

func (m *manager) Pipeliner() redis.Pipeliner {
	return m.cluster.TxPipeline()
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

func (m *manager) AcquireLock(uuid, currency string) (bool, error) {
	var key string

	if testAccounts[uuid] {
		return true, nil
	}
	key = getAccountLockKey(uuid, currency)
	return m.cluster.SetNX(key, lock, time.Second*10)
}

func (m *manager) ReleaseLock(uuid, currency string) error {
	var key string

	if testAccounts[uuid] {
		return nil
	}
	key = getAccountLockKey(uuid, currency)
	return m.cluster.Del(key)
}

func (m *manager) GetBalance(uuid, currency string) (decimal.Decimal, error) {
	var key string
	var value string
	var decimalValue decimal.Decimal
	var err error

	if testAccounts[uuid] {
		return decimal.NewFromFloat(math.MaxFloat64), nil
	}
	key = getBalanceKey(uuid, currency)
	value, err = m.cluster.Get(key)
	if err != nil {
		return decimal.Zero, wrapErr(ErrKeyNotFound, err)
	}

	decimalValue, err = decimal.NewFromString(value)
	if err != nil {
		return decimal.Zero, wrapErr(ErrDecimalParse, err)
	}

	return decimalValue, nil
}

func (m *manager) SetBalance(uuid, currency string, amount decimal.Decimal) error {
	var key string

	key = getBalanceKey(uuid, currency)
	return m.cluster.Set(key, amount.String(), 0)
}

func (m *manager) PlusBalance(uuid, currency string, amount decimal.Decimal) error {
	var key, value string
	var decimalValue decimal.Decimal
	var err error

	if testAccounts[uuid] {
		return nil
	}

	key = getBalanceKey(uuid, currency)
	value, err = m.cluster.Get(key)
	if err != nil {
		return wrapErr(ErrKeyNotFound, err)
	}
	decimalValue, err = decimal.NewFromString(value)
	if err != nil {
		return wrapErr(ErrDecimalParse, err)
	}

	decimalValue = decimalValue.Add(amount)
	return m.SetBalance(uuid, currency, decimalValue)
}

func (m *manager) MinusBalance(uuid, currency string, amount decimal.Decimal) error {
	var key, value string
	var decimalValue decimal.Decimal
	var err error

	if testAccounts[uuid] {
		return nil
	}
	key = getBalanceKey(uuid, currency)
	value, err = m.cluster.Get(key)
	if err != nil {
		return wrapErr(ErrKeyNotFound, err)
	}
	decimalValue, err = decimal.NewFromString(value)
	if err != nil {
		return wrapErr(ErrDecimalParse, err)
	}

	decimalValue = decimalValue.Sub(amount)
	if decimalValue.LessThan(decimal.Zero) {
		return wrapErr(ErrNegativeBalance, err)
	}
	return m.SetBalance(uuid, currency, decimalValue)
}

func (m *manager) SetOrder(orderUUID string, side string) error {
	return m.cluster.Set(getOrderKey(orderUUID), side, 0)
}

func (m *manager) GetOrder(orderUUID string) (string, error) {
	return m.cluster.Get(getOrderKey(orderUUID))
}

func (m *manager) DeleteOrder(orderUUID string) error {
	return m.cluster.Del(getOrderKey(orderUUID))
}

func (m *manager) StreamsInit() error {
	var err error

	group, _ := errgroup.WithContext(context.TODO())

	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(MatchStream.String(), MatchGroup.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(ErrorStream.String(), ErrorGroup.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(OrderCancellationStream.String(), OrderCancellationGroup.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(BalanceUpdateStream.String(), BalanceUpdateGroup.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(OrderPlacementStream.String(), EventGroup.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(OrderInitializeStream.String(), EventGroup.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(OrderFulfillmentStream.String(), EventGroup.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(OrderPartialFillStream.String(), EventGroup.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(OrderMatchingStream.String(), EventGroup.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateMkStream(MarketOrderMatchingStream.String(), EventGroup.String())
	})

	if err = group.Wait(); err != nil {
		return err
	}

	group, _ = errgroup.WithContext(context.TODO())

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(MatchStream.String(), MatchGroup.String(), MatchConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(MatchStream.String(), MatchGroup.String(), MatchClaimer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderPlacementStream.String(), OrderPlacementGroup.String(), OrderPlacementConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderPlacementStream.String(), OrderPlacementGroup.String(), OrderPlacementClaimer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(ErrorStream.String(), ErrorGroup.String(), ErrorConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(ErrorStream.String(), ErrorGroup.String(), ErrorClaimer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderInitializeStream.String(), OrderInitializeGroup.String(), OrderInitializeConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderInitializeStream.String(), OrderInitializeGroup.String(), OrderInitializeClaimer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(BalanceUpdateStream.String(), BalanceUpdateGroup.String(), BalanceUpdateConsumer.String())

	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(BalanceUpdateStream.String(), BalanceUpdateGroup.String(), BalanceUpdateClaimer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderPlacementStream.String(), EventGroup.String(), EventConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderPlacementStream.String(), EventGroup.String(), EventClaimer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderInitializeStream.String(), EventGroup.String(), EventConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderInitializeStream.String(), EventGroup.String(), EventClaimer.String())
	})
	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderFulfillmentStream.String(), EventGroup.String(), EventConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderFulfillmentStream.String(), EventGroup.String(), EventClaimer.String())
	})
	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderPartialFillStream.String(), EventGroup.String(), EventConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderPartialFillStream.String(), EventGroup.String(), EventClaimer.String())
	})
	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderMatchingStream.String(), EventGroup.String(), EventConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(OrderMatchingStream.String(), EventGroup.String(), EventClaimer.String())
	})
	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(MarketOrderMatchingStream.String(), EventGroup.String(), EventConsumer.String())
	})

	group.Go(func() error {
		return m.cluster.XGroupCreateConsumer(MarketOrderMatchingStream.String(), EventGroup.String(), EventClaimer.String())
	})

	return group.Wait()
}

func newManager(cluster goredis.Service) *manager {
	return &manager{cluster: cluster}
}