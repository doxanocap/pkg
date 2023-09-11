package aws

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"io"
)

const contentType = "multipart/form-data"

type S3Client struct {
	client     *minio.Client
	log        *zap.Logger
	bucketName string
}

func InitS3() (*S3Client, error) {
	ctx := context.TODO()

	minioClient, err := minio.New(viper.GetString("S3_ENDPOINT"), &minio.Options{
		Creds:  credentials.NewStaticV4(viper.GetString("S3_ACCESS_KEY"), viper.GetString("S3_SECRET_KEY"), ""),
		Secure: viper.GetBool("S3_USE_SSL"),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to S3: %v", err)
	}
	s3Client := &S3Client{
		client: minioClient,
		//log:        logger.Log.Named("[STORAGE][S3]"),
		bucketName: viper.GetString("S3_BUCKET_NAME"),
	}

	exists, err := s3Client.checkIfBucketExists(ctx, s3Client.bucketName)
	if err != nil {
		return nil, fmt.Errorf("failed to find bucket: %v", err)
	}
	if !exists {
		return nil, fmt.Errorf("such bucket does not exists")
	}

	return s3Client, nil
}

func (s *S3Client) checkIfBucketExists(ctx context.Context, bucketName string) (bool, error) {
	exists, err := s.client.BucketExists(ctx, bucketName)
	if err != nil {
		return false, err
	}
	return exists, err
}

func (s *S3Client) Save(ctx context.Context, name string, file io.Reader, size int64) error {
	log := s.log.Named("Save").With(
		zap.String("objectName", name),
		zap.Int64("size", size))

	uploadInfo, err := s.client.PutObject(ctx, s.bucketName, name, file, size, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}

	log.Info("UploadInfo", zap.String("uploadKey", uploadInfo.Key))
	return nil
}
