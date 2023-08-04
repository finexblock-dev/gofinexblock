package client

import (
	"github.com/finexblock-dev/gofinexblock/pkg/gen/bitcoin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewBitcoinProxyClient(host string) (bitcoin.BitcoinProxyClient, error) {
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return bitcoin.NewBitcoinProxyClient(conn), nil
}