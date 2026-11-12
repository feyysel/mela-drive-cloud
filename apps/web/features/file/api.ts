import { api } from "@/store/api"
import type {
  DriveFile,
  MoveFileRequest,
  RenameFileRequest,
  UploadFileRequest,
} from "./types"

export const fileApi = api.injectEndpoints({
  endpoints: (build) => ({
    getFile: build.query<DriveFile, string>({
      query: (fileId) => `/files/${fileId}`,
      providesTags: (_result, _error, fileId) => [{ type: "File", id: fileId }],
    }),
    listFiles: build.query<DriveFile[], string | undefined>({
      query: (folderId) => ({
        url: "/files",
        params: { folderId },
      }),
      providesTags: [{ type: "File", id: "LIST" }],
    }),
    uploadFile: build.mutation<DriveFile, UploadFileRequest>({
      query: ({ folderId, file }) => {
        const body = new FormData()
        body.append("file", file)
        if (folderId) {
          body.append("folderId", folderId)
        }

        return {
          url: "/files",
          method: "POST",
          body,
        }
      },
      invalidatesTags: [{ type: "File", id: "LIST" }, "User"],
    }),
    renameFile: build.mutation<
      DriveFile,
      { fileId: string } & RenameFileRequest
    >({
      query: ({ fileId, ...body }) => ({
        url: `/files/${fileId}/name`,
        method: "PATCH",
        body,
      }),
      invalidatesTags: (_result, _error, { fileId }) => [
        { type: "File", id: fileId },
        { type: "File", id: "LIST" },
      ],
    }),
    moveFile: build.mutation<DriveFile, { fileId: string } & MoveFileRequest>({
      query: ({ fileId, ...body }) => ({
        url: `/files/${fileId}/folder`,
        method: "PATCH",
        body,
      }),
      invalidatesTags: (_result, _error, { fileId }) => [
        { type: "File", id: fileId },
        { type: "File", id: "LIST" },
      ],
    }),
    listTrashedFiles: build.query<DriveFile[], void>({
      query: () => "/files/trash",
      providesTags: [{ type: "File", id: "TRASH_LIST" }],
    }),
    trashFile: build.mutation<void, string>({
      query: (fileId) => ({
        url: `/files/${fileId}`,
        method: "DELETE",
      }),
      invalidatesTags: (_result, _error, fileId) => [
        { type: "File", id: fileId },
        { type: "File", id: "LIST" },
        { type: "File", id: "TRASH_LIST" },
      ],
    }),
    restoreFile: build.mutation<void, string>({
      query: (fileId) => ({
        url: `/files/${fileId}/restore`,
        method: "POST",
      }),
      invalidatesTags: [
        { type: "File", id: "LIST" },
        { type: "File", id: "TRASH_LIST" },
      ],
    }),
    purgeFile: build.mutation<void, string>({
      query: (fileId) => ({
        url: `/files/${fileId}/purge`,
        method: "DELETE",
      }),
      invalidatesTags: [{ type: "File", id: "TRASH_LIST" }, "User"],
    }),
  }),
})

export const {
  useGetFileQuery,
  useListFilesQuery,
  useUploadFileMutation,
  useRenameFileMutation,
  useMoveFileMutation,
  useListTrashedFilesQuery,
  useTrashFileMutation,
  useRestoreFileMutation,
  usePurgeFileMutation,
} = fileApi
