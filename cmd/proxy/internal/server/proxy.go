package server

import (
	"context"
	"fmt"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/bitcoin"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/ethereum"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/health"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/polygon"
	"github.com/finexblock-dev/gofinexblock/pkg/goaws"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

type Proxy struct {
	bitcoin  *BitcoinProxy
	ethereum *EthereumProxy
	polygon  *PolygonProxy

	health.UnimplementedHealthCheckServer
}

func (p *Proxy) Check(ctx context.Context, input *health.HealthCheckInput) (*health.HealthCheckOutput, error) {
	return &health.HealthCheckOutput{Message: fmt.Sprintf("Hello %s", input.Name)}, nil
}

func (p *Proxy) WhoAmI(ctx context.Context, input *health.WhoAmIInput) (*health.WhoAmIOutput, error) {
	privateIP, err := goaws.OwnPrivateIP()
	if err != nil {
		return nil, status.Errorf(codes.Unknown, err.Error())
	}
	return &health.WhoAmIOutput{Message: fmt.Sprintf("Hello %s, I am %s", input.Name, privateIP)}, nil
}

func (p *Proxy) SetBitcoin(bitcoin *BitcoinProxy) {
	p.bitcoin = bitcoin
}

func (p *Proxy) SetEthereum(ethereum *EthereumProxy) {
	p.ethereum = ethereum
}

func (p *Proxy) SetPolygon(polygon *PolygonProxy) {
	p.polygon = polygon
}

func (p *Proxy) Listen(gRPCServer *grpc.Server, port string) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Error occurred while listening port on %v : %v", port, err)
	}
	log.Println("GRPC SERVER START")
	if err := gRPCServer.Serve(listener); err != nil {
		log.Fatalf("Error occurred while serve listener... : %v", err)
	}
}

func (p *Proxy) Register(server *grpc.Server) {
	ethereum.RegisterEthereumProxyServer(server, p.ethereum)
	bitcoin.RegisterBitcoinProxyServer(server, p.bitcoin)
	polygon.RegisterPolygonProxyServer(server, p.polygon)
	health.RegisterHealthCheckServer(server, p)
}

func (p *Proxy) mustEmbedUnimplementedHealthCheckServer() {
	//TODO implement me
	panic("implement me")
}

func NewProxy() *Proxy {
	return &Proxy{}
}