package bootstrap

import (
	"fmt"
	"github.com/finexblock-dev/gofinexblock/cmd/proxy/internal/config"
	"github.com/finexblock-dev/gofinexblock/cmd/proxy/internal/server"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"os"
)

func ClientConnections(proxy *server.Proxy, configure *config.ProxyConfiguration) {
	btc, err := grpc.Dial(fmt.Sprintf("%v:%v", configure.BitcoinKeyServer, configure.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to key server: %v", err)
	}
	bitcoin := server.NewBitcoinProxy(btc)

	eth, err := grpc.Dial(fmt.Sprintf("%v:%v", configure.EthereumKeyServer, configure.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to key server: %v", err)
	}
	ethereum := server.NewEthereumProxy(eth)

	matic, err := grpc.Dial(fmt.Sprintf("%v:%v", configure.PolygonKeyServer, configure.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to key server: %v", err)
	}
	polygon := server.NewPolygonProxy(matic)

	proxy.SetBitcoin(bitcoin)
	proxy.SetEthereum(ethereum)
	proxy.SetPolygon(polygon)
}

func RegisterWalletProxyServer(server *grpc.Server, proxy *server.Proxy) {
	proxy.Register(server)
}

func ProxyBootstrap() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file : %v", err)
	}

	configuration, err := config.NewProxyConfig()
	if err != nil {
		log.Fatal(err)
	}

	var target *config.ProxyConfiguration
	if os.Getenv("APPMODE") == "LOCAL" {
		target = &config.ProxyConfiguration{
			EthereumKeyServer: "localhost:40041",
			PolygonKeyServer:  "localhost:30031",
			BitcoinKeyServer:  "localhost:20021",
			Port:              "50051",
		}
	} else {
		configuration.Credentials(os.Getenv("SECRET_NAME"), &target)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.ChainUnaryServer(
			prometheus.UnaryServerInterceptor,
			recovery.UnaryServerInterceptor(),
		)),
	)

	proxy := server.NewProxy()

	ClientConnections(proxy, target)

	RegisterWalletProxyServer(grpcServer, proxy)

	// For health check
	go func(grpcServer *grpc.Server) {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%v", 80))
		if err != nil {
			log.Fatalf("Error occurred while listening port on %v : %v", 80, err)
		}
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Error occurred while serve listener... : %v", err)
		}
		log.Println("HTTP SERVER START")

	}(grpcServer)

	proxy.Listen(grpcServer, target.Port)
}
