
import { api } from "@/store/api"
import type {
    CreateFolderRequest,
    Folder,
    RenameFolderArgs,
} from "./types"

export const folderApi = api.injectEndpoints({
    endpoints: (build) => ({
        listFolders: build.query<Folder[], string | void>({
            query: (parentId) => ({
                url: "/folders",
                params: parentId ? { parentId } : undefined,
            }),
            providesTags: (result) => [
                { type: "Folder", id: "LIST" },
                ...(result ?? []).map(({ id }) => ({
                    type: "Folder" as const,
                    id,
                })),
            ],
        }),

        getFolder: build.query<Folder, string>({
            query: (folderId) => `/folders/${folderId}`,
            providesTags: (_result, _error, folderId) => [
                { type: "Folder", id: folderId },
            ],
        }),

        createFolder: build.mutation<Folder, CreateFolderRequest>({
            query: (body) => ({
                url: "/folders",
                method: "POST",
                body,
            }),
            invalidatesTags: [{ type: "Folder", id: "LIST" }],
        }),

        renameFolder: build.mutation<Folder, RenameFolderArgs>({
            query: ({ folderId, name }) => ({
                url: `/folders/${folderId}/name`,
                method: "PATCH",
                body: { name },
            }),
            invalidatesTags: (_result, _error, { folderId }) => [
                { type: "Folder", id: "LIST" },
                { type: "Folder", id: folderId },
            ],
        }),
    }),
})

export const {
    useListFoldersQuery,
    useGetFolderQuery,
    useCreateFolderMutation,
    useRenameFolderMutation,
} = folderApi
