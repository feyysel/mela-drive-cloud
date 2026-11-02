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
import { useLoginMutation } from "../api"
import type { LoginRequest } from "../types"

export function Login() {
  const t = useTranslations("auth.login")
  const router = useRouter()
  const [login, { isLoading }] = useLoginMutation()
  const [form, setForm] = useState<LoginRequest>({
    email: "",
    password: "",
  })
  const [showPassword, setShowPassword] = useState(false)

  const canSubmit = !!form.email && !!form.password && form.password.length >= 8

  const handleSubmit = async (e: SubmitEvent<HTMLFormElement>) => {
    e.preventDefault()

    if (!canSubmit) return

    try {
      await login(form).unwrap()
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
                <Field>
                  <FieldLabel htmlFor="email">{t("email")}</FieldLabel>
                  <Input
                    id="email"
                    name="email"
                    type="email"
                    inputMode="email"
                    autoComplete="email"
                    placeholder={t("emailPlaceholder")}
                    value={form.email}
                    required
                    disabled={isLoading}
                    autoFocus
                    onChange={(e) =>
                      setForm({ ...form, email: e.target.value })
                    }
                  />
                </Field>

                <Field>
                  <div className="flex items-center justify-between">
                    <FieldLabel htmlFor="password">{t("password")}</FieldLabel>
                    <Button
                      nativeButton={false}
                      variant="link"
                      render={<Link href="/forgot-password" />}
                      className="text-muted-foreground"
                    >
                      {t("forgotPassword")}
                    </Button>
                  </div>
                  <InputGroup>
                    <InputGroupInput
                      id="password"
                      name="password"
                      type={showPassword ? "text" : "password"}
                      autoComplete="current-password"
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
                          showPassword ? t("hidePassword") : t("showPassword")
                        }
                        aria-pressed={showPassword}
                        onClick={() => setShowPassword(!showPassword)}
                      >
                        {showPassword ? <EyeOffIcon /> : <EyeIcon />}
                      </InputGroupButton>
                    </InputGroupAddon>
                  </InputGroup>
                </Field>

                <Button
                  size="lg"
                  type="submit"
                  disabled={isLoading || !canSubmit}
                >
                  {isLoading ? <Spinner /> : t("submit")}
                </Button>

                <FieldDescription className="text-center">
                  {t("noAccount")} <Link href="/signup">{t("signup")}</Link>
                </FieldDescription>
              </FieldGroup>
            </FieldSet>
          </form>
        </CardContent>
      </Card>
    </main>
  )
}
