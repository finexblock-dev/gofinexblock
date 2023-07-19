package match

import (
	"context"
	"github.com/finexblock-dev/gofinexblock/finexblock/constant"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/utils"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (e *engine) LimitAskBigger(pair *grpc_order.BidAsk) (err error) {
	var bidQuantity, mul, filledQuantity, unitPrice decimal.Decimal
	var askFulfillment *grpc_order.OrderFulfillment
	var bidPartialFill *grpc_order.OrderPartialFill
	var orderMatching *grpc_order.OrderMatching
	var bidIncrease *grpc_order.BalanceUpdate
	var askIncrease *grpc_order.BalanceUpdate
	var fee *grpc_order.Fee
	var prefix = "LimitAskBigger"
	var bid, ask = pair.Bid, pair.Ask
	var currency grpc_order.Currency
	var tx redis.Pipeliner
	var ctx context.Context

	bidQuantity = decimal.NewFromFloat(bid.Quantity)

	currency = utils.OpponentCurrency(ask.Symbol)
	mul = utils.CoinDecimal(currency)

	// filled quantity = ask.Quantity
	filledQuantity = decimal.NewFromFloat(ask.Quantity)

	// unit price = bid.UnitPrice
	unitPrice = decimal.NewFromFloat(bid.UnitPrice)

	// fee = filled quantity * unit price * fee ratio
	fee = &grpc_order.Fee{Amount: filledQuantity.Mul(unitPrice).Mul(constant.FeeRatio).InexactFloat64()}

	// ask fulfillment
	askFulfillment = utils.NewOrderFulfillment(ask.UserUUID, ask.OrderUUID, filledQuantity, unitPrice, ask.Symbol, ask.OrderType, ask.MakeTime, fee)

	// bid partial fill
	bidPartialFill = utils.NewOrderPartialFill(bid.UserUUID, bid.OrderUUID, bidQuantity, filledQuantity, unitPrice, bid.Symbol, bid.OrderType, bid.MakeTime, fee)

	// matched order
	orderMatching = utils.NewOrderMatching(unitPrice, filledQuantity, ask.OrderType, ask.Symbol)

	// bid increase
	bidIncrease = utils.NewBalanceUpdate(bid.UserUUID, filledQuantity.Mul(constant.ReverseFeeRatio).Mul(mul), currency, grpc_order.Reason_MAKE)

	// ask increase
	askIncrease = utils.NewBalanceUpdate(ask.UserUUID, filledQuantity.Mul(unitPrice).Mul(constant.ReverseFeeRatio), grpc_order.Currency_BTC, grpc_order.Reason_TAKE)

	// Send stream
	tx = e.tradeManager.Pipeliner()
	ctx = context.Background()

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, bidIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, askIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderMatchingStreamPipeline(tx, ctx, orderMatching); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order matching stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderFulfillmentStreamPipeline(tx, ctx, askFulfillment); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order fulfillment stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderPartialFillStreamPipeline(tx, ctx, bidPartialFill); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order partial fill stream, %v", prefix, err)
	}

	_ = e.tradeManager.DeleteOrder(ask.OrderUUID)

	if _, err = tx.Exec(ctx); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to execute pipeline, %v", prefix, err)
	}

	return nil
}

