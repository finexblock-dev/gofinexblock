package main

import (
	"context"
	"crypto/tls"
	"fmt"
	_ "github.com/finexblock-dev/gofinexblock/pkg/admin"
	_ "github.com/finexblock-dev/gofinexblock/pkg/announcement"
	_ "github.com/finexblock-dev/gofinexblock/pkg/auth"
	_ "github.com/finexblock-dev/gofinexblock/pkg/btcd"
	_ "github.com/finexblock-dev/gofinexblock/pkg/cache"
	_ "github.com/finexblock-dev/gofinexblock/pkg/constant"
	_ "github.com/finexblock-dev/gofinexblock/pkg/daemon"
	_ "github.com/finexblock-dev/gofinexblock/pkg/database"
	_ "github.com/finexblock-dev/gofinexblock/pkg/engine/cancellation"
	_ "github.com/finexblock-dev/gofinexblock/pkg/engine/event"
	_ "github.com/finexblock-dev/gofinexblock/pkg/engine/match"
	_ "github.com/finexblock-dev/gofinexblock/pkg/engine/refund"
	_ "github.com/finexblock-dev/gofinexblock/pkg/entity"
	_ "github.com/finexblock-dev/gofinexblock/pkg/ethereum"
	_ "github.com/finexblock-dev/gofinexblock/pkg/files"
	_ "github.com/finexblock-dev/gofinexblock/pkg/gen/bitcoin"
	_ "github.com/finexblock-dev/gofinexblock/pkg/gen/contracts"
	_ "github.com/finexblock-dev/gofinexblock/pkg/gen/erc20"
	_ "github.com/finexblock-dev/gofinexblock/pkg/gen/ethereum"
	_ "github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/health"
	_ "github.com/finexblock-dev/gofinexblock/pkg/gen/polygon"
	_ "github.com/finexblock-dev/gofinexblock/pkg/goaws"
	_ "github.com/finexblock-dev/gofinexblock/pkg/goredis"
	_ "github.com/finexblock-dev/gofinexblock/pkg/image"
	_ "github.com/finexblock-dev/gofinexblock/pkg/instance"
	_ "github.com/finexblock-dev/gofinexblock/pkg/order"
	_ "github.com/finexblock-dev/gofinexblock/pkg/orderbook"
	_ "github.com/finexblock-dev/gofinexblock/pkg/safety"
	_ "github.com/finexblock-dev/gofinexblock/pkg/secure"
	_ "github.com/finexblock-dev/gofinexblock/pkg/trade"
	_ "github.com/finexblock-dev/gofinexblock/pkg/types"
	_ "github.com/finexblock-dev/gofinexblock/pkg/user"
	_ "github.com/finexblock-dev/gofinexblock/pkg/utils"
	_ "github.com/finexblock-dev/gofinexblock/pkg/wallet"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func test(i int) (err error) {

	if i == 0 {
		return nil
	}

	return fmt.Errorf("test error %d", i)
}

func main() {
	conn, _ := grpc.Dial("backoffice-api-dev.finexblock.com:50051", grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
		InsecureSkipVerify: false,
	})))

	client := health.NewHealthCheckClient(conn)

	output, err := client.Check(context.Background(), &health.HealthCheckInput{Name: "test"})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println(output)
}