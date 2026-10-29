package user

import (
	"github.com/gin-gonic/gin"

	"github.com/feyysel/mela-drive-cloud/internal/ratelimit"
)

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, limit *ratelimit.Middleware) {
	general := limit.RateLimit("user_general", ratelimit.General)
	restricted := limit.RateLimit("user_restricted", ratelimit.Restricted)
	sensitive := limit.RateLimit("user_sensitive", ratelimit.Sensitive)

	me := router.Group("/me")

	me.GET("", general, handler.Get)
	me.PATCH("/name", restricted, handler.UpdateName)
	me.PATCH("/email", sensitive, handler.UpdateEmail)
	me.DELETE("", sensitive, handler.Delete)
}
