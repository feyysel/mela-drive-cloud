package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery

		ctx.Next()

		latency := time.Since(start)
		status := ctx.Writer.Status()

		attrs := []slog.Attr{
			slog.String("requestId", GetRequestID(ctx)),
			slog.Int("status", status),
			slog.String("method", ctx.Request.Method),
			slog.String("path", path),
			slog.Duration("latency", latency),
			slog.String("ip", ctx.ClientIP()),
		}

		if query != "" {
			attrs = append(attrs, slog.String("query", query))
		}

		level := slog.LevelInfo

		switch {
		case status >= 500:
			level = slog.LevelError
		case status >= 400:
			level = slog.LevelWarn
		}

		if len(ctx.Errors) > 0 {
			attrs = append(attrs, slog.String("error", ctx.Errors.String()))

			if level == slog.LevelInfo {
				level = slog.LevelError
			}
		}

		slog.LogAttrs(ctx.Request.Context(), level, "http request completed", attrs...)
	}
}
