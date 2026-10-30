export type LoginRequest = {
  email: string
  password: string
}

export type SignupRequest = {
  firstName: string
  lastName?: string | null
  email: string
  password: string
}

export type ForgotPasswordRequest = {
  email: string
}

export type ResetPasswordRequest = {
  token: string
  password: string
}
