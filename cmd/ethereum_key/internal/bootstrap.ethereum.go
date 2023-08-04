package bootstrap

import (
	"fmt"
	"github.com/finexblock-dev/gofinexblock/cmd/ethereum_key/internal/config"
	"github.com/finexblock-dev/gofinexblock/cmd/ethereum_key/internal/server"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func RegisterEthereumKeyServer(server *grpc.Server, ethereum *server.EthereumKey) {
	ethereum.Register(server)
}

func EthereumKeyBootstrap() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file : %v", err)
	}

	configuration, err := config.NewEthereumKeyConfig()
	if err != nil {
		log.Fatalf("Error creating configuration: %v", err)
	}

	var target *config.EthereumKeyConfiguration
	configuration.Credentials(os.Getenv("SECRET_NAME"), &target)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.ChainUnaryServer(
			prometheus.UnaryServerInterceptor,
			recovery.UnaryServerInterceptor(),
		)),
	)

	ethereumServer, err := server.NewEthereumKey(target)
	if err != nil {
		log.Fatal(err)
	}

	RegisterEthereumKeyServer(grpcServer, ethereumServer)

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

	ethereumServer.Listen(grpcServer, os.Getenv("PORT"))
}
