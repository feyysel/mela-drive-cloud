package folder

import (
	"github.com/gin-gonic/gin"

	"github.com/feyysel/mela-drive-cloud/internal/ratelimit"
)

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, limit *ratelimit.Middleware) {
	general := limit.RateLimit("folder_general", ratelimit.General)
	restricted := limit.RateLimit("folder_restricted", ratelimit.Restricted)

	folders := router.Group("/folders")

	folders.POST("", restricted, handler.Create)
	folders.GET("", general, handler.List)
	folders.GET("/trash", general, handler.ListTrashed)
	folders.DELETE("/trash", restricted, handler.EmptyTrash)
	folders.GET("/:folderID", general, handler.Get)
	folders.PATCH("/:folderID/name", restricted, handler.Rename)
	folders.PATCH("/:folderID/parent", restricted, handler.Move)
	folders.DELETE("/:folderID", restricted, handler.Trash)
	folders.POST("/:folderID/restore", restricted, handler.Restore)
	folders.DELETE("/:folderID/purge", restricted, handler.Purge)
}
