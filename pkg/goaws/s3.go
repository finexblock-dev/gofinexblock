package goaws

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/finexblock-dev/gofinexblock/pkg/entity"
	"mime/multipart"
	"strings"
	"time"
)

func NewS3Client(sess *session.Session) *s3.S3 {
	return s3.New(sess)
}

func Upload(client *s3.S3, input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	return client.PutObject(input)
}

func UploadErrorLog(client *s3.S3, errorLog *entity.FinexblockErrorLog, bucket string) error {
	jsonData, err := json.Marshal(errorLog)
	if err != nil {
		errors.New("Failed to convert log data to JSON!")
	}

	currentDate := time.Now().Format("2006-01-02")
	currentTime := time.Now().Format("15:04:05")
	currentDate = strings.ReplaceAll(currentDate, "-", "")
	currentTime = strings.ReplaceAll(currentTime, ":", "")
	path := "finexblockErrorLogs/" + currentDate + currentTime + ".json"

	if _, err = Upload(client, &s3.PutObjectInput{
		Body:        aws.ReadSeekCloser(bytes.NewReader(jsonData)),
		Bucket:      aws.String(bucket),
		ContentType: aws.String("application/json"),
		Key:         aws.String(path),
		ACL:         aws.String("public-read"),
	}); err != nil {
		return err
	}
	return nil
}

func UploadBatch(client *s3.S3, files []*multipart.FileHeader, bucket, basePath string) (map[string]string, error) {
	var result = make(map[string]string)
	var open multipart.File
	var objectUrl, path string
	var content []byte
	var err error

	for _, file := range files {
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
	return result, nil
}
