package user

import (
	"time"

	"github.com/feyysel/mela-drive-cloud/internal/emailx"
	"github.com/feyysel/mela-drive-cloud/internal/stringsx"
	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID `json:"id"`
	FirstName        string    `json:"firstName"`
	LastName         *string   `json:"lastName"`
	Email            string    `json:"email"`
	PasswordHash     string    `json:"-"`
	StorageUsedBytes int64     `json:"storageUsedBytes"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

func NewUser(firstName string, lastName *string, email string, passwordHash string) (*User, error) {
	trimmedFirstName := stringsx.Trim(firstName)
	if err := requireValidFirstName(trimmedFirstName); err != nil {
		return nil, err
	}

	trimmedLastName := stringsx.TrimOptional(lastName)
	if err := requireValidLastName(trimmedLastName); err != nil {
		return nil, err
	}

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
		FirstName:    trimmedFirstName,
		LastName:     trimmedLastName,
		Email:        emailx.Normalize(email),
		PasswordHash: passwordHash,
	}, nil
}
