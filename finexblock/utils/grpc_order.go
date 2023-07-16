package utils

import (
	"encoding/json"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
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

func MessagesToJson(s proto.Message) (map[string]interface{}, error) {
	var marshaller protojson.MarshalOptions
	var mapData map[string]interface{}
	var jsonData []byte
	var err error

	marshaller = protojson.MarshalOptions{UseProtoNames: true}

	jsonData, err = marshaller.Marshal(s)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonData, &mapData)
	if err != nil {
		return nil, err
	}

	return mapData, nil
}
