package config

import (
	"errors"
	"fmt"
	"strings"
)

const (
	envDev  = "dev"
	envProd = "prod"
)

type AppConfig struct {
	Name           string
	Env            string
	Port           string
	Domain         string
	OwnerEmail     string
	OwnerPassword  string
	AllowedOrigins []string
}

func (c *AppConfig) IsDev() bool {
	return c.Env == envDev
}

func (c *AppConfig) IsProd() bool {
	return c.Env == envProd
}

func LoadApp() (AppConfig, error) {
	name, err := get("APP_NAME")
	if err != nil {
		return AppConfig{}, err
	}

	env, err := get("APP_ENV")
	if err != nil {
		return AppConfig{}, err
	}

	if env != envDev && env != envProd {
		return AppConfig{}, fmt.Errorf("APP_ENV must be %s or %s, got %q", envDev, envProd, env)
	}

	port, err := get("APP_PORT")
	if err != nil {
		return AppConfig{}, err
	}

	domain, err := get("APP_DOMAIN")
	if err != nil {
		return AppConfig{}, err
	}

	ownerEmail, err := get("APP_OWNER_EMAIL")
	if err != nil {
		return AppConfig{}, err
	}

	ownerPassword, err := get("APP_OWNER_PASSWORD")
	if err != nil {
		return AppConfig{}, err
	}

	originsRaw, err := get("APP_ALLOWED_ORIGINS")
	if err != nil {
		return AppConfig{}, err
	}

	var origins []string

	for origin := range strings.SplitSeq(originsRaw, ",") {
		origin = strings.TrimRight(strings.TrimSpace(origin), "/")
		if origin == "" {
			continue
		}

		if strings.Contains(origin, "*") {
			return AppConfig{}, fmt.Errorf("APP_ALLOWED_ORIGINS must not contain wildcards, got %q", origin)
		}

		origins = append(origins, origin)
	}

	if len(origins) == 0 {
		return AppConfig{}, errors.New("APP_ALLOWED_ORIGINS must contain at least one origin")
	}

	return AppConfig{
		Name:           name,
		Env:            env,
		Port:           port,
		Domain:         domain,
		OwnerEmail:     ownerEmail,
		OwnerPassword:  ownerPassword,
		AllowedOrigins: origins,
	}, nil
}
