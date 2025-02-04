package storage

import (
    "bytes"
    "context"
    "fmt"
    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
)

type Storage interface {
    UploadFile(ctx context.Context, bucketName, objectName string, data []byte) error
    DownloadFile(ctx context.Context, bucketName, objectName string) ([]byte, error)
    DeleteFile(ctx context.Context, bucketName, objectName string) error
}

type MinioStorage struct {
    client *minio.Client
}

func NewMinioStorage(endpoint, accessKeyID, secretAccessKey string, useSSL bool) (*MinioStorage, error) {
    client, err := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
        Secure: useSSL,
    })
    if err != nil {
        return nil, fmt.Errorf("failed to create minio client: %w", err)
    }
    
    return &MinioStorage{client: client}, nil
}

func (s *MinioStorage) UploadFile(ctx context.Context, bucketName, objectName string, data []byte) error {
    reader := bytes.NewReader(data)
    _, err := s.client.PutObject(ctx, bucketName, objectName, reader, int64(len(data)), minio.PutObjectOptions{
        ContentType: "application/pdf",
    })
    if err != nil {
        return fmt.Errorf("failed to upload file: %w", err)
    }
    return nil
}

func (s *MinioStorage) DownloadFile(ctx context.Context, bucketName, objectName string) ([]byte, error) {
    obj, err := s.client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
    if err != nil {
        return nil, fmt.Errorf("failed to get object: %w", err)
    }
    
    buf := new(bytes.Buffer)
    if _, err := buf.ReadFrom(obj); err != nil {
        return nil, fmt.Errorf("failed to read object: %w", err)
    }
    
    return buf.Bytes(), nil
}

func (s *MinioStorage) DeleteFile(ctx context.Context, bucketName, objectName string) error {
    err := s.client.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
    if err != nil {
        return fmt.Errorf("failed to delete object: %w", err)
    }
    return nil
}