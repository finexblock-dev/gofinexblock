package config

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/finexblock-dev/gofinexblock/pkg/secure"
	"log"
	"os"
)

type PolygonKeyConfig struct {
	sess *session.Session
}

type PolygonKeyAssumeRoleStruct struct {
	RoleArn      string `json:"ROLE_ARN,omitempty"`
	SecretFirst  string `json:"SECRET_FIRST,omitempty"`
	SecretSecond string `json:"SECRET_SECOND,omitempty"`
	Credentials  string `json:"SECRET_ENV,omitempty"`
	SessionName  string `json:"SESSION_NAME,omitempty"`
	KeyID        string `json:"KEY_ID,omitempty"`
}

type PolygonKeyConfiguration struct {
	First    string `json:"FIRST,omitempty"`
	Second   string `json:"SECOND,omitempty"`
	Endpoint string `json:"ENDPOINT,omitempty"`
	Token    string `json:"TOKEN,omitempty"`
}

func NewPolygonKeyConfig() (*PolygonKeyConfig, error) {
	sess, err := secure.GetSessionFromEnv()
	if err != nil {
		return nil, err
	}
	return &PolygonKeyConfig{sess: sess}, nil
}

func (e *PolygonKeyConfig) SetSess(sess *session.Session) {
	e.sess = sess
}

func (e *PolygonKeyConfig) AssumeRole(roleArn string, sessionName string) (*sts.AssumeRoleOutput, error) {
	stsClient := secure.NewSTSClient(e.sess)
	return secure.AssumeRole(stsClient, roleArn, sessionName, 900)
}

func (e *PolygonKeyConfig) GetSecretValue(secretName string) (*secretsmanager.GetSecretValueOutput, error) {
	secretClient := secure.NewSecretMangerClient(e.sess)
	return secure.GetSecretValue(secretClient, secretName)
}

func (e *PolygonKeyConfig) Decrypt(cipherText []byte, keyID string) (*kms.DecryptOutput, error) {
	kmsClient := secure.NewKMSClient(e.sess)
	return secure.Decrypt(kmsClient, cipherText, keyID)
}

func (e *PolygonKeyConfig) Credentials(secretName string, target interface{}) {
	var assumeRoleStruct *PolygonKeyAssumeRoleStruct

	if os.Getenv("APPMODE") == "LOCAL" {
		return
	}

	// Get secret value for assume role
	value, err := e.GetSecretValue(secretName)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal to assume role struct
	if err := json.Unmarshal([]byte(*value.SecretString), &assumeRoleStruct); err != nil {
		log.Fatal(err)
	}

	// Assume role
	role, err := e.AssumeRole(assumeRoleStruct.RoleArn, assumeRoleStruct.SessionName)
	if err != nil {
		log.Fatal(err)
	}

	// Set new session
	sess, err := secure.GetSession(*role.Credentials.AccessKeyId, *role.Credentials.SecretAccessKey, *role.Credentials.SessionToken)
	if err != nil {
		log.Fatal(err)
	}
	e.SetSess(sess)

	// Get secret first
	value, err = e.GetSecretValue(assumeRoleStruct.SecretFirst)
	if err != nil {
		log.Fatal(err)
	}

	// Decrypt secret first
	decrypt, err := e.Decrypt(value.SecretBinary, assumeRoleStruct.KeyID)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(decrypt.Plaintext, &target); err != nil {
		log.Fatal(err)
	}

	// Get secret second
	value, err = e.GetSecretValue(assumeRoleStruct.SecretSecond)
	if err != nil {
		log.Fatal(err)
	}

	// Decrypt secret second
	decrypt, err = e.Decrypt(value.SecretBinary, assumeRoleStruct.KeyID)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(decrypt.Plaintext, &target); err != nil {
		log.Fatal(err)
	}

	// Get credentials
	value, err = e.GetSecretValue(assumeRoleStruct.Credentials)
	if err != nil {
		log.Fatal(err)
	}

	// Decrypt credentials
	decrypt, err = e.Decrypt(value.SecretBinary, assumeRoleStruct.KeyID)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(decrypt.Plaintext, &target); err != nil {
		log.Fatal(err)
	}
}