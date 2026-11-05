export type Folder = {
    id: string
    ownerId: string
    parentId: string | null
    name: string
    deletedAt?: string
    createdAt: string
    updatedAt: string
}

export type CreateFolderRequest = {
    name: string
    parentId?: string | null
}

export type RenameFolderArgs = {
    folderId: string
    name: string
}