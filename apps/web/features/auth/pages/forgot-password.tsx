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
import { Field, FieldGroup, FieldLabel, FieldSet } from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { Spinner } from "@/components/ui/spinner"
import { toast } from "@/components/ui/toast"
import { parseErr } from "@/store/utils"
import { MailOpenIcon } from "lucide-react"
import { SubmitEvent, useState } from "react"
import { useForgotPasswordMutation } from "../api"

export function ForgotPassword() {
  const [forgotPassword, { isLoading, isSuccess }] = useForgotPasswordMutation()
  const [email, setEmail] = useState("")

  const handleSubmit = async (e: SubmitEvent<HTMLFormElement>) => {
    e.preventDefault()

    if (!email) return

    try {
      await forgotPassword({ email }).unwrap()
    } catch (error: unknown) {
      toast.add({ type: "error", title: parseErr(error).message })
    }
  }

  if (isSuccess)
    return (
      <Empty className="h-dvh">
        <EmptyHeader>
          <EmptyMedia>
            <MailOpenIcon className="size-8 text-muted-foreground" />
          </EmptyMedia>
          <EmptyTitle>Check your email</EmptyTitle>
          <EmptyDescription>
            We&apos;ve sent a password reset email to your inbox.
          </EmptyDescription>
        </EmptyHeader>
      </Empty>
    )

  return (
    <main className="flex min-h-dvh flex-col px-4 py-10">
      <Card className="m-auto w-full max-w-md">
        <CardHeader>
          <CardTitle>Forgot your password?</CardTitle>
          <CardDescription>
            Enter your email to reset your password.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit}>
            <FieldSet>
              <FieldGroup>
                <Field>
                  <FieldLabel htmlFor="email">Email</FieldLabel>
                  <Input
                    id="email"
                    name="email"
                    type="email"
                    inputMode="email"
                    autoComplete="email"
                    placeholder="e.g. you@example.com"
                    value={email}
                    required
                    disabled={isLoading}
                    autoFocus
                    onChange={(e) => setEmail(e.target.value)}
                  />
                </Field>

                <Button size="lg" type="submit" disabled={isLoading || !email}>
                  {isLoading ? <Spinner /> : "Send reset email"}
                </Button>
              </FieldGroup>
            </FieldSet>
          </form>
        </CardContent>
      </Card>
    </main>
  )
}
