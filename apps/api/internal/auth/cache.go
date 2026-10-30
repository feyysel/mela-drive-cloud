package auth

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const resetTokenTTL = 30 * time.Minute

type Cache interface {
	Get(ctx context.Context, key string) (uuid.UUID, error)
	Exists(ctx context.Context, key string) (bool, error)
	Set(ctx context.Context, key string, userID uuid.UUID, ttl time.Duration) error
	Del(ctx context.Context, key string) error
}

func resetTokenKey(tokenHash string) string {
	return "auth:reset:" + tokenHash
}
