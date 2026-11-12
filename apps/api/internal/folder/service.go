package folder

import (
	"context"
	"errors"
	"log/slog"

	"github.com/feyysel/mela-drive-cloud/internal/errs"
	"github.com/feyysel/mela-drive-cloud/internal/stringsx"
	"github.com/google/uuid"
)

const (
	msgFolderNotFound = "Folder not found"
	msgParentNotFound = "Parent folder not found"
	msgNameInUse      = "A folder with this name already exists here"
	msgMoveIntoSelf   = "A folder cannot be moved into itself or its subfolders"
)

type Service struct {
	repo       Repository
	userSvc    UserService
	storageSvc StorageService
}

func NewService(repo Repository, userSvc UserService, storageSvc StorageService) *Service {
	return &Service{
		repo:       repo,
		userSvc:    userSvc,
		storageSvc: storageSvc,
	}
}

func (s *Service) Create(ctx context.Context, ownerID uuid.UUID, parentID *uuid.UUID, name string) (*Folder, error) {
	folder, err := NewFolder(ownerID, parentID, name)
	if err != nil {
		return nil, err
	}

	if parentID != nil {
		if _, err := s.repo.Get(ctx, ownerID, *parentID); err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				return nil, errs.Wrap(errs.ErrNotFound, msgParentNotFound)
			}

			return nil, err
		}
	}

	created, err := s.repo.Create(ctx, folder)
	if err != nil {
		if errors.Is(err, errs.ErrConflict) {
			return nil, errs.Wrap(errs.ErrConflict, msgNameInUse)
		}

		return nil, err
	}

	return created, nil
}

func (s *Service) Get(ctx context.Context, ownerID, folderID uuid.UUID) (*Folder, error) {
	folder, err := s.repo.Get(ctx, ownerID, folderID)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, errs.Wrap(errs.ErrNotFound, msgFolderNotFound)
		}

		return nil, err
	}

	return folder, nil
}

func (s *Service) List(ctx context.Context, ownerID uuid.UUID, parentID *uuid.UUID) ([]*Folder, error) {
	if parentID != nil {
		if _, err := s.repo.Get(ctx, ownerID, *parentID); err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				return nil, errs.Wrap(errs.ErrNotFound, msgParentNotFound)
			}

			return nil, err
		}
	}

	return s.repo.List(ctx, ownerID, parentID)
}

func (s *Service) Rename(ctx context.Context, ownerID, folderID uuid.UUID, name string) (*Folder, error) {
	trimmedName := stringsx.Trim(name)
	if err := requireValidName(trimmedName); err != nil {
		return nil, err
	}

	renamed, err := s.repo.Rename(ctx, ownerID, folderID, trimmedName)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, errs.Wrap(errs.ErrNotFound, msgFolderNotFound)
		}

		if errors.Is(err, errs.ErrConflict) {
			return nil, errs.Wrap(errs.ErrConflict, msgNameInUse)
		}

		return nil, err
	}

	return renamed, nil
}

func (s *Service) Move(ctx context.Context, ownerID, folderID uuid.UUID, newParentID *uuid.UUID) (*Folder, error) {
	if newParentID != nil {
		if *newParentID == folderID {
			return nil, errs.Wrap(errs.ErrInvalidInput, msgMoveIntoSelf)
		}

		parent, err := s.repo.Get(ctx, ownerID, *newParentID)
		if err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				return nil, errs.Wrap(errs.ErrNotFound, msgParentNotFound)
			}

			return nil, err
		}

		for parent.ParentID != nil {
			if *parent.ParentID == folderID {
				return nil, errs.Wrap(errs.ErrInvalidInput, msgMoveIntoSelf)
			}

			parent, err = s.repo.Get(ctx, ownerID, *parent.ParentID)
			if err != nil {
				return nil, err
			}
		}
	}

	moved, err := s.repo.Move(ctx, ownerID, folderID, newParentID)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, errs.Wrap(errs.ErrNotFound, msgFolderNotFound)
		}

		if errors.Is(err, errs.ErrConflict) {
			return nil, errs.Wrap(errs.ErrConflict, msgNameInUse)
		}

		return nil, err
	}

	return moved, nil
}

func (s *Service) Trash(ctx context.Context, ownerID, folderID uuid.UUID) error {
	if err := s.repo.Trash(ctx, ownerID, folderID); err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.Wrap(errs.ErrNotFound, msgFolderNotFound)
		}

		return err
	}

	return nil
}

func (s *Service) ListTrashed(ctx context.Context, ownerID uuid.UUID) ([]*Folder, error) {
	return s.repo.ListTrashed(ctx, ownerID)
}

func (s *Service) Restore(ctx context.Context, ownerID, folderID uuid.UUID) error {
	trashed, err := s.repo.GetTrashed(ctx, ownerID, folderID)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.Wrap(errs.ErrNotFound, msgFolderNotFound)
		}

		return err
	}

	parentID := trashed.ParentID

	if parentID != nil {
		if _, err := s.repo.Get(ctx, ownerID, *parentID); err != nil {
			if !errors.Is(err, errs.ErrNotFound) {
				return err
			}

			parentID = nil
		}
	}

	if err := s.repo.Restore(ctx, ownerID, folderID, parentID); err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.Wrap(errs.ErrNotFound, msgFolderNotFound)
		}

		if errors.Is(err, errs.ErrConflict) {
			return errs.Wrap(errs.ErrConflict, msgNameInUse)
		}

		return err
	}

	return nil
}

func (s *Service) Purge(ctx context.Context, ownerID, folderID uuid.UUID) error {
	if _, err := s.repo.GetTrashed(ctx, ownerID, folderID); err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.Wrap(errs.ErrNotFound, msgFolderNotFound)
		}

		return err
	}

	keys, totalSize, err := s.repo.ListFileKeys(ctx, ownerID, folderID)
	if err != nil {
		return err
	}

	for _, key := range keys {
		if err := s.storageSvc.Delete(ctx, key); err != nil {
			return err
		}
	}

	if err := s.repo.Delete(ctx, ownerID, folderID); err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.Wrap(errs.ErrNotFound, msgFolderNotFound)
		}

		return err
	}

	if err := s.userSvc.ReleaseStorage(ctx, ownerID, totalSize); err != nil {
		slog.ErrorContext(ctx, "failed to release storage after folder purge", "user_id", ownerID, "size_bytes", totalSize, "err", err)
	}

	return nil
}
