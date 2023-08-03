package bootstrap

import (
	"fmt"
	"github.com/finexblock-dev/gofinexblock/cmd/bitcoin_key/internal/config"
	"github.com/finexblock-dev/gofinexblock/cmd/bitcoin_key/internal/server"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func RegisterBitcoinKeyServer(bitcoinKey *server.BitcoinKey, grpcServer *grpc.Server) {
	bitcoinKey.Register(grpcServer)
}

func BitcoinKeyBoostrap() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file : %v", err)
	}

	configuration, err := config.NewBitcoinKeyConfig()
	if err != nil {
		log.Fatalf("Error creating configuration: %v", err)
	}

	var target *config.BitcoinKeyConfiguration
	configuration.Credentials(os.Getenv("SECRET_NAME"), &target)

	bitcoinKeyServer, err := server.NewBitcoinKey(target)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(target.Mnemonic)
	log.Println(target.WalletAccount)
	log.Println(target.RpcHost)
	log.Println(target.RpcPort)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.ChainUnaryServer(
			prometheus.UnaryServerInterceptor,
			recovery.UnaryServerInterceptor(),
		)),
	)

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

	RegisterBitcoinKeyServer(bitcoinKeyServer, grpcServer)

	bitcoinKeyServer.Listen(grpcServer, os.Getenv("PORT"))
}
