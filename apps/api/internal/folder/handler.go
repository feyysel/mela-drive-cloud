package folder

import (
	"github.com/feyysel/mela-drive-cloud/internal/errs"
	"github.com/feyysel/mela-drive-cloud/internal/ginx"
	"github.com/feyysel/mela-drive-cloud/internal/res"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Create(ctx *gin.Context) {
	ownerID, err := ginx.UserID(ctx)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	var req CreateFolderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res.Error(ctx, err)
		return
	}

	folder, err := h.svc.Create(ctx, ownerID, req.ParentID, req.Name)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	res.Created(ctx, folder)
}

func (h *Handler) Get(ctx *gin.Context) {
	ownerID, err := ginx.UserID(ctx)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	folderID, err := ginx.Param(ctx, "folderID")
	if err != nil {
		res.Error(ctx, errs.Wrap(errs.ErrInvalidInput, "Invalid folder ID"))
		return
	}

	folder, err := h.svc.Get(ctx, ownerID, folderID)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	res.OK(ctx, folder)
}

func (h *Handler) List(ctx *gin.Context) {
	ownerID, err := ginx.UserID(ctx)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	var parentID *uuid.UUID

	if value := ctx.Query("parentId"); value != "" {
		parsed, err := uuid.Parse(value)
		if err != nil {
			res.Error(ctx, errs.Wrap(errs.ErrInvalidInput, "Invalid parent folder ID"))
			return
		}

		parentID = &parsed
	}

	folders, err := h.svc.List(ctx, ownerID, parentID)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	res.OK(ctx, folders)
}

func (h *Handler) Rename(ctx *gin.Context) {
	ownerID, err := ginx.UserID(ctx)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	folderID, err := uuid.Parse(ctx.Param("folderID"))
	if err != nil {
		res.Error(ctx, errs.Wrap(errs.ErrInvalidInput, "Invalid folder ID"))
		return
	}

	var req RenameFolderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res.Error(ctx, err)
		return
	}

	folder, err := h.svc.Rename(ctx, ownerID, folderID, req.Name)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	res.OK(ctx, folder)
}

func (h *Handler) Move(ctx *gin.Context) {
	ownerID, err := ginx.UserID(ctx)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	folderID, err := ginx.Param(ctx, "folderID")
	if err != nil {
		res.Error(ctx, err)
		return
	}

	var req MoveFolderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res.Error(ctx, err)
		return
	}

	folder, err := h.svc.Move(ctx, ownerID, folderID, req.ParentID)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	res.OK(ctx, folder)
}
