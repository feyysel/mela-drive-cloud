"use client"

import { MenuIcon } from "lucide-react"
import { LocaleSwitcher } from "../locale-switcher"
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
        <MenuIcon />
      </Button>

      <LocaleSwitcher />
    </Header>
  )
}
