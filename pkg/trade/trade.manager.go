package trade

import (
	"context"
	"github.com/finexblock-dev/gofinexblock/pkg/goredis"
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
		return ErrNegativeBalance
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
		return ErrNegativeBalance
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

	m.matchStreamInit(group)

	m.errorStreamInit(group)

	m.cancellationStreamInit(group)

	m.balanceUpdateStreamInit(group)

	m.placementStreamInit(group)

	m.initializeStreamInit(group)

	m.fulfillmentStreamInit(group)

	m.partialFillStreamInit(group)

	m.matchingStreamInit(group)

	m.marketOrderMatchingStreamInit(group)

	if err = group.Wait(); err != nil {
		return err
	}

	group, _ = errgroup.WithContext(context.TODO())

	m.cancellationConsumerInit(group)

	m.matchConsumerInit(group)

	m.errorConsumerInit(group)

	m.initializeConsumerInit(group)

	m.balanceUpdateConsumerInit(group)

	m.placementConsumerInit(group)

	m.fulfillmentConsumerInit(group)

	m.partialFillConsumerInit(group)

	m.matchingConsumerInit(group)

	m.marketOrderMatchingConsumerInit(group)

	return group.Wait()
}

func (m *manager) Pipeliner() redis.Pipeliner {
	return m.cluster.TxPipeline()
}

func newManager(cluster goredis.Service) *manager {
	return &manager{cluster: cluster}
}