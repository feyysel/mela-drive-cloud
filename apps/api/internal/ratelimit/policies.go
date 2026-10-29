package ratelimit

import "github.com/go-redis/redis_rate/v10"

var (
	General    = redis_rate.PerMinute(60)
	Restricted = redis_rate.PerMinute(10)
	Sensitive  = redis_rate.PerMinute(3)
)
