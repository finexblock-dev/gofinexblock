package structs

import "github.com/shopspring/decimal"

type (
	Asset struct {
		CoinID  uint            `json:"coinId" binding:"required" body:"coinId"`
		Balance decimal.Decimal `json:"balance" binding:"required" body:"balance"`
	}
)