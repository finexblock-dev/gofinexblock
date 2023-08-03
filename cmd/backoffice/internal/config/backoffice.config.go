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

type BackOfficeConfiguration struct {
	RedisHost     string `json:"REDIS_HOST,omitempty"`
	RedisPort     string `json:"REDIS_PORT,omitempty"`
	RedisUser     string `json:"REDIS_USER,omitempty"`
	RedisPass     string `json:"REDIS_PASS,omitempty"`
	MysqlHost     string `json:"MYSQL_HOST,omitempty"`
	MysqlPort     string `json:"MYSQL_PORT,omitempty"`
	MysqlUser     string `json:"MYSQL_USER,omitempty"`
	MysqlPass     string `json:"MYSQL_PASS,omitempty"`
	MysqlDB       string `json:"MYSQL_DB,omitempty"`
	JwtSecret     string `json:"JWT_SECRET,omitempty"`
	S3ImagePath   string `json:"S3_IMAGE_PATH,omitempty"`
	S3ImageBucket string `json:"S3_IMAGE_BUCKET,omitempty"`
}

type ConfigClient struct {
	sess *session.Session
}

func (c *ConfigClient) SetSess(sess *session.Session) {
	c.sess = sess
}

func (c *ConfigClient) AssumeRole(roleArn string, sessionName string) (*sts.AssumeRoleOutput, error) {
	svc := sts.New(c.sess)
	return secure.AssumeRole(svc, roleArn, sessionName, 900)
}

func (c *ConfigClient) GetSecretValue(secretName string) (*secretsmanager.GetSecretValueOutput, error) {
	svc := secretsmanager.New(c.sess)
	return secure.GetSecretValue(svc, secretName)
}

func (c *ConfigClient) Decrypt(cipherText []byte, keyID string) (*kms.DecryptOutput, error) {
	svc := kms.New(c.sess)
	return secure.Decrypt(svc, cipherText, keyID)
}

func (c *ConfigClient) Credentials(secretName string, target interface{}) {
	var assumeRoleStruct *secure.AssumeRoleStruct
	var secret *secretsmanager.GetSecretValueOutput
	var output *kms.DecryptOutput
	var role *sts.AssumeRoleOutput
	var sess *session.Session
	var err error

	sess, err = secure.GetSessionFromEnv()
	if err != nil {
		log.Panicln(err)
	}

	secret, err = c.GetSecretValue(secretName)
	if err != nil {
		log.Panicln(err)
	}

	if err = json.Unmarshal([]byte(*secret.SecretString), &assumeRoleStruct); err != nil {
		log.Panicln(err)
	}

	role, err = c.AssumeRole(assumeRoleStruct.RoleArn, assumeRoleStruct.SessionName)
	if err != nil {
		log.Panicln(err)
	}

	sess, err = secure.GetSession(*role.Credentials.AccessKeyId, *role.Credentials.SecretAccessKey, *role.Credentials.SessionToken)
	if err != nil {
		log.Panicln(err)
	}

	c.SetSess(sess)

	secret, err = c.GetSecretValue(assumeRoleStruct.SecretName)
	if err != nil {
		log.Panicln(err)
	}

	output, err = c.Decrypt(secret.SecretBinary, assumeRoleStruct.KeyID)
	if err != nil {
		log.Panicln(err)
	}

	if err = json.Unmarshal(output.Plaintext, &target); err != nil {
		log.Panicln(err)
	}
}