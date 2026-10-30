package auth

import (
	"context"
	"errors"
	"time"

	"github.com/feyysel/mela-drive-cloud/internal/errs"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	client *redis.Client
}

func NewCache(client *redis.Client) Cache {
	return &redisCache{client: client}
}

func (c *redisCache) Get(ctx context.Context, key string) (uuid.UUID, error) {
	value, err := c.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return uuid.Nil, errs.ErrNotFound
		}

		return uuid.Nil, err
	}

	return uuid.Parse(value)
}

func (c *redisCache) Set(ctx context.Context, key string, userID uuid.UUID, ttl time.Duration) error {
	if userID == uuid.Nil {
		return errs.ErrInvalidInput
	}

	return c.client.Set(ctx, key, userID.String(), ttl).Err()
}

func (c *redisCache) Del(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}
