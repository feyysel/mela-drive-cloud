package folder

import (
	"context"

	"github.com/google/uuid"
)

type UserService interface {
	ReleaseStorage(ctx context.Context, userID uuid.UUID, sizeBytes int64) error
}

type StorageService interface {
	Delete(ctx context.Context, key string) error
}