func (e *engine) LimitAskEqual(pair *grpc_order.BidAsk) (err error) {
	var mul, filledQuantity, unitPrice decimal.Decimal
	var askFulfillment *grpc_order.OrderFulfillment
	var bidFulfillment *grpc_order.OrderFulfillment
	var orderMatching *grpc_order.OrderMatching
	var bidIncrease *grpc_order.BalanceUpdate
	var askIncrease *grpc_order.BalanceUpdate
	var fee *grpc_order.Fee
	var prefix = "LimitAskEqual"
	var bid, ask = pair.Bid, pair.Ask
	var currency grpc_order.Currency
	var tx redis.Pipeliner
	var ctx context.Context

	currency = utils.OpponentCurrency(ask.Symbol)
	mul = utils.CoinDecimal(currency)

	// filled quantity = ask.Quantity
	filledQuantity = decimal.NewFromFloat(ask.Quantity)

	// unit price = bid.UnitPrice
	unitPrice = decimal.NewFromFloat(bid.UnitPrice)

	// fee = filled quantity * unit price * fee ratio
	fee = &grpc_order.Fee{Amount: filledQuantity.Mul(unitPrice).Mul(constant.FeeRatio).InexactFloat64()}

	// ask fulfillment
	askFulfillment = utils.NewOrderFulfillment(ask.UserUUID, ask.OrderUUID, filledQuantity, unitPrice, ask.Symbol, ask.OrderType, ask.MakeTime, fee)

	// bid fulfillment
	bidFulfillment = utils.NewOrderFulfillment(bid.UserUUID, bid.OrderUUID, filledQuantity, unitPrice, bid.Symbol, bid.OrderType, bid.MakeTime, fee)

	// matched order
	orderMatching = utils.NewOrderMatching(unitPrice, filledQuantity, ask.OrderType, ask.Symbol)

	// bid increase
	bidIncrease = utils.NewBalanceUpdate(bid.UserUUID, filledQuantity.Mul(constant.ReverseFeeRatio).Mul(mul), currency, grpc_order.Reason_MAKE)

	// ask increase
	askIncrease = utils.NewBalanceUpdate(ask.UserUUID, filledQuantity.Mul(unitPrice).Mul(constant.ReverseFeeRatio), grpc_order.Currency_BTC, grpc_order.Reason_TAKE)

	// Send stream
	tx = e.tradeManager.Pipeliner()
	ctx = context.Background()

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, bidIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, askIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderMatchingStreamPipeline(tx, ctx, orderMatching); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order matching stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderFulfillmentStreamPipeline(tx, ctx, askFulfillment); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order fulfillment stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderFulfillmentStreamPipeline(tx, ctx, bidFulfillment); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order fulfillment stream, %v", prefix, err)
	}

	if _, err = tx.Exec(ctx); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to execute pipeline, %v", prefix, err)
	}

	_ = e.tradeManager.DeleteOrder(ask.OrderUUID)
	_ = e.tradeManager.DeleteOrder(bid.OrderUUID)

	return nil
}

func (e *engine) LimitAskSmaller(pair *grpc_order.BidAsk) (err error) {
	var askQuantity, mul, filledQuantity, unitPrice decimal.Decimal
	var bidFulfillment *grpc_order.OrderFulfillment
	var askPartialFill *grpc_order.OrderPartialFill
	var orderMatching *grpc_order.OrderMatching
	var bidIncrease *grpc_order.BalanceUpdate
	var askIncrease *grpc_order.BalanceUpdate
	var fee *grpc_order.Fee
	var prefix = "LimitAskSmaller"
	var bid, ask = pair.Bid, pair.Ask
	var currency grpc_order.Currency
	var tx redis.Pipeliner
	var ctx context.Context

	askQuantity = decimal.NewFromFloat(ask.Quantity)

	currency = utils.OpponentCurrency(ask.Symbol)
	mul = utils.CoinDecimal(currency)

	// filled quantity = bid.Quantity
	filledQuantity = decimal.NewFromFloat(bid.Quantity)

	// unit price = bid.UnitPrice
	unitPrice = decimal.NewFromFloat(bid.UnitPrice)

	// fee = filled quantity * unit price * fee ratio
	fee = &grpc_order.Fee{Amount: filledQuantity.Mul(unitPrice).Mul(constant.FeeRatio).InexactFloat64()}

	// ask partial fill
	askPartialFill = utils.NewOrderPartialFill(ask.UserUUID, ask.OrderUUID, askQuantity, filledQuantity, unitPrice, ask.Symbol, ask.OrderType, ask.MakeTime, fee)

	// bid fulfillment
	bidFulfillment = utils.NewOrderFulfillment(bid.UserUUID, bid.OrderUUID, filledQuantity, unitPrice, bid.Symbol, bid.OrderType, bid.MakeTime, fee)

	// matched order
	orderMatching = utils.NewOrderMatching(unitPrice, filledQuantity, ask.OrderType, ask.Symbol)

	// bid increase
	bidIncrease = utils.NewBalanceUpdate(bid.UserUUID, filledQuantity.Mul(constant.ReverseFeeRatio).Mul(mul), currency, grpc_order.Reason_MAKE)

	// ask increase
	askIncrease = utils.NewBalanceUpdate(ask.UserUUID, filledQuantity.Mul(unitPrice).Mul(constant.ReverseFeeRatio), grpc_order.Currency_BTC, grpc_order.Reason_TAKE)

	// Send stream
	tx = e.tradeManager.Pipeliner()
	ctx = context.Background()

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, bidIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, askIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderMatchingStreamPipeline(tx, ctx, orderMatching); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order matching stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderPartialFillStreamPipeline(tx, ctx, askPartialFill); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order fulfillment stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderFulfillmentStreamPipeline(tx, ctx, bidFulfillment); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order fulfillment stream, %v", prefix, err)
	}

	_ = e.tradeManager.DeleteOrder(bid.OrderUUID)

	if _, err = tx.Exec(ctx); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to execute pipeline, %v", prefix, err)
	}

	return nil
}

