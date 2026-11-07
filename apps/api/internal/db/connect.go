package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net"
	"time"

	"github.com/feyysel/mela-drive-cloud/internal/config"
	"github.com/go-sql-driver/mysql"
)

func Connect(ctx context.Context, cfg config.DBConfig) (*sql.DB, error) {
	dsn := mysql.NewConfig()
	dsn.User = cfg.User
	dsn.Passwd = cfg.Password
	dsn.Net = "tcp"
	dsn.Addr = net.JoinHostPort(cfg.Host, cfg.Port)
	dsn.DBName = cfg.Name
	dsn.ParseTime = true
	dsn.Loc = time.UTC

	connector, err := mysql.NewConnector(dsn)
	if err != nil {
		return nil, fmt.Errorf("create connector: %w", err)
	}

	pool := sql.OpenDB(connector)

	if err := pool.PingContext(ctx); err != nil {
		_ = pool.Close()

		return nil, fmt.Errorf("ping database: %w", err)
	}

	slog.InfoContext(ctx, "db connected", "host", cfg.Host, "name", cfg.Name)

	return pool, nil
}
