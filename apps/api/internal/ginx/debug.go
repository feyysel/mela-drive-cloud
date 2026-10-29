package ginx

import "github.com/gin-gonic/gin"

const keyDebug = "debug"

// IsDebug reports whether debug mode is enabled in ctx.
func IsDebug(ctx *gin.Context) bool {
	return ctx.GetBool(keyDebug)
}

// SetDebug sets debug mode in ctx.
func SetDebug(ctx *gin.Context, debug bool) {
	ctx.Set(keyDebug, debug)
}
