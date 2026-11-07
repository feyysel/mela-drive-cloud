package user

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

func (h *Handler) Get(ctx *gin.Context) {
	userID, err := ginx.UserID(ctx)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	snapshot, err := h.svc.GetSnapshot(ctx, userID)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	res.OK(ctx, snapshot)
}
