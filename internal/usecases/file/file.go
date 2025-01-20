package file

import (
	"context"
	"mime/multipart"

	"github.com/fatjan/fitbyte/internal/config"
	s3uploader "github.com/fatjan/fitbyte/internal/pkg/s3_uploader"
)

type useCase struct {
	config config.Config
	s3     *s3uploader.Uploader
}

func NewUseCase(config config.Config) UseCase {
	s3Config := &s3uploader.Config{
		BucketName:      config.Aws.BucketName,
		AccessKeyID:     config.Aws.AccessKeyID,
		AccessKeySecret: config.Aws.SecretAccessKey,
		Region:          config.Aws.Region,
		AccountID:       config.Aws.AccountID,
	}

	s3, _ := s3uploader.NewUploader(s3Config)

	return &useCase{
		config: config,
		s3:     s3,
	}
}

func (uc *useCase) UploadFile(ctx context.Context, file multipart.File, fileName string) (string, error) {
	uploadChan := uc.s3.UploadFile(ctx, file, fileName)
	uploadResult := <-uploadChan

	if uploadResult.Err != nil {
		return "", uploadResult.Err
	}

	return uploadResult.URL, nil
}
