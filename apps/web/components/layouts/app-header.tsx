"use client"

import { HamburgerIcon } from "lucide-react"
import { Button } from "../ui/button"
import { useSidebar } from "../ui/sidebar"
import { Header } from "./header"

export function AppHeader() {
  const { toggleSidebar } = useSidebar()

  return (
    <Header className="border-b bg-background px-2">
      <Button
        variant="ghost"
        className="md:invisible"
        aria-label="open menu"
        onClick={toggleSidebar}
      >
        <HamburgerIcon className="size-6" />
      </Button>
    </Header>
  )
}
