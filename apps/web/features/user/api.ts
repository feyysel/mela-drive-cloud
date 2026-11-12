import { api } from "@/store/api"
import type { User } from "./types"

export const userApi = api.injectEndpoints({
  endpoints: (build) => ({
    me: build.query<User, void>({
      query: () => "/users/me",
      providesTags: ["User"],
    }),
  }),
})

export const { useMeQuery } = userApi
