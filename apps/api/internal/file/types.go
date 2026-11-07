package file

import "github.com/google/uuid"

type RenameFileRequest struct {
	Name string `json:"name" binding:"required"`
}

type MoveFileRequest struct {
	FolderID *uuid.UUID `json:"folderId"`
}
