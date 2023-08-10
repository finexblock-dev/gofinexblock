package internal

import (
	"fmt"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/config"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/handler"
	"github.com/finexblock-dev/gofinexblock/cmd/backoffice/internal/router"
	"github.com/finexblock-dev/gofinexblock/pkg/database"
	"github.com/finexblock-dev/gofinexblock/pkg/gen/health"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
)

func Bootstrap() {
	//var configClient = new(config.ConfigClient)
	//var configuration = new(config.BackOfficeConfiguration)
	var db *gorm.DB
	var cluster *redis.ClusterClient

	redisCfg := config.RedisConfig()
	mysqlCfg := config.MySQLConfig()

	if os.Getenv("APPMODE") == "LOCAL" {
		//configClient.Credentials(os.Getenv("SECRET_NAME"), configuration)
		sshCfg := config.SSHConfig()
		db = database.GetTunnelledMySQLV2(sshCfg, mysqlCfg)
		cluster = database.CreateRedisClusterWithSSH(sshCfg, redisCfg)
	} else {
		db = database.CreateMySQLClient(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASS"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DB"))
		cluster = database.CreateRedisCluster(redisCfg.RedisHost, redisCfg.RedisUser, redisCfg.RedisPass)
	}

	go func() {
		conn, err := net.Listen("tcp", fmt.Sprintf(":%v", 50051))
		if err != nil {
			log.Fatalf("failed to listen on port `%v` : %v", 50051, err)
		}

		grpcServer := grpc.NewServer()
		server := handler.NewServer()
		health.RegisterHealthCheckServer(grpcServer, server)
		reflection.Register(grpcServer)

		if err = grpcServer.Serve(conn); err != nil {
			log.Fatalf("failed to listen on port `%v` : %v", 50051, err)
		}
	}()

	app := router.Router(db, cluster)
	log.Println(os.Getenv("PORT"))
	if err := app.Listen(fmt.Sprintf(":%v", os.Getenv("PORT"))); err != nil {
		log.Fatalf("failed to listen on port `%v` : %v", os.Getenv("PORT"), err)
	}
}
