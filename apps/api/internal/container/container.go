package container

import (
	"context"
	"time"

	"github.com/feyysel/mela-drive-cloud/internal/auth"
	"github.com/feyysel/mela-drive-cloud/internal/cache"
	"github.com/feyysel/mela-drive-cloud/internal/config"
	"github.com/feyysel/mela-drive-cloud/internal/db"
	"github.com/feyysel/mela-drive-cloud/internal/healthz"
	"github.com/feyysel/mela-drive-cloud/internal/mail"
	"github.com/feyysel/mela-drive-cloud/internal/ratelimit"
	"github.com/feyysel/mela-drive-cloud/internal/session"
	"github.com/feyysel/mela-drive-cloud/internal/user"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/resend/resend-go/v4"
)

type Container struct {
	Cfg                 config.Config
	Pool                *pgxpool.Pool
	Redis               *redis.Client
	RateLimitMiddleware *ratelimit.Middleware
	AuthMiddleware      *auth.Middleware
	UserSvc             *user.Service
	HealthzHandler      *healthz.Handler
	UserHandler         *user.Handler
	AuthHandler         *auth.Handler
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

	resendClient := resend.NewClient(cfg.Mail.APIKey)

	healthSvc := healthz.NewService(cfg.App)
	userSvc := user.NewService(user.NewRepository(pool), user.NewCache(redisClient), cfg.App)
	sessSvc := session.NewService(session.NewCache(redisClient), cfg.Session)
	mailSvc := mail.NewService(resendClient)
	authSvc := auth.NewService(auth.NewCache(redisClient), userSvc, sessSvc, mailSvc, cfg.App, cfg.Mail)

	ok = true

	return &Container{
		Cfg:                 cfg,
		Pool:                pool,
		Redis:               redisClient,
		RateLimitMiddleware: ratelimit.NewMiddleware(redisClient),
		AuthMiddleware:      auth.NewMiddleware(sessSvc, userSvc, cfg.Session),
		UserSvc:             userSvc,
		HealthzHandler:      healthz.NewHandler(healthSvc),
		UserHandler:         user.NewHandler(userSvc),
		AuthHandler:         auth.NewHandler(authSvc, cfg.Session),
	}, nil
}

func (c *Container) Close() {
	c.Pool.Close()
	_ = c.Redis.Close()
}
