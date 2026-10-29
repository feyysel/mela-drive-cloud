package res

import (
	"errors"
	"net/http"

	"github.com/feyysel/mela-drive-cloud/internal/errs"
	"github.com/feyysel/mela-drive-cloud/internal/ginx"
	"github.com/gin-gonic/gin"
)

func Error(ctx *gin.Context, err error) {
	if err == nil {
		err = errs.ErrInternal
	}

	ctx.AbortWithStatusJSON(status(err), ErrorResponse{
		Message: message(ctx, err),
	})
}

func Created(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusCreated, data)
}

func OK(ctx *gin.Context, data any) {
	ctx.JSON(http.StatusOK, data)
}

func NoContent(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

var sentinels = []error{
	errs.ErrNotFound,
	errs.ErrConflict,
	errs.ErrUnauthorized,
	errs.ErrForbidden,
	errs.ErrInvalidInput,
	errs.ErrTooManyRequests,
	errs.ErrPayloadTooLarge,
	errs.ErrUnavailable,
	errs.ErrInternal,
}

func sentinel(err error) error {
	for _, s := range sentinels {
		if errors.Is(err, s) {
			return s
		}
	}

	return nil
}

func status(err error) int {
	switch sentinel(err) {
	case errs.ErrNotFound:
		return http.StatusNotFound
	case errs.ErrConflict:
		return http.StatusConflict
	case errs.ErrUnauthorized:
		return http.StatusUnauthorized
	case errs.ErrForbidden:
		return http.StatusForbidden
	case errs.ErrInvalidInput:
		return http.StatusUnprocessableEntity
	case errs.ErrTooManyRequests:
		return http.StatusTooManyRequests
	case errs.ErrPayloadTooLarge:
		return http.StatusRequestEntityTooLarge
	case errs.ErrUnavailable:
		return http.StatusServiceUnavailable
	default:
		return http.StatusInternalServerError
	}
}

func message(ctx *gin.Context, err error) string {
	if ginx.IsDebug(ctx) {
		return err.Error()
	}

	s := sentinel(err)
	if s == nil || s == errs.ErrInternal {
		return "Something went wrong"
	}

	if msg := errs.Message(err); msg != "" {
		return msg
	}

	return s.Error()
}
