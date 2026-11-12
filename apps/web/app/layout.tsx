import { Geist_Mono, Inter, Noto_Sans_Ethiopic } from "next/font/google"

import { ThemeProvider } from "@/components/theme-provider"
import { Toaster } from "@/components/ui/toast"
import { TooltipProvider } from "@/components/ui/tooltip"
import { cn } from "@/lib/utils"
import { StoreProvider } from "@/store/provider"
import { Metadata } from "next"
import { NextIntlClientProvider } from "next-intl"
import { getLocale, getTranslations } from "next-intl/server"
import { NuqsAdapter } from "nuqs/adapters/react"
import "./globals.css"

const inter = Inter({ subsets: ["latin"], variable: "--font-sans" })

const fontMono = Geist_Mono({
  subsets: ["latin"],
  variable: "--font-mono",
})

const fontEthiopic = Noto_Sans_Ethiopic({
  subsets: ["ethiopic"],
  variable: "--font-ethiopic",
})

export async function generateMetadata(): Promise<Metadata> {
  const t = await getTranslations("site")

  return {
    title: t("name"),
    description: t("description"),
    icons: {
      icon: [
        {
          url: "/favicon-dark.svg",
          media: "(prefers-color-scheme: light)",
          type: "image/svg+xml",
        },
        {
          url: "/favicon-light.svg",
          media: "(prefers-color-scheme: dark)",
          type: "image/svg+xml",
        },
      ],
    },
  }
}

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  const locale = await getLocale()

  return (
    <html
      lang={locale}
      suppressHydrationWarning
      className={cn(
        "antialiased",
        "font-sans",
        inter.variable,
        fontMono.variable,
        fontEthiopic.variable
      )}
    >
      <body>
        <NextIntlClientProvider>
          <ThemeProvider>
            <StoreProvider>
              <TooltipProvider>
                <NuqsAdapter>{children}</NuqsAdapter>
              </TooltipProvider>
            </StoreProvider>
          </ThemeProvider>
        </NextIntlClientProvider>
        <Toaster />
      </body>
    </html>
  )
}
