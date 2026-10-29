package ginx

import (
	"fmt"

	"github.com/feyysel/mela-drive-cloud/internal/errs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Param(ctx *gin.Context, name string) (uuid.UUID, error) {
	id, err := uuid.Parse(ctx.Param(name))
	if err != nil {
		return uuid.Nil, errs.Wrap(errs.ErrInvalidInput, fmt.Sprintf("Invalid %s", name))
	}

	return id, nil
}
