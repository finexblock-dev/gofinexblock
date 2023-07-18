package refund

import (
	"context"
	"encoding/json"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
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

func (e *engine) Do(event *grpc_order.BalanceUpdate) (err error) {
	var amount = decimal.NewFromFloat(event.GetDiff())
	var user = event.GetUserUUID()
	var currency = event.GetCurrency()
	var ctx = context.Background()
	var lock bool

	if lock, err = e.tradeManager.AcquireLock(user, currency.String()); err != nil || !lock {
		return status.Errorf(status.Code(err), "acquire lock failed: %v", err)
	}

	defer func() {
		if err = e.tradeManager.ReleaseLock(user, currency.String()); err != nil {
			err = status.Errorf(status.Code(err), "release lock failed: %v", err)
		}
	}()

	if err = e.tradeManager.PlusBalance(user, currency.String(), amount); err != nil {
		return status.Errorf(status.Code(err), "plus balance failed: %v", err)
	}

	if _, err = e.eventSubscriber.BalanceUpdateEvent(ctx, event); err != nil {
		return status.Errorf(status.Code(err), "send balance update event failed: %v", err)
	}

	if _, err = e.chartServer.BalanceUpdateEvent(ctx, event); err != nil {
		return status.Errorf(status.Code(err), "send balance update event failed: %v", err)
	}

	return nil
}