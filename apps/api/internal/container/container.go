package container

import (
	"context"
	"time"

	"github.com/feyysel/mela-drive-cloud/internal/cache"
	"github.com/feyysel/mela-drive-cloud/internal/config"
	"github.com/feyysel/mela-drive-cloud/internal/db"
	"github.com/feyysel/mela-drive-cloud/internal/ratelimit"
	"github.com/feyysel/mela-drive-cloud/internal/user"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Container struct {
	Cfg                 config.Config
	Pool                *pgxpool.Pool
	Redis               *redis.Client
	RateLimitMiddleware *ratelimit.Middleware
	UserSvc             *user.Service
	UserHandler         *user.Handler
}

func New(cfg config.Config) (*Container, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := db.Connect(ctx, cfg.DB)
	if err != nil {
		return nil, err
	}

	redisClient, err := cache.Connect(ctx, cfg.Redis)
	if err != nil {
		pool.Close()
		return nil, err
	}

	ok := false
	defer func() {
		if !ok {
			pool.Close()
			_ = redisClient.Close()
		}
	}()

	userSvc := user.NewService(user.NewRepository(pool), user.NewCache(redisClient), cfg.App)

	ok = true

	return &Container{
		Cfg:                 cfg,
		Pool:                pool,
		Redis:               redisClient,
		RateLimitMiddleware: ratelimit.NewMiddleware(redisClient),
		UserSvc:             userSvc,
		UserHandler:         user.NewHandler(userSvc),
	}, nil
}

func (c *Container) Close() {
	c.Pool.Close()
	_ = c.Redis.Close()
}
