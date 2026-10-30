import { api } from "@/store/api"
import type { UpdateEmailRequest, UpdateUserNameRequest, User } from "./types"

export const userApi = api.injectEndpoints({
  endpoints: (build) => ({
    me: build.query<User, void>({
      query: () => "/me",
      providesTags: ["User"],
    }),
    updateName: build.mutation<User, UpdateUserNameRequest>({
      query: (body) => ({
        url: "/me/name",
        method: "PATCH",
        body,
      }),
      invalidatesTags: ["User"],
    }),
    updateEmail: build.mutation<User, UpdateEmailRequest>({
      query: (body) => ({
        url: "/me/email",
        method: "PATCH",
        body,
      }),
      invalidatesTags: ["User"],
    }),
    deleteMe: build.mutation<void, void>({
      query: () => ({
        url: "/me",
        method: "DELETE",
      }),
      invalidatesTags: ["User"],
    }),
  }),
})

export const {
  useMeQuery,
  useUpdateNameMutation,
  useUpdateEmailMutation,
  useDeleteMeMutation,
} = userApi
