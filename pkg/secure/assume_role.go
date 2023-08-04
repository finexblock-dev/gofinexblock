package secure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func NewSTSClient(sess *session.Session) *sts.STS {
	return sts.New(sess)
}

func AssumeRole(svc *sts.STS, roleArn, sessionName string, duration int64) (*sts.AssumeRoleOutput, error) {
	return svc.AssumeRole(&sts.AssumeRoleInput{
		RoleArn:         aws.String(roleArn),
		RoleSessionName: aws.String(sessionName),
	})
}