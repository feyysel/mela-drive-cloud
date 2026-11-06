package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/feyysel/mela-drive-cloud/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func Connect(ctx context.Context, cfg config.DBConfig) (*sql.DB, error) {
	pool, err := sql.Open("mysql", cfg.URL)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	if err := pool.PingContext(ctx); err != nil {
		pool.Close()

		return nil, fmt.Errorf("ping database: %w", err)
	}

	slog.InfoContext(ctx, "db connected")

	return pool, nil
}
