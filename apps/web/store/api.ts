import { config } from "@/lib/config";
import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const api = createApi({
  reducerPath: "api",

  baseQuery: fetchBaseQuery({
    baseUrl: config.API_BASE_URL,
    credentials: "include",
  }),

  tagTypes: [],

  endpoints: () => ({}),
});
