package config

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/finexblock-dev/gofinexblock/pkg/secure"
	"os"
)

type EventSubscriberConfiguration struct {
	MysqlHost     string `json:"MYSQL_HOST"`
	MysqlUser     string `json:"MYSQL_USER"`
	MysqlPass     string `json:"MYSQL_PASS"`
	MysqlPort     string `json:"MYSQL_PORT"`
	MysqlDatabase string `json:"MYSQL_DB"`
	RedisHost     string `json:"REDIS_HOST"`
	RedisPort     string `json:"REDIS_PORT"`
	RedisUser     string `json:"REDIS_USER"`
	RedisPass     string `json:"REDIS_PASS"`
}

type EventSubscriberConfig struct {
	sess *session.Session
}

func NewEventSubscriberConfig() (*EventSubscriberConfig, error) {
	sess, err := secure.GetSessionFromEnv()
	if err != nil {
		return nil, err
	}
	return &EventSubscriberConfig{sess: sess}, nil
}

func (e *EventSubscriberConfig) SetSess(sess *session.Session) {
	e.sess = sess
}

func (e *EventSubscriberConfig) AssumeRole(roleArn string, sessionName string) (*sts.AssumeRoleOutput, error) {
	svc := secure.NewSTSClient(e.sess)
	return secure.AssumeRole(svc, roleArn, sessionName, 900)
}

func (e *EventSubscriberConfig) GetSecretValue(secretName string) (*secretsmanager.GetSecretValueOutput, error) {
	svc := secure.NewSecretMangerClient(e.sess)
	return secure.GetSecretValue(svc, secretName)
}

func (e *EventSubscriberConfig) Decrypt(cipherText []byte, keyID string) (*kms.DecryptOutput, error) {
	svc := secure.NewKMSClient(e.sess)
	return secure.Decrypt(svc, cipherText, keyID)
}

func (e *EventSubscriberConfig) Credentials(secretName string, target interface{}) {
	var assumeRoleStruct *secure.AssumeRoleStruct
	var secret *secretsmanager.GetSecretValueOutput
	var assumedRole *sts.AssumeRoleOutput
	var decrypted *kms.DecryptOutput
	var sess *session.Session
	var err error

	if os.Getenv("APPMODE") == "LOCAL" {
		target = &EventSubscriberConfiguration{
			MysqlHost:     os.Getenv("MYSQL_HOST"),
			MysqlUser:     os.Getenv("MYSQL_USER"),
			MysqlPass:     os.Getenv("MYSQL_PASS"),
			MysqlPort:     os.Getenv("MYSQL_PORT"),
			MysqlDatabase: os.Getenv("MYSQL_DB"),
		}
		return
	}

	if secret, err = e.GetSecretValue(secretName); err != nil {
		panic(err)
	}

	if err = json.Unmarshal([]byte(*secret.SecretString), &assumeRoleStruct); err != nil {
		panic(err)
	}

	if assumedRole, err = e.AssumeRole(assumeRoleStruct.RoleArn, assumeRoleStruct.SessionName); err != nil {
		panic(err)
	}

	if sess, err = secure.GetSession(*assumedRole.Credentials.AccessKeyId, *assumedRole.Credentials.SecretAccessKey, *assumedRole.Credentials.SessionToken); err != nil {
		panic(err)
	}

	e.SetSess(sess)

	if secret, err = e.GetSecretValue(assumeRoleStruct.SecretName); err != nil {
		panic(err)
	}

	if decrypted, err = e.Decrypt(secret.SecretBinary, assumeRoleStruct.KeyID); err != nil {
		panic(err)
	}

	if err = json.Unmarshal(decrypted.Plaintext, target); err != nil {
		panic(err)
	}
}