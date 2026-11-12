import { api } from "@/store/api"

export const trashApi = api.injectEndpoints({
  endpoints: (build) => ({
    emptyTrash: build.mutation<void, void>({
      query: () => ({
        url: "/trash",
        method: "DELETE",
      }),
      invalidatesTags: [
        { type: "Folder", id: "TRASH_LIST" },
        { type: "Folder", id: "LIST" },
        { type: "File", id: "TRASH_LIST" },
        { type: "File", id: "LIST" },
        "User",
      ],
    }),
  }),
})

export const { useEmptyTrashMutation } = trashApi
