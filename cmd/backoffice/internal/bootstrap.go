package internal

import (
	"fmt"
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
	"strings"
)

func Bootstrap() {
	//var configClient = new(config.ConfigClient)
	//var configuration = new(config.BackOfficeConfiguration)
	var db *gorm.DB
	var cluster *redis.ClusterClient

	//configClient.Credentials(os.Getenv("SECRET_NAME"), configuration)

	if os.Getenv("APPMODE") == "LOCAL" {
		db = database.GetTunnelledMySQL(
			os.Getenv("SSH_HOST"),
			os.Getenv("SSH_USER"),
			os.Getenv("PEM"),
			os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASS"),
			os.Getenv("MYSQL_DB"),
			22,
			6033,
		)
	} else {
		db = database.CreateMySQLClient(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASS"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DB"))
		cluster = database.CreateRedisCluster(strings.Split(os.Getenv("REDIS_HOST"), ","), os.Getenv("REDIS_USER"), os.Getenv("REDIS_PASS"))
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