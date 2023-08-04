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

type BitcoinConfig struct {
	sess *session.Session
}

type BitcoinConfiguration struct {
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

func NewBitcoinConfig() (*BitcoinConfig, error) {
	sess, err := secure.GetSessionFromEnv()
	if err != nil {
		return nil, err
	}
	return &BitcoinConfig{sess: sess}, nil
}

func (b *BitcoinConfig) SetSess(sess *session.Session) {
	b.sess = sess
}

func (b *BitcoinConfig) AssumeRole(roleArn string, sessionName string) (*sts.AssumeRoleOutput, error) {
	stsClient := secure.NewSTSClient(b.sess)
	return secure.AssumeRole(stsClient, roleArn, sessionName, 900)
}

func (b *BitcoinConfig) GetSecretValue(secretName string) (*secretsmanager.GetSecretValueOutput, error) {
	secretClient := secure.NewSecretMangerClient(b.sess)
	return secure.GetSecretValue(secretClient, secretName)
}

func (b *BitcoinConfig) Decrypt(cipherText []byte, keyID string) (*kms.DecryptOutput, error) {
	kmsClient := secure.NewKMSClient(b.sess)
	return secure.Decrypt(kmsClient, cipherText, keyID)
}

func (b *BitcoinConfig) Credentials(secretName string, target interface{}) {
	var assumeRoleStruct *secure.AssumeRoleStruct
	var value *secretsmanager.GetSecretValueOutput
	var err error
	// Get secret value for assume role struct
	value, err = b.GetSecretValue(secretName)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal
	if err = json.Unmarshal([]byte(*value.SecretString), &assumeRoleStruct); err != nil {
		log.Fatal(err)
	}

	// Assume role
	role, err := b.AssumeRole(assumeRoleStruct.RoleArn, assumeRoleStruct.SessionName)
	if err != nil {
		log.Fatal(err)
	}

	// Set new session using assumed role credentials
	sess, err := secure.GetSession(*role.Credentials.AccessKeyId, *role.Credentials.SecretAccessKey, *role.Credentials.SessionToken)
	if err != nil {
		log.Fatal(err)
	}

	// Set session
	b.SetSess(sess)

	// Get secret value for credentials
	value, err = b.GetSecretValue(assumeRoleStruct.SecretName)
	if err != nil {
		log.Fatal(err)
	}

	// Decrypt credentials
	decrypted, err := b.Decrypt(value.SecretBinary, assumeRoleStruct.KeyID)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal
	if err = json.Unmarshal(decrypted.Plaintext, &target); err != nil {
		log.Fatal(err)
	}
}