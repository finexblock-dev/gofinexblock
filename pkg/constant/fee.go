package constant

import "github.com/shopspring/decimal"

var (
	FeeRatio        decimal.Decimal = decimal.Zero
	ReverseFeeRatio decimal.Decimal = decimal.NewFromFloat(1).Sub(FeeRatio)
)
