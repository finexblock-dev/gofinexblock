package database

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

type SSHConfig struct {
	SSHHost string
	SSHUser string
	SSHPem  string
	SSHPort int
}

func NewSSHConfig(SSHHost string, SSHUser string, SSHPem string, SSHPort int) *SSHConfig {
	return &SSHConfig{SSHHost: SSHHost, SSHUser: SSHUser, SSHPem: SSHPem, SSHPort: SSHPort}
}

type ViaSSHDialer struct {
	client *ssh.Client
}

func (d *ViaSSHDialer) Dial(addr string) (net.Conn, error) {
	return d.client.Dial("tcp", addr)
}

func (d *ViaSSHDialer) SetDeadline(t time.Time) error {
	return nil
}

func (d *ViaSSHDialer) SetReadDeadline(t time.Time) error {
	return nil
}

func (d *ViaSSHDialer) SetWriteDeadline(t time.Time) error {
	return nil
}

func clientConfig(sshUser, pem string) *ssh.ClientConfig {

	key, err := ioutil.ReadFile(pem)
	if err != nil {
		log.Panicf("failed to open pem key file : %v", err)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Panicf("failed to parse private key : %v", err)
	}
	return &ssh.ClientConfig{
		User: sshUser,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second * 5,
	}
}

func setAgentClient(cfg *SSHConfig, sshConfig *ssh.ClientConfig) {
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
	if cfg.SSHPem != "" {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PasswordCallback(func() (string, error) {
			return cfg.SSHPem, nil
		}))
	}
}

func sshConnection(cfg *SSHConfig, sshClientConfig *ssh.ClientConfig) *ssh.Client {
	conn, err := ssh.Dial("tcp", fmt.Sprintf("%v:%v", cfg.SSHHost, cfg.SSHPort), sshClientConfig)
	if err != nil {
		log.Fatalf("Failed to dial through tcp : %v", err)
	}

	return conn
}
