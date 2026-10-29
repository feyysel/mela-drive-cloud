package ratelimit

import (
	"fmt"
	"log/slog"
	"strconv"

	"github.com/feyysel/mela-drive-cloud/internal/errs"
	"github.com/feyysel/mela-drive-cloud/internal/res"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

type Middleware struct {
	client *redis.Client
}

func NewMiddleware(client *redis.Client) *Middleware {
	return &Middleware{client: client}
}

func (m *Middleware) RateLimit(key string, limit redis_rate.Limit) gin.HandlerFunc {
	limiter := redis_rate.NewLimiter(m.client)

	return func(ctx *gin.Context) {
		key := fmt.Sprintf("%s:%s", key, ctx.ClientIP())

		result, err := limiter.Allow(ctx.Request.Context(), key, limit)
		if err != nil {
			slog.Error("rate limit check failed, allowing request", "key", key, "err", err.Error())
			ctx.Next()

			return
		}

		ctx.Header("X-RateLimit-Limit", strconv.Itoa(limit.Rate))
		ctx.Header("X-RateLimit-Remaining", strconv.Itoa(result.Remaining))

		if result.Allowed == 0 {
			ctx.Header("Retry-After", strconv.Itoa(int(result.RetryAfter.Seconds())))
			res.Error(ctx, errs.Wrap(errs.ErrTooManyRequests, "Too many requests"))

			return
		}

		ctx.Next()
	}
}
