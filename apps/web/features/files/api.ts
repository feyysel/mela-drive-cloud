import { api } from "@/store/api"
import type {
  DriveFile,
  MoveFileRequest,
  RenameFileRequest,
  UploadFileRequest,
} from "./types"

// Each folder's file list gets its own tag, with "root" for files outside
// any folder, so a change only refetches the lists it affects.
const listTag = (folderId: string | null | void) => ({
  type: "File" as const,
  id: `LIST-${folderId ?? "root"}`,
})

export const fileApi = api.injectEndpoints({
  endpoints: (build) => ({
    listFiles: build.query<DriveFile[], string | null | void>({
      query: (folderId) => ({
        url: "/files",
        params: folderId ? { folderId } : undefined,
      }),
      providesTags: (result, _error, folderId) => [
        listTag(folderId),
        ...(result ?? []).map(({ id }) => ({ type: "File" as const, id })),
      ],
    }),
    getFile: build.query<DriveFile, string>({
      query: (fileId) => `/files/${fileId}`,
      providesTags: (_result, _error, fileId) => [{ type: "File", id: fileId }],
    }),
    uploadFile: build.mutation<DriveFile, UploadFileRequest>({
      query: ({ file, folderId }) => {
        const body = new FormData()
        body.append("file", file)
        if (folderId) body.append("folderId", folderId)

        // The browser sets the multipart Content-Type and boundary.
        return { url: "/files", method: "POST", body }
      },
      invalidatesTags: (_result, _error, { folderId }) => [listTag(folderId)],
    }),
    renameFile: build.mutation<DriveFile, RenameFileRequest>({
      query: ({ fileId, name }) => ({
        url: `/files/${fileId}/name`,
        method: "PATCH",
        body: { name },
      }),
      // Lists provide each file's id tag, so this also refetches its list.
      invalidatesTags: (_result, _error, { fileId }) => [
        { type: "File", id: fileId },
      ],
    }),
    moveFile: build.mutation<DriveFile, MoveFileRequest>({
      query: ({ fileId, folderId }) => ({
        url: `/files/${fileId}/folder`,
        method: "PATCH",
        body: { folderId },
      }),
      // The file's id tag refetches the source list; the destination list
      // doesn't contain it yet, so it is invalidated explicitly.
      invalidatesTags: (_result, _error, { fileId, folderId }) => [
        { type: "File", id: fileId },
        listTag(folderId),
      ],
    }),
  }),
})

export const {
  useListFilesQuery,
  useGetFileQuery,
  useUploadFileMutation,
  useRenameFileMutation,
  useMoveFileMutation,
} = fileApi
