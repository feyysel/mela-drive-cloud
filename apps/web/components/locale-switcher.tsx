"use client"

import { useLocale } from "next-intl"
import { useRouter } from "next/navigation"
import { useTransition } from "react"

import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import { setLocale } from "@/i18n/actions"
import type { Locale } from "@/i18n/config"

const items: { label: string; value: Locale }[] = [
  { label: "English", value: "en" },
  { label: "አማርኛ", value: "am" },
]

export function LocaleSwitcher() {
  const locale = useLocale()
  const router = useRouter()
  const [pending, startTransition] = useTransition()

  function change(next: Locale) {
    startTransition(async () => {
      await setLocale(next)
      router.refresh()
    })
  }

  return (
    <Select
      items={items}
      value={locale}
      disabled={pending}
      onValueChange={(value) => {
        if (value) change(value as Locale)
      }}
    >
      <SelectTrigger className="w-36">
        <SelectValue />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          {items.map((item) => (
            <SelectItem key={item.value} value={item.value}>
              {item.label}
            </SelectItem>
          ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  )
}
