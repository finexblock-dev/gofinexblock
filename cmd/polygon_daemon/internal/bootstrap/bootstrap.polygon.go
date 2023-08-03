package bootstrap

import (
	"github.com/finexblock-dev/gofinexblock/cmd/polygon_daemon/internal/client"
	"github.com/finexblock-dev/gofinexblock/cmd/polygon_daemon/internal/config"
	polygonDaemon "github.com/finexblock-dev/gofinexblock/cmd/polygon_daemon/internal/task"
	"github.com/finexblock-dev/gofinexblock/finexblock/daemon"
	"github.com/finexblock-dev/gofinexblock/finexblock/database"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/polygon"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
	"time"
)

func PolygonBootstrap() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed to load env file : %v", err)
	}

	var configurations *config.PolygonConfiguration
	var configClient *config.PolygonConfig
	var proxyClient polygon.PolygonProxyClient
	var cluster *redis.ClusterClient
	var db *gorm.DB
	var err error

	configClient, err = config.NewPolygonConfig()
	if err != nil {
		log.Fatal(err)
	}

	configClient.Credentials(os.Getenv("SECRET_NAME"), &configurations)
	cluster = database.CreateRedisCluster(strings.Split(configurations.RedisHost, ","), configurations.RedisUser, configurations.RedisPass)
	if os.Getenv("APPMODE") == "LOCAL" {
		configurations.ProxyHost = "localhost:50051"
	}

	log.Println(configurations)

	// Database connection
	db = database.Mysql(configurations.MysqlUser, configurations.MysqlPass, configurations.MysqlDatabase, configurations.MysqlHost, configurations.MysqlPort)

	// gRPC connection
	proxyClient, err = client.NewPolygonProxyClient(configurations.ProxyHost)

	// Start polygonDaemon
	deposit := polygonDaemon.NewDeposit(proxyClient, configurations.HotWallet, db, cluster, time.Second*10)
	withdrawal := polygonDaemon.NewWithdrawal(proxyClient, configurations.HotWallet, db, cluster, time.Second*10)

	go daemon.Run(deposit)
	go daemon.Run(withdrawal)
}
