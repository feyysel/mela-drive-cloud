package trash

import (
	"github.com/gin-gonic/gin"

	"github.com/feyysel/mela-drive-cloud/internal/ratelimit"
)

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, limit *ratelimit.Middleware) {
	restricted := limit.RateLimit("trash_restricted", ratelimit.Restricted)

	trash := router.Group("/trash")

	trash.DELETE("", restricted, handler.EmptyTrash)
}
