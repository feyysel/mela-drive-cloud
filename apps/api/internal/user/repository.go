package user

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, user *User) (*User, error)
	GetByID(ctx context.Context, userID uuid.UUID) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	CountByRole(ctx context.Context, role UserRole) (int, error)
	UpdateName(ctx context.Context, userID uuid.UUID, firstName string, lastName *string) (*User, error)
	UpdateEmail(ctx context.Context, userID uuid.UUID, email string) (*User, error)
	UpdatePassword(ctx context.Context, userID uuid.UUID, passwordHash string) (*User, error)
	Delete(ctx context.Context, userID uuid.UUID) error
}
