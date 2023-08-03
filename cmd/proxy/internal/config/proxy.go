package config

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/finexblock-dev/gofinexblock/finexblock/secure"
	"log"
	"os"
)

type ProxyConfig struct {
	sess *session.Session
}

type ProxyConfiguration struct {
	BitcoinKeyServer  string `json:"BITCOIN_KEY,omitempty"`
	EthereumKeyServer string `json:"ETHEREUM_KEY,omitempty"`
	PolygonKeyServer  string `json:"POLYGON_KEY,omitempty"`
	Port              string `json:"PORT,omitempty"`
}

func (p *ProxyConfig) SetSess(sess *session.Session) {
	p.sess = sess
}

func NewProxyConfig() (*ProxyConfig, error) {
	sess, err := secure.GetSessionFromEnv()
	if err != nil {
		return nil, err
	}
	return &ProxyConfig{sess: sess}, nil
}

func (p *ProxyConfig) AssumeRole(roleArn string, sessionName string) (*sts.AssumeRoleOutput, error) {
	stsClient := secure.NewSTSClient(p.sess)
	return secure.AssumeRole(stsClient, roleArn, sessionName, 900)
}

func (p *ProxyConfig) GetSecretValue(secretName string) (*secretsmanager.GetSecretValueOutput, error) {
	secretClient := secure.NewSecretMangerClient(p.sess)
	return secure.GetSecretValue(secretClient, secretName)
}

func (p *ProxyConfig) Decrypt(cipherText []byte, keyID string) (*kms.DecryptOutput, error) {
	kmsClient := secure.NewKMSClient(p.sess)
	return secure.Decrypt(kmsClient, cipherText, keyID)
}

func (p *ProxyConfig) Credentials(secretName string, target interface{}) {
	var assumeRoleStruct *secure.AssumeRoleStruct

	if os.Getenv("APPMODE") == "LOCAL" {
		target = &ProxyConfiguration{
			EthereumKeyServer: "localhost:40041",
			PolygonKeyServer:  "localhost:30031",
			BitcoinKeyServer:  "localhost:20021",
			Port:              "50051",
		}
		return
	}
	// Get secret value for assume role
	value, err := p.GetSecretValue(secretName)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal to assume role struct
	if err := json.Unmarshal([]byte(*value.SecretString), &assumeRoleStruct); err != nil {
		log.Fatal(err)
	}

	// Assume role
	role, err := p.AssumeRole(assumeRoleStruct.RoleArn, assumeRoleStruct.SessionName)
	if err != nil {
		log.Fatal(err)
	}

	// Set new session
	sess, err := secure.GetSession(*role.Credentials.AccessKeyId, *role.Credentials.SecretAccessKey, *role.Credentials.SessionToken)
	if err != nil {
		log.Fatal(err)
	}
	p.SetSess(sess)

	// Get real secret value
	value, err = p.GetSecretValue(assumeRoleStruct.SecretName)
	if err != nil {
		log.Fatal(err)
	}

	decrypt, err := p.Decrypt(value.SecretBinary, assumeRoleStruct.KeyID)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(decrypt.Plaintext, &target); err != nil {
		log.Fatal(err)
	}
}
