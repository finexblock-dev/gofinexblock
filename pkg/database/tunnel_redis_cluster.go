package database

import (
	"context"
	"crypto/tls"
	"github.com/redis/go-redis/v9"
	"net"
)

func CreateRedisClusterWithSSH(sshCfg *SSHConfig, redisCfg *RedisClusterConfig) *redis.ClusterClient {
	sshConfig := clientConfig(sshCfg.SSHUser, sshCfg.SSHPem)

	setAgentClient(sshCfg, sshConfig)

	conn := sshConnection(sshCfg, sshConfig)

	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:     redisCfg.RedisHost,
		Username:  redisCfg.RedisUser,
		Password:  redisCfg.RedisPass,
		TLSConfig: &tls.Config{InsecureSkipVerify: true},
		Dialer: func(c context.Context, network, addr string) (net.Conn, error) {
			dialer := &ViaSSHDialer{client: conn}
			return dialer.Dial(addr)
		},
	})
}
