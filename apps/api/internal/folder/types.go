package folder

import "github.com/google/uuid"

type CreateFolderRequest struct {
	ParentID *uuid.UUID `json:"parentId"`
	Name     string     `json:"name" binding:"required"`
}

type RenameFolderRequest struct {
	Name string `json:"name" binding:"required"`
}

type MoveFolderRequest struct {
	ParentID *uuid.UUID `json:"parentId"`
}
