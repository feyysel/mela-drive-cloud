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
    parentId: string | null
    name: string
}

export type RenameFolderRequest = {
    name: string
}

export type MoveFolderRequest = {
    parentId: string | null
}
