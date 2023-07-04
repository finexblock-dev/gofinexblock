package secure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func NewSecretMangerClient(sess *session.Session) *secretsmanager.SecretsManager {
	return secretsmanager.New(sess)
}

func GetSecretValue(svc *secretsmanager.SecretsManager, secretName string) (*secretsmanager.GetSecretValueOutput, error) {
	return svc.GetSecretValue(&secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
}

func CreateSecretValue(svc *secretsmanager.SecretsManager, name, value string) (*secretsmanager.CreateSecretOutput, error) {
	return svc.CreateSecret(&secretsmanager.CreateSecretInput{
		Name:         aws.String(name),
		SecretString: aws.String(value),
	})
}