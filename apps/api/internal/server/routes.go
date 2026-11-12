package server

import (
	"github.com/feyysel/mela-drive-cloud/internal/auth"
	"github.com/feyysel/mela-drive-cloud/internal/container"
	"github.com/feyysel/mela-drive-cloud/internal/file"
	"github.com/feyysel/mela-drive-cloud/internal/folder"
	"github.com/feyysel/mela-drive-cloud/internal/healthz"
	"github.com/feyysel/mela-drive-cloud/internal/trash"
	"github.com/feyysel/mela-drive-cloud/internal/user"
	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine, c *container.Container) {
	healthz.RegisterRoutes(router, c.HealthzHandler)

	v1 := router.Group("/v1")

	auth.RegisterRoutes(v1, c.AuthHandler, c.RateLimitMiddleware, c.AuthMiddleware)

	protected := v1.Group("", c.AuthMiddleware.RequireAuth())

	users := protected.Group("/users")

	user.RegisterRoutes(users, c.UserHandler, c.RateLimitMiddleware)

	folder.RegisterRoutes(protected, c.FolderHandler, c.RateLimitMiddleware)
	file.RegisterRoutes(protected, c.FileHandler, c.RateLimitMiddleware)
	trash.RegisterRoutes(protected, c.TrashHandler, c.RateLimitMiddleware)
}
