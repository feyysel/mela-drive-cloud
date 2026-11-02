package auth

import (
	"github.com/feyysel/mela-drive-cloud/internal/ratelimit"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, limit *ratelimit.Middleware, authMw *Middleware) {
	general := limit.RateLimit("auth_general", ratelimit.General)
	sensitive := limit.RateLimit("auth_sensitive", ratelimit.Sensitive)

	auth := router.Group("/auth")

	auth.POST("/signup", sensitive, handler.Signup)
	auth.POST("/login", sensitive, handler.Login)
	auth.POST("/forgot-password", sensitive, handler.ForgotPassword)
	auth.POST("/reset-password", sensitive, handler.ResetPassword)
	auth.POST("/reset-password/validate", sensitive, handler.ValidateResetToken)
	auth.POST("/logout", authMw.RequireAuth(), general, handler.Logout)
}
