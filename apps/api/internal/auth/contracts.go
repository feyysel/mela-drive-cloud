package auth

import (
	"context"

	"github.com/feyysel/mela-drive-cloud/internal/session"
	"github.com/feyysel/mela-drive-cloud/internal/user"
	"github.com/google/uuid"
)

type UserService interface {
	Create(ctx context.Context, firstName string, lastName *string, email string, password string, role user.UserRole) (*user.User, error)
	GetByEmail(ctx context.Context, email string) (*user.User, error)
	GetSnapshot(ctx context.Context, userID uuid.UUID) (*user.User, error)
	UpdatePassword(ctx context.Context, userID uuid.UUID, password string) (*user.User, error)
}

type SessionService interface {
	Create(ctx context.Context, userID uuid.UUID, ip, userAgent string) (*session.Session, error)
	Get(ctx context.Context, sessionID uuid.UUID) (*session.Session, error)
	Revoke(ctx context.Context, userID, sessionID uuid.UUID) error
	RevokeAll(ctx context.Context, userID uuid.UUID) error
}

type MailService interface {
	SendTemplateEmail(from string, to []string, templateID string, variables map[string]any) error
}
