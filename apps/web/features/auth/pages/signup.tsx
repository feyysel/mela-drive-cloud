"use client"

import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {
  Field,
  FieldDescription,
  FieldError,
  FieldGroup,
  FieldLabel,
  FieldSet,
} from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import {
  InputGroup,
  InputGroupAddon,
  InputGroupButton,
  InputGroupInput,
} from "@/components/ui/input-group"
import { Spinner } from "@/components/ui/spinner"
import { toast } from "@/components/ui/toast"
import { parseErr } from "@/store/utils"
import { EyeIcon, EyeOffIcon } from "lucide-react"
import Link from "next/link"
import { useRouter } from "next/navigation"
import { ChangeEvent, SubmitEvent, useRef, useState } from "react"
import { useSignupMutation } from "../api"
import {
  signupFields,
  validateSignup,
  type SignupErrors,
  type SignupField,
  type SignupValues,
} from "../validation"

const initialValues: SignupValues = {
  firstName: "",
  lastName: "",
  email: "",
  password: "",
  confirmPassword: "",
}

export function Signup() {
  const router = useRouter()
  const [signup, { isLoading }] = useSignupMutation()
  const [values, setValues] = useState(initialValues)
  const [errors, setErrors] = useState<SignupErrors>({})
  const [submitted, setSubmitted] = useState(false)
  const [showPassword, setShowPassword] = useState(false)
  const formRef = useRef<HTMLFormElement>(null)

  const focusField = (field: SignupField) => {
    formRef.current
      ?.querySelector<HTMLInputElement>(`[name="${field}"]`)
      ?.focus()
  }

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const next = { ...values, [e.target.name]: e.target.value }
    setValues(next)
    // Only re-validate live once the user has tried to submit.
    if (submitted) setErrors(validateSignup(next))
  }

  const handleSubmit = async (e: SubmitEvent<HTMLFormElement>) => {
    e.preventDefault()
    setSubmitted(true)

    const found = validateSignup(values)
    setErrors(found)
    const firstInvalid = signupFields.find((field) => found[field])
    if (firstInvalid) {
      focusField(firstInvalid)
      return
    }

    try {
      await signup({
        firstName: values.firstName.trim(),
        lastName: values.lastName.trim() || undefined,
        email: values.email.trim(),
        password: values.password,
      }).unwrap()
      router.replace("/")
    } catch (error: unknown) {
      const { status, message } = parseErr(error)

      if (status === 409) {
        setErrors({ email: "An account with this email already exists" })
        focusField("email")
        return
      }

      toast.add({
        type: "error",
        title:
          status === 429
            ? "Too many attempts. Please wait a moment and try again."
            : message,
      })
    }
  }

  const fieldProps = (field: SignupField) => ({
    id: field,
    name: field,
    value: values[field],
    onChange: handleChange,
    // readOnly, not disabled: disabled inputs can't take focus, which
    // would break focusing the email field after a 409.
    readOnly: isLoading,
    "aria-invalid": errors[field] ? true : undefined,
    "aria-describedby": errors[field] ? `${field}-error` : undefined,
  })

  const fieldError = (field: SignupField) =>
    errors[field] && (
      <FieldError id={`${field}-error`}>{errors[field]}</FieldError>
    )

  return (
    <main className="flex min-h-dvh flex-col px-4 py-10">
      <Card className="m-auto w-full max-w-md">
        <CardHeader>
          <CardTitle>Create your account</CardTitle>
          <CardDescription>
            Store, organize, and share your files with Mela Drive.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form ref={formRef} noValidate onSubmit={handleSubmit}>
            <FieldSet>
              <FieldGroup>
                <div className="grid gap-6 sm:grid-cols-2">
                  <Field data-invalid={!!errors.firstName}>
                    <FieldLabel htmlFor="firstName">First name</FieldLabel>
                    <Input
                      {...fieldProps("firstName")}
                      autoComplete="given-name"
                      autoFocus
                    />
                    {fieldError("firstName")}
                  </Field>

                  <Field data-invalid={!!errors.lastName}>
                    <FieldLabel htmlFor="lastName">
                      Last name
                      <span className="font-normal text-muted-foreground">
                        (optional)
                      </span>
                    </FieldLabel>
                    <Input
                      {...fieldProps("lastName")}
                      autoComplete="family-name"
                    />
                    {fieldError("lastName")}
                  </Field>
                </div>

                <Field data-invalid={!!errors.email}>
                  <FieldLabel htmlFor="email">Email</FieldLabel>
                  <Input
                    {...fieldProps("email")}
                    type="email"
                    inputMode="email"
                    autoComplete="email"
                    placeholder="you@example.com"
                  />
                  {fieldError("email")}
                </Field>

                <Field data-invalid={!!errors.password}>
                  <FieldLabel htmlFor="password">Password</FieldLabel>
                  <InputGroup>
                    <InputGroupInput
                      {...fieldProps("password")}
                      type={showPassword ? "text" : "password"}
                      autoComplete="new-password"
                    />
                    <InputGroupAddon align="inline-end">
                      <InputGroupButton
                        type="button"
                        size="icon-xs"
                        aria-label={
                          showPassword ? "Hide passwords" : "Show passwords"
                        }
                        aria-pressed={showPassword}
                        onClick={() => setShowPassword(!showPassword)}
                      >
                        {showPassword ? <EyeOffIcon /> : <EyeIcon />}
                      </InputGroupButton>
                    </InputGroupAddon>
                  </InputGroup>
                  {errors.password ? (
                    fieldError("password")
                  ) : (
                    <FieldDescription>At least 8 characters.</FieldDescription>
                  )}
                </Field>

                <Field data-invalid={!!errors.confirmPassword}>
                  <FieldLabel htmlFor="confirmPassword">
                    Confirm password
                  </FieldLabel>
                  <InputGroup>
                    <InputGroupInput
                      {...fieldProps("confirmPassword")}
                      type={showPassword ? "text" : "password"}
                      autoComplete="new-password"
                    />
                  </InputGroup>
                  {fieldError("confirmPassword")}
                </Field>

                <Button size="lg" type="submit" disabled={isLoading}>
                  {isLoading ? <Spinner /> : "Create account"}
                </Button>

                <FieldDescription className="text-center">
                  Already have an account? <Link href="/login">Log in</Link>
                </FieldDescription>
              </FieldGroup>
            </FieldSet>
          </form>
        </CardContent>
      </Card>
    </main>
  )
}
