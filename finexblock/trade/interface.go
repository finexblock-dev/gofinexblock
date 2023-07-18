package trade

import (
	"context"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/goredis"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
	"time"
)

type Manager interface {
	AcquireLock(uuid, currency string) (bool, error)
	ReleaseLock(uuid, currency string) error
	GetBalance(uuid, currency string) (decimal.Decimal, error)
	SetBalance(uuid, currency string, amount decimal.Decimal) error
	PlusBalance(uuid, currency string, amount decimal.Decimal) error
	MinusBalance(uuid, currency string, amount decimal.Decimal) error
	SetOrder(orderUUID string, side string) error
	GetOrder(orderUUID string) (string, error)
	DeleteOrder(orderUUID string) error

	StreamsInit() error

	SendMatchStream(matchCase types.Case, pair *grpc_order.BidAsk) error
	SendErrorStream(input *grpc_order.ErrorInput) error
	SendInitializeStream(order *grpc_order.OrderInitialize) error
	SendPlacementStream(order *grpc_order.OrderPlacement) error
	SendCancellationStream(order *grpc_order.OrderCancelled) error
	SendBalanceUpdateStream(event *grpc_order.BalanceUpdate) error
	SendOrderFulfillmentStream(event *grpc_order.OrderFulfillment) error
	SendOrderPartialFillStream(event *grpc_order.OrderPartialFill) error
	SendOrderMatchingStream(event *grpc_order.OrderMatching) error
	SendMarketOrderMatchingStream(event *grpc_order.MarketOrderMatching) error

	Pipeliner() redis.Pipeliner

	SendMatchStreamPipeline(tx redis.Pipeliner, ctx context.Context, matchCase types.Case, pair *grpc_order.BidAsk) error
	SendErrorStreamPipeline(tx redis.Pipeliner, ctx context.Context, input *grpc_order.ErrorInput) error
	SendInitializeStreamPipeline(tx redis.Pipeliner, ctx context.Context, event *grpc_order.OrderInitialize) error
	SendPlacementStreamPipeline(tx redis.Pipeliner, ctx context.Context, event *grpc_order.OrderPlacement) error
	SendCancellationStreamPipeline(tx redis.Pipeliner, ctx context.Context, event *grpc_order.OrderCancelled) error
	SendBalanceUpdateStreamPipeline(tx redis.Pipeliner, ctx context.Context, event *grpc_order.BalanceUpdate) error
	SendOrderFulfillmentStreamPipeline(tx redis.Pipeliner, ctx context.Context, event *grpc_order.OrderFulfillment) error
	SendOrderPartialFillStreamPipeline(tx redis.Pipeliner, ctx context.Context, event *grpc_order.OrderPartialFill) error
	SendOrderMatchingStreamPipeline(tx redis.Pipeliner, ctx context.Context, event *grpc_order.OrderMatching) error
	SendMarketOrderMatchingStreamPipeline(tx redis.Pipeliner, ctx context.Context, event *grpc_order.MarketOrderMatching) error

	ReadStream(stream types.Stream, group types.Group, consumer types.Consumer, count int64, block time.Duration) ([]redis.XStream, error)
	ReadStreams(stream []types.Stream, group types.Group, consumer types.Consumer, count int64, block time.Duration) ([]redis.XStream, error)
	ReadPendingStream(stream types.Stream, group types.Group) (*redis.XPending, error)
	ReadStreamInfo(stream types.Stream) (*redis.XInfoStream, error)

	AckStream(stream types.Stream, group types.Group, id string) (err error)
	ClaimStream(stream types.Stream, group types.Group, consumer types.Consumer, minIdleTime time.Duration, ids []string) ([]redis.XMessage, error)
}

func New(redisClient *redis.ClusterClient) Manager {
	return newManager(goredis.NewService(redisClient))
}