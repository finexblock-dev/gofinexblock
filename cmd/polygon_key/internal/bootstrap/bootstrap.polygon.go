package bootstrap

import (
	"fmt"
	"github.com/finexblock-dev/gofinexblock/cmd/polygon_key/internal/config"
	"github.com/finexblock-dev/gofinexblock/cmd/polygon_key/internal/server"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func RegisterPolygonKeyServer(server *grpc.Server, polygon *server.PolygonKey) {
	polygon.Register(server)
}

func PolygonKeyBootstrap() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file : %v", err)
	}

	// Same as PolygonKeyBootstrap() but with different config
	configuration, err := config.NewPolygonKeyConfig()
	if err != nil {
		log.Fatalf("Error creating configuration: %v", err)
	}

	// Same as PolygonKeyBootstrap() but with different config
	var target *config.PolygonKeyConfiguration
	configuration.Credentials(os.Getenv("SECRET_NAME"), &target)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.ChainUnaryServer(
			prometheus.UnaryServerInterceptor,
			recovery.UnaryServerInterceptor(),
		)),
	)

	// Same as PolygonKeyBootstrap() but with different config
	polygonServer, err := server.NewPolygonKey(target)
	if err != nil {
		log.Fatal(err)
	}

	// Same as PolygonKeyBootstrap() but with different config
	RegisterPolygonKeyServer(grpcServer, polygonServer)

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

	polygonServer.Listen(grpcServer, os.Getenv("PORT"))
}
