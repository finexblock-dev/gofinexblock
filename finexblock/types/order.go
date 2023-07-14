package types

import (
	"github.com/shopspring/decimal"
	"math"
)

type Reason string

const (
	Cancel Reason = "CANCEL"
	Fill   Reason = "FILL"
	Place  Reason = "PLACE"
)

type OrderType string

const (
	Bid OrderType = "BID"
	Ask OrderType = "ASK"
)

type OrderStatus string

const (
	Cancelled     OrderStatus = "CANCELLED"
	Placed        OrderStatus = "PLACED"
	Fulfilled     OrderStatus = "FULFILLED"
	PartialFilled OrderStatus = "PARTIAL_FILLED"
)

type Duration string

const (
	OneMinute     Duration = "ONE_MINUTE"
	FiveMinute    Duration = "FIVE_MINUTE"
	ThreeMinute            = "THREE_MINUTE"
	FifteenMinute          = "FIFTEEN_MINUTE"
	ThirtyMinute           = "THIRTY_MINUTE"
	OneHour                = "ONE_HOUR"
	TwoHour                = "TWO_HOUR"
	FourHour               = "FOUR_HOUR"
	SixHour                = "SIX_HOUR"
	EightHour              = "EIGHT_HOUR"
	TwelveHour             = "TWELVE_HOUR"
	OneDay                 = "ONE_DAY"
	ThreeDay               = "THREE_DAY"
	OneWeek                = "ONE_WEEK"
	OneMonth               = "ONE_MONTH"
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
