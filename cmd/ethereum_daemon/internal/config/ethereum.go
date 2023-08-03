package config

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/finexblock-dev/gofinexblock/pkg/secure"
	"log"
)

type EthereumConfig struct {
	sess *session.Session
}

type EthereumConfiguration struct {
	MysqlUser     string `json:"MYSQL_USER"`
	MysqlPass     string `json:"MYSQL_PASS"`
	MysqlHost     string `json:"MYSQL_HOST"`
	MysqlPort     string `json:"MYSQL_PORT"`
	MysqlDatabase string `json:"MYSQL_DATABASE"`

	RedisHost string `json:"REDIS_HOST"`
	RedisPort string `json:"REDIS_PORT"`
	RedisUser string `json:"REDIS_USER"`
	RedisPass string `json:"REDIS_PASS"`

	ProxyHost string `json:"PROXY_HOST"`

	HotWallet string `json:"HOT_WALLET"`
}

func NewEthereumConfig() (*EthereumConfig, error) {
	sess, err := secure.GetSessionFromEnv()
	if err != nil {
		return nil, err
	}
	return &EthereumConfig{sess: sess}, nil
}

func (e *EthereumConfig) SetSess(sess *session.Session) {
	e.sess = sess
}

func (e *EthereumConfig) AssumeRole(roleArn string, sessionName string) (*sts.AssumeRoleOutput, error) {
	stsClient := secure.NewSTSClient(e.sess)
	return secure.AssumeRole(stsClient, roleArn, sessionName, 900)
}

func (e *EthereumConfig) GetSecretValue(secretName string) (*secretsmanager.GetSecretValueOutput, error) {
	secretClient := secure.NewSecretMangerClient(e.sess)
	return secure.GetSecretValue(secretClient, secretName)
}

func (e *EthereumConfig) Decrypt(cipherText []byte, keyID string) (*kms.DecryptOutput, error) {
	kmsClient := secure.NewKMSClient(e.sess)
	return secure.Decrypt(kmsClient, cipherText, keyID)
}

func (e *EthereumConfig) Credentials(secretName string, target interface{}) {
	var assumeRoleStruct *secure.AssumeRoleStruct
	var value *secretsmanager.GetSecretValueOutput
	var err error
	// Get secret value for assume role struct
	value, err = e.GetSecretValue(secretName)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal
	if err = json.Unmarshal([]byte(*value.SecretString), &assumeRoleStruct); err != nil {
		log.Fatal(err)
	}

	// Assume role
	role, err := e.AssumeRole(assumeRoleStruct.RoleArn, assumeRoleStruct.SessionName)
	if err != nil {
		log.Fatal(err)
	}

	// Set new session using assumed role credentials
	sess, err := secure.GetSession(*role.Credentials.AccessKeyId, *role.Credentials.SecretAccessKey, *role.Credentials.SessionToken)
	if err != nil {
		log.Fatal(err)
	}

	// Set session
	e.SetSess(sess)

	// Get secret value for credentials
	value, err = e.GetSecretValue(assumeRoleStruct.SecretName)
	if err != nil {
		log.Fatal(err)
	}

	// Decrypt credentials
	decrypted, err := e.Decrypt(value.SecretBinary, assumeRoleStruct.KeyID)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal
	if err = json.Unmarshal(decrypted.Plaintext, &target); err != nil {
		log.Fatal(err)
	}
}