package utils

import (
	"github.com/finexblock-dev/gofinexblock/pkg/constant"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
	"github.com/shopspring/decimal"
)

var (
	SymbolToCurrency = map[grpc_order.SymbolType]grpc_order.Currency{
		grpc_order.SymbolType_BTCETH:   grpc_order.Currency_ETH,
		grpc_order.SymbolType_BTCETC:   grpc_order.Currency_ETC,
		grpc_order.SymbolType_BTCMATIC: grpc_order.Currency_MATIC,
		grpc_order.SymbolType_BTCLPT:   grpc_order.Currency_LPT,
		grpc_order.SymbolType_BTCMANA:  grpc_order.Currency_MANA,
		grpc_order.SymbolType_BTCAXS:   grpc_order.Currency_AXS,
		grpc_order.SymbolType_BTCAUDIO: grpc_order.Currency_AUDIO,
		grpc_order.SymbolType_BTCSAND:  grpc_order.Currency_SAND,
		grpc_order.SymbolType_BTCCOMP:  grpc_order.Currency_COMP,
		grpc_order.SymbolType_BTCLINK:  grpc_order.Currency_LINK,
		grpc_order.SymbolType_BTCDYDX:  grpc_order.Currency_DYDX,
		grpc_order.SymbolType_BTCBNB:   grpc_order.Currency_BNB,
		grpc_order.SymbolType_BTCOP:    grpc_order.Currency_OP,
		grpc_order.SymbolType_BTCAVAX:  grpc_order.Currency_AVAX,
		grpc_order.SymbolType_BTCARB:   grpc_order.Currency_ARB,
	}
	SymbolToCurrencyString = map[grpc_order.SymbolType]string{
		grpc_order.SymbolType_BTCETH:   grpc_order.Currency_ETH.String(),
		grpc_order.SymbolType_BTCETC:   grpc_order.Currency_ETC.String(),
		grpc_order.SymbolType_BTCMATIC: grpc_order.Currency_MATIC.String(),
		grpc_order.SymbolType_BTCLPT:   grpc_order.Currency_LPT.String(),
		grpc_order.SymbolType_BTCMANA:  grpc_order.Currency_MANA.String(),
		grpc_order.SymbolType_BTCAXS:   grpc_order.Currency_AXS.String(),
		grpc_order.SymbolType_BTCAUDIO: grpc_order.Currency_AUDIO.String(),
		grpc_order.SymbolType_BTCSAND:  grpc_order.Currency_SAND.String(),
		grpc_order.SymbolType_BTCCOMP:  grpc_order.Currency_COMP.String(),
		grpc_order.SymbolType_BTCLINK:  grpc_order.Currency_LINK.String(),
		grpc_order.SymbolType_BTCDYDX:  grpc_order.Currency_DYDX.String(),
		grpc_order.SymbolType_BTCBNB:   grpc_order.Currency_BNB.String(),
		grpc_order.SymbolType_BTCOP:    grpc_order.Currency_OP.String(),
		grpc_order.SymbolType_BTCAVAX:  grpc_order.Currency_AVAX.String(),
		grpc_order.SymbolType_BTCARB:   grpc_order.Currency_ARB.String(),
	}

	DecimalByCurrency = map[grpc_order.Currency]decimal.Decimal{
		grpc_order.Currency_ETH:   decimal.NewFromFloat(constant.EthDecimal),
		grpc_order.Currency_ETC:   decimal.NewFromFloat(constant.EtcDecimal),
		grpc_order.Currency_MATIC: decimal.NewFromFloat(constant.MaticDecimal),
		grpc_order.Currency_LPT:   decimal.NewFromFloat(constant.LptDecimal),
		grpc_order.Currency_MANA:  decimal.NewFromFloat(constant.ManaDecimal),
		grpc_order.Currency_AXS:   decimal.NewFromFloat(constant.AxsDecimal),
		grpc_order.Currency_AUDIO: decimal.NewFromFloat(constant.AudioDecimal),
		grpc_order.Currency_SAND:  decimal.NewFromFloat(constant.SandDecimal),
		grpc_order.Currency_COMP:  decimal.NewFromFloat(constant.CompDecimal),
		grpc_order.Currency_LINK:  decimal.NewFromFloat(constant.LinkDecimal),
		grpc_order.Currency_DYDX:  decimal.NewFromFloat(constant.DydxDecimal),
		grpc_order.Currency_BNB:   decimal.NewFromFloat(constant.EngDecimal),
		grpc_order.Currency_OP:    decimal.NewFromFloat(constant.OpDecimal),
		grpc_order.Currency_AVAX:  decimal.NewFromFloat(constant.AvaxDecimal),
		grpc_order.Currency_ARB:   decimal.NewFromFloat(constant.ArbDecimal),
		grpc_order.Currency_BTC:   decimal.NewFromFloat(constant.BtcDecimal),
	}
)

func OpponentCurrency(symbol grpc_order.SymbolType) grpc_order.Currency {
	return SymbolToCurrency[symbol]
}

func OpponentCurrencyToString(symbol grpc_order.SymbolType) string {
	return SymbolToCurrencyString[symbol]
}

func CoinDecimal(currency grpc_order.Currency) decimal.Decimal {
	return DecimalByCurrency[currency]
}