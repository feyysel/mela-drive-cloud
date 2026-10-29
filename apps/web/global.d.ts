import type { locales } from "@/i18n/config"
import en from "./i18n/messages/en.json"

declare module "next-intl" {
  interface AppConfig {
    Locale: (typeof locales)[number]
    Messages: typeof en
  }
}
