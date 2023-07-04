package database

import (
	"crypto/tls"
	"github.com/redis/go-redis/v9"
	"time"
)

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

func CreateRedisClient(addr, password string, dbNumber int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Network:     "tcp",
		Addr:        addr,
		Password:    password,
		DB:          dbNumber,
		ReadTimeout: 5 * time.Second,
	})
}