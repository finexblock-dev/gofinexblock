package trade

import (
	"fmt"
	"github.com/finexblock-dev/gofinexblock/finexblock/goredis"
	"github.com/shopspring/decimal"
	"time"
)

type service struct {
	cluster goredis.Service
}

func newService(cluster goredis.Service) *service {
	return &service{cluster: cluster}
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
	return s.cluster.Set(fmt.Sprintf("order:%v", orderUUID), side, 0)
}

func (s *service) GetOrder(orderUUID string) (string, error) {
	return s.cluster.Get(fmt.Sprintf("order:%v", orderUUID))
}

func (s *service) DeleteOrder(orderUUID string) error {
	return s.cluster.Del(fmt.Sprintf("order:%v", orderUUID))
}
