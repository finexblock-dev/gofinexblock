package client

import (
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/ethereum"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewEthereumProxyClient(host string) (ethereum.EthereumProxyClient, error) {
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return ethereum.NewEthereumProxyClient(conn), nil
}
