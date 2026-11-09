export type Folder = {
  id: string
  ownerId: string
  parentId?: string
  name: string
  deletedAt?: string
  createdAt: string
  updatedAt: string
}

export type CreateFolderRequest = {
  parentId?: string
  name: string
}

export type RenameFolderRequest = {
  name: string
}

export type MoveFolderRequest = {
  parentId?: string
}
