package constant

import "github.com/shopspring/decimal"

const (
	BtcDecimal   = 100_000_000
	EthDecimal   = 1_000_000_000_000_000_000
	EtcDecimal   = 1_000_000_000_000_000_000
	MaticDecimal = 1_000_000_000_000_000_000
	LptDecimal   = 1_000_000_000_000_000_000
	ManaDecimal  = 1_000_000_000_000_000_000
	AxsDecimal   = 1_000_000_000_000_000_000
	AudioDecimal = 1_000_000_000_000_000_000
	SandDecimal  = 1_000_000_000_000_000_000
	CompDecimal  = 1_000_000_000_000_000_000
	LinkDecimal  = 1_000_000_000_000_000_000
	DydxDecimal  = 1_000_000_000_000_000_000
	EngDecimal   = 1_000_000_000_000_000_000
	OpDecimal    = 1_000_000_000_000_000_000
	AvaxDecimal  = 1_000_000_000_000_000_000
	ArbDecimal   = 1_000_000_000_000_000_000
)

var (
	BitcoinDecimal  = decimal.NewFromFloat(100_000_000)
	EthereumDecimal = decimal.NewFromFloat(1_000_000_000_000_000_000)
	PolygonDecimal  = decimal.NewFromFloat(1_000_000_000_000_000_000)
)

var (
	PolygonMinimumGatheringAmount  = decimal.NewFromFloat(40000000000000000)
	EthereumMinimumGatheringAmount = decimal.NewFromFloat(40000000000000000)
)
