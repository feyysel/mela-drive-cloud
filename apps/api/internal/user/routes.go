package user

import (
	"github.com/gin-gonic/gin"

	"github.com/feyysel/mela-drive-cloud/internal/ratelimit"
)

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, limit *ratelimit.Middleware) {
	general := limit.RateLimit("user_general", ratelimit.General)

	me := router.Group("/me")

	me.GET("", general, handler.Get)
}
