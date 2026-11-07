package user

import (
	"time"

	"github.com/feyysel/mela-drive-cloud/internal/emailx"
	"github.com/google/uuid"
)

type User struct {
	ID                uuid.UUID `json:"id"`
	Email             string    `json:"email"`
	PasswordHash      string    `json:"-"`
	StorageUsedBytes  int64     `json:"storageUsedBytes"`
	StorageQuotaBytes int64     `json:"storageQuotaBytes"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

func NewUser(email string, passwordHash string) (*User, error) {
	if err := requireValidEmail(email); err != nil {
		return nil, err
	}

	if err := requireValidPasswordHash(passwordHash); err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &User{
		ID:           id,
		Email:        emailx.Normalize(email),
		PasswordHash: passwordHash,
	}, nil
}
