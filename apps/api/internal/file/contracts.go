package file

import (
	"context"
	"io"

	"github.com/feyysel/mela-drive-cloud/internal/folder"
	"github.com/feyysel/mela-drive-cloud/internal/user"
	"github.com/google/uuid"
)

type UserService interface {
	GetByID(ctx context.Context, userID uuid.UUID) (*user.User, error)
}

type FolderService interface {
	Get(ctx context.Context, ownerID, folderID uuid.UUID) (*folder.Folder, error)
}

type StorageService interface {
	Upload(ctx context.Context, key, contentType string, body io.Reader) error
	Delete(ctx context.Context, key string) error
	PublicURL(key string) string
}
