export type DriveFile = {
  id: string
  ownerId: string
  folderId: string | null
  name: string
  sizeBytes: number
  mimeType: string
  url: string
  deletedAt?: string
  createdAt: string
  updatedAt: string
}

export type UploadFileRequest = {
  folderId?: string
  file: Blob
}

export type RenameFileRequest = {
  name: string
}

export type MoveFileRequest = {
  folderId: string | null
}
