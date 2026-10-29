package server

import (
	"github.com/feyysel/mela-drive-cloud/internal/container"
	"github.com/feyysel/mela-drive-cloud/internal/user"
	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine, c *container.Container) {
	v1 := router.Group("/v1")

	protected := v1.Group("")

	users := protected.Group("/users")

	user.RegisterRoutes(users, c.UserHandler, c.RateLimitMiddleware)
}
