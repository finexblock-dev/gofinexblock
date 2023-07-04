package secure

import (
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/sts"
)

type Interface interface {
	AssumeRole(roleArn string, sessionName string) (*sts.AssumeRoleOutput, error)   // AssumeRole is a method that assumes a role and returns a token that can be used to access the role for a limited time.
	GetSecretValue(secretName string) (*secretsmanager.GetSecretValueOutput, error) // GetSecretValue is a method that retrieves the contents of the encrypted fields SecretString or SecretBinary from the specified version of a secret, whichever contains content.
	Decrypt(cipherText []byte, keyID string) (*kms.DecryptOutput, error)            // Decrypt is a method that decrypts the given encrypted text.
	Credentials(secretName string, target interface{})                              // Credentials is a method that returns the credentials using given info.
}

type AssumeRoleStruct struct {
	RoleArn     string `json:"ROLE_ARN,omitempty"`
	SessionName string `json:"SESSION_NAME,omitempty"`
	SecretName  string `json:"SECRET_NAME,omitempty"`
	KeyID       string `json:"KEY_ID,omitempty"`
}

type AwsCredentials struct {
	AccessKey string `json:"access_key,omitempty"`
	SecretKey string `json:"secret_key,omitempty"`
	Region    string `json:"region,omitempty"`
}
