package storage

import (
	"context"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/transfermanager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/feyysel/mela-drive-cloud/internal/config"
)

type Service struct {
	client  *s3.Client
	uploads *transfermanager.Client

	cfg config.StorageConfig
}

func NewService(cfg config.StorageConfig) *Service {
	client := s3.New(s3.Options{
		Region:                     "auto",
		BaseEndpoint:               aws.String(cfg.R2Endpoint),
		Credentials:                credentials.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		RequestChecksumCalculation: aws.RequestChecksumCalculationWhenRequired,
		ResponseChecksumValidation: aws.ResponseChecksumValidationWhenRequired,
	})

	return &Service{
		client:  client,
		uploads: transfermanager.New(client),
		cfg:     cfg,
	}
}

func (s *Service) Upload(ctx context.Context, key, contentType string, body io.Reader) error {
	if err := requireValidKey(key); err != nil {
		return err
	}

	if err := requireValidContentType(contentType); err != nil {
		return err
	}

	_, err := s.uploads.UploadObject(ctx, &transfermanager.UploadObjectInput{
		Bucket:      aws.String(s.cfg.Bucket),
		Key:         aws.String(key),
		Body:        body,
		ContentType: aws.String(contentType),
	})

	return err
}

func (s *Service) Delete(ctx context.Context, key string) error {
	if err := requireValidKey(key); err != nil {
		return err
	}

	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.cfg.Bucket),
		Key:    aws.String(key),
	})

	return err
}

func (s *Service) DeleteByPrefix(ctx context.Context, prefix string) error {
	if err := requireValidPrefix(prefix); err != nil {
		return err
	}

	paginator := s3.NewListObjectsV2Paginator(s.client, &s3.ListObjectsV2Input{
		Bucket: aws.String(s.cfg.Bucket),
		Prefix: aws.String(prefix),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}

		if len(page.Contents) == 0 {
			continue
		}

		objects := make([]types.ObjectIdentifier, len(page.Contents))
		for i, obj := range page.Contents {
			objects[i] = types.ObjectIdentifier{Key: obj.Key}
		}

		_, err = s.client.DeleteObjects(ctx, &s3.DeleteObjectsInput{
			Bucket: aws.String(s.cfg.Bucket),
			Delete: &types.Delete{Objects: objects},
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) PublicURL(key string) string {
	return s.cfg.PublicBaseURL + "/" + strings.TrimLeft(key, "/")
}

func (s *Service) PublicURLPtr(key *string) *string {
	if key == nil {
		return nil
	}

	url := s.PublicURL(*key)

	return &url
}
