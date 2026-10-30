package session

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Cache interface {
	Get(ctx context.Context, key string) (*Session, error)
	Set(ctx context.Context, key string, session *Session, ttl time.Duration) error
	Del(ctx context.Context, key string) (int64, error)

	SAdd(ctx context.Context, key, member string) error
	SRem(ctx context.Context, key string, members ...string) error
	SMembers(ctx context.Context, key string) ([]string, error)
	Expire(ctx context.Context, key string, ttl time.Duration) error
	TTL(ctx context.Context, key string) (time.Duration, error)
}

func sessionKey(sessionID uuid.UUID) string {
	return fmt.Sprintf("session:%s", sessionID)
}

func userSessionsKey(userID uuid.UUID) string {
	return fmt.Sprintf("user_sessions:%s", userID)
}
