package db

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/feyysel/mela-drive-cloud/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect(ctx context.Context, cfg config.DBConfig) (*pgxpool.Pool, error) {
	poolCfg, err := pgxpool.ParseConfig(cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("parse database url: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolCfg)
	if err != nil {
		return nil, fmt.Errorf("create pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()

		return nil, fmt.Errorf("ping database: %w", err)
	}

	slog.InfoContext(ctx, "db connected", "host", poolCfg.ConnConfig.Host, "name", poolCfg.ConnConfig.Database)

	return pool, nil
}
