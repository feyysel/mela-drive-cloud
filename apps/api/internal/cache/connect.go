package cache

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/feyysel/mela-drive-cloud/internal/config"
	"github.com/redis/go-redis/v9"
)

func Connect(ctx context.Context, cfg config.RedisConfig) (*redis.Client, error) {
	opts, err := redis.ParseURL(cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("parse redis url: %w", err)
	}

	client := redis.NewClient(opts)

	if err := client.Ping(ctx).Err(); err != nil {
		_ = client.Close()

		return nil, fmt.Errorf("ping redis: %w", err)
	}

	slog.InfoContext(ctx, "cache connected", "host", opts.Addr, "db", opts.DB)

	return client, nil
}
