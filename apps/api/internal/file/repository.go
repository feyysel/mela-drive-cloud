package file

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, file *File) (*File, error)
	Get(ctx context.Context, ownerID, fileID uuid.UUID) (*File, error)
	List(ctx context.Context, ownerID uuid.UUID, folderID *uuid.UUID) ([]*File, error)
	Rename(ctx context.Context, ownerID, fileID uuid.UUID, name string) (*File, error)
	Move(ctx context.Context, ownerID, fileID uuid.UUID, newFolderID *uuid.UUID) (*File, error)
}
