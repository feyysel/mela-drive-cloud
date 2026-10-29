import type { ComponentProps } from "react"

import { cn } from "cn"

export function Header({ className, ...props }: ComponentProps<"header">) {
  return (
    <header
      className={cn(
        "sticky top-0 z-40 flex h-16 w-full shrink-0 items-center justify-between px-5",
        className
      )}
      {...props}
    />
  )
} 