package user

import (
	"context"
	"errors"
	"log/slog"

	"github.com/feyysel/mela-drive-cloud/internal/config"
	"github.com/feyysel/mela-drive-cloud/internal/emailx"
	"github.com/feyysel/mela-drive-cloud/internal/errs"
	"github.com/feyysel/mela-drive-cloud/internal/stringsx"
	"github.com/google/uuid"
)

const (
	msgUserNotFound = "User not found"
	msgEmailInUse   = "Email already in use"
)

type Service struct {
	repo   Repository
	cache  Cache
	appCfg config.AppConfig
}

func NewService(repo Repository, cache Cache, appCfg config.AppConfig) *Service {
	return &Service{
		repo:   repo,
		cache:  cache,
		appCfg: appCfg,
	}
}

func (s *Service) Create(ctx context.Context, firstName string, lastName *string, email string, passwordHash string, role UserRole) (*User, error) {
	user, err := NewUser(firstName, lastName, email, "", role)
	if err != nil {
		return nil, err
	}

	created, err := s.repo.Create(ctx, user)
	if err != nil {
		if errors.Is(err, errs.ErrConflict) {
			return nil, errs.Wrap(errs.ErrConflict, msgEmailInUse)
		}

		return nil, err
	}

	return created, nil
}

func (s *Service) GetSnapshot(ctx context.Context, userID uuid.UUID) (*User, error) {
	snapshot, err := s.cache.Get(ctx, userSnapshotKey(userID))
	if err == nil {
		return snapshot, nil
	}

	if !errors.Is(err, errs.ErrNotFound) {
		slog.WarnContext(ctx, "failed to read user snapshot from cache", "user_id", userID, "err", err)
	}

	user, err := s.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if err := s.cache.Set(ctx, userSnapshotKey(userID), user, userSnapshotTTL); err != nil {
		slog.WarnContext(ctx, "failed to cache user snapshot", "user_id", userID, "err", err)
	}

	return user, nil
}

func (s *Service) GetByID(ctx context.Context, userID uuid.UUID) (*User, error) {
	user, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, errs.Wrap(errs.ErrNotFound, msgUserNotFound)
		}

		return nil, err
	}

	return user, nil
}

func (s *Service) GetByEmail(ctx context.Context, email string) (*User, error) {
	user, err := s.repo.GetByEmail(ctx, emailx.Normalize(stringsx.Trim(email)))
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, errs.Wrap(errs.ErrNotFound, msgUserNotFound)
		}

		return nil, err
	}

	return user, nil
}

func (s *Service) UpdateName(ctx context.Context, userID uuid.UUID, firstName string, lastName *string) (*User, error) {
	trimmedFirstName := stringsx.Trim(firstName)
	if err := requireValidFirstName(trimmedFirstName); err != nil {
		return nil, err
	}

	trimmedLastName := stringsx.TrimOptional(lastName)
	if err := requireValidLastName(trimmedLastName); err != nil {
		return nil, err
	}

	updated, err := s.repo.UpdateName(ctx, userID, trimmedFirstName, trimmedLastName)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, errs.Wrap(errs.ErrNotFound, msgUserNotFound)
		}

		return nil, err
	}

	return updated, nil
}

func (s *Service) UpdateEmail(ctx context.Context, userID uuid.UUID, email string) (*User, error) {
	trimmedEmail := stringsx.Trim(email)
	if err := requireValidEmail(trimmedEmail); err != nil {
		return nil, err
	}

	updated, err := s.repo.UpdateEmail(ctx, userID, emailx.Normalize(trimmedEmail))
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, errs.Wrap(errs.ErrNotFound, msgUserNotFound)
		}

		if errors.Is(err, errs.ErrConflict) {
			return nil, errs.Wrap(errs.ErrConflict, msgEmailInUse)
		}

		return nil, err
	}

	return updated, nil
}

func (s *Service) Delete(ctx context.Context, userID uuid.UUID) error {
	if err := s.repo.Delete(ctx, userID); err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.Wrap(errs.ErrNotFound, msgUserNotFound)
		}

		return err
	}

	return nil
}

func (s *Service) SeedOwner(ctx context.Context) error {
	count, err := s.repo.CountByRole(ctx, UserRoleOwner)
	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	user, err := NewUser("Owner", nil, s.appCfg.OwnerEmail, "", UserRoleOwner)
	if err != nil {
		return err
	}

	owner, err := s.repo.Create(ctx, user)
	if err != nil {
		if errors.Is(err, errs.ErrConflict) {
			return errs.Wrap(errs.ErrConflict, msgEmailInUse)
		}

		return err
	}

	slog.InfoContext(ctx, "owner seeded", "email", emailx.Mask(owner.Email))

	return nil
}
