import { api } from "@/store/api"
import type {
    CreateFolderRequest,
    Folder,
    MoveFolderRequest,
    RenameFolderRequest,
} from "./types"

export const folderApi = api.injectEndpoints({
    endpoints: (build) => ({
        getFolder: build.query<Folder, string>({
            query: (folderId) => `/folders/${folderId}`,
        }),
        listFolders: build.query<Folder[], string | undefined>({
            query: (parentId) => ({
                url: "/folders",
                params: { parentId },
            }),
        }),
        createFolder: build.mutation<Folder, CreateFolderRequest>({
            query: (body) => ({
                url: "/folders",
                method: "POST",
                body,
            }),
        }),
        renameFolder: build.mutation<
            Folder,
            { folderId: string } & RenameFolderRequest
        >({
            query: ({ folderId, ...body }) => ({
                url: `/folders/${folderId}/name`,
                method: "PATCH",
                body,
            }),
        }),
        moveFolder: build.mutation<
            Folder,
            { folderId: string } & MoveFolderRequest
        >({
            query: ({ folderId, ...body }) => ({
                url: `/folders/${folderId}/parent`,
                method: "PATCH",
                body,
            }),
        }),
    }),
})

export const {
    useGetFolderQuery,
    useListFoldersQuery,
    useCreateFolderMutation,
    useRenameFolderMutation,
    useMoveFolderMutation,
} = folderApi
