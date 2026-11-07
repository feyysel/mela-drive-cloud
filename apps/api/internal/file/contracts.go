package file

import (
	"context"
	"io"

	"github.com/feyysel/mela-drive-cloud/internal/folder"
	"github.com/google/uuid"
)

type UserService interface {
	ReserveStorage(ctx context.Context, userID uuid.UUID, sizeBytes int64) error
	ReleaseStorage(ctx context.Context, userID uuid.UUID, sizeBytes int64) error
}

type FolderService interface {
	Get(ctx context.Context, ownerID, folderID uuid.UUID) (*folder.Folder, error)
}

type StorageService interface {
	Upload(ctx context.Context, key, contentType string, body io.Reader) error
	Delete(ctx context.Context, key string) error
	PublicURL(key string) string
}
