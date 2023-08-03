package internal

import (
	"github.com/finexblock-dev/gofinexblock/cmd/event_subscriber/internal/channel"
	"github.com/finexblock-dev/gofinexblock/cmd/event_subscriber/internal/config"
	"github.com/finexblock-dev/gofinexblock/cmd/event_subscriber/internal/server"
	"github.com/finexblock-dev/gofinexblock/pkg/database"
	"github.com/finexblock-dev/gofinexblock/pkg/order"
	"github.com/finexblock-dev/gofinexblock/pkg/safety"
	"github.com/finexblock-dev/gofinexblock/pkg/wallet"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"gorm.io/gorm"
	"os"
	"strings"
)

func Bootstrap() {
	var configuration *config.EventSubscriberConfig
	var target *config.EventSubscriberConfiguration
	var cluster *redis.ClusterClient
	var db *gorm.DB
	var err error

	configuration, err = config.NewEventSubscriberConfig()
	if err != nil {
		panic(err)
	}

	configuration.Credentials(os.Getenv("SECRET_NAME"), &target)

	if os.Getenv("APPMODE") == "LOCAL" {
		db = database.GetTunnelledMySQL(os.Getenv("SSH_HOST"), os.Getenv("SSH_USER"), os.Getenv("PEM"),
			os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASS"), os.Getenv("MYSQL_DB"), 22, 6033)
	} else {
		db = database.CreateMySQLClient(target.MysqlUser, target.MysqlPass, target.MysqlHost, target.MysqlPort, target.MysqlDatabase)
		cluster = database.CreateRedisCluster(strings.Split(target.RedisHost, ","), target.RedisUser, target.RedisPass)
	}

	// service
	orderService := order.NewService(db)
	walletService := wallet.NewService(db, cluster)

	marketOrderMatching := channel.NewMarketOrderMatchingChannel(orderService)
	go safety.InfinitySubscribe(marketOrderMatching)

	balanceUpdate := channel.NewBalanceUpdateChannel(walletService)
	go safety.InfinitySubscribe(balanceUpdate)

	chart := channel.NewChartDrawerChannel(orderService)
	go safety.InfinitySubscribe(chart)

	interval := channel.NewIntervalChannel(orderService)
	go safety.InfinitySubscribe(interval)

	cancellation := channel.NewOrderCancellationChannel(orderService)
	go safety.InfinitySubscribe(cancellation)

	fulfillment := channel.NewOrderFulfillmentChannel(orderService)
	go safety.InfinitySubscribe(fulfillment)

	initialize := channel.NewOrderInitializeChannel(orderService)
	go safety.InfinitySubscribe(initialize)

	matching := channel.NewOrderMatchingChannel(orderService)
	go safety.InfinitySubscribe(matching)

	partial := channel.NewOrderPartialFillChannel(orderService)
	go safety.InfinitySubscribe(partial)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.ChainUnaryServer(
			recovery.UnaryServerInterceptor(),
		)),
	)

	app := server.NewServer(balanceUpdate, cancellation, fulfillment, partial, matching, initialize, chart, marketOrderMatching)

	app.Register(grpcServer)
	app.Listen(grpcServer, os.Getenv("PORT"))
}