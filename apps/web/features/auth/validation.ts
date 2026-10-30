// Mirrors the rules in apps/api/internal/user/validation.go so users get
// feedback before a round trip. The API remains the source of truth.

export interface SignupValues {
  firstName: string
  lastName: string
  email: string
  password: string
  confirmPassword: string
}

export type SignupField = keyof SignupValues

export type SignupErrors = Partial<Record<SignupField, string>>

export const signupFields: SignupField[] = [
  "firstName",
  "lastName",
  "email",
  "password",
  "confirmPassword",
]

const NAME_MAX = 50
const FIRST_NAME_MIN = 2
const PASSWORD_MIN = 8
// bcrypt ignores input beyond 72 bytes, so the API rejects longer passwords.
const PASSWORD_MAX_BYTES = 72
const EMAIL_PATTERN = /^[^\s@]+@[^\s@]+\.[^\s@]+$/

// Counts code points, matching Go's utf8.RuneCountInString.
function runeCount(value: string) {
  return Array.from(value).length
}

export function validateSignup(values: SignupValues): SignupErrors {
  const errors: SignupErrors = {}

  const firstName = values.firstName.trim()
  if (!firstName) {
    errors.firstName = "First name is required"
  } else if (runeCount(firstName) < FIRST_NAME_MIN) {
    errors.firstName = "First name must be at least 2 characters"
  } else if (runeCount(firstName) > NAME_MAX) {
    errors.firstName = "First name must be at most 50 characters"
  }

  if (runeCount(values.lastName.trim()) > NAME_MAX) {
    errors.lastName = "Last name must be at most 50 characters"
  }

  const email = values.email.trim()
  if (!email) {
    errors.email = "Email is required"
  } else if (!EMAIL_PATTERN.test(email)) {
    errors.email = "Enter a valid email address"
  }

  if (!values.password.trim()) {
    errors.password = "Password is required"
  } else if (runeCount(values.password) < PASSWORD_MIN) {
    errors.password = "Password must be at least 8 characters"
  } else if (
    new TextEncoder().encode(values.password).length > PASSWORD_MAX_BYTES
  ) {
    errors.password = "Password is too long"
  }

  if (values.confirmPassword !== values.password) {
    errors.confirmPassword = "Passwords do not match"
  }

  return errors
}
