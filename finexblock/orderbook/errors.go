package orderbook

import "errors"

var (
	ErrOrderNotFound = errors.New("order not found")

	ErrOrderTypeNotFound = errors.New("order type not found")

	ErrOrderCancelFailed = errors.New("order cancel failed")
)
