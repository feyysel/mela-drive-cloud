package folder

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, folder *Folder) (*Folder, error)
	Get(ctx context.Context, ownerID, folderID uuid.UUID) (*Folder, error)
	List(ctx context.Context, ownerID uuid.UUID, parentID *uuid.UUID) ([]*Folder, error)
	Rename(ctx context.Context, ownerID, folderID uuid.UUID, name string) (*Folder, error)
	Move(ctx context.Context, ownerID, folderID uuid.UUID, newParentID *uuid.UUID) (*Folder, error)
}
