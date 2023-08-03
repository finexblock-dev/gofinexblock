package database

import (
	"context"
	"fmt"
	gomysql "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net"
	"os"
	"time"
)

func GetTunnelledMySQL(sshHost, sshUser, sshPassword, remoteHost, remoteUser, remotePassword, remoteDatabase string, sshPort, remotePort int) *gorm.DB {
	sshConfig := getSSHConfig(sshUser, sshPassword)

	var agentClient agent.Agent
	// Establish a connection to the local ssh-agent
	if conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
				log.Fatalf("Failed to close connection")
			}
		}(conn)

		// Create a new instance of the ssh agent
		agentClient = agent.NewClient(conn)
	}

	// When the agentClient connection succeeded, add them as AuthMethod
	if agentClient != nil {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PublicKeysCallback(agentClient.Signers))
	}
	// When there's a non-empty password add the password AuthMethod
	if sshPassword != "" {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PasswordCallback(func() (string, error) {
			return sshPassword, nil
		}))
	}

	sshConnection, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", sshHost, sshPort), sshConfig)
	if err != nil {
		log.Fatalf("Failed to dial through tcp : %v", err)
	}

	gomysql.RegisterDialContext("mysql+tcp", func(ctx context.Context, addr string) (net.Conn, error) {
		dialer := &ViaSSHDialer{client: sshConnection}
		return dialer.Dial(addr)
	})

	dsn := fmt.Sprintf("%v:%v@mysql+tcp(%v:%v)/%v?parseTime=true", remoteUser, remotePassword, remoteHost, remotePort, remoteDatabase)
	timezone := "Asia/Seoul"
	_, err = time.LoadLocation(timezone)
	if err != nil {
		log.Panicf("Error loading location: %v", err.Error())
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info, // Log level
				Colorful: true,        // Disable color
			}),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		log.Panicf("Error opening connection: %v", err.Error())
	}
	log.Println("GET INSTANCE DONE")
	return db
}