package healthz

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine, handler *Handler) {
	router.GET("/", handler.Info)
	router.GET("/healthz", handler.Healthz)
}
