package user

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const (
	userSnapshotTTL = 5 * time.Minute
)

type Cache interface {
	Get(ctx context.Context, key string) (*User, error)
	Set(ctx context.Context, key string, snapshot *User, ttl time.Duration) error
	Del(ctx context.Context, key string) error
}

func userSnapshotKey(userID uuid.UUID) string {
	return "user_snapshot:" + userID.String()
}
