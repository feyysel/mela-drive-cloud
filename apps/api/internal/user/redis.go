package user

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/feyysel/mela-drive-cloud/internal/errs"
	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	client *redis.Client
}

func NewCache(client *redis.Client) Cache {
	return &redisCache{client: client}
}

func (c *redisCache) Get(ctx context.Context, key string) (*User, error) {
	value, err := c.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, errs.ErrNotFound
		}

		return nil, err
	}

	var user User
	if err := json.Unmarshal([]byte(value), &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (c *redisCache) Set(ctx context.Context, key string, user *User, ttl time.Duration) error {
	if user == nil {
		return errs.ErrInvalidInput
	}

	value, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, value, ttl).Err()
}

func (c *redisCache) Del(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}
