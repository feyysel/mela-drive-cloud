package trash

import (
	"github.com/feyysel/mela-drive-cloud/internal/ginx"
	"github.com/feyysel/mela-drive-cloud/internal/res"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) EmptyTrash(ctx *gin.Context) {
	ownerID, err := ginx.UserID(ctx)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	if err := h.svc.EmptyTrash(ctx, ownerID); err != nil {
		res.Error(ctx, err)
		return
	}

	res.NoContent(ctx)
}
