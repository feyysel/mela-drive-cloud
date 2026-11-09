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
      providesTags: (_result, _error, folderId) => [
        { type: "Folder", id: folderId },
      ],
    }),
    listFolders: build.query<Folder[], string | undefined>({
      query: (parentId) => ({
        url: "/folders",
        params: { parentId },
      }),
      providesTags: [{ type: "Folder", id: "LIST" }],
    }),
    createFolder: build.mutation<Folder, CreateFolderRequest>({
      query: (body) => ({
        url: "/folders",
        method: "POST",
        body,
      }),
      invalidatesTags: [{ type: "Folder", id: "LIST" }],
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
      invalidatesTags: (_result, _error, { folderId }) => [
        { type: "Folder", id: folderId },
        { type: "Folder", id: "LIST" },
      ],
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
      invalidatesTags: (_result, _error, { folderId }) => [
        { type: "Folder", id: folderId },
        { type: "Folder", id: "LIST" },
      ],
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
