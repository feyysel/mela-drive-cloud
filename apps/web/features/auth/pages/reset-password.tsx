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
  Empty,
  EmptyDescription,
  EmptyHeader,
  EmptyMedia,
  EmptyTitle,
} from "@/components/ui/empty"
import { Field, FieldError, FieldGroup, FieldSet } from "@/components/ui/field"
import {
  InputGroup,
  InputGroupAddon,
  InputGroupButton,
  InputGroupInput,
} from "@/components/ui/input-group"
import { Spinner } from "@/components/ui/spinner"
import { toast } from "@/components/ui/toast"
import { parseErr } from "@/store/utils"
import { skipToken } from "@reduxjs/toolkit/query"
import { AlertCircleIcon, EyeIcon, EyeOffIcon } from "lucide-react"
import { useTranslations } from "next-intl"
import { useRouter } from "next/navigation"
import { useQueryState } from "nuqs"
import { type SubmitEvent, useState } from "react"
import { useResetPasswordMutation, useValidateResetTokenQuery } from "../api"

export function ResetPassword() {
  const t = useTranslations("auth.resetPassword")
  const router = useRouter()
  const [token] = useQueryState("token")
  const { isError, isLoading } = useValidateResetTokenQuery(
    token ?? skipToken,
    { refetchOnMountOrArgChange: true }
  )
  const [resetPassword, { isLoading: isResetting, isSuccess: isResetSuccess }] =
    useResetPasswordMutation()
  const [password, setPassword] = useState("")
  const [confirmPassword, setConfirmPassword] = useState("")
  const [showPassword, setShowPassword] = useState(false)

  const mismatch = confirmPassword.length > 0 && password !== confirmPassword

  const handleSubmit = async (e: SubmitEvent<HTMLFormElement>) => {
    e.preventDefault()

    if (!token || password !== confirmPassword) return

    try {
      await resetPassword({ token, password }).unwrap()
      router.replace("/login")
    } catch (error: unknown) {
      toast.add({
        type: "error",
        title: parseErr(error).message,
      })
    }
  }

  if (isLoading)
    return (
      <Empty className="h-dvh">
        <Spinner className="size-6" />
      </Empty>
    )

  if (isResetSuccess)
    return (
      <Empty className="h-dvh">
        <EmptyHeader>
          <EmptyTitle className="shimmer text-muted-foreground shimmer-duration-1300 shimmer-spread-5">
            {t("redirecting")}
          </EmptyTitle>
          <EmptyDescription>{t("pleaseWait")}</EmptyDescription>
        </EmptyHeader>
      </Empty>
    )

  if (!token || isError)
    return (
      <Empty className="h-dvh">
        <EmptyHeader>
          <EmptyMedia>
            <AlertCircleIcon className="size-8 text-muted-foreground" />
          </EmptyMedia>
          <EmptyTitle>{t("invalidTitle")}</EmptyTitle>
          <EmptyDescription>{t("invalidDescription")}</EmptyDescription>
        </EmptyHeader>
      </Empty>
    )

  return (
    <main className="flex h-dvh flex-col">
      <Card className="m-auto w-full max-w-sm">
        <CardHeader>
          <CardTitle>{t("title")}</CardTitle>
          <CardDescription>{t("description")}</CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit}>
            <FieldSet>
              <FieldGroup>
                <Field>
                  <InputGroup>
                    <InputGroupInput
                      id="password"
                      name="password"
                      type={showPassword ? "text" : "password"}
                      autoComplete="new-password"
                      placeholder={t("passwordPlaceholder")}
                      value={password}
                      required
                      disabled={isResetting}
                      onChange={(e) => setPassword(e.target.value)}
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
                </Field>

                <Field data-invalid={mismatch}>
                  <InputGroup>
                    <InputGroupInput
                      id="confirm-password"
                      name="confirmPassword"
                      type={showPassword ? "text" : "password"}
                      autoComplete="new-password"
                      placeholder={t("confirmPasswordPlaceholder")}
                      value={confirmPassword}
                      required
                      disabled={isResetting}
                      aria-invalid={mismatch}
                      aria-describedby={
                        mismatch ? "confirm-password-error" : undefined
                      }
                      onChange={(e) => setConfirmPassword(e.target.value)}
                    />
                  </InputGroup>
                  {mismatch && (
                    <FieldError id="confirm-password-error">
                      {t("passwordMismatch")}
                    </FieldError>
                  )}
                </Field>

                <Button
                  size="lg"
                  type="submit"
                  disabled={
                    isResetting || mismatch || !password || !confirmPassword
                  }
                >
                  {isResetting ? <Spinner /> : t("submit")}
                </Button>
              </FieldGroup>
            </FieldSet>
          </form>
        </CardContent>
      </Card>
    </main>
  )
}