func (e *engine) LimitBidBigger(pair *grpc_order.BidAsk) (err error) {
	var totalQuantity, mul, filledQuantity, unitPrice decimal.Decimal
	var bidFulfillment *grpc_order.OrderFulfillment
	var askPartialFill *grpc_order.OrderPartialFill
	var orderMatching *grpc_order.OrderMatching
	var bidIncrease *grpc_order.BalanceUpdate
	var askIncrease *grpc_order.BalanceUpdate
	var fee *grpc_order.Fee
	var prefix = "LimitBidBigger"
	var currency grpc_order.Currency
	var tx redis.Pipeliner
	var ctx context.Context

	var bid, ask = pair.Bid, pair.Ask

	currency = utils.OpponentCurrency(bid.Symbol)
	mul = utils.CoinDecimal(currency)

	// filled quantity = bid.Quantity
	filledQuantity = decimal.NewFromFloat(bid.Quantity)
	totalQuantity = decimal.NewFromFloat(ask.Quantity)

	// unit price = ask.UnitPrice
	unitPrice = decimal.NewFromFloat(ask.UnitPrice)

	// fee = filled quantity * unit price * fee ratio
	fee = &grpc_order.Fee{Amount: filledQuantity.Mul(unitPrice).Mul(constant.FeeRatio).InexactFloat64()}

	// bid fulfillment
	bidFulfillment = utils.NewOrderFulfillment(bid.UserUUID, bid.OrderUUID, filledQuantity, unitPrice, bid.Symbol, bid.OrderType, bid.MakeTime, fee)

	// ask partial fill
	askPartialFill = utils.NewOrderPartialFill(ask.UserUUID, ask.OrderUUID, totalQuantity, filledQuantity, unitPrice, ask.Symbol, ask.OrderType, ask.MakeTime, fee)

	// matched order
	orderMatching = utils.NewOrderMatching(unitPrice, filledQuantity, bid.OrderType, bid.Symbol)

	// bid increase
	bidIncrease = utils.NewBalanceUpdate(bid.UserUUID, filledQuantity.Mul(constant.ReverseFeeRatio).Mul(mul), currency, grpc_order.Reason_TAKE)

	// ask increase
	askIncrease = utils.NewBalanceUpdate(ask.UserUUID, filledQuantity.Mul(unitPrice).Mul(constant.ReverseFeeRatio), grpc_order.Currency_BTC, grpc_order.Reason_MAKE)

	// Send stream
	tx = e.tradeManager.Pipeliner()
	ctx = context.Background()

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, bidIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, askIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderMatchingStreamPipeline(tx, ctx, orderMatching); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order matching stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderPartialFillStreamPipeline(tx, ctx, askPartialFill); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order fulfillment stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderFulfillmentStreamPipeline(tx, ctx, bidFulfillment); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order fulfillment stream, %v", prefix, err)
	}

	_ = e.tradeManager.DeleteOrder(bid.OrderUUID)

	if _, err = tx.Exec(ctx); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to execute pipeline, %v", prefix, err)
	}

	return nil
}

