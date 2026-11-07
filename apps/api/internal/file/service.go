package file

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/feyysel/mela-drive-cloud/internal/errs"
	"github.com/feyysel/mela-drive-cloud/internal/stringsx"
	"github.com/google/uuid"
)

const (
	msgFileNotFound   = "File not found"
	msgFolderNotFound = "Folder not found"
	msgNameInUse      = "A file with this name already exists here"

	quotaBytes     = 1 << 30
	sniffSize      = 512
	cleanupTimeout = 10 * time.Second
)

type Service struct {
	repo       Repository
	userSvc    UserService
	folderSvc  FolderService
	storageSvc StorageService
}

func NewService(repo Repository, userSvc UserService, folderSvc FolderService, storageSvc StorageService) *Service {
	return &Service{
		repo:       repo,
		userSvc:    userSvc,
		folderSvc:  folderSvc,
		storageSvc: storageSvc,
	}
}

func (s *Service) Upload(ctx context.Context, ownerID uuid.UUID, folderID *uuid.UUID, name string, sizeBytes int64, body io.Reader) (*File, error) {
	reader := bufio.NewReaderSize(body, sniffSize)

	head, err := reader.Peek(sniffSize)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}

	objectID, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	storageKey := fmt.Sprintf("files/%s/%s", ownerID, objectID)

	file, err := NewFile(ownerID, folderID, name, storageKey, sizeBytes, http.DetectContentType(head))
	if err != nil {
		return nil, err
	}

	if folderID != nil {
		if _, err := s.folderSvc.Get(ctx, ownerID, *folderID); err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				return nil, errs.Wrap(errs.ErrNotFound, msgFolderNotFound)
			}

			return nil, err
		}
	}

	if err := s.userSvc.ReserveStorage(ctx, ownerID, sizeBytes, quotaBytes); err != nil {
		return nil, err
	}

	if err := s.storageSvc.Upload(ctx, file.StorageKey, file.MimeType, reader); err != nil {
		cleanupCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), cleanupTimeout)
		defer cancel()

		if releaseErr := s.userSvc.ReleaseStorage(cleanupCtx, ownerID, sizeBytes); releaseErr != nil {
			slog.ErrorContext(ctx, "failed to release reserved storage", "user_id", ownerID, "size_bytes", sizeBytes, "err", releaseErr)
		}

		return nil, err
	}

	created, err := s.repo.Create(ctx, file)
	if err != nil {
		cleanupCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), cleanupTimeout)
		defer cancel()

		if deleteErr := s.storageSvc.Delete(cleanupCtx, file.StorageKey); deleteErr != nil {
			slog.ErrorContext(ctx, "failed to delete orphaned object", "storage_key", file.StorageKey, "err", deleteErr)
		}

		if releaseErr := s.userSvc.ReleaseStorage(cleanupCtx, ownerID, sizeBytes); releaseErr != nil {
			slog.ErrorContext(ctx, "failed to release reserved storage", "user_id", ownerID, "size_bytes", sizeBytes, "err", releaseErr)
		}

		if errors.Is(err, errs.ErrConflict) {
			return nil, errs.Wrap(errs.ErrConflict, msgNameInUse)
		}

		return nil, err
	}

	created.URL = s.storageSvc.PublicURL(created.StorageKey)

	return created, nil
}

func (s *Service) Get(ctx context.Context, ownerID, fileID uuid.UUID) (*File, error) {
	file, err := s.repo.Get(ctx, ownerID, fileID)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, errs.Wrap(errs.ErrNotFound, msgFileNotFound)
		}

		return nil, err
	}

	file.URL = s.storageSvc.PublicURL(file.StorageKey)

	return file, nil
}

func (s *Service) List(ctx context.Context, ownerID uuid.UUID, folderID *uuid.UUID) ([]*File, error) {
	if folderID != nil {
		if _, err := s.folderSvc.Get(ctx, ownerID, *folderID); err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				return nil, errs.Wrap(errs.ErrNotFound, msgFolderNotFound)
			}

			return nil, err
		}
	}

	files, err := s.repo.List(ctx, ownerID, folderID)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		file.URL = s.storageSvc.PublicURL(file.StorageKey)
	}

	return files, nil
}

func (s *Service) Rename(ctx context.Context, ownerID, fileID uuid.UUID, name string) (*File, error) {
	trimmedName := stringsx.Trim(name)
	if err := requireValidName(trimmedName); err != nil {
		return nil, err
	}

	renamed, err := s.repo.Rename(ctx, ownerID, fileID, trimmedName)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, errs.Wrap(errs.ErrNotFound, msgFileNotFound)
		}

		if errors.Is(err, errs.ErrConflict) {
			return nil, errs.Wrap(errs.ErrConflict, msgNameInUse)
		}

		return nil, err
	}

	renamed.URL = s.storageSvc.PublicURL(renamed.StorageKey)

	return renamed, nil
}

func (s *Service) Move(ctx context.Context, ownerID, fileID uuid.UUID, newFolderID *uuid.UUID) (*File, error) {
	if newFolderID != nil {
		if _, err := s.folderSvc.Get(ctx, ownerID, *newFolderID); err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				return nil, errs.Wrap(errs.ErrNotFound, msgFolderNotFound)
			}

			return nil, err
		}
	}

	moved, err := s.repo.Move(ctx, ownerID, fileID, newFolderID)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return nil, errs.Wrap(errs.ErrNotFound, msgFileNotFound)
		}

		if errors.Is(err, errs.ErrConflict) {
			return nil, errs.Wrap(errs.ErrConflict, msgNameInUse)
		}

		return nil, err
	}

	moved.URL = s.storageSvc.PublicURL(moved.StorageKey)

	return moved, nil
}
