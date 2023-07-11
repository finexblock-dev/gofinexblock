package goaws

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"mime/multipart"
)

func NewS3Client(sess *session.Session) *s3.S3 {
	return s3.New(sess)
}

func Upload(client *s3.S3, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return client.PutObject(input)
}

func UploadBatch(client *s3.S3, f *multipart.Form, bucket, basePath string) (map[string]string, error) {
	var result = make(map[string]string)
	var open multipart.File
	var objectUrl, path string
	var content []byte
	var err error

	for _, header := range f.File {
		for _, file := range header {
			open, err = file.Open()
			if err != nil {
				return nil, err
			}

			content = make([]byte, file.Size)
			if _, err = open.Read(content); err != nil {
				return nil, err
			}

			path = basePath + file.Filename

			if _, err = Upload(client, &s3.PutObjectInput{
				Body:        bytes.NewReader(content),
				Bucket:      aws.String(bucket),
				ContentType: aws.String(file.Header.Get("Content-Type")),
				Key:         aws.String(path),
				ACL:         aws.String("public-read"),
			}); err != nil {
				return nil, err
			}

			objectUrl = fmt.Sprintf("https://%s.s3.ap-northeast-2.amazonaws.com/%s", bucket, path)

			result[file.Filename] = objectUrl
		}
	}
	return result, nil
}
