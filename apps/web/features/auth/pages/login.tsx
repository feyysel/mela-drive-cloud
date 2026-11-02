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
import Link from "next/link"
import { useRouter } from "next/navigation"
import { SubmitEvent, useState } from "react"
import { useLoginMutation } from "../api"
import type { LoginRequest } from "../types"

export function Login() {
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
          <CardTitle>Welcome back</CardTitle>
          <CardDescription>Log in to your Mela Drive account.</CardDescription>
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
                    <FieldLabel htmlFor="password">Password</FieldLabel>
                    <Button
                      nativeButton={false}
                      variant="link"
                      render={<Link href="/forgot-password" />}
                      className="text-muted-foreground"
                    >
                      Forgot password?
                    </Button>
                  </div>
                  <InputGroup>
                    <InputGroupInput
                      id="password"
                      name="password"
                      type={showPassword ? "text" : "password"}
                      autoComplete="current-password"
                      placeholder="Enter your password"
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
                </Field>

                <Button
                  size="lg"
                  type="submit"
                  disabled={isLoading || !canSubmit}
                >
                  {isLoading ? <Spinner /> : "Log in"}
                </Button>

                <FieldDescription className="text-center">
                  Don&apos;t have an account?{" "}
                  <Link href="/signup">Sign up</Link>
                </FieldDescription>
              </FieldGroup>
            </FieldSet>
          </form>
        </CardContent>
      </Card>
    </main>
  )
}
