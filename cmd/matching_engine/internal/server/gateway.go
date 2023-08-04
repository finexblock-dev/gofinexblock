package server

import (
	"context"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/pkg/orderbook"
	"github.com/finexblock-dev/gofinexblock/pkg/trade"
	"github.com/finexblock-dev/gofinexblock/pkg/utils"
	"github.com/redis/go-redis/v9"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"log"
	"math"
	"sort"
)

type Gateway struct {
	tradeManager trade.Manager
	orderBook    orderbook.Manager
	event        grpc_order.EventClient

	grpc_order.UnimplementedCancelOrderServer
	grpc_order.UnimplementedLimitOrderServer
	grpc_order.UnimplementedMarketOrderServer
	grpc_order.UnimplementedOrderBookServer
}

func New(
	cluster *redis.ClusterClient,
	db *gorm.DB,
	conn *grpc.ClientConn,
) *Gateway {
	return &Gateway{
		tradeManager:                   trade.New(cluster),
		orderBook:                      orderbook.New(cluster, db),
		event:                          grpc_order.NewEventClient(conn),
		UnimplementedCancelOrderServer: grpc_order.UnimplementedCancelOrderServer{},
		UnimplementedLimitOrderServer:  grpc_order.UnimplementedLimitOrderServer{},
		UnimplementedMarketOrderServer: grpc_order.UnimplementedMarketOrderServer{},
		UnimplementedOrderBookServer:   grpc_order.UnimplementedOrderBookServer{},
	}
}

func (g *Gateway) CancelOrder(_ context.Context, in *grpc_order.OrderCancellation) (*grpc_order.Ack, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("recovered from panic", r)
		}
	}()
	_, err := g.orderBook.CancelOrder(in.OrderUUID)
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Error())
	}
	return &grpc_order.Ack{Success: true}, nil
}

func (g *Gateway) LimitOrderInit(_ context.Context, in *grpc_order.LimitOrderInput) (*grpc_order.Ack, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("recovered from panic", r)
		}
	}()
	done := make(chan error)
	defer close(done)

	go func(done chan error, in *grpc_order.LimitOrderInput) {
		var wallet string
		var target *grpc_order.Order
		var currency grpc_order.Currency
		var queue func(order *grpc_order.Order) (*grpc_order.Order, error)
		var amount float64

		// 수량
		quantity := decimal.NewFromFloat(in.Quantity)

		// 가격
		unitPrice := decimal.NewFromFloat(in.UnitPrice)

		// 지정가 매수/매도 주문 생성
		target = &grpc_order.Order{
			UserUUID:  in.UserUUID,
			OrderUUID: in.OrderUUID,
			Quantity:  in.Quantity,
			UnitPrice: in.UnitPrice,
			OrderType: in.OrderType,
			Symbol:    in.Symbol,
			MakeTime:  timestamppb.Now(),
		}

		switch in.OrderType {
		case grpc_order.OrderType_ASK:
			currency = utils.OpponentCurrency(target.Symbol)
			wallet = utils.OpponentCurrencyToString(target.Symbol)
			amount = utils.CoinDecimal(currency).Mul(quantity).InexactFloat64() // 지정가 매도 시 수량 * decimal
			queue = g.orderBook.LimitAskInsert
		case grpc_order.OrderType_BID:
			currency = grpc_order.Currency_BTC
			wallet = grpc_order.Currency_BTC.String()
			amount = decimal.NewFromFloat(in.Quantity).Mul(decimal.NewFromFloat(in.UnitPrice)).InexactFloat64() // 지정가 매수 시 수량 * unit_price
			queue = g.orderBook.LimitBidInsert
		default:
			done <- status.Errorf(codes.InvalidArgument, "invalid order type")
			return
		}

		// 계정 Lock 획득
		if ok, err := g.tradeManager.AcquireLock(in.UserUUID, currency.String()); err != nil || !ok {
			done <- status.Errorf(codes.ResourceExhausted, "failed to acquire lock: [%v]", err)
			return
		}

		// 계정 Lock 해제
		defer func(tradeManager trade.Manager, uuid, currency string) {
			_ = tradeManager.ReleaseLock(uuid, currency)
		}(g.tradeManager, in.UserUUID, currency.String())

		// Redis 트랜잭션
		tx := g.tradeManager.Pipeliner()
		ctx := context.Background()

		// 잔고 차감
		if err := g.tradeManager.MinusBalanceWithTx(tx, ctx, in.UserUUID, wallet, decimal.NewFromFloat(amount)); err != nil {
			done <- status.Errorf(codes.Internal, "failed to minus balance : [%v]", err)
			return
		}

		// 잔고 변동 event
		balanceUpdate := utils.NewBalanceUpdate(in.UserUUID, decimal.NewFromFloat(amount).Neg(), currency, grpc_order.Reason_ADVANCE_PAYMENT)
		if err := g.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, balanceUpdate); err != nil {
			done <- status.Errorf(codes.Internal, "failed to send balance update stream: [%v]", err)
			return
		}

		// 주문 DB Insert
		orderInitialized := utils.NewOrderInitialized(target.UserUUID, target.OrderUUID, quantity, unitPrice, target.OrderType, target.Symbol)
		if err := g.tradeManager.SendInitializeStreamPipeline(tx, ctx, orderInitialized); err != nil {
			done <- status.Errorf(codes.Internal, "failed to send initialize stream")
			return
		}

		// Redis EXEC
		if _, err := tx.Exec(ctx); err != nil {
			done <- status.Errorf(codes.Internal, "failed to exec pipeline: [%v]", err)
			return
		}

		go func(queue func(order *grpc_order.Order) (*grpc_order.Order, error), target *grpc_order.Order) {
			var _err error

			if _, _err = queue(target); _err != nil {
				// logging error
				log.Println("failed to insert order to orderbook", _err)
				log.Println("failed order", target)
			}
		}(queue, target)

		done <- nil
		return
	}(done, in)

	if err := <-done; err != nil {
		return nil, err
	}

	return &grpc_order.Ack{Success: true}, nil
}

func (g *Gateway) MarketOrderInit(_ context.Context, in *grpc_order.MarketOrderInput) (*grpc_order.Ack, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("recovered from panic", r)
		}
	}()
	done := make(chan error)
	defer close(done)

	go func(done chan error, in *grpc_order.MarketOrderInput) {
		var wallet string
		var target *grpc_order.Order
		var currency grpc_order.Currency
		var queue func(order *grpc_order.Order) (*grpc_order.Order, error)

		target = &grpc_order.Order{
			UserUUID:  in.UserUUID,
			OrderUUID: in.OrderUUID,
			Quantity:  in.Quantity,
			UnitPrice: math.MaxFloat64,
			OrderType: in.OrderType,
			Symbol:    in.Symbol,
			MakeTime:  timestamppb.Now(),
		}

		switch in.OrderType {
		case grpc_order.OrderType_BID:
			// currency
			currency = grpc_order.Currency_BTC
			// wallet
			wallet = currency.String()
			// push queue
			queue = g.orderBook.MarketBidInsert
		case grpc_order.OrderType_ASK:
			// currency
			currency = utils.OpponentCurrency(target.Symbol)
			// wallet
			wallet = currency.String()
			// push queue
			queue = g.orderBook.MarketAskInsert
		default:
			done <- status.Errorf(codes.InvalidArgument, "invalid order type")
			return
		}

		if ok, err := g.tradeManager.AcquireLock(in.UserUUID, currency.String()); err != nil || !ok {
			done <- status.Errorf(codes.ResourceExhausted, "failed to acquire lock: [%v]", err)
			return
		}
		defer func(tradeManager trade.Manager, uuid, currency string) {
			_ = tradeManager.ReleaseLock(uuid, currency)
		}(g.tradeManager, in.UserUUID, currency.String())

		// Redis 트랜잭션
		tx := g.tradeManager.Pipeliner()

		ctx := context.Background()

		// advance payment
		amount := decimal.NewFromFloat(in.Quantity)
		if err := g.tradeManager.MinusBalanceWithTx(tx, ctx, in.UserUUID, wallet, amount); err != nil {
			done <- status.Errorf(codes.Internal, "failed to minus balance : [%v]", err)
			return
		}

		// stack event
		balanceUpdate := utils.NewBalanceUpdate(in.UserUUID, amount.Neg(), currency, grpc_order.Reason_ADVANCE_PAYMENT)
		if err := g.tradeManager.SendBalanceUpdateStreamPipeline(tx, ctx, balanceUpdate); err != nil {
			done <- status.Errorf(codes.Internal, "failed to send balance update stream: [%v]", err)
			return
		}

		if _, err := tx.Exec(ctx); err != nil {
			done <- status.Errorf(codes.Internal, "failed to exec pipeline: [%v]", err)
			return
		}

		go func(queue func(order *grpc_order.Order) (*grpc_order.Order, error), target *grpc_order.Order) {
			var _err error

			if _, _err = queue(target); _err != nil {
				// logging error
				log.Println("failed to insert order to orderbook", _err)
				log.Println("failed order", target)
			}
		}(queue, target)

		done <- nil
		return
	}(done, in)

	if err := <-done; err != nil {
		return nil, err
	}

	return &grpc_order.Ack{Success: true}, nil
}

func (g *Gateway) GetOrderBook(_ context.Context, _ *grpc_order.GetOrderBookInput) (*grpc_order.GetOrderBookOutput, error) {
	var bidResult, askResult []*grpc_order.OrderBookData
	var bidList, askList []*grpc_order.Order
	var bidOrder, askOrder = make(map[float64]float64), make(map[float64]float64)
	var err error

	bidList, err = g.orderBook.BidOrder()
	if err != nil && err != orderbook.ErrOrderBookEmpty {
		return nil, status.Errorf(codes.Internal, "failed to get bid order: [%v]", err)
	}

	if err == orderbook.ErrOrderBookEmpty {
		bidList = []*grpc_order.Order{}
	}

	askList, err = g.orderBook.AskOrder()
	if err != nil && err != orderbook.ErrOrderBookEmpty {
		return nil, status.Errorf(codes.Internal, "failed to get ask order: [%v]", err)
	}

	if err == orderbook.ErrOrderBookEmpty {
		askList = []*grpc_order.Order{}
	}

	if len(bidList) > 0 {
		for _, v := range bidList {
			bidOrder[v.UnitPrice] += v.Quantity
		}
	}

	if len(askList) > 0 {
		for _, v := range askList {
			askOrder[v.UnitPrice] += v.Quantity
		}
	}

	for price, quantity := range bidOrder {
		bidResult = append(bidResult, &grpc_order.OrderBookData{
			Price:  price,
			Volume: quantity,
		})
	}

	for price, quantity := range askOrder {
		askResult = append(askResult, &grpc_order.OrderBookData{
			Price:  price,
			Volume: quantity,
		})
	}

	sort.Slice(bidResult, func(i, j int) bool {
		return bidResult[i].Price > bidResult[j].Price
	})

	sort.Slice(askResult, func(i, j int) bool {
		return askResult[i].Price < askResult[j].Price
	})

	if len(bidResult) > 25 {
		bidResult = bidResult[:25]
	}

	if len(askResult) > 25 {
		askResult = askResult[:25]
	}

	return &grpc_order.GetOrderBookOutput{
		Bids: bidResult,
		Asks: askResult,
	}, nil
}

func (g *Gateway) mustEmbedUnimplementedOrderBookServer() {
	//TODO implement me
	panic("implement me")
}

func (g *Gateway) mustEmbedUnimplementedCancelOrderServer() {
	//TODO implement me
	panic("implement me")
}

func (g *Gateway) mustEmbedUnimplementedLimitOrderServer() {
	//TODO implement me
	panic("implement me")
}

func (g *Gateway) mustEmbedUnimplementedMarketOrderServer() {
	//TODO implement me
	panic("implement me")
}