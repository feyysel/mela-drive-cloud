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
import { useTranslations } from "next-intl"
import Link from "next/link"
import { useRouter } from "next/navigation"
import { type SubmitEvent, useState } from "react"
import { useSignupMutation } from "../api"
import type { SignupRequest } from "../types"

export function Signup() {
  const t = useTranslations("auth.signup")
  const auth = useTranslations("auth.login")
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
          <CardTitle>{t("title")}</CardTitle>
          <CardDescription>{t("description")}</CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit}>
            <FieldSet>
              <FieldGroup>
                <Field orientation="horizontal">
                  <Field>
                    <FieldLabel htmlFor="firstName">
                      {t("firstName")}
                    </FieldLabel>
                    <Input
                      id="firstName"
                      name="firstName"
                      autoComplete="given-name"
                      placeholder={t("firstNamePlaceholder")}
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
                      {t("lastName")}
                      <Badge variant="outline">{t("optional")}</Badge>
                    </FieldLabel>
                    <Input
                      id="lastName"
                      name="lastName"
                      autoComplete="family-name"
                      placeholder={t("lastNamePlaceholder")}
                      value={form.lastName ?? ""}
                      disabled={isLoading}
                      onChange={(e) =>
                        setForm({ ...form, lastName: e.target.value })
                      }
                    />
                  </Field>
                </Field>

                <Field>
                  <FieldLabel htmlFor="email">{auth("email")}</FieldLabel>
                  <Input
                    id="email"
                    name="email"
                    type="email"
                    inputMode="email"
                    autoComplete="email"
                    placeholder={auth("emailPlaceholder")}
                    value={form.email}
                    required
                    disabled={isLoading}
                    onChange={(e) =>
                      setForm({ ...form, email: e.target.value })
                    }
                  />
                </Field>

                <Field>
                  <FieldLabel htmlFor="password">{auth("password")}</FieldLabel>
                  <InputGroup>
                    <InputGroupInput
                      id="password"
                      name="password"
                      type={showPassword ? "text" : "password"}
                      autoComplete="new-password"
                      placeholder={t("passwordPlaceholder")}
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
                          showPassword
                            ? auth("hidePassword")
                            : auth("showPassword")
                        }
                        aria-pressed={showPassword}
                        onClick={() => setShowPassword(!showPassword)}
                      >
                        {showPassword ? <EyeOffIcon /> : <EyeIcon />}
                      </InputGroupButton>
                    </InputGroupAddon>
                  </InputGroup>
                  <FieldDescription>{t("passwordHint")}</FieldDescription>
                </Field>

                <Button
                  size="lg"
                  type="submit"
                  disabled={isLoading || !canSubmit}
                >
                  {isLoading ? <Spinner /> : t("submit")}
                </Button>

                <FieldDescription className="text-center">
                  {t("alreadyHaveAccount")}{" "}
                  <Link href="/login">{t("login")}</Link>
                </FieldDescription>
              </FieldGroup>
            </FieldSet>
          </form>
        </CardContent>
      </Card>
    </main>
  )
}
