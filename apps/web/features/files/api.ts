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
    }),
    listFiles: build.query<DriveFile[], string | undefined>({
      query: (folderId) => ({
        url: "/files",
        params: { folderId },
      }),
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
    }),
    moveFile: build.mutation<DriveFile, { fileId: string } & MoveFileRequest>({
      query: ({ fileId, ...body }) => ({
        url: `/files/${fileId}/folder`,
        method: "PATCH",
        body,
      }),
    }),
  }),
})

export const {
  useGetFileQuery,
  useListFilesQuery,
  useUploadFileMutation,
  useRenameFileMutation,
  useMoveFileMutation,
} = fileApi
