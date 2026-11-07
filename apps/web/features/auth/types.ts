export type LoginRequest = {
  email: string
  password: string
}

export type SignupRequest = {
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

export type ValidateResetTokenRequest = {
  token: string
}
