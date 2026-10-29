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

func (h *Handler) UpdateName(ctx *gin.Context) {
	userID, err := ginx.Param(ctx, "userId")
	if err != nil {
		res.Error(ctx, err)
		return
	}

	var req UpdateUserNameRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res.Error(ctx, err)
		return
	}

	user, err := h.svc.UpdateName(ctx, userID, req.FirstName, req.LastName)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	res.OK(ctx, user)
}

func (h *Handler) UpdateEmail(ctx *gin.Context) {
	userID, err := ginx.Param(ctx, "userId")
	if err != nil {
		res.Error(ctx, err)
		return
	}

	var req UpdateEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res.Error(ctx, err)
		return
	}

	user, err := h.svc.UpdateEmail(ctx, userID, req.Email)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	res.OK(ctx, user)
}

func (h *Handler) Delete(ctx *gin.Context) {
	userID, err := ginx.Param(ctx, "userId")
	if err != nil {
		res.Error(ctx, err)
		return
	}

	if err := h.svc.Delete(ctx, userID); err != nil {
		res.Error(ctx, err)
		return
	}

	res.NoContent(ctx)
}
