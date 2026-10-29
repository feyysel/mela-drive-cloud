package config

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

func get(key string) (string, error) {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return "", fmt.Errorf("missing env var %s", key)
	}

	return value, nil
}

func getURL(key string) (string, error) {
	value, err := get(key)
	if err != nil {
		return "", err
	}

	parsed, err := url.Parse(value)
	if err != nil || parsed.Scheme != "https" || parsed.Host == "" {
		return "", fmt.Errorf("%s is not a valid https url", key)
	}

	return strings.TrimRight(value, "/"), nil
}
