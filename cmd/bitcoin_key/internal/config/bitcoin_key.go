package config

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/finexblock-dev/gofinexblock/pkg/secure"
	"log"
	"strings"
)

type BitcoinKeyConfig struct {
	sess *session.Session
}

type BitcoinKeyConfiguration struct {
	WalletType    string `json:"WALLET_TYPE,omitempty"`
	WalletAccount string `json:"WALLET_ACCOUNT,omitempty"`
	RpcPort       int    `json:"RPC_PORT,omitempty"`
	RpcHost       string `json:"RPC_HOST,omitempty"`
	RpcUser       string `json:"RPC_USER,omitempty"`
	RpcPassword   string `json:"RPC_PASSWORD,omitempty"`
	Mnemonic      string `json:"MNEMONIC,omitempty"`
}

func NewBitcoinKeyConfig() (*BitcoinKeyConfig, error) {
	sess, err := secure.GetSessionFromEnv()
	if err != nil {
		return nil, err
	}
	return &BitcoinKeyConfig{sess: sess}, nil
}

func (b *BitcoinKeyConfig) SetSess(sess *session.Session) {
	b.sess = sess
}

func (b *BitcoinKeyConfig) AssumeRole(roleArn string, sessionName string) (*sts.AssumeRoleOutput, error) {
	stsClient := secure.NewSTSClient(b.sess)
	return secure.AssumeRole(stsClient, roleArn, sessionName, 900)
}

func (b *BitcoinKeyConfig) GetSecretValue(secretName string) (*secretsmanager.GetSecretValueOutput, error) {
	secretClient := secure.NewSecretMangerClient(b.sess)
	return secure.GetSecretValue(secretClient, secretName)
}

func (b *BitcoinKeyConfig) Decrypt(cipherText []byte, keyID string) (*kms.DecryptOutput, error) {
	kmsClient := secure.NewKMSClient(b.sess)
	return secure.Decrypt(kmsClient, cipherText, keyID)
}

func (b *BitcoinKeyConfig) Credentials(secretName string, target interface{}) {
	var assumeRoleStruct *secure.AssumeRoleStruct

	// Get secret value for assume role
	value, err := b.GetSecretValue(secretName)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal to assume role struct
	if err := json.Unmarshal([]byte(*value.SecretString), &assumeRoleStruct); err != nil {
		log.Fatal(err)
	}

	// Assume role
	role, err := b.AssumeRole(assumeRoleStruct.RoleArn, assumeRoleStruct.SessionName)
	if err != nil {
		log.Fatal(err)
	}

	// Set new session
	sess, err := secure.GetSession(*role.Credentials.AccessKeyId, *role.Credentials.SecretAccessKey, *role.Credentials.SessionToken)
	if err != nil {
		log.Fatal(err)
	}
	b.SetSess(sess)

	var cred, secret string
	cred, secret = strings.Split(assumeRoleStruct.SecretName, ",")[0], strings.Split(assumeRoleStruct.SecretName, ",")[1]

	// Get real secret value
	value, err = b.GetSecretValue(cred)
	if err != nil {
		log.Fatal(err)
	}

	decrypt, err := b.Decrypt(value.SecretBinary, assumeRoleStruct.KeyID)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(decrypt.Plaintext, &target); err != nil {
		log.Fatal(err)
	}

	// Get real secret value
	value, err = b.GetSecretValue(secret)
	if err != nil {
		log.Fatal(err)
	}

	decrypt, err = b.Decrypt(value.SecretBinary, assumeRoleStruct.KeyID)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(decrypt.Plaintext, &target); err != nil {
		log.Fatal(err)
	}
}