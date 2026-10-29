package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/feyysel/mela-drive-cloud/internal/config"
	"github.com/feyysel/mela-drive-cloud/internal/container"
	"github.com/feyysel/mela-drive-cloud/internal/middleware"
	"github.com/feyysel/mela-drive-cloud/internal/slogx"
	"github.com/gin-gonic/gin"
)

func Run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	slogx.Init(cfg.App.IsDev())

	c, err := container.New(cfg)
	if err != nil {
		return err
	}
	defer c.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := c.UserSvc.SeedOwner(ctx); err != nil {
		return err
	}

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	_ = router.SetTrustedProxies(nil)

	router.Use(gin.Recovery())
	router.Use(middleware.CORS(c.Cfg.App))
	router.Use(middleware.RequestID())
	router.Use(middleware.Logger())
	router.Use(middleware.Debug(c.Cfg.App.IsDev()))

	registerRoutes(router, c)

	srv := &http.Server{
		Addr:              ":" + c.Cfg.App.Port,
		Handler:           router,
		ReadHeaderTimeout: 10 * time.Second,
	}

	serverErr := make(chan error, 1)

	go func() {
		slog.Info("server started", "env", c.Cfg.App.Env, "port", c.Cfg.App.Port)

		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	select {
	case err := <-serverErr:
		return fmt.Errorf("server error: %w", err)

	case <-quit:
		slog.Info("shutting down")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("graceful shutdown: %w", err)
	}

	slog.Info("shutdown complete")

	return nil
}
