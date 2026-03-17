package s3

import (
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type StorageService interface {
	Upload(key string, file io.Reader, contentType string) (string, error)
	Delete(key string) error
}

type s3StorageService struct {
	client  *s3.Client
	bucket  string
	baseURL string
}

func NewS3StorageService(client *s3.Client, bucket, baseURL string) *s3StorageService {
	return &s3StorageService{client: client, bucket: bucket, baseURL: baseURL}
}

func (s *s3StorageService) Upload(ctx context.Context, key string, file io.Reader, contentType string) (string, error) {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", s.baseURL, key), nil
}

func (s *s3StorageService) Delete(ctx context.Context, key string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	return err
}
