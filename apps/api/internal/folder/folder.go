package folder

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Folder struct {
	ID        uuid.UUID  `json:"id"`
	OwnerID   uuid.UUID  `json:"ownerId"`
	ParentID  *uuid.UUID `json:"parentId"`
	Name      string     `json:"name"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}

func NewFolder(ownerID uuid.UUID, parentID *uuid.UUID, name string) (*Folder, error) {
	trimmedName := strings.TrimSpace(name)
	if err := requireValidName(trimmedName); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &Folder{
		ID:       id,
		OwnerID:  ownerID,
		ParentID: parentID,
		Name:     trimmedName,
	}, nil
}
