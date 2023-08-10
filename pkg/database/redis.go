package database

import (
	"crypto/tls"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisClusterConfig struct {
	RedisHost []string
	RedisUser string
	RedisPass string
}

type RedisConfig struct {
	RedisHost      string
	RedisUser      string
	RedisPass      string
	DatabaseNumber int
}

func CreateRedisCluster(addresses []string, user, password string) *redis.ClusterClient {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:       addresses,
		Username:    user,
		Password:    password,
		TLSConfig:   &tls.Config{InsecureSkipVerify: true},
		PoolSize:    2500,
		PoolTimeout: time.Second,
	})
}

func CreateRedisClusterV2(cfg *RedisClusterConfig) *redis.ClusterClient {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:       cfg.RedisHost,
		Username:    cfg.RedisUser,
		Password:    cfg.RedisPass,
		TLSConfig:   &tls.Config{InsecureSkipVerify: true},
		PoolSize:    2500,
		PoolTimeout: time.Second,
	})
}

func CreateRedisClient(addr, password string, dbNumber int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Network:     "tcp",
		Addr:        addr,
		Password:    password,
		DB:          dbNumber,
		ReadTimeout: 5 * time.Second,
	})
}

func CreateRedisClientV2(cfg *RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Network:     "tcp",
		Addr:        cfg.RedisHost,
		Username:    cfg.RedisUser,
		Password:    cfg.RedisPass,
		DB:          cfg.DatabaseNumber,
		ReadTimeout: 5 * time.Second,
	})
}
