package cancellation

import (
	"context"
	"encoding/json"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/utils"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

func (e *engine) ParseMessage(message redis.XMessage) (_event *grpc_order.OrderCancelled, err error) {
	var bytes []byte
	var data = make(map[string]string)

	bytes, err = json.Marshal(message.Values)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	if err = protojson.Unmarshal([]byte(data["_event"]), _event); err != nil {
		return nil, err
	}

	return _event, nil
}

func (e *engine) Do(event *grpc_order.OrderCancelled) (err error) {
	var amount decimal.Decimal
	var currency grpc_order.Currency
	var balanceUpdate *grpc_order.BalanceUpdate
	var ctx = context.Background()

	var userUUID = event.GetUserUUID()
	var quantity = event.GetQuantity()
	var unitPrice = event.GetUnitPrice()
	var symbol = event.GetSymbol()

	switch event.GetOrderType() {
	case grpc_order.OrderType_BID:
		currency = grpc_order.Currency_BTC
	case grpc_order.OrderType_ASK:
		currency = utils.OpponentCurrency(symbol)
	default:
		return status.Error(codes.InvalidArgument, "invalid order type")
	}

	amount = decimal.NewFromFloat(quantity).Mul(decimal.NewFromFloat(unitPrice))

	balanceUpdate = utils.NewBalanceUpdate(userUUID, amount, currency, grpc_order.Reason_REFUND)

	if e.tradeManager.SendBalanceUpdateStream(balanceUpdate) != nil {
		return err
	}

	if _, err = e.eventManager.OrderCancellationEvent(ctx, event); err != nil {
		// FIXME: How to fix this?
		return err
	}

	return nil
}