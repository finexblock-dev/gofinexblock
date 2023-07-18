package utils

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/shopspring/decimal"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrderStructs interface {
	*grpc_order.BidAsk | *grpc_order.Order | *grpc_order.BalanceUpdate |
		*grpc_order.OrderMatching | *grpc_order.OrderMatchingFailed |
		*grpc_order.OrderFulfillment | *grpc_order.OrderPartialFill |
		*grpc_order.OrderCancelled | *grpc_order.OrderInitialize |
		*grpc_order.OrderCancellationFailed | *grpc_order.OrderCancellation |
		*grpc_order.MarketOrderInput | *grpc_order.LimitOrderInput |
		*grpc_order.ErrorInput | *grpc_order.GetOrderBookOutput
}

func ParseLimitOrderInput(input *grpc_order.LimitOrderInput) *grpc_order.Order {
	return &grpc_order.Order{
		UserUUID:  input.GetUserUUID(),
		OrderUUID: input.GetOrderUUID(),
		Quantity:  input.GetQuantity(),
		UnitPrice: input.GetUnitPrice(),
		OrderType: input.GetOrderType(),
		Symbol:    input.GetSymbol(),
		MakeTime:  input.GetMakeTime(),
	}
}

func ParseMarketOrderInput(input *grpc_order.MarketOrderInput) *grpc_order.Order {
	return &grpc_order.Order{
		UserUUID:  input.GetUserUUID(),
		OrderUUID: input.GetOrderUUID(),
		Quantity:  input.GetQuantity(),
		OrderType: input.GetOrderType(),
		Symbol:    input.GetSymbol(),
		MakeTime:  input.GetMakeTime(),
	}
}

func NewOrderFulfillment(
	userUUID, orderUUID string,
	filledQuantity, unitPrice decimal.Decimal,
	symbol grpc_order.SymbolType,
	orderType grpc_order.OrderType,
	makeTime *timestamppb.Timestamp,
	fee *grpc_order.Fee,
) *grpc_order.OrderFulfillment {
	return &grpc_order.OrderFulfillment{
		UserUUID:       userUUID,
		OrderUUID:      orderUUID,
		FilledQuantity: filledQuantity.InexactFloat64(),
		UnitPrice:      unitPrice.InexactFloat64(),
		Symbol:         symbol,
		OrderType:      orderType,
		MakeTime:       makeTime,
		TakeTime:       timestamppb.Now(),
		Fee:            fee,
	}
}

func NewOrderPartialFill(
	userUUID, orderUUID string,
	totalQuantity, filledQuantity, unitPrice decimal.Decimal,
	symbol grpc_order.SymbolType,
	orderType grpc_order.OrderType,
	makeTime *timestamppb.Timestamp,
	fee *grpc_order.Fee,
) *grpc_order.OrderPartialFill {
	return &grpc_order.OrderPartialFill{
		UserUUID:       userUUID,
		OrderUUID:      orderUUID,
		TotalQuantity:  totalQuantity.InexactFloat64(),
		FilledQuantity: filledQuantity.InexactFloat64(),
		UnitPrice:      unitPrice.InexactFloat64(),
		Symbol:         symbol,
		OrderType:      orderType,
		MakeTime:       makeTime,
		TakeTime:       timestamppb.Now(),
		Fee:            fee,
	}
}

func NewBalanceUpdate(
	userUUID string,
	diff decimal.Decimal,
	currency grpc_order.Currency,
	reason grpc_order.Reason,
) *grpc_order.BalanceUpdate {
	return &grpc_order.BalanceUpdate{
		UserUUID: userUUID,
		Diff:     diff.InexactFloat64(),
		Currency: currency,
		Reason:   reason,
	}
}

func NewOrderMatchingFailed(
	userUUID, orderUUID string,
	quantity decimal.Decimal,
	orderType grpc_order.OrderType,
	symbol grpc_order.SymbolType,
	makeTime *timestamppb.Timestamp,
	msg string,
) *grpc_order.OrderMatchingFailed {
	return &grpc_order.OrderMatchingFailed{
		UserUUID:  userUUID,
		OrderUUID: orderUUID,
		Quantity:  quantity.InexactFloat64(),
		OrderType: orderType,
		Symbol:    symbol,
		MakeTime:  makeTime,
		Msg:       msg,
	}
}

func NewOrderMatching(
	unitPrice, quantity decimal.Decimal,
	orderType grpc_order.OrderType,
	symbol grpc_order.SymbolType,
) *grpc_order.OrderMatching {
	return &grpc_order.OrderMatching{
		UnitPrice: unitPrice.InexactFloat64(),
		Quantity:  quantity.InexactFloat64(),
		Timestamp: timestamppb.Now(),
		OrderType: orderType,
		Symbol:    symbol,
	}
}

func NewOrderCancelled(
	userUUID, orderUUID string,
	quantity, unitPrice decimal.Decimal,
	orderType grpc_order.OrderType,
	symbol grpc_order.SymbolType,
) *grpc_order.OrderCancelled {
	return &grpc_order.OrderCancelled{
		UserUUID:  userUUID,
		OrderUUID: orderUUID,
		Quantity:  quantity.InexactFloat64(),
		UnitPrice: unitPrice.InexactFloat64(),
		OrderType: orderType,
		Symbol:    symbol,
	}
}

func NewOrderInitialized(userUUID, orderUUID string, quantity, unitPrice decimal.Decimal, orderType grpc_order.OrderType, symbol grpc_order.SymbolType) *grpc_order.OrderInitialize {
	return &grpc_order.OrderInitialize{
		UserUUID:  userUUID,
		OrderUUID: orderUUID,
		Quantity:  quantity.InexactFloat64(),
		UnitPrice: unitPrice.InexactFloat64(),
		Symbol:    symbol,
		OrderType: orderType,
		MakeTime:  timestamppb.Now(),
	}
}

func NewMarketOrderMatching(
	userUUID, orderUUID string,
	quantity, unitPrice decimal.Decimal,
	orderType grpc_order.OrderType,
	symbol grpc_order.SymbolType,
	makeTime *timestamppb.Timestamp,
	fee *grpc_order.Fee,
) *grpc_order.MarketOrderMatching {
	return &grpc_order.MarketOrderMatching{
		UserUUID:  userUUID,
		OrderUUID: orderUUID,
		Quantity:  quantity.InexactFloat64(),
		UnitPrice: unitPrice.InexactFloat64(),
		Symbol:    symbol,
		OrderType: orderType,
		MakeTime:  makeTime,
		TakeTime:  timestamppb.Now(),
		Fee:       fee,
	}
}

func NewOrderPlacement(userUUID, orderUUID string, quantity, unitPrice decimal.Decimal, orderType grpc_order.OrderType, symbol grpc_order.SymbolType) *grpc_order.OrderPlacement {
	return &grpc_order.OrderPlacement{
		UserUUID:  userUUID,
		OrderUUID: orderUUID,
		Quantity:  quantity.InexactFloat64(),
		UnitPrice: unitPrice.InexactFloat64(),
		Symbol:    symbol,
		OrderType: orderType,
		MakeTime:  timestamppb.Now(),
	}
}