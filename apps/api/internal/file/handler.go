package file

import (
	"errors"
	"net/http"

	"github.com/feyysel/mela-drive-cloud/internal/errs"
	"github.com/feyysel/mela-drive-cloud/internal/ginx"
	"github.com/feyysel/mela-drive-cloud/internal/res"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const multipartOverhead = 1 << 20

type Handler struct {
	svc *Service
}

func NewHandler(svc *Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) Upload(ctx *gin.Context) {
	ownerID, err := ginx.UserID(ctx)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	ctx.Request.Body = http.MaxBytesReader(ctx.Writer, ctx.Request.Body, maxFileSize+multipartOverhead)

	header, err := ctx.FormFile("file")
	if err != nil {
		if _, ok := errors.AsType[*http.MaxBytesError](err); ok {
			res.Error(ctx, errs.Wrap(errs.ErrInvalidInput, "File is too large"))
			return
		}

		res.Error(ctx, errs.Wrap(errs.ErrInvalidInput, "File is required"))
		return
	}

	var folderID *uuid.UUID

	if value := ctx.PostForm("folderId"); value != "" {
		parsed, err := uuid.Parse(value)
		if err != nil {
			res.Error(ctx, errs.Wrap(errs.ErrInvalidInput, "Invalid folder ID"))
			return
		}

		folderID = &parsed
	}

	source, err := header.Open()
	if err != nil {
		res.Error(ctx, err)
		return
	}
	defer func() {
		_ = source.Close()
	}()

	file, err := h.svc.Upload(ctx, ownerID, folderID, header.Filename, header.Size, source)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	res.Created(ctx, file)
}

func (h *Handler) Get(ctx *gin.Context) {
	ownerID, err := ginx.UserID(ctx)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	fileID, err := ginx.Param(ctx, "fileID")
	if err != nil {
		res.Error(ctx, err)
		return
	}

	file, err := h.svc.Get(ctx, ownerID, fileID)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	res.OK(ctx, file)
}

func (h *Handler) List(ctx *gin.Context) {
	ownerID, err := ginx.UserID(ctx)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	var folderID *uuid.UUID

	if value := ctx.Query("folderId"); value != "" {
		parsed, err := uuid.Parse(value)
		if err != nil {
			res.Error(ctx, errs.Wrap(errs.ErrInvalidInput, "Invalid folder ID"))
			return
		}

		folderID = &parsed
	}

	files, err := h.svc.List(ctx, ownerID, folderID)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	res.OK(ctx, files)
}

func (h *Handler) Rename(ctx *gin.Context) {
	ownerID, err := ginx.UserID(ctx)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	fileID, err := ginx.Param(ctx, "fileID")
	if err != nil {
		res.Error(ctx, err)
		return
	}

	var req RenameFileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res.Error(ctx, err)
		return
	}

	file, err := h.svc.Rename(ctx, ownerID, fileID, req.Name)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	res.OK(ctx, file)
}

func (h *Handler) Move(ctx *gin.Context) {
	ownerID, err := ginx.UserID(ctx)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	fileID, err := ginx.Param(ctx, "fileID")
	if err != nil {
		res.Error(ctx, err)
		return
	}

	var req MoveFileRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		res.Error(ctx, err)
		return
	}

	file, err := h.svc.Move(ctx, ownerID, fileID, req.FolderID)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	res.OK(ctx, file)
}

func (h *Handler) Trash(ctx *gin.Context) {
	ownerID, err := ginx.UserID(ctx)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	fileID, err := ginx.Param(ctx, "fileID")
	if err != nil {
		res.Error(ctx, err)
		return
	}

	if err := h.svc.Trash(ctx, ownerID, fileID); err != nil {
		res.Error(ctx, err)
		return
	}

	res.NoContent(ctx)
}

func (h *Handler) ListTrashed(ctx *gin.Context) {
	ownerID, err := ginx.UserID(ctx)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	files, err := h.svc.ListTrashed(ctx, ownerID)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	res.OK(ctx, files)
}

func (h *Handler) Restore(ctx *gin.Context) {
	ownerID, err := ginx.UserID(ctx)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	fileID, err := ginx.Param(ctx, "fileID")
	if err != nil {
		res.Error(ctx, err)
		return
	}

	if err := h.svc.Restore(ctx, ownerID, fileID); err != nil {
		res.Error(ctx, err)
		return
	}

	res.NoContent(ctx)
}

func (h *Handler) Purge(ctx *gin.Context) {
	ownerID, err := ginx.UserID(ctx)
	if err != nil {
		res.Error(ctx, err)
		return
	}

	fileID, err := ginx.Param(ctx, "fileID")
	if err != nil {
		res.Error(ctx, err)
		return
	}

	if err := h.svc.Purge(ctx, ownerID, fileID); err != nil {
		res.Error(ctx, err)
		return
	}

	res.NoContent(ctx)
}
