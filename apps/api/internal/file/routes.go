package file

import (
	"github.com/gin-gonic/gin"

	"github.com/feyysel/mela-drive-cloud/internal/ratelimit"
)

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, limit *ratelimit.Middleware) {
	general := limit.RateLimit("file_general", ratelimit.General)
	restricted := limit.RateLimit("file_restricted", ratelimit.Restricted)

	files := router.Group("/files")

	files.POST("", restricted, handler.Upload)
	files.GET("", general, handler.List)
	files.GET("/:fileID", general, handler.Get)
	files.PATCH("/:fileID/name", restricted, handler.Rename)
	files.PATCH("/:fileID/folder", restricted, handler.Move)
}
