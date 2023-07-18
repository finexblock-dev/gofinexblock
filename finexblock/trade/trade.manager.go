package trade

import (
	"context"
	"github.com/finexblock-dev/gofinexblock/finexblock/goredis"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
	"golang.org/x/sync/errgroup"
	"math"
	"time"
)

type manager struct {
	cluster goredis.Service
}

func (m *manager) SetBalanceWithTx(tx redis.Pipeliner, ctx context.Context, uuid, currency string, amount decimal.Decimal) error {
	var key string

	key = getBalanceKey(uuid, currency)
	return tx.Set(ctx, key, amount.String(), 0).Err()
}

func (m *manager) MinusBalanceWithTx(tx redis.Pipeliner, ctx context.Context, uuid, currency string, amount decimal.Decimal) error {
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
	return m.SetBalanceWithTx(tx, ctx, uuid, currency, decimalValue)
}

func (m *manager) PlusBalanceWithTx(tx redis.Pipeliner, ctx context.Context, uuid, currency string, amount decimal.Decimal) error {
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
	return m.SetBalanceWithTx(tx, ctx, uuid, currency, decimalValue)
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

func (m *manager) Pipeliner() redis.Pipeliner {
	return m.cluster.TxPipeline()
}