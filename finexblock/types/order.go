package types

import "math"

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

type PoleData struct {
	LowPrice     float64 `json:"low_price,omitempty"`
	HighPrice    float64 `json:"high_price,omitempty"`
	ClosePrice   float64 `json:"close_price,omitempty"`
	Volume       float64 `json:"volume,omitempty"`
	TradingValue float64 `json:"trading_value,omitempty"`
}

func NewPriceSet() *PoleData {
	return &PoleData{
		LowPrice:     math.MaxFloat64,
		HighPrice:    0,
		ClosePrice:   0,
		Volume:       0,
		TradingValue: 0,
	}
}
