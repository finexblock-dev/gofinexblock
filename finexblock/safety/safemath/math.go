package safemath

import (
	"github.com/shopspring/decimal"
	"math/big"
)

var (
	BitcoinDecimal  = decimal.NewFromFloat(100000000)
	EthereumDecimal = decimal.NewFromFloat(1000000000000000000)
)

func ParseWei(value *big.Float) float64 {
	f64, _ := value.Float64()
	return f64
}

func ParseBtcToSatoshi(value *big.Float) string {
	f64, _ := value.Float64()
	return decimal.NewFromFloat(f64).Mul(BitcoinDecimal).String()
}

func ParseWeiToEther(value *big.Float) string {
	f64, _ := value.Float64()
	return decimal.NewFromFloat(f64).Div(EthereumDecimal).String()
}
