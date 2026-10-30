package session

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

func (c *redisCache) Get(ctx context.Context, key string) (*Session, error) {
	value, err := c.client.Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, errs.ErrNotFound
		}

		return nil, err
	}

	var session Session
	if err := json.Unmarshal([]byte(value), &session); err != nil {
		return nil, err
	}

	return &session, nil
}

func (c *redisCache) Set(ctx context.Context, key string, session *Session, ttl time.Duration) error {
	if session == nil {
		return errs.ErrInvalidInput
	}

	value, err := json.Marshal(session)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, key, value, ttl).Err()
}

func (c *redisCache) Del(ctx context.Context, key string) (int64, error) {
	return c.client.Del(ctx, key).Result()
}

func (c *redisCache) SAdd(ctx context.Context, key, member string) error {
	return c.client.SAdd(ctx, key, member).Err()
}

func (c *redisCache) SRem(ctx context.Context, key string, members ...string) error {
	if len(members) == 0 {
		return nil
	}

	args := make([]interface{}, len(members))
	for i, m := range members {
		args[i] = m
	}

	return c.client.SRem(ctx, key, args...).Err()
}

func (c *redisCache) SMembers(ctx context.Context, key string) ([]string, error) {
	return c.client.SMembers(ctx, key).Result()
}

func (c *redisCache) Expire(ctx context.Context, key string, ttl time.Duration) error {
	return c.client.Expire(ctx, key, ttl).Err()
}

func (c *redisCache) TTL(ctx context.Context, key string) (time.Duration, error) {
	remaining, err := c.client.TTL(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	if remaining < 0 {
		return 0, errs.ErrNotFound
	}

	return remaining, nil
}
