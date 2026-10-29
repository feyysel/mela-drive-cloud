package config

import "strings"

type StorageConfig struct {
	R2Endpoint      string
	PublicBaseURL   string
	AccessKeyID     string
	SecretAccessKey string
	Bucket          string
}

func LoadStorage() (StorageConfig, error) {
	r2Endpoint, err := getURL("STORAGE_R2_ENDPOINT")
	if err != nil {
		return StorageConfig{}, err
	}

	publicBaseURL, err := getURL("STORAGE_PUBLIC_BASE_URL")
	if err != nil {
		return StorageConfig{}, err
	}

	accessKeyID, err := get("STORAGE_ACCESS_KEY_ID")
	if err != nil {
		return StorageConfig{}, err
	}

	secretAccessKey, err := get("STORAGE_SECRET_ACCESS_KEY")
	if err != nil {
		return StorageConfig{}, err
	}

	bucket, err := get("STORAGE_BUCKET")
	if err != nil {
		return StorageConfig{}, err
	}

	return StorageConfig{
		R2Endpoint:      r2Endpoint,
		PublicBaseURL:   publicBaseURL,
		AccessKeyID:     accessKeyID,
		SecretAccessKey: secretAccessKey,
		Bucket:          bucket,
	}, nil
}

func (c StorageConfig) PublicURL(key string) string {
	return c.PublicBaseURL + "/" + strings.TrimLeft(key, "/")
}
