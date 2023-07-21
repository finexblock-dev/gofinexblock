package main

import (
	"fmt"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/admin"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/announcement"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/auth"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/btcd"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/cache"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/constant"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/daemon"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/database"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/engine/cancellation"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/engine/event"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/engine/match"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/engine/refund"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/entity"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/ethereum"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/files"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/gen/bitcoin"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/gen/contracts"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/gen/erc20"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/gen/ethereum"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/gen/grpc_order"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/gen/polygon"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/goaws"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/goredis"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/image"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/instance"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/order"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/orderbook"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/safety"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/secure"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/stream"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/structure"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/trade"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/types"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/user"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/utils"
	_ "github.com/finexblock-dev/gofinexblock/finexblock/wallet"
	"net"
)

func test(i int) (err error) {

	if i == 0 {
		return nil
	}

	return fmt.Errorf("test error %d", i)
}

func main() {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, i := range interfaces {
		addrs, err := i.Addrs()
		if err != nil {
			panic(err)
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip != nil && ip.IsPrivate() {
				fmt.Println(ip)
			}
		}
	}
}