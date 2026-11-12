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

	GetTrashed(ctx context.Context, ownerID, fileID uuid.UUID) (*File, error)
	ListTrashed(ctx context.Context, ownerID uuid.UUID) ([]*File, error)
	Trash(ctx context.Context, ownerID, fileID uuid.UUID) error
	Restore(ctx context.Context, ownerID, fileID uuid.UUID, folderID *uuid.UUID) error
	Delete(ctx context.Context, ownerID, fileID uuid.UUID) error
}
