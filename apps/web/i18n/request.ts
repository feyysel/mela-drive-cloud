import { getRequestConfig } from "next-intl/server"
import { cookies } from "next/headers"
import { defaultLocale, LOCALE_COOKIE, locales, type Locale } from "./config"

export default getRequestConfig(async () => {
  const stored = (await cookies()).get(LOCALE_COOKIE)?.value
  const locale = locales.includes(stored as Locale)
    ? (stored as Locale)
    : defaultLocale

  return {
    locale,
    messages: (await import(`./messages/${locale}.json`)).default,
  }
})
