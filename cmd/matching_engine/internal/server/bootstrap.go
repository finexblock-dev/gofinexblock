package server

import (
	"fmt"
	"github.com/finexblock-dev/gofinexblock/cmd/matching_engine/internal/config"
	"github.com/finexblock-dev/gofinexblock/cmd/matching_engine/internal/engine"
	"github.com/finexblock-dev/gofinexblock/pkg/database"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/grpc_order"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/health"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
	"strings"
)

func RegisterServer(server *Gateway, grpcServer *grpc.Server) {
	grpc_order.RegisterCancelOrderServer(grpcServer, server)
	grpc_order.RegisterMarketOrderServer(grpcServer, server)
	grpc_order.RegisterLimitOrderServer(grpcServer, server)
	grpc_order.RegisterOrderBookServer(grpcServer, server)
	health.RegisterHealthCheckServer(grpcServer, server)
	reflection.Register(grpcServer)
}

func Bootstrap() {

	var configuration *config.MatchingEngineConfiguration
	var configClient *config.MatchingEngineConfig
	var eventSubscriber *grpc.ClientConn
	var chartServer *grpc.ClientConn
	var cluster *redis.ClusterClient
	var db *gorm.DB
	var err error

	configClient, err = config.NewGrpcServerConfig()
	if err != nil {
		panic(err)
	}

	configClient.Credentials(os.Getenv("SECRET_NAME"), &configuration)

	log.Println(configuration)

	db = database.CreateMySQLClient(configuration.MysqlUser, configuration.MysqlPass, configuration.MysqlHost, configuration.MysqlPort, configuration.MysqlDB)
	cluster = database.CreateRedisCluster(strings.Split(configuration.RedisHost, ","), configuration.RedisUser, configuration.RedisPass)

	eventSubscriber, err = grpc.Dial(os.Getenv("EVENT_SUBSCRIBER"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error occurred while dialing to event subscriber : %v", err)
	}

	chartServer, err = grpc.Dial(os.Getenv("CHART_SERVER"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error occurred while dialing to event subscriber : %v", err)
	}

	// gRPC server
	server := New(cluster, db, eventSubscriber)

	LoadOrderBook(server.orderBook)
	LoadStream(server.tradeManager)

	// Engine start
	matchingEngine := engine.NewEngine(cluster, eventSubscriber, chartServer)
	matchingEngine.Run()

	grpcServer := grpc.NewServer()

	RegisterServer(server, grpcServer)

	port := os.Getenv("PORT")
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Error occurred while listening port on %v : %v", port, err)
	}

	log.Println("GRPC SERVER START")

	if err = grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error occurred while serve listener.. : %v", err)
	}
}