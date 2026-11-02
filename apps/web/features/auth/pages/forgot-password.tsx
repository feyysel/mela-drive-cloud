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
import { useTranslations } from "next-intl"
import { type SubmitEvent, useState } from "react"
import { useForgotPasswordMutation } from "../api"

export function ForgotPassword() {
  const t = useTranslations("auth.forgotPassword")
  const auth = useTranslations("auth.login")
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
          <EmptyTitle>{t("successTitle")}</EmptyTitle>
          <EmptyDescription>{t("successDescription")}</EmptyDescription>
        </EmptyHeader>
      </Empty>
    )

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
                  <FieldLabel htmlFor="email">{auth("email")}</FieldLabel>
                  <Input
                    id="email"
                    name="email"
                    type="email"
                    inputMode="email"
                    autoComplete="email"
                    placeholder={auth("emailPlaceholder")}
                    value={email}
                    required
                    disabled={isLoading}
                    autoFocus
                    onChange={(e) => setEmail(e.target.value)}
                  />
                </Field>

                <Button size="lg" type="submit" disabled={isLoading || !email}>
                  {isLoading ? <Spinner /> : t("submit")}
                </Button>
              </FieldGroup>
            </FieldSet>
          </form>
        </CardContent>
      </Card>
    </main>
  )
}
