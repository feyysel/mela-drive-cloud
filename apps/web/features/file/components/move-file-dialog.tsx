import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog"
import {
  Breadcrumb,
  BreadcrumbEllipsis,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb"
import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { Spinner } from "@/components/ui/spinner"
import { toast } from "@/components/ui/toast"
import { useMoveFileMutation } from "@/features/file/api"
import { DriveFile } from "@/features/file/types"
import { useListFoldersQuery } from "@/features/folder/api"
import { Folder as FolderType } from "@/features/folder/types"
import { parseErr } from "@/store/utils"
import FolderOpenBoldIcon from "@iconify-react/solar/folder-open-bold"
import { useTranslations } from "next-intl"
import { Fragment, useState } from "react"

export function MoveFileDialog({
  open,
  onOpenChange,
  file,
}: {
  open: boolean
  onOpenChange: (open: boolean) => void
  file: DriveFile
}) {
  const t = useTranslations()
  const [stack, setStack] = useState<FolderType[]>([])
  const [moveFile, { isLoading: isMoving }] = useMoveFileMutation()

  const current = stack.at(-1)
  const { data, isLoading } = useListFoldersQuery(current?.id, { skip: !open })

  const folders = data ?? []
  const isSameFolder = (current?.id ?? null) === file.folderId

  const crumbs = [{ name: t("navigation.myDrive") }, ...stack]
  const last = crumbs.length - 1
  const isCollapsed = crumbs.length > 3

  const goTo = (index: number) => setStack((s) => s.slice(0, index))

  function handleOpenChange(next: boolean) {
    if (!next) setStack([])
    onOpenChange(next)
  }

  const handleMove = async () => {
    try {
      await moveFile({
        fileId: file.id,
        folderId: current?.id,
      }).unwrap()
      handleOpenChange(false)
    } catch (error: unknown) {
      toast.add({ type: "error", title: parseErr(error).message })
    }
  }

  return (
    <AlertDialog open={open} onOpenChange={handleOpenChange}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>{t("common.move")}</AlertDialogTitle>
        </AlertDialogHeader>

        <Breadcrumb>
          <BreadcrumbList>
            {crumbs.map((crumb, i) => {
              if (isCollapsed && i > 0 && i < last) {
                if (i !== 1) return null

                return (
                  <Fragment key="ellipsis">
                    <BreadcrumbSeparator />
                    <BreadcrumbItem>
                      <DropdownMenu>
                        <DropdownMenuTrigger
                          render={
                            <Button
                              size="icon-sm"
                              variant="ghost"
                              disabled={isMoving}
                            >
                              <BreadcrumbEllipsis />
                              <span className="sr-only">
                                {t("common.moreActions")}
                              </span>
                            </Button>
                          }
                        />
                        <DropdownMenuContent align="start">
                          <DropdownMenuGroup>
                            {crumbs.slice(1, last).map((hidden, j) => (
                              <DropdownMenuItem
                                key={j}
                                onClick={() => goTo(j + 1)}
                              >
                                {hidden.name}
                              </DropdownMenuItem>
                            ))}
                          </DropdownMenuGroup>
                        </DropdownMenuContent>
                      </DropdownMenu>
                    </BreadcrumbItem>
                  </Fragment>
                )
              }

              return (
                <Fragment key={i}>
                  {i > 0 && <BreadcrumbSeparator />}
                  <BreadcrumbItem>
                    {i === last ? (
                      <BreadcrumbPage className="max-w-40 truncate">
                        {crumb.name}
                      </BreadcrumbPage>
                    ) : (
                      <BreadcrumbLink
                        render={
                          <button
                            type="button"
                            disabled={isMoving}
                            onClick={() => goTo(i)}
                            className="max-w-40 truncate"
                          />
                        }
                      >
                        {crumb.name}
                      </BreadcrumbLink>
                    )}
                  </BreadcrumbItem>
                </Fragment>
              )
            })}
          </BreadcrumbList>
        </Breadcrumb>

        <div className="h-64 overflow-y-auto rounded-md border">
          {isLoading ? (
            <div className="flex h-full items-center justify-center">
              <Spinner className="size-6" />
            </div>
          ) : folders.length === 0 ? (
            <div className="flex h-full items-center justify-center text-sm text-muted-foreground">
              {t("drive.empty.title")}
            </div>
          ) : (
            folders.map((f) => (
              <button
                key={f.id}
                type="button"
                disabled={isMoving}
                className="flex w-full items-center gap-3 px-3 py-2 text-left hover:bg-muted/50"
                onClick={() => setStack((s) => [...s, f])}
              >
                <FolderOpenBoldIcon className="size-5 shrink-0" />
                <span className="truncate">{f.name}</span>
              </button>
            ))
          )}
        </div>

        <AlertDialogFooter>
          <AlertDialogCancel disabled={isMoving}>
            {t("common.cancel")}
          </AlertDialogCancel>
          <AlertDialogAction
            disabled={isSameFolder || isMoving}
            onClick={handleMove}
            className="md:w-24!"
          >
            {isMoving ? <Spinner /> : t("common.move")}
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}
