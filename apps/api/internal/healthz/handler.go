package healthz

import (
	"github.com/feyysel/mela-drive-cloud/internal/res"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Info(ctx *gin.Context) {
	res.OK(ctx, h.svc.Info())
}

func (h *Handler) Healthz(ctx *gin.Context) {
	res.OK(ctx, h.svc.Healthz())
}
