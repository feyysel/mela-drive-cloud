package trash

import (
	"context"

	"github.com/google/uuid"
)

type FolderService interface {
	EmptyTrash(ctx context.Context, ownerID uuid.UUID) error
}

type FileService interface {
	EmptyTrash(ctx context.Context, ownerID uuid.UUID) error
}
