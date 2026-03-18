package service

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/hoshina-dev/pasta/internal/infra/s3"
)

type StorageService struct {
	storage storage.StorageService
}

func NewStorageService(storage storage.StorageService) *StorageService {
	return &StorageService{storage: storage}
}

func (s *StorageService) GenerateUploadURL(ctx context.Context, fileName, contentType string) (uploadURL, fileKey string, err error) {
	fileKey = s.generateFileKey(fileName)

	uploadURL, err = s.storage.GeneratePresignedUploadURL(ctx, fileKey, contentType)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate upload URL: %w", err)
	}

	return uploadURL, fileKey, nil
}

func (s *StorageService) generateFileKey(fileName string) string {
	timestamp := time.Now().Unix()
	uniqueID := uuid.New().String()
	ext := filepath.Ext(fileName)

	return fmt.Sprintf("uploads/%d-%s%s", timestamp, uniqueID, ext)
}

func (s *StorageService) DeleteFile(ctx context.Context, fileKey string) error {
	return s.storage.Delete(ctx, fileKey)
}
