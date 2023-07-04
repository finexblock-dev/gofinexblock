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

type ViaSSHDialer struct {
	client *ssh.Client
}

func (d *ViaSSHDialer) Dial(addr string) (net.Conn, error) {
	return d.client.Dial("tcp", addr)
}

func unixSSHDial() (agent.Agent, error) {
	var client agent.Agent
	conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		log.Fatalf("Failed to unix dial : %v", err)
	}
	client = agent.NewClient(conn)
	return client, nil
}

func tcpSSHDial(sshHost string, sshPort int, sshConfig *ssh.ClientConfig) (*ssh.Client, error) {
	return ssh.Dial("tcp", fmt.Sprintf("%v:%v", sshHost, sshPort), sshConfig)
}

func getSSHConfig(sshUser, pem string) *ssh.ClientConfig {

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

func getSSHConnection(sshHost, sshUser, sshPassword string, sshPort int) (*ssh.Client, error) {
	sshConfig := getSSHConfig(sshUser, sshPassword)

	var agentClient agent.Agent
	// Establish a connection to the local ssh-agent
	if conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		defer conn.Close()

		// Create a new instance of the ssh agent
		agentClient = agent.NewClient(conn)
	}

	// When the agentClient connection succeeded, add them as AuthMethod
	if agentClient != nil {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PublicKeysCallback(agentClient.Signers))
	}
	// When there's a non empty password add the password AuthMethod
	if sshPassword != "" {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PasswordCallback(func() (string, error) {
			return sshPassword, nil
		}))
	}

	return ssh.Dial("tcp", fmt.Sprintf("%v:%v", sshHost, sshPort), sshConfig)
}