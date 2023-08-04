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

type PolygonConfig struct {
	sess *session.Session
}

type PolygonConfiguration struct {
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

func NewPolygonConfig() (*PolygonConfig, error) {
	sess, err := secure.GetSessionFromEnv()
	if err != nil {
		return nil, err
	}
	return &PolygonConfig{sess: sess}, nil
}

func (p *PolygonConfig) SetSess(sess *session.Session) {
	p.sess = sess
}

func (p *PolygonConfig) AssumeRole(roleArn string, sessionName string) (*sts.AssumeRoleOutput, error) {
	stsClient := secure.NewSTSClient(p.sess)
	return secure.AssumeRole(stsClient, roleArn, sessionName, 900)
}

func (p *PolygonConfig) GetSecretValue(secretName string) (*secretsmanager.GetSecretValueOutput, error) {
	secretClient := secure.NewSecretMangerClient(p.sess)
	return secure.GetSecretValue(secretClient, secretName)
}

func (p *PolygonConfig) Decrypt(cipherText []byte, keyID string) (*kms.DecryptOutput, error) {
	kmsClient := secure.NewKMSClient(p.sess)
	return secure.Decrypt(kmsClient, cipherText, keyID)
}

func (p *PolygonConfig) Credentials(secretName string, target interface{}) {
	var assumeRoleStruct *secure.AssumeRoleStruct
	var value *secretsmanager.GetSecretValueOutput
	var err error

	// Get secret value for assume role struct
	value, err = p.GetSecretValue(secretName)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal
	if err = json.Unmarshal([]byte(*value.SecretString), &assumeRoleStruct); err != nil {
		log.Fatal(err)
	}

	// Assume role
	role, err := p.AssumeRole(assumeRoleStruct.RoleArn, assumeRoleStruct.SessionName)
	if err != nil {
		log.Fatal(err)
	}

	// Set new session using assumed role credentials
	sess, err := secure.GetSession(*role.Credentials.AccessKeyId, *role.Credentials.SecretAccessKey, *role.Credentials.SessionToken)
	if err != nil {
		log.Fatal(err)
	}

	// Set session
	p.SetSess(sess)

	// Get secret value for credentials
	value, err = p.GetSecretValue(assumeRoleStruct.SecretName)
	if err != nil {
		log.Fatal(err)
	}

	// Decrypt credentials
	decrypted, err := p.Decrypt(value.SecretBinary, assumeRoleStruct.KeyID)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal
	if err = json.Unmarshal(decrypted.Plaintext, &target); err != nil {
		log.Fatal(err)
	}
}