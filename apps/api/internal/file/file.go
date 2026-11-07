package file

import (
	"time"

	"github.com/feyysel/mela-drive-cloud/internal/stringsx"
	"github.com/google/uuid"
)

type File struct {
	ID         uuid.UUID  `json:"id"`
	OwnerID    uuid.UUID  `json:"ownerId"`
	FolderID   *uuid.UUID `json:"folderId"`
	Name       string     `json:"name"`
	StorageKey string     `json:"-"`
	SizeBytes  int64      `json:"sizeBytes"`
	MimeType   string     `json:"mimeType"`
	URL        string     `json:"url"`
	DeletedAt  *time.Time `json:"deletedAt,omitempty"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time  `json:"updatedAt"`
}

func NewFile(ownerID uuid.UUID, folderID *uuid.UUID, name string, storageKey string, sizeBytes int64, mimeType string) (*File, error) {
	trimmedName := stringsx.Trim(name)
	if err := requireValidName(trimmedName); err != nil {
		return nil, err
	}

	if err := requireValidStorageKey(storageKey); err != nil {
		return nil, err
	}

	if err := requireValidSize(sizeBytes); err != nil {
		return nil, err
	}

	if err := requireValidMimeType(mimeType); err != nil {
		return nil, err
	}

	fileID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &File{
		ID:         fileID,
		OwnerID:    ownerID,
		FolderID:   folderID,
		Name:       trimmedName,
		StorageKey: storageKey,
		SizeBytes:  sizeBytes,
		MimeType:   mimeType,
	}, nil
}