func (e *engine) LimitBidEqual(pair *grpc_order.BidAsk) (err error) {
	var mul, filledQuantity, unitPrice decimal.Decimal
	var bidFulfillment *grpc_order.OrderFulfillment
	var askFulfillment *grpc_order.OrderFulfillment
	var orderMatching *grpc_order.OrderMatching
	var bidIncrease *grpc_order.BalanceUpdate
	var askIncrease *grpc_order.BalanceUpdate
	var fee *grpc_order.Fee
	var prefix = "LimitBidEqual"
	var currency grpc_order.Currency
	var tx redis.Pipeliner
	var ctx context.Context

	var bid, ask = pair.Bid, pair.Ask

	currency = utils.OpponentCurrency(bid.Symbol)
	mul = utils.CoinDecimal(currency)

	// filled quantity = bid.Quantity
	filledQuantity = decimal.NewFromFloat(bid.Quantity)

	// unit price = ask.UnitPrice
	unitPrice = decimal.NewFromFloat(ask.UnitPrice)

	// fee = filled quantity * unit price * fee ratio
	fee = &grpc_order.Fee{Amount: filledQuantity.Mul(unitPrice).Mul(constant.FeeRatio).InexactFloat64()}

	// bid fulfillment
	bidFulfillment = utils.NewOrderFulfillment(bid.UserUUID, bid.OrderUUID, filledQuantity, unitPrice, bid.Symbol, bid.OrderType, bid.MakeTime, fee)

	// ask fulfillment
	askFulfillment = utils.NewOrderFulfillment(ask.UserUUID, ask.OrderUUID, filledQuantity, unitPrice, ask.Symbol, ask.OrderType, ask.MakeTime, fee)

	// matched order
	orderMatching = utils.NewOrderMatching(unitPrice, filledQuantity, bid.OrderType, bid.Symbol)

	// bid increase
	bidIncrease = utils.NewBalanceUpdate(bid.UserUUID, filledQuantity.Mul(constant.ReverseFeeRatio).Mul(mul), currency, grpc_order.Reason_TAKE)

	// ask increase
	askIncrease = utils.NewBalanceUpdate(ask.UserUUID, filledQuantity.Mul(unitPrice).Mul(constant.ReverseFeeRatio), grpc_order.Currency_BTC, grpc_order.Reason_MAKE)

	// Send stream
	tx = e.tradeManager.Pipeliner()
	ctx = context.Background()

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, bidIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, askIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderMatchingStreamPipeline(tx, ctx, orderMatching); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order matching stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderFulfillmentStreamPipeline(tx, ctx, bidFulfillment); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order fulfillment stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderFulfillmentStreamPipeline(tx, ctx, askFulfillment); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order fulfillment stream, %v", prefix, err)
	}

	if _, err = tx.Exec(ctx); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to execute pipeline, %v", prefix, err)
	}

	_ = e.tradeManager.DeleteOrder(ask.OrderUUID)
	_ = e.tradeManager.DeleteOrder(bid.OrderUUID)

	return nil
}

func (e *engine) LimitBidSmaller(pair *grpc_order.BidAsk) (err error) {
	var totalQuantity, mul, filledQuantity, unitPrice decimal.Decimal
	var bidPartialFill *grpc_order.OrderPartialFill
	var askFulfillment *grpc_order.OrderFulfillment
	var orderMatching *grpc_order.OrderMatching
	var bidIncrease *grpc_order.BalanceUpdate
	var askIncrease *grpc_order.BalanceUpdate
	var fee *grpc_order.Fee
	var prefix = "LimitBidSmaller"
	var currency grpc_order.Currency
	var tx redis.Pipeliner
	var ctx context.Context

	var bid, ask = pair.Bid, pair.Ask

	currency = utils.OpponentCurrency(bid.Symbol)
	mul = utils.CoinDecimal(currency)

	// filled quantity = ask.Quantity
	filledQuantity = decimal.NewFromFloat(ask.Quantity)

	// total quantity = bid.Quantity
	totalQuantity = decimal.NewFromFloat(bid.Quantity)

	// unit price = ask.UnitPrice
	unitPrice = decimal.NewFromFloat(ask.UnitPrice)

	// fee = filled quantity * unit price * fee ratio
	fee = &grpc_order.Fee{Amount: filledQuantity.Mul(unitPrice).Mul(constant.FeeRatio).InexactFloat64()}

	// bid partial fill
	bidPartialFill = utils.NewOrderPartialFill(bid.UserUUID, bid.OrderUUID, totalQuantity, filledQuantity, unitPrice, bid.Symbol, bid.OrderType, bid.MakeTime, fee)

	// ask fulfillment
	askFulfillment = utils.NewOrderFulfillment(ask.UserUUID, ask.OrderUUID, filledQuantity, unitPrice, ask.Symbol, ask.OrderType, ask.MakeTime, fee)

	// matched order
	orderMatching = utils.NewOrderMatching(unitPrice, filledQuantity, bid.OrderType, bid.Symbol)

	// bid increase
	bidIncrease = utils.NewBalanceUpdate(bid.UserUUID, filledQuantity.Mul(constant.ReverseFeeRatio).Mul(mul), currency, grpc_order.Reason_TAKE)

	// ask increase
	askIncrease = utils.NewBalanceUpdate(ask.UserUUID, filledQuantity.Mul(unitPrice).Mul(constant.ReverseFeeRatio), grpc_order.Currency_BTC, grpc_order.Reason_MAKE)

	// Send stream
	tx = e.tradeManager.Pipeliner()
	ctx = context.Background()

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, bidIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, askIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderMatchingStreamPipeline(tx, ctx, orderMatching); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order matching stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderPartialFillStreamPipeline(tx, ctx, bidPartialFill); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order partial fill stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderFulfillmentStreamPipeline(tx, ctx, askFulfillment); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order fulfillment stream, %v", prefix, err)
	}

	_ = e.tradeManager.DeleteOrder(ask.OrderUUID)

	if _, err = tx.Exec(ctx); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to execute pipeline, %v", prefix, err)
	}

	return nil
}

