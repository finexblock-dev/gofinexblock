package client

import (
	"github.com/finexblock-dev/gofinexblock/pkg/gen/polygon"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewPolygonProxyClient(host string) (polygon.PolygonProxyClient, error) {
	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return polygon.NewPolygonProxyClient(conn), nil
}