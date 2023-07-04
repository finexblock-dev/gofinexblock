package secure

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

func NewKMSClient(sess *session.Session) *kms.KMS {
	return kms.New(sess)
}

func Encrypt(svc *kms.KMS, keyID, plainText string) (*kms.EncryptOutput, error) {
	return svc.Encrypt(&kms.EncryptInput{
		EncryptionAlgorithm: aws.String("SYMMETRIC_DEFAULT"),
		KeyId:               aws.String(keyID),
		Plaintext:           []byte(plainText),
	})
}

func Decrypt(svc *kms.KMS, cipherText []byte, keyID string) (*kms.DecryptOutput, error) {
	return svc.Decrypt(&kms.DecryptInput{
		CiphertextBlob:      cipherText,
		EncryptionAlgorithm: aws.String("SYMMETRIC_DEFAULT"),
		KeyId:               aws.String(keyID),
	})
}