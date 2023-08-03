package config

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/finexblock-dev/gofinexblock/pkg/secure"
)

type MatchingEngineConfiguration struct {
	RedisHost string `json:"REDIS_HOST,omitempty"`
	RedisPort string `json:"REDIS_PORT,omitempty"`
	RedisUser string `json:"REDIS_USER,omitempty"`
	RedisPass string `json:"REDIS_PASS,omitempty"`
	MysqlHost string `json:"MYSQL_HOST,omitempty"`
	MysqlPort string `json:"MYSQL_PORT,omitempty"`
	MysqlUser string `json:"MYSQL_USER,omitempty"`
	MysqlPass string `json:"MYSQL_PASS,omitempty"`
	MysqlDB   string `json:"MYSQL_DB,omitempty"`
}

type MatchingEngineConfig struct {
	sess *session.Session
}

func NewGrpcServerConfig() (*MatchingEngineConfig, error) {
	sess, err := secure.GetSessionFromEnv()
	if err != nil {
		return nil, err
	}
	return &MatchingEngineConfig{sess: sess}, nil
}

func (m *MatchingEngineConfig) SetSess(sess *session.Session) {
	m.sess = sess
}

func (m *MatchingEngineConfig) AssumeRole(roleArn string, sessionName string) (*sts.AssumeRoleOutput, error) {
	svc := secure.NewSTSClient(m.sess)
	return secure.AssumeRole(svc, roleArn, sessionName, 900)
}

func (m *MatchingEngineConfig) GetSecretValue(secretName string) (*secretsmanager.GetSecretValueOutput, error) {
	svc := secure.NewSecretMangerClient(m.sess)
	return secure.GetSecretValue(svc, secretName)
}

func (m *MatchingEngineConfig) Decrypt(cipherText []byte, keyID string) (*kms.DecryptOutput, error) {
	svc := secure.NewKMSClient(m.sess)
	return secure.Decrypt(svc, cipherText, keyID)
}

func (m *MatchingEngineConfig) Credentials(secretName string, target interface{}) {
	var assumeRoleStruct *secure.AssumeRoleStruct
	var secret *secretsmanager.GetSecretValueOutput
	var assumedRole *sts.AssumeRoleOutput
	var decrypted *kms.DecryptOutput
	var sess *session.Session
	var err error

	if secret, err = m.GetSecretValue(secretName); err != nil {
		panic(err)
	}

	if err = json.Unmarshal([]byte(*secret.SecretString), &assumeRoleStruct); err != nil {
		panic(err)
	}

	if assumedRole, err = m.AssumeRole(assumeRoleStruct.RoleArn, assumeRoleStruct.SessionName); err != nil {
		panic(err)
	}

	if sess, err = secure.GetSession(*assumedRole.Credentials.AccessKeyId, *assumedRole.Credentials.SecretAccessKey, *assumedRole.Credentials.SessionToken); err != nil {
		panic(err)
	}

	m.SetSess(sess)

	if secret, err = m.GetSecretValue(assumeRoleStruct.SecretName); err != nil {
		panic(err)
	}

	if decrypted, err = m.Decrypt(secret.SecretBinary, assumeRoleStruct.KeyID); err != nil {
		panic(err)
	}

	if err = json.Unmarshal(decrypted.Plaintext, &target); err != nil {
		panic(err)
	}
}