package secure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
)

func GetSessionFromEnv() (*session.Session, error) {
	return session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
			AccessKeyID:     os.Getenv("AWS_ACCESS_KEY"),
			SecretAccessKey: os.Getenv("AWS_SECRET_KEY"),
		}),
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
}

func GetSession(accessKey, secretKey, sessionToken string) (*session.Session, error) {
	return session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
			AccessKeyID:     accessKey,
			SecretAccessKey: secretKey,
			SessionToken:    sessionToken,
		}),
	})
}
