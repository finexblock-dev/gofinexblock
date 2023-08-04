package server

import (
	"context"
	"fmt"
	"github.com/finexblock-dev/gofinexblock/cmd/event_subscriber/internal/channel"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	grpc_order.UnimplementedEventServer

	balanceUpdate       *channel.BalanceUpdateChannel
	orderCancellation   *channel.OrderCancellationChannel
	orderFulfillment    *channel.OrderFulfillmentChannel
	orderPartialFill    *channel.OrderPartialFillChannel
	orderMatching       *channel.OrderMatchingChannel
	orderInitialize     *channel.OrderInitializeChannel
	chartDrawer         *channel.ChartDrawerChannel
	marketOrderMatching *channel.MarketOrderMatchingChannel
}

func (s *Server) Listen(gRPCServer *grpc.Server, port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Error occurred while listening port on %v : %v", port, err)
	}
	log.Println("GRPC SERVER START")
	if err := gRPCServer.Serve(listener); err != nil {
		log.Fatalf("Error occurred while serve listener... : %v", err)
	}
}

func (s *Server) Register(grpcServer *grpc.Server) {
	grpc_order.RegisterEventServer(grpcServer, s)
}

func NewServer(
	balanceUpdate *channel.BalanceUpdateChannel,
	orderCancellation *channel.OrderCancellationChannel,
	orderFulfillment *channel.OrderFulfillmentChannel,
	orderPartialFill *channel.OrderPartialFillChannel,
	orderMatching *channel.OrderMatchingChannel,
	orderInitialize *channel.OrderInitializeChannel,
	chartDrawer *channel.ChartDrawerChannel,
	marketOrderMatching *channel.MarketOrderMatchingChannel,
) *Server {
	return &Server{
		UnimplementedEventServer: grpc_order.UnimplementedEventServer{},
		balanceUpdate:            balanceUpdate,
		orderCancellation:        orderCancellation,
		orderFulfillment:         orderFulfillment,
		orderPartialFill:         orderPartialFill,
		orderMatching:            orderMatching,
		orderInitialize:          orderInitialize,
		chartDrawer:              chartDrawer,
		marketOrderMatching:      marketOrderMatching,
	}
}

func (s *Server) OrderPlacementEvent(ctx context.Context, input *grpc_order.OrderPlacement) (*grpc_order.Ack, error) {
	log.Println(input)
	return &grpc_order.Ack{Success: true}, nil
}

func (s *Server) MarketOrderMatchingEvent(ctx context.Context, input *grpc_order.MarketOrderMatching) (*grpc_order.Ack, error) {
	s.marketOrderMatching.Receive(input)
	return &grpc_order.Ack{Success: true}, nil
}

func (s *Server) OrderPlacementFailedEvent(ctx context.Context, input *grpc_order.OrderPlacementFailed) (*grpc_order.Ack, error) {
	log.Println(input)
	return &grpc_order.Ack{Success: true}, nil
}

func (s *Server) OrderCancellationEvent(ctx context.Context, input *grpc_order.OrderCancelled) (*grpc_order.Ack, error) {
	s.orderCancellation.Receive(input)
	return &grpc_order.Ack{Success: true}, nil
}

func (s *Server) OrderCancellationFailedEvent(ctx context.Context, input *grpc_order.OrderCancellationFailed) (*grpc_order.Ack, error) {
	log.Println(input)
	return &grpc_order.Ack{Success: true}, nil
}

func (s *Server) OrderMatchingEvent(ctx context.Context, input *grpc_order.OrderMatching) (*grpc_order.Ack, error) {
	s.orderMatching.Receive(input)
	s.chartDrawer.Receive(input)
	return &grpc_order.Ack{Success: true}, nil
}

func (s *Server) OrderMatchingFailedEvent(ctx context.Context, input *grpc_order.OrderMatchingFailed) (*grpc_order.Ack, error) {
	log.Println(input)
	return &grpc_order.Ack{Success: true}, nil
}

func (s *Server) OrderPartialFillEvent(ctx context.Context, input *grpc_order.OrderPartialFill) (*grpc_order.Ack, error) {
	s.orderPartialFill.Receive(input)
	return &grpc_order.Ack{Success: true}, nil
}

func (s *Server) OrderFulfillmentEvent(ctx context.Context, input *grpc_order.OrderFulfillment) (*grpc_order.Ack, error) {
	s.orderFulfillment.Receive(input)
	return &grpc_order.Ack{Success: true}, nil
}

func (s *Server) OrderInitializeEvent(ctx context.Context, input *grpc_order.OrderInitialize) (*grpc_order.Ack, error) {
	s.orderInitialize.Receive(input)
	return &grpc_order.Ack{Success: true}, nil
}

func (s *Server) BalanceUpdateEvent(ctx context.Context, input *grpc_order.BalanceUpdate) (*grpc_order.Ack, error) {
	s.balanceUpdate.Receive(input)
	return &grpc_order.Ack{Success: true}, nil
}

func (s *Server) mustEmbedUnimplementedEventServer() {}