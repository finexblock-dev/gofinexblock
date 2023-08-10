package config

import (
	"github.com/finexblock-dev/gofinexblock/pkg/database"
	"os"
	"strings"
)

func SSHConfig() *database.SSHConfig {
	return &database.SSHConfig{
		SSHHost: os.Getenv("SSH_HOST"),
		SSHUser: os.Getenv("SSH_USER"),
		SSHPem:  os.Getenv("PEM"),
		SSHPort: 22,
	}
}

func MySQLConfig() *database.MySqlConfig {
	return &database.MySqlConfig{
		MySqlHost: os.Getenv("MYSQL_HOST"),
		MySqlUser: os.Getenv("MYSQL_USER"),
		MySqlPass: os.Getenv("MYSQL_PASS"),
		MySqlDB:   os.Getenv("MYSQL_DB"),
		MySqlPort: "6033",
	}
}

func RedisConfig() *database.RedisClusterConfig {
	return &database.RedisClusterConfig{
		RedisHost: strings.Split(os.Getenv("REDIS_HOST"), ","),
		RedisUser: os.Getenv("REDIS_USER"),
		RedisPass: os.Getenv("REDIS_PASS"),
	}
}
