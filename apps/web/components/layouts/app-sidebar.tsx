"use client"

import {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  useSidebar,
} from "@/components/ui/sidebar"
import { useLogoutMutation } from "@/features/auth/api"
import { useMeQuery } from "@/features/user/api"
import { useIsMobile } from "@/hooks/use-mobile"
import { parseErr } from "@/store/utils"
import CloudOutlineIcon from "@iconify-react/solar/cloud-outline"
import Folder2OutlineIcon from "@iconify-react/solar/folder-2-outline"
import TrashBinTrashOutlineIcon from "@iconify-react/solar/trash-bin-trash-outline"
import { filesize } from "filesize"
import { LogOutIcon } from "lucide-react"
import { useTranslations } from "next-intl"
import Link from "next/link"
import { usePathname, useRouter } from "next/navigation"
import { parseAsStringLiteral, useQueryState } from "nuqs"
import { ComponentType } from "react"
import { Logo } from "../ui/logo"
import { Progress } from "../ui/progress"
import { Skeleton } from "../ui/skeleton"
import { Spinner } from "../ui/spinner"
import { toast } from "../ui/toast"

type NavLink = {
  title: string
  url: string
  icon: ComponentType
}

const views = ["grid", "list"] as const

export function AppSidebar() {
  const router = useRouter()
  const pathname = usePathname()
  const t = useTranslations()
  const [view] = useQueryState(
    "view",
    parseAsStringLiteral(views).withDefault("grid")
  )
  const { data, isLoading } = useMeQuery()
  const { toggleSidebar } = useSidebar()
  const isMobile = useIsMobile()
  const [logout, { isLoading: isLoggingOut }] = useLogoutMutation()

  const navLinks: NavLink[] = [
    {
      title: t("navigation.myDrive"),
      url: view ? `/drive?view=${view}` : "/drive",
      icon: Folder2OutlineIcon,
    },
    {
      title: t("navigation.trash"),
      url: view ? `/trash?view=${view}` : "/trash",
      icon: TrashBinTrashOutlineIcon,
    },
  ]

  const handleLogout = async () => {
    try {
      await logout().unwrap()
      router.replace("/login")
    } catch (error: unknown) {
      toast.add({ type: "error", title: parseErr(error).message })
    }
  }

  return (
    <Sidebar collapsible="icon">
      <SidebarHeader className="min-h-16 bg-background">
        <SidebarMenu className="my-auto">
          <SidebarMenuItem>
            <div className="flex items-center gap-2">
              <Logo className="size-5 shrink-0" />
              <h1 className="text-lg font-medium whitespace-nowrap">
                {t("site.name")}
              </h1>
            </div>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarHeader>

      <SidebarContent className="bg-background">
        <SidebarGroup>
          <SidebarMenu>
            {navLinks.map((navLink) => (
              <SidebarMenuItem key={navLink.title}>
                <SidebarMenuButton
                  isActive={
                    pathname === navLink.url.split("?")[0] ||
                    pathname.startsWith(`${navLink.url.split("?")[0]}/`)
                  }
                  render={
                    <Link
                      href={navLink.url}
                      onClick={isMobile ? toggleSidebar : undefined}
                    />
                  }
                >
                  <navLink.icon />
                  <span className="text-sm">{navLink.title}</span>
                </SidebarMenuButton>
              </SidebarMenuItem>
            ))}
          </SidebarMenu>
        </SidebarGroup>
      </SidebarContent>

      <SidebarFooter className="bg-background">
        {isLoading ? (
          <div className="space-y-2 p-2 text-sm">
            <div className="flex items-center gap-2">
              <Skeleton className="size-5 rounded-full" />
              <Skeleton className="h-5 w-24" />
            </div>
            <Skeleton className="h-2" />
            <Skeleton className="h-5 w-50" />
          </div>
        ) : data ? (
          <div className="space-y-2 p-2 text-sm">
            <div className="flex items-center gap-2 text-muted-foreground">
              <CloudOutlineIcon className="size-5" />
              <span>{t("navigation.storage")}</span>
            </div>

            <Progress
              value={(data.storageUsedBytes / data.storageQuotaBytes) * 100}
            />

            <p className="text-muted-foreground">
              {t("navigation.storageUsed", {
                used: filesize(data.storageUsedBytes),
                quota: filesize(data.storageQuotaBytes),
              })}
            </p>
          </div>
        ) : null}
        <SidebarMenu className="pb-3">
          <SidebarMenuItem>
            <SidebarMenuButton onClick={handleLogout}>
              {isLoggingOut ? <Spinner /> : <LogOutIcon />}
              <span>{t("common.logout")}</span>
            </SidebarMenuButton>
          </SidebarMenuItem>
        </SidebarMenu>
      </SidebarFooter>
    </Sidebar>
  )
}
