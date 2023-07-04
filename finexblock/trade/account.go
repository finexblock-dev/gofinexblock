package trade

import (
	"context"
	"github.com/shopspring/decimal"
)

func (a *accountService) GetBalance(ctx context.Context, uuid, currency string) (decimal.Decimal, error) {
	var key string
	var value string
	var decimalValue decimal.Decimal
	var err error

	key = getKey(uuid, currency)
	value, err = a.redisClient.Get(ctx, key).Result()
	if err != nil {
		return decimal.Zero, wrapErr(ErrKeyNotFound, err)
	}

	decimalValue, err = decimal.NewFromString(value)
	if err != nil {
		return decimal.Zero, wrapErr(ErrDecimalParse, err)
	}

	return decimalValue, nil
}

func (a *accountService) SetBalance(ctx context.Context, uuid, currency string, amount decimal.Decimal) error {
	var key string

	key = getKey(uuid, currency)
	return a.redisClient.Set(ctx, key, amount.String(), 0).Err()
}

func (a *accountService) PlusBalance(ctx context.Context, uuid, currency string, amount decimal.Decimal) error {
	var key, value string
	var decimalValue decimal.Decimal
	var err error
	key = getKey(uuid, currency)
	value, err = a.redisClient.Get(ctx, key).Result()
	if err != nil {
		return wrapErr(ErrKeyNotFound, err)
	}
	decimalValue, err = decimal.NewFromString(value)
	if err != nil {
		return wrapErr(ErrDecimalParse, err)
	}

	decimalValue = decimalValue.Add(amount)
	return a.SetBalance(ctx, uuid, currency, decimalValue)
}

func (a *accountService) MinusBalance(ctx context.Context, uuid, currency string, amount decimal.Decimal) error {
	var key, value string
	var decimalValue decimal.Decimal
	var err error

	key = getKey(uuid, currency)
	value, err = a.redisClient.Get(ctx, key).Result()
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
	return a.SetBalance(ctx, uuid, currency, decimalValue)
}