package server

import (
	"fmt"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/bitcoin"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/ethereum"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/polygon"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Proxy struct {
	bitcoin  *BitcoinProxy
	ethereum *EthereumProxy
	polygon  *PolygonProxy
}

func NewProxy() *Proxy {
	return &Proxy{}
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
}
