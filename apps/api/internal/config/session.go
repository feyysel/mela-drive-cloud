package config

import "time"

const (
	sessionTTL        = 7 * 24 * time.Hour
	sessionCookieTTL  = 30 * 24 * time.Hour
	sessionCookieName = "__Secure-session"
)

type SessionConfig struct {
	TTL            time.Duration
	RenewThreshold time.Duration
	CookieName     string
	CookieDomain   string
	CookieMaxAge   int
	CookieSecure   bool
}

func NewSessionConfig(app AppConfig) SessionConfig {
	cfg := SessionConfig{
		TTL:            sessionTTL,
		RenewThreshold: sessionTTL / 4,
		CookieName:     sessionCookieName,
		CookieDomain:   app.Domain,
		CookieMaxAge:   int(sessionCookieTTL.Seconds()),
		CookieSecure:   true,
	}

	return cfg
}
