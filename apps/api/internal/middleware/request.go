package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	RequestIDHeader     = "X-Request-ID"
	requestIDContextKey = "requestID"
)

func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.GetHeader(RequestIDHeader)
		if id == "" {
			id = uuid.NewString()
		}

		ctx.Set(requestIDContextKey, id)
		ctx.Header(RequestIDHeader, id)
		ctx.Next()
	}
}

func GetRequestID(ctx *gin.Context) string {
	value, exists := ctx.Get(requestIDContextKey)
	if !exists {
		return ""
	}

	id, ok := value.(string)
	if !ok {
		return ""
	}

	return id
}
