import { NextRequest, NextResponse } from "next/server"

const GUEST_ONLY_PATHS = new Set(["/login", "/signup", "/forgot-password"])

const PUBLIC_PATHS = new Set([...GUEST_ONLY_PATHS, "/reset-password"])

export default function proxy(req: NextRequest) {
  const token = req.cookies.get("__Secure-session")?.value
  const { pathname } = req.nextUrl

  if (!token && !PUBLIC_PATHS.has(pathname)) {
    const url = new URL("/login", req.url)

    return NextResponse.redirect(url)
  }

  if (token && GUEST_ONLY_PATHS.has(pathname)) {
    return NextResponse.redirect(new URL("/", req.url))
  }

  return NextResponse.next()
}

export const config = {
  matcher: [
    "/((?!_next/static|_next/image|icons|favicon|manifest.webmanifest|.*\\.(?:png|svg|jpg|ico|webp)$).*)",
  ],
}
