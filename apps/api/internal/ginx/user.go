package ginx

import (
	"github.com/feyysel/mela-drive-cloud/internal/errs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	keyUserID    = "userID"
	keyUserEmail = "userEmail"
)

func SetUserContext(ctx *gin.Context, userID uuid.UUID, email string) {
	ctx.Set(keyUserID, userID)
	ctx.Set(keyUserEmail, email)
}

func UserID(ctx *gin.Context) (uuid.UUID, error) {
	value, exists := ctx.Get(keyUserID)
	if !exists {
		return uuid.Nil, errs.Wrap(errs.ErrInternal, "user id missing from context")
	}

	id, ok := value.(uuid.UUID)
	if !ok {
		return uuid.Nil, errs.Wrap(errs.ErrInternal, "user id has unexpected type")
	}

	return id, nil
}

func UserEmail(ctx *gin.Context) (string, error) {
	value, exists := ctx.Get(keyUserEmail)
	if !exists {
		return "", errs.Wrap(errs.ErrInternal, "user email missing from context")
	}

	email, ok := value.(string)
	if !ok {
		return "", errs.Wrap(errs.ErrInternal, "user email has unexpected type")
	}

	return email, nil
}
