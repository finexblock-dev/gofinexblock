package match

import (
	"encoding/json"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/finexblock/trade"
	"github.com/finexblock-dev/gofinexblock/finexblock/types"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

func (e *engine) ParseMessage(message redis.XMessage) (_case types.Case, pair *grpc_order.BidAsk, err error) {
	var bytes []byte
	var data = make(map[string]string)
	pair = new(grpc_order.BidAsk)

	bytes, err = json.Marshal(message.Values)
	if err != nil {
		return "", nil, err
	}

	if err = json.Unmarshal(bytes, &data); err != nil {
		return "", nil, err
	}

	if err = protojson.Unmarshal([]byte(data["pair"]), pair); err != nil {
		return "", nil, err
	}

	return types.Case(data["case"]), pair, nil
}

func (e *engine) Do(_case types.Case, pair *grpc_order.BidAsk) (err error) {
	switch types.Case(_case) {
	case trade.CaseLimitAskBigger:
		return e.LimitAskBigger(pair)
	case trade.CaseLimitAskEqual:
		return e.LimitAskEqual(pair)
	case trade.CaseLimitAskSmaller:
		return e.LimitAskSmaller(pair)
	case trade.CaseLimitBidBigger:
		return e.LimitBidBigger(pair)
	case trade.CaseLimitBidEqual:
		return e.LimitBidEqual(pair)
	case trade.CaseLimitBidSmaller:
		return e.LimitBidSmaller(pair)
	case trade.CaseMarketAskBigger:
		return e.MarketAskBigger(pair)
	case trade.CaseMarketAskEqual:
		return e.MarketAskEqual(pair)
	case trade.CaseMarketAskSmaller:
		return e.MarketAskSmaller(pair)
	case trade.CaseMarketBidBigger:
		return e.MarketBidBigger(pair)
	case trade.CaseMarketBidEqual:
		return e.MarketBidEqual(pair)
	case trade.CaseMarketBidSmaller:
		return e.MarketBidSmaller(pair)
	default:
		return status.Errorf(codes.InvalidArgument, "invalid case: %s", _case)
	}
}