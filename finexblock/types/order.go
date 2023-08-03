package types

import (
	"errors"
	"github.com/shopspring/decimal"
	"math"
)

type Reason string

func (r Reason) String() string {
	return string(r)
}

func (r Reason) Validate() error {
	switch r {
	case Cancel, Fill, Place:
		return nil
	default:
		return errors.New("invalid reason")
	}
}

const (
	Cancel Reason = "CANCEL"
	Fill   Reason = "FILL"
	Place  Reason = "PLACE"
)

type OrderType string

func (o OrderType) String() string {
	return string(o)
}

func (o OrderType) Validate() error {
	switch o {
	case Bid, Ask:
		return nil
	default:
		return errors.New("invalid order type")
	}
}

const (
	Bid OrderType = "BID"
	Ask OrderType = "ASK"
)

type OrderStatus string

func (o OrderStatus) String() string {
	return string(o)
}

func (o OrderStatus) Validate() error {
	switch o {
	case Cancelled, Placed, Fulfilled, PartialFilled:
		return nil
	default:
		return errors.New("invalid order status")
	}
}

const (
	Cancelled     OrderStatus = "CANCELLED"
	Placed        OrderStatus = "PLACED"
	Fulfilled     OrderStatus = "FULFILLED"
	PartialFilled OrderStatus = "PARTIAL_FILLED"
)

type Duration string

func (d Duration) String() string {
	return string(d)
}

func (d Duration) Validate() error {
	switch d {
	case OneMinute, FiveMinute, ThreeMinute, FifteenMinute, ThirtyMinute, OneHour, TwoHour, FourHour, SixHour, EightHour, TwelveHour, OneDay, ThreeDay, OneWeek, OneMonth:
		return nil
	default:
		return errors.New("invalid duration")
	}
}

const (
	OneMinute     Duration = "ONE_MINUTE"
	FiveMinute    Duration = "FIVE_MINUTE"
	ThreeMinute   Duration = "THREE_MINUTE"
	FifteenMinute Duration = "FIFTEEN_MINUTE"
	ThirtyMinute  Duration = "THIRTY_MINUTE"
	OneHour       Duration = "ONE_HOUR"
	TwoHour       Duration = "TWO_HOUR"
	FourHour      Duration = "FOUR_HOUR"
	SixHour       Duration = "SIX_HOUR"
	EightHour     Duration = "EIGHT_HOUR"
	TwelveHour    Duration = "TWELVE_HOUR"
	OneDay        Duration = "ONE_DAY"
	ThreeDay      Duration = "THREE_DAY"
	OneWeek       Duration = "ONE_WEEK"
	OneMonth      Duration = "ONE_MONTH"
)

type PoleData struct {
	LowPrice     decimal.Decimal `json:"low_price,omitempty"`
	HighPrice    decimal.Decimal `json:"high_price,omitempty"`
	ClosePrice   decimal.Decimal `json:"close_price,omitempty"`
	Volume       decimal.Decimal `json:"volume,omitempty"`
	TradingValue decimal.Decimal `json:"trading_value,omitempty"`
}

func NewPriceSet() *PoleData {
	return &PoleData{
		LowPrice:     decimal.NewFromFloat(math.MaxFloat64),
		HighPrice:    decimal.Zero,
		ClosePrice:   decimal.Zero,
		Volume:       decimal.Zero,
		TradingValue: decimal.Zero,
	}
}