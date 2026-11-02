"use client"

import { Badge } from "@/components/ui/badge"
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
import { SubmitEvent, useState } from "react"
import { useSignupMutation } from "../api"
import type { SignupRequest } from "../types"

export function Signup() {
  const router = useRouter()
  const [signup, { isLoading }] = useSignupMutation()
  const [form, setForm] = useState<SignupRequest>({
    firstName: "",
    lastName: "",
    email: "",
    password: "",
  })
  const [showPassword, setShowPassword] = useState(false)

  const canSubmit =
    !!form.firstName &&
    !!form.email &&
    !!form.password &&
    form.password.length >= 8

  const handleSubmit = async (e: SubmitEvent<HTMLFormElement>) => {
    e.preventDefault()

    if (!canSubmit) return

    try {
      await signup({
        ...form,
        lastName: form.lastName || null,
      }).unwrap()
      router.replace("/")
    } catch (error: unknown) {
      toast.add({ type: "error", title: parseErr(error).message })
    }
  }

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
          <form onSubmit={handleSubmit}>
            <FieldSet>
              <FieldGroup>
                <Field orientation="horizontal">
                  <Field>
                    <FieldLabel htmlFor="firstName">First name</FieldLabel>
                    <Input
                      id="firstName"
                      name="firstName"
                      autoComplete="given-name"
                      placeholder="e.g. Abebe"
                      value={form.firstName}
                      required
                      disabled={isLoading}
                      autoFocus
                      onChange={(e) =>
                        setForm({ ...form, firstName: e.target.value })
                      }
                    />
                  </Field>

                  <Field>
                    <FieldLabel htmlFor="lastName">
                      Last name
                      <Badge variant="outline">optional</Badge>
                    </FieldLabel>
                    <Input
                      id="lastName"
                      name="lastName"
                      autoComplete="family-name"
                      placeholder="e.g. Kebede"
                      value={form.lastName ?? ""}
                      disabled={isLoading}
                      onChange={(e) =>
                        setForm({ ...form, lastName: e.target.value })
                      }
                    />
                  </Field>
                </Field>

                <Field>
                  <FieldLabel htmlFor="email">Email</FieldLabel>
                  <Input
                    id="email"
                    name="email"
                    type="email"
                    inputMode="email"
                    autoComplete="email"
                    placeholder="e.g. you@example.com"
                    value={form.email}
                    required
                    disabled={isLoading}
                    onChange={(e) =>
                      setForm({ ...form, email: e.target.value })
                    }
                  />
                </Field>

                <Field>
                  <FieldLabel htmlFor="password">Password</FieldLabel>
                  <InputGroup>
                    <InputGroupInput
                      id="password"
                      name="password"
                      type={showPassword ? "text" : "password"}
                      autoComplete="new-password"
                      placeholder="Create a password"
                      value={form.password}
                      required
                      disabled={isLoading}
                      onChange={(e) =>
                        setForm({ ...form, password: e.target.value })
                      }
                    />
                    <InputGroupAddon align="inline-end">
                      <InputGroupButton
                        type="button"
                        size="icon-xs"
                        aria-label={
                          showPassword ? "Hide password" : "Show password"
                        }
                        aria-pressed={showPassword}
                        onClick={() => setShowPassword(!showPassword)}
                      >
                        {showPassword ? <EyeOffIcon /> : <EyeIcon />}
                      </InputGroupButton>
                    </InputGroupAddon>
                  </InputGroup>
                  <FieldDescription>At least 8 characters.</FieldDescription>
                </Field>

                <Button
                  size="lg"
                  type="submit"
                  disabled={isLoading || !canSubmit}
                >
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
