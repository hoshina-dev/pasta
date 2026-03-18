package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type StorageService interface {
	GeneratePresignedUploadURL(ctx context.Context, key string, contentType string) (string, error)
	GeneratePresignedDownloadURL(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}

type s3StorageService struct {
	client        *s3.Client
	presignClient *s3.PresignClient
	bucket        string
	baseURL       string
	presignExpiry time.Duration
}

func NewS3StorageService(client *s3.Client, bucket, baseURL string) *s3StorageService {
	presignClient := s3.NewPresignClient(client)
	return &s3StorageService{
		client:        client,
		presignClient: presignClient,
		bucket:        bucket,
		baseURL:       baseURL,
		presignExpiry: 15 * time.Minute,
	}
}

func (s *s3StorageService) GeneratePresignedUploadURL(ctx context.Context, key string, contentType string) (string, error) {
	presignResult, err := s.presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		ContentType: aws.String(contentType),
	}, s3.WithPresignExpires(s.presignExpiry))

	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	return presignResult.URL, nil
}

func (s *s3StorageService) GeneratePresignedDownloadURL(ctx context.Context, key string) (string, error) {
	presignResult, err := s.presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(s.presignExpiry))

	if err != nil {
		return "", fmt.Errorf("failed to generate presigned download URL: %w", err)
	}

	return presignResult.URL, nil
}

func (s *s3StorageService) Delete(ctx context.Context, key string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	return err
}
