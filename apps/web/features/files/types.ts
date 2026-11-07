// Named DriveFile to avoid clashing with the browser's built-in File type.
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
  file: File
  folderId?: string | null
}

export type RenameFileRequest = {
  fileId: string
  name: string
}

export type MoveFileRequest = {
  fileId: string
  folderId: string | null
}