func (e *engine) MarketAskBigger(pair *grpc_order.BidAsk) (err error) {
	var bidQuantity, askQuantity, mul, filledQuantity, unitPrice decimal.Decimal
	var bidPartialFill *grpc_order.OrderPartialFill
	var askFulfillment *grpc_order.MarketOrderMatching
	var orderMatching *grpc_order.OrderMatching
	var bidIncrease *grpc_order.BalanceUpdate
	var askIncrease *grpc_order.BalanceUpdate
	var fee *grpc_order.Fee
	var prefix = "MarketAskBigger"
	var currency grpc_order.Currency
	var tx redis.Pipeliner
	var ctx context.Context

	var bid, ask = pair.Bid, pair.Ask

	currency = utils.OpponentCurrency(bid.Symbol)
	mul = utils.CoinDecimal(currency)

	bidQuantity = decimal.NewFromFloat(bid.Quantity)
	askQuantity = decimal.NewFromFloat(ask.Quantity)

	// filled quantity = ask.Quantity / mul
	filledQuantity = askQuantity.Div(mul)

	// unit price = bid.UnitPrice
	unitPrice = decimal.NewFromFloat(bid.UnitPrice)

	// fee = filled quantity * unit price * fee ratio
	fee = &grpc_order.Fee{Amount: filledQuantity.Mul(unitPrice).Mul(constant.FeeRatio).InexactFloat64()}

	// bid partial fill
	bidPartialFill = utils.NewOrderPartialFill(bid.UserUUID, bid.OrderUUID, bidQuantity, filledQuantity, unitPrice, bid.Symbol, bid.OrderType, bid.MakeTime, fee)

	// ask fulfillment
	askFulfillment = utils.NewMarketOrderMatching(ask.UserUUID, ask.OrderUUID, filledQuantity, unitPrice, ask.OrderType, ask.Symbol, ask.MakeTime, fee)

	// matched order
	orderMatching = utils.NewOrderMatching(unitPrice, filledQuantity, ask.OrderType, ask.Symbol)

	// bid increase
	bidIncrease = utils.NewBalanceUpdate(bid.UserUUID, constant.ReverseFeeRatio.Mul(filledQuantity).Mul(unitPrice), currency, grpc_order.Reason_MAKE)

	// ask increase
	askIncrease = utils.NewBalanceUpdate(ask.UserUUID, filledQuantity.Mul(unitPrice).Mul(constant.ReverseFeeRatio), grpc_order.Currency_BTC, grpc_order.Reason_TAKE)

	// Send stream

	tx = e.tradeManager.Pipeliner()
	ctx = context.Background()

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, bidIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, askIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderMatchingStreamPipeline(tx, ctx, orderMatching); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order matching stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderPartialFillStreamPipeline(tx, ctx, bidPartialFill); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order partial fill stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendMarketOrderMatchingStreamPipeline(tx, ctx, askFulfillment); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order fulfillment stream, %v", prefix, err)
	}

	if _, err = tx.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (e *engine) MarketAskEqual(pair *grpc_order.BidAsk) (err error) {
	var bidQuantity, askQuantity, filledQuantity, unitPrice decimal.Decimal
	var bidFulfillment *grpc_order.OrderFulfillment
	var askFulfillment *grpc_order.MarketOrderMatching
	var orderMatching *grpc_order.OrderMatching
	var bidIncrease *grpc_order.BalanceUpdate
	var askIncrease *grpc_order.BalanceUpdate
	var fee *grpc_order.Fee
	var prefix = "MarketAskEqual"
	var currency grpc_order.Currency
	var tx redis.Pipeliner
	var ctx context.Context

	var bid, ask = pair.Bid, pair.Ask

	currency = utils.OpponentCurrency(bid.Symbol)

	bidQuantity = decimal.NewFromFloat(bid.Quantity)
	askQuantity = decimal.NewFromFloat(ask.Quantity)

	// filled quantity = bid.Quantity
	filledQuantity = bidQuantity

	// unit price = bid.UnitPrice
	unitPrice = decimal.NewFromFloat(bid.UnitPrice)

	// fee = filled quantity * unit price * fee ratio
	fee = &grpc_order.Fee{Amount: filledQuantity.Mul(unitPrice).Mul(constant.FeeRatio).InexactFloat64()}

	// bid fulfillment
	bidFulfillment = utils.NewOrderFulfillment(bid.UserUUID, bid.OrderUUID, filledQuantity, unitPrice, bid.Symbol, bid.OrderType, bid.MakeTime, fee)

	// ask fulfillment
	askFulfillment = utils.NewMarketOrderMatching(ask.UserUUID, ask.OrderUUID, filledQuantity, unitPrice, ask.OrderType, ask.Symbol, ask.MakeTime, fee)

	// matched order
	orderMatching = utils.NewOrderMatching(unitPrice, filledQuantity, ask.OrderType, ask.Symbol)

	// bid increase
	bidIncrease = utils.NewBalanceUpdate(bid.UserUUID, constant.ReverseFeeRatio.Mul(askQuantity), currency, grpc_order.Reason_MAKE)

	// ask increase
	askIncrease = utils.NewBalanceUpdate(ask.UserUUID, filledQuantity.Mul(unitPrice).Mul(constant.ReverseFeeRatio), grpc_order.Currency_BTC, grpc_order.Reason_TAKE)

	// Send stream
	tx = e.tradeManager.Pipeliner()
	ctx = context.Background()

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, bidIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, askIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderMatchingStreamPipeline(tx, ctx, orderMatching); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order matching stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderFulfillmentStreamPipeline(tx, ctx, bidFulfillment); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order fulfillment stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendMarketOrderMatchingStreamPipeline(tx, ctx, askFulfillment); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order fulfillment stream, %v", prefix, err)
	}

	_ = e.tradeManager.DeleteOrder(bid.OrderUUID)

	if _, err = tx.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (e *engine) MarketAskSmaller(pair *grpc_order.BidAsk) (err error) {
	var bidQuantity, filledQuantity, unitPrice decimal.Decimal
	var bidFulfillment *grpc_order.OrderFulfillment
	var askPartialFill *grpc_order.MarketOrderMatching
	var orderMatching *grpc_order.OrderMatching
	var bidIncrease *grpc_order.BalanceUpdate
	var askIncrease *grpc_order.BalanceUpdate
	var fee *grpc_order.Fee
	var prefix = "MarketAskSmaller"
	var currency grpc_order.Currency
	var tx redis.Pipeliner
	var ctx context.Context

	var bid, ask = pair.Bid, pair.Ask

	currency = utils.OpponentCurrency(bid.Symbol)

	bidQuantity = decimal.NewFromFloat(bid.Quantity)

	// filled quantity = bid.Quantity
	filledQuantity = bidQuantity

	// unit price = bid.UnitPrice
	unitPrice = decimal.NewFromFloat(bid.UnitPrice)

	// fee = filled quantity * unit price * fee ratio
	fee = &grpc_order.Fee{Amount: filledQuantity.Mul(unitPrice).Mul(constant.FeeRatio).InexactFloat64()}

	// bid fulfillment
	bidFulfillment = utils.NewOrderFulfillment(bid.UserUUID, bid.OrderUUID, filledQuantity, unitPrice, bid.Symbol, bid.OrderType, bid.MakeTime, fee)

	// ask partial fill
	askPartialFill = utils.NewMarketOrderMatching(ask.UserUUID, ask.OrderUUID, filledQuantity, unitPrice, ask.OrderType, ask.Symbol, ask.MakeTime, fee)

	// matched order
	orderMatching = utils.NewOrderMatching(unitPrice, filledQuantity, ask.OrderType, ask.Symbol)

	// bid increase
	bidIncrease = utils.NewBalanceUpdate(bid.UserUUID, filledQuantity.Mul(unitPrice).Mul(constant.ReverseFeeRatio), currency, grpc_order.Reason_MAKE)

	// ask increase
	askIncrease = utils.NewBalanceUpdate(ask.UserUUID, filledQuantity.Mul(constant.ReverseFeeRatio), grpc_order.Currency_BTC, grpc_order.Reason_TAKE)

	// Send stream
	tx = e.tradeManager.Pipeliner()
	ctx = context.Background()

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, bidIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, askIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderMatchingStreamPipeline(tx, ctx, orderMatching); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order matching stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderFulfillmentStreamPipeline(tx, ctx, bidFulfillment); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order fulfillment stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendMarketOrderMatchingStreamPipeline(tx, ctx, askPartialFill); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order partial fill stream, %v", prefix, err)
	}

	_ = e.tradeManager.DeleteOrder(bid.OrderUUID)

	if _, err = tx.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (e *engine) MarketBidBigger(pair *grpc_order.BidAsk) (err error) {
	var mul, bidQuantity, askQuantity, filledQuantity, unitPrice decimal.Decimal
	var bidFulfillment *grpc_order.MarketOrderMatching
	var askPartialFill *grpc_order.OrderPartialFill
	var orderMatching *grpc_order.OrderMatching
	var bidIncrease *grpc_order.BalanceUpdate
	var askIncrease *grpc_order.BalanceUpdate
	var fee *grpc_order.Fee
	var prefix = "MarketBidBigger"
	var currency grpc_order.Currency
	var tx redis.Pipeliner
	var ctx context.Context

	var bid, ask = pair.Bid, pair.Ask

	currency = utils.OpponentCurrency(bid.Symbol)
	mul = utils.CoinDecimal(currency)

	bidQuantity = decimal.NewFromFloat(bid.Quantity)
	askQuantity = decimal.NewFromFloat(ask.Quantity)

	// unit price = ask.UnitPrice
	unitPrice = decimal.NewFromFloat(ask.UnitPrice)

	// filled quantity = bid.Quantity / ask.UnitPrice
	filledQuantity = bidQuantity.Div(unitPrice)

	// fee = filled quantity * unit price * fee ratio
	fee = &grpc_order.Fee{Amount: bidQuantity.Mul(unitPrice).Mul(constant.FeeRatio).InexactFloat64()}

	// bid fulfillment
	bidFulfillment = utils.NewMarketOrderMatching(bid.UserUUID, bid.OrderUUID, filledQuantity, unitPrice, bid.OrderType, bid.Symbol, bid.MakeTime, fee)

	// ask partial fill
	askPartialFill = utils.NewOrderPartialFill(ask.UserUUID, ask.OrderUUID, askQuantity, filledQuantity, unitPrice, ask.Symbol, ask.OrderType, ask.MakeTime, fee)

	// matched order
	orderMatching = utils.NewOrderMatching(unitPrice, filledQuantity, bid.OrderType, bid.Symbol)

	// bid increase
	bidIncrease = utils.NewBalanceUpdate(bid.UserUUID, filledQuantity.Mul(mul).Mul(constant.ReverseFeeRatio), currency, grpc_order.Reason_TAKE)

	// ask increase
	askIncrease = utils.NewBalanceUpdate(ask.UserUUID, bidQuantity.Mul(constant.ReverseFeeRatio), grpc_order.Currency_BTC, grpc_order.Reason_MAKE)

	// Send stream
	tx = e.tradeManager.Pipeliner()
	ctx = context.Background()

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, bidIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, askIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderMatchingStreamPipeline(tx, ctx, orderMatching); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order matching stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendMarketOrderMatchingStreamPipeline(tx, ctx, bidFulfillment); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order fulfillment stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderPartialFillStreamPipeline(tx, ctx, askPartialFill); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order partial fill stream, %v", prefix, err)
	}

	if _, err = tx.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (e *engine) MarketBidEqual(pair *grpc_order.BidAsk) (err error) {
	var mul, bidQuantity, askQuantity, filledQuantity, unitPrice decimal.Decimal
	var bidFulfillment *grpc_order.MarketOrderMatching
	var askFulfillment *grpc_order.OrderFulfillment
	var orderMatching *grpc_order.OrderMatching
	var bidIncrease *grpc_order.BalanceUpdate
	var askIncrease *grpc_order.BalanceUpdate
	var fee *grpc_order.Fee
	var prefix = "MarketBidEqual"
	var currency grpc_order.Currency
	var tx redis.Pipeliner
	var ctx context.Context

	var bid, ask = pair.Bid, pair.Ask

	currency = utils.OpponentCurrency(bid.Symbol)
	mul = utils.CoinDecimal(currency)

	bidQuantity = decimal.NewFromFloat(bid.Quantity)
	askQuantity = decimal.NewFromFloat(ask.Quantity)

	// unit price = ask.UnitPrice
	unitPrice = decimal.NewFromFloat(ask.UnitPrice)

	// filled quantity = ask.Quantity
	filledQuantity = askQuantity

	// fee = filled quantity * unit price * fee ratio
	fee = &grpc_order.Fee{Amount: filledQuantity.Mul(unitPrice).Mul(constant.FeeRatio).InexactFloat64()}

	// bid fulfillment
	bidFulfillment = utils.NewMarketOrderMatching(bid.UserUUID, bid.OrderUUID, filledQuantity, unitPrice, bid.OrderType, bid.Symbol, bid.MakeTime, fee)

	// ask fulfillment
	askFulfillment = utils.NewOrderFulfillment(ask.UserUUID, ask.OrderUUID, filledQuantity, unitPrice, ask.Symbol, ask.OrderType, ask.MakeTime, fee)

	// matched order
	orderMatching = utils.NewOrderMatching(unitPrice, filledQuantity, bid.OrderType, bid.Symbol)

	// bid increase
	bidIncrease = utils.NewBalanceUpdate(bid.UserUUID, filledQuantity.Mul(mul).Mul(constant.ReverseFeeRatio), currency, grpc_order.Reason_TAKE)

	// ask increase
	askIncrease = utils.NewBalanceUpdate(ask.UserUUID, bidQuantity.Mul(constant.ReverseFeeRatio), grpc_order.Currency_BTC, grpc_order.Reason_MAKE)

	// Send stream
	tx = e.tradeManager.Pipeliner()
	ctx = context.Background()

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, bidIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, askIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderMatchingStreamPipeline(tx, ctx, orderMatching); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order matching stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendMarketOrderMatchingStreamPipeline(tx, ctx, bidFulfillment); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order fulfillment stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderFulfillmentStreamPipeline(tx, ctx, askFulfillment); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order fulfillment stream, %v", prefix, err)
	}

	_ = e.tradeManager.DeleteOrder(ask.OrderUUID)

	if _, err = tx.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (e *engine) MarketBidSmaller(pair *grpc_order.BidAsk) (err error) {
	var mul, askQuantity, filledQuantity, unitPrice decimal.Decimal
	var bidPartialFill *grpc_order.MarketOrderMatching
	var askFulfillment *grpc_order.OrderFulfillment
	var orderMatching *grpc_order.OrderMatching
	var bidIncrease *grpc_order.BalanceUpdate
	var askIncrease *grpc_order.BalanceUpdate
	var fee *grpc_order.Fee
	var prefix = "MarketBidBigger"
	var currency grpc_order.Currency
	var tx redis.Pipeliner
	var ctx context.Context

	var bid, ask = pair.Bid, pair.Ask

	currency = utils.OpponentCurrency(bid.Symbol)
	mul = utils.CoinDecimal(currency)

	askQuantity = decimal.NewFromFloat(ask.Quantity)

	// unit price = ask.UnitPrice
	unitPrice = decimal.NewFromFloat(ask.UnitPrice)

	// filled quantity = ask.Quantity
	filledQuantity = askQuantity

	// fee = filled quantity * unit price * fee ratio
	fee = &grpc_order.Fee{Amount: filledQuantity.Mul(unitPrice).Mul(constant.FeeRatio).InexactFloat64()}

	// bid partial fill
	bidPartialFill = utils.NewMarketOrderMatching(bid.UserUUID, bid.OrderUUID, filledQuantity, unitPrice, bid.OrderType, bid.Symbol, bid.MakeTime, fee)

	// ask fulfillment
	askFulfillment = utils.NewOrderFulfillment(ask.UserUUID, ask.OrderUUID, filledQuantity, unitPrice, ask.Symbol, ask.OrderType, ask.MakeTime, fee)

	// matched order
	orderMatching = utils.NewOrderMatching(unitPrice, filledQuantity, bid.OrderType, bid.Symbol)

	// bid increase
	bidIncrease = utils.NewBalanceUpdate(bid.UserUUID, filledQuantity.Mul(mul).Mul(constant.ReverseFeeRatio), currency, grpc_order.Reason_TAKE)

	// ask increase
	askIncrease = utils.NewBalanceUpdate(ask.UserUUID, filledQuantity.Mul(unitPrice).Mul(constant.ReverseFeeRatio), grpc_order.Currency_BTC, grpc_order.Reason_MAKE)

	// Send stream
	tx = e.tradeManager.Pipeliner()
	ctx = context.Background()

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, bidIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, askIncrease); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send balance update stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderMatchingStreamPipeline(tx, ctx, orderMatching); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order matching stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendMarketOrderMatchingStreamPipeline(tx, ctx, bidPartialFill); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order partial fill stream, %v", prefix, err)
	}

	if err = e.tradeManager.SendOrderFulfillmentStreamPipeline(tx, ctx, askFulfillment); err != nil {
		return status.Errorf(codes.Internal, "%v: failed to send order fulfillment stream, %v", prefix, err)
	}

	_ = e.tradeManager.DeleteOrder(ask.OrderUUID)

	if _, err = tx.Exec(ctx); err != nil {
		return err
	}

	return nil
}