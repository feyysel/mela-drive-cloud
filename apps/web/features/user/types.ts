export type UserRole = "Owner" | "Admin" | "Member"

export type User = {
  id: string
  firstName: string
  lastName: string | null
  email: string
  role: UserRole
  createdAt: string
  updatedAt: string
}

export type UpdateUserNameRequest = {
  firstName: string
  lastName?: string | null
}

export type UpdateEmailRequest = {
  email: string
}
