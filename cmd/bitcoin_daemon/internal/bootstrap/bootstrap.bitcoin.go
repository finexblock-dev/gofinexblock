package bootstrap

import (
	"github.com/finexblock-dev/gofinexblock/cmd/bitcoin_daemon/internal/client"
	"github.com/finexblock-dev/gofinexblock/cmd/bitcoin_daemon/internal/config"
	bitcoinDaemon "github.com/finexblock-dev/gofinexblock/cmd/bitcoin_daemon/internal/task"
	"github.com/finexblock-dev/gofinexblock/finexblock/daemon"
	"github.com/finexblock-dev/gofinexblock/finexblock/database"
	"github.com/finexblock-dev/gofinexblock/finexblock/gen/bitcoin"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
	"time"
)

func BitcoinBootstrap() {

	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("failed to load env file : %v", err)
	}

	var configurations *config.BitcoinConfiguration
	var configClient *config.BitcoinConfig
	var proxyClient bitcoin.BitcoinProxyClient
	var cluster *redis.ClusterClient
	var db *gorm.DB
	var err error

	configClient, err = config.NewBitcoinConfig()
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
	proxyClient, err = client.NewBitcoinProxyClient(configurations.ProxyHost)

	deposit := bitcoinDaemon.NewDeposit(proxyClient, cluster, db, time.Second*10, configurations.HotWallet)
	withdrawal := bitcoinDaemon.NewWithdrawal(proxyClient, cluster, db, time.Second*10)

	go daemon.Run(deposit)
	go daemon.Run(withdrawal)
}
