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

type service struct {
	cluster goredis.Service
}

func (s *service) ReadStream(args *redis.XReadGroupArgs) ([]redis.XStream, error) {
	return s.cluster.XReadGroup(args)
}

func (s *service) SendCancellationStream(order *grpc_order.Order) error {
	var stream map[string]interface{}
	var err error

	stream, err = utils.MessagesToJson(order)
	if err != nil {
		return ErrMarshalFailed
	}

	return s.cluster.XAdd(&redis.XAddArgs{
		Stream: CancelStream.String(),
		ID:     "*",
		Values: stream,
	})
}

func (s *service) SendMatchStream(matchCase types.Case, pair *grpc_order.BidAsk) error {
	var stream map[string]interface{}
	var err error

	stream, err = utils.MessagesToJson(pair)
	if err != nil {
		return ErrMarshalFailed
	}

	stream["case"] = matchCase.String()

	return s.cluster.XAdd(&redis.XAddArgs{
		Stream: MatchStream.String(),
		ID:     "*",
		Values: stream,
	})
}

func (s *service) SendPlacementStream(order *grpc_order.Order) error {
	var stream map[string]interface{}
	var err error

	stream, err = utils.MessagesToJson(order)
	if err != nil {
		return ErrMarshalFailed
	}

	return s.cluster.XAdd(&redis.XAddArgs{
		Stream: PlaceStream.String(),
		ID:     "*",
		Values: stream,
	})
}

func (s *service) SendRefundStream(order *grpc_order.Order) error {
	var stream map[string]interface{}
	var err error

	stream, err = utils.MessagesToJson(order)
	if err != nil {
		return ErrMarshalFailed
	}

	return s.cluster.XAdd(&redis.XAddArgs{
		Stream: RefundStream.String(),
		ID:     "*",
		Values: stream,
	})
}

func (s *service) SendErrorStream(input *grpc_order.ErrorInput) error {
	var stream map[string]interface{}
	var err error

	stream, err = utils.MessagesToJson(input)
	if err != nil {
		return ErrMarshalFailed
	}

	return s.cluster.XAdd(&redis.XAddArgs{
		Stream: ErrorStream.String(),
		ID:     "*",
		Values: stream,
	})
}

func (s *service) AcquireLock(uuid, currency string) (bool, error) {
	var key string

	if testAccounts[uuid] {
		return true, nil
	}
	key = getAccountLockKey(uuid, currency)
	return s.cluster.SetNX(key, lock, time.Second*10)
}

func (s *service) ReleaseLock(uuid, currency string) error {
	var key string

	if testAccounts[uuid] {
		return nil
	}
	key = getAccountLockKey(uuid, currency)
	return s.cluster.Del(key)
}

func (s *service) GetBalance(uuid, currency string) (decimal.Decimal, error) {
	var key string
	var value string
	var decimalValue decimal.Decimal
	var err error

	if testAccounts[uuid] {
		return decimal.NewFromFloat(math.MaxFloat64), nil
	}
	key = getBalanceKey(uuid, currency)
	value, err = s.cluster.Get(key)
	if err != nil {
		return decimal.Zero, wrapErr(ErrKeyNotFound, err)
	}

	decimalValue, err = decimal.NewFromString(value)
	if err != nil {
		return decimal.Zero, wrapErr(ErrDecimalParse, err)
	}

	return decimalValue, nil
}

func (s *service) SetBalance(uuid, currency string, amount decimal.Decimal) error {
	var key string

	key = getBalanceKey(uuid, currency)
	return s.cluster.Set(key, amount.String(), 0)
}

func (s *service) PlusBalance(uuid, currency string, amount decimal.Decimal) error {
	var key, value string
	var decimalValue decimal.Decimal
	var err error

	if testAccounts[uuid] {
		return nil
	}

	key = getBalanceKey(uuid, currency)
	value, err = s.cluster.Get(key)
	if err != nil {
		return wrapErr(ErrKeyNotFound, err)
	}
	decimalValue, err = decimal.NewFromString(value)
	if err != nil {
		return wrapErr(ErrDecimalParse, err)
	}

	decimalValue = decimalValue.Add(amount)
	return s.SetBalance(uuid, currency, decimalValue)
}

func (s *service) MinusBalance(uuid, currency string, amount decimal.Decimal) error {
	var key, value string
	var decimalValue decimal.Decimal
	var err error

	if testAccounts[uuid] {
		return nil
	}
	key = getBalanceKey(uuid, currency)
	value, err = s.cluster.Get(key)
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
	return s.SetBalance(uuid, currency, decimalValue)
}

func (s *service) SetOrder(orderUUID string, side string) error {
	return s.cluster.Set(getOrderKey(orderUUID), side, 0)
}

func (s *service) GetOrder(orderUUID string) (string, error) {
	return s.cluster.Get(getOrderKey(orderUUID))
}

func (s *service) DeleteOrder(orderUUID string) error {
	return s.cluster.Del(getOrderKey(orderUUID))
}

func (s *service) StreamsInit() error {
	var err error

	group, _ := errgroup.WithContext(context.TODO())

	group.Go(func() error {
		return s.cluster.XGroupCreateMkStream(MatchStream.String(), MatchGroup.String())
	})
	group.Go(func() error {
		return s.cluster.XGroupCreateMkStream(PlaceStream.String(), PlaceGroup.String())
	})
	group.Go(func() error {
		return s.cluster.XGroupCreateMkStream(RefundStream.String(), RefundGroup.String())
	})
	group.Go(func() error {
		return s.cluster.XGroupCreateMkStream(ErrorStream.String(), ErrorGroup.String())
	})
	group.Go(func() error {
		return s.cluster.XGroupCreateMkStream(CancelStream.String(), CancelGroup.String())
	})

	if err = group.Wait(); err != nil {
		return err
	}

	group, _ = errgroup.WithContext(context.TODO())

	group.Go(func() error {
		return s.cluster.XGroupCreateConsumer(MatchStream.String(), MatchGroup.String(), MatchConsumer.String())
	})
	group.Go(func() error {
		return s.cluster.XGroupCreateConsumer(PlaceStream.String(), PlaceGroup.String(), PlaceConsumer.String())
	})
	group.Go(func() error {
		return s.cluster.XGroupCreateConsumer(RefundStream.String(), RefundGroup.String(), RefundConsumer.String())
	})
	group.Go(func() error {
		return s.cluster.XGroupCreateConsumer(ErrorStream.String(), ErrorGroup.String(), ErrorConsumer.String())
	})
	group.Go(func() error {
		return s.cluster.XGroupCreateConsumer(CancelStream.String(), CancelGroup.String(), CancelConsumer.String())
	})

	return group.Wait()
}

func newService(cluster goredis.Service) *service {
	return &service{cluster: cluster}
}