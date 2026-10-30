import { Geist_Mono, Inter, Noto_Sans_Ethiopic } from "next/font/google"

import { ThemeProvider } from "@/components/theme-provider"
import { cn } from "@/lib/utils"
import { NextIntlClientProvider } from "next-intl"
import { getLocale } from "next-intl/server"
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
          <NuqsAdapter>
            <ThemeProvider>{children}</ThemeProvider>
          </NuqsAdapter>
        </NextIntlClientProvider>
      </body>
    </html>
  )
}
