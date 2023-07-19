package refund

import (
	"context"
	"encoding/json"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

func (e *engine) ParseMessage(message redis.XMessage) (event *grpc_order.BalanceUpdate, err error) {
	var bytes []byte
	var data = make(map[string]string)

	bytes, err = json.Marshal(message.Values)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	event = new(grpc_order.BalanceUpdate)

	if err = protojson.Unmarshal([]byte(data["event"]), event); err != nil {
		return nil, err
	}

	return event, nil
}

func (e *engine) Do(event *grpc_order.BalanceUpdate) (_err error) {
	var amount = decimal.NewFromFloat(event.GetDiff())
	var user = event.GetUserUUID()
	var currency = event.GetCurrency()
	var ctx = context.Background()
	var lock bool

	if event.Reason != grpc_order.Reason_ADVANCE_PAYMENT {
		if lock, _err = e.tradeManager.AcquireLock(user, currency.String()); _err != nil || !lock {
			return status.Errorf(codes.ResourceExhausted, "acquire lock failed: %v", _err)
		}

		defer func() {
			if _err = e.tradeManager.ReleaseLock(user, currency.String()); _err != nil {
				_err = status.Errorf(codes.Internal, "release lock failed: %v", _err)
			}
		}()

		if _err = e.tradeManager.PlusBalance(user, currency.String(), amount); _err != nil {
			return status.Errorf(codes.Internal, "plus balance failed: %v", _err)
		}
	}

	if _, _err = e.eventSubscriber.BalanceUpdateEvent(ctx, event); _err != nil {
		return status.Errorf(status.Code(_err), "send balance update event failed: %v", _err)
	}

	if _, _err = e.chartServer.BalanceUpdateEvent(ctx, event); _err != nil {
		return status.Errorf(status.Code(_err), "send balance update event failed: %v", _err)
	}

	return nil
}