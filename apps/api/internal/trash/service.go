package trash

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	folderSvc FolderService
	fileSvc   FileService
}

func NewService(folderSvc FolderService, fileSvc FileService) *Service {
	return &Service{
		folderSvc: folderSvc,
		fileSvc:   fileSvc,
	}
}

func (s *Service) EmptyTrash(ctx context.Context, ownerID uuid.UUID) error {
	if err := s.fileSvc.EmptyTrash(ctx, ownerID); err != nil {
		return err
	}

	if err := s.folderSvc.EmptyTrash(ctx, ownerID); err != nil {
		return err
	}

	return nil
}
