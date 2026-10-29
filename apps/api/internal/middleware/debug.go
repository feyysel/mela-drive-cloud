package middleware

import (
	"github.com/feyysel/mela-drive-cloud/internal/ginx"
	"github.com/gin-gonic/gin"
)

func Debug(debug bool) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ginx.SetDebug(ctx, debug)
		ctx.Next()
	}
}
