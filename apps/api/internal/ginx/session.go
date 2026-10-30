package ginx

import (
	"github.com/feyysel/mela-drive-cloud/internal/errs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const keySessionID = "sessionID"

func SessionID(ctx *gin.Context) (uuid.UUID, error) {
	value, exists := ctx.Get(keySessionID)
	if !exists {
		return uuid.Nil, errs.Wrap(errs.ErrInternal, "session id missing from context")
	}

	id, ok := value.(uuid.UUID)
	if !ok {
		return uuid.Nil, errs.Wrap(errs.ErrInternal, "session id has unexpected type")
	}

	return id, nil
}

func SetSessionID(ctx *gin.Context, sessionID uuid.UUID) {
	ctx.Set(keySessionID, sessionID)
}
