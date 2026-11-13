import { api } from "@/store/api"
import type {
  ForgotPasswordRequest,
  LoginRequest,
  ResetPasswordRequest,
  SignupRequest,
} from "./types"

export const authApi = api.injectEndpoints({
  endpoints: (build) => ({
    signup: build.mutation<void, SignupRequest>({
      query: (body) => ({
        url: "/auth/signup",
        method: "POST",
        body,
      }),
      invalidatesTags: ["User", "File", "Folder"],
    }),
    login: build.mutation<void, LoginRequest>({
      query: (body) => ({
        url: "/auth/login",
        method: "POST",
        body,
      }),
      invalidatesTags: ["User", "File", "Folder"],
    }),
    logout: build.mutation<void, void>({
      query: () => ({
        url: "/auth/logout",
        method: "POST",
      }),
      invalidatesTags: ["User", "File", "Folder"],
    }),
    forgotPassword: build.mutation<void, ForgotPasswordRequest>({
      query: (body) => ({
        url: "/auth/forgot-password",
        method: "POST",
        body,
      }),
    }),
    resetPassword: build.mutation<void, ResetPasswordRequest>({
      query: (body) => ({
        url: "/auth/reset-password",
        method: "POST",
        body,
      }),
    }),
    validateResetToken: build.query<void, string>({
      query: (token) => ({
        url: "/auth/reset-password/validate",
        params: { token },
      }),
      keepUnusedDataFor: 0,
    }),
  }),
})

export const {
  useSignupMutation,
  useLoginMutation,
  useLogoutMutation,
  useForgotPasswordMutation,
  useResetPasswordMutation,
  useValidateResetTokenQuery,
} = authApi
