package user

import (
"context"
"uuid"
)

type Repository interface {
Create(ctx context.Context, user *User) (*User, error)
GetByID(ctx context.Context, userID uuid.UUID) (*User, error)
GetByEmail(ctx context.Context, userID uuid.UUID) (*User, error)
UpdateName(ctx context.Context, userID uuid.UUID, firstName string, lastName *string) (*User, error)
UpdateEmail(ctx context.Context, userID uuid.UUID, email string) (*User, error)
Delete(ctx context.Context, userID uuid.UUID) error
}
