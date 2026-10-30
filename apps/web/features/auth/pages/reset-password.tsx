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
import { Field, FieldGroup, FieldSet } from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { Spinner } from "@/components/ui/spinner"
import { toast } from "@/components/ui/toast"
import { parseErr } from "@/store/utils"
import { AlertCircleIcon } from "lucide-react"
import { useRouter } from "next/navigation"
import { useQueryState } from "nuqs"
import { SubmitEvent, useEffect, useState } from "react"
import { useResetPasswordMutation, useValidateResetTokenMutation } from "../api"

export function ResetPassword() {
  const router = useRouter()
  const [token] = useQueryState("token")
  const [validateToken, { isSuccess, isError }] =
    useValidateResetTokenMutation()
  const [resetPassword, { isLoading: isResetting, isSuccess: isResetSuccess }] =
    useResetPasswordMutation()
  const [password, setPassword] = useState("")
  const [confirmPassword, setConfirmPassword] = useState("")

  const mismatch = confirmPassword.length > 0 && password !== confirmPassword

  useEffect(() => {
    if (!token) return

    validateToken({ token })
  }, [token, validateToken])

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

  if (!token || isError)
    return (
      <Empty className="h-dvh">
        <EmptyHeader>
          <EmptyMedia>
            <AlertCircleIcon className="size-8 text-muted-foreground" />
          </EmptyMedia>
          <EmptyTitle>Invalid or expired link</EmptyTitle>
          <EmptyDescription>
            This password reset link is invalid or has expired.
          </EmptyDescription>
        </EmptyHeader>
      </Empty>
    )

  if (!isSuccess)
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
            Redirecting you to login
          </EmptyTitle>
          <EmptyDescription>Please wait</EmptyDescription>
        </EmptyHeader>
      </Empty>
    )

  return (
    <main className="flex h-dvh flex-col">
      <Card className="m-auto w-full max-w-sm">
        <CardHeader>
          <CardTitle>Set a new password</CardTitle>
          <CardDescription>
            Choose a new password for your account.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit}>
            <FieldSet>
              <FieldGroup>
                <Field>
                  <Input
                    id="password"
                    name="password"
                    type="password"
                    autoComplete="new-password"
                    placeholder="New password"
                    value={password}
                    required
                    disabled={isResetting}
                    onChange={(e) => setPassword(e.target.value)}
                  />
                </Field>

                <Field>
                  <Input
                    id="confirm-password"
                    name="confirmPassword"
                    type="password"
                    autoComplete="new-password"
                    placeholder="Confirm new password"
                    value={confirmPassword}
                    required
                    disabled={isResetting}
                    aria-invalid={mismatch}
                    aria-describedby={
                      mismatch ? "confirm-password-error" : undefined
                    }
                    onChange={(e) => setConfirmPassword(e.target.value)}
                  />
                </Field>

                {mismatch && (
                  <p
                    id="confirm-password-error"
                    className="text-sm text-destructive"
                  >
                    Passwords do not match
                  </p>
                )}

                <Button
                  size="lg"
                  type="submit"
                  disabled={
                    isResetting || mismatch || !password || !confirmPassword
                  }
                >
                  {isResetting ? <Spinner /> : "Reset password"}
                </Button>
              </FieldGroup>
            </FieldSet>
          </form>
        </CardContent>
      </Card>
    </main>
  )
}
