import {
  Attachment,
  AttachmentAction,
  AttachmentActions,
  AttachmentContent,
  AttachmentMedia,
  AttachmentTitle,
} from "@/components/ui/attachment"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { Spinner } from "@/components/ui/spinner"
import { toast } from "@/components/ui/toast"
import { parseErr } from "@/store/utils"
import FolderOpenBoldIcon from "@iconify-react/solar/folder-open-bold"
import InfoCircleOutlineIcon from "@iconify-react/solar/info-circle-outline"
import MoveToFolderOutlineIcon from "@iconify-react/solar/move-to-folder-outline"
import PenNewSquareOutlineIcon from "@iconify-react/solar/pen-new-square-outline"
import RefreshCircleOutlineIcon from "@iconify-react/solar/refresh-circle-outline"
import TrashBinTrashOutlineIcon from "@iconify-react/solar/trash-bin-trash-outline"
import { cn } from "cn"
import { MoreVerticalIcon } from "lucide-react"
import { useTranslations } from "next-intl"
import { useRouter } from "next/navigation"
import { useState } from "react"
import { usePurgeFolderMutation, useRestoreFolderMutation } from "../api"
import { Folder as FolderType } from "../types"
import { FolderDetailsDialog } from "./folder-details-dialog"
import { MoveFolderDialog } from "./move-folder-dialog"
import { RenameFolderDialog } from "./rename-folder-dialog"
import { TrashFolderDialog } from "./trash-folder-dialog"

export function Folder({
  folder,
  index,
  view,
}: {
  folder: FolderType
  index: number
  view: "list" | "grid"
}) {
  const router = useRouter()
  const t = useTranslations()
  const [isRenameOpen, setIsRenameOpen] = useState(false)
  const [isMoveOpen, setIsMoveOpen] = useState(false)
  const [isDetailsOpen, setIsDetailsOpen] = useState(false)
  const [isTrashOpen, setIsTrashOpen] = useState(false)

  const [restoreFolder, { isLoading: isRestoringFolder }] =
    useRestoreFolderMutation()
  const [purgeFolder, { isLoading: isPurgingFolder }] = usePurgeFolderMutation()

  const handleRestore = async () => {
    try {
      await restoreFolder(folder.id).unwrap()
    } catch (error: unknown) {
      toast.add({ type: "error", title: parseErr(error).message })
    }
  }

  const handlePurge = async () => {
    try {
      await purgeFolder(folder.id).unwrap()
    } catch (error: unknown) {
      toast.add({ type: "error", title: parseErr(error).message })
    }
  }

  return (
    <>
      <Attachment
        className={cn(
          "w-full cursor-pointer",
          view === "list"
            ? "rounded-none border-0 bg-transparent hover:bg-muted/50"
            : "bg-muted/50 hover:bg-muted",
          view === "list" && index > 0 && "border-t"
        )}
        onClick={() =>
          router.push(
            view === "list"
              ? `/drive/folders/${folder.id}?view=list`
              : `/drive/folders/${folder.id}`
          )
        }
      >
        <AttachmentMedia className="bg-transparent">
          {isRestoringFolder || isPurgingFolder ? (
            <Spinner className="size-6" />
          ) : (
            <FolderOpenBoldIcon className="size-6" />
          )}
        </AttachmentMedia>
        <AttachmentContent>
          <AttachmentTitle>{folder.name}</AttachmentTitle>
        </AttachmentContent>
        <AttachmentActions onClick={(e) => e.stopPropagation()}>
          <DropdownMenu>
            <DropdownMenuTrigger
              render={
                <AttachmentAction
                  size="icon-lg"
                  aria-label={t("common.moreActions")}
                />
              }
            >
              <MoreVerticalIcon />
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
              {folder.deletedAt ? (
                <>
                  <DropdownMenuItem onClick={handleRestore}>
                    <RefreshCircleOutlineIcon />
                    {t("common.restore")}
                  </DropdownMenuItem>
                  <DropdownMenuItem variant="destructive" onClick={handlePurge}>
                    <TrashBinTrashOutlineIcon />
                    {t("common.deletePermanently")}
                  </DropdownMenuItem>
                </>
              ) : (
                <>
                  <DropdownMenuItem onClick={() => setIsDetailsOpen(true)}>
                    <InfoCircleOutlineIcon />
                    {t("common.details")}
                  </DropdownMenuItem>
                  <DropdownMenuSeparator />
                  <DropdownMenuItem onClick={() => setIsRenameOpen(true)}>
                    <PenNewSquareOutlineIcon />
                    {t("common.rename")}
                  </DropdownMenuItem>
                  <DropdownMenuItem onClick={() => setIsMoveOpen(true)}>
                    <MoveToFolderOutlineIcon />
                    {t("common.move")}
                  </DropdownMenuItem>
                  <DropdownMenuSeparator />
                  <DropdownMenuItem
                    variant="destructive"
                    onClick={() => setIsTrashOpen(true)}
                  >
                    <TrashBinTrashOutlineIcon />
                    {t("common.trash")}
                  </DropdownMenuItem>
                </>
              )}
            </DropdownMenuContent>
          </DropdownMenu>
        </AttachmentActions>
      </Attachment>

      <FolderDetailsDialog
        open={isDetailsOpen}
        onOpenChange={setIsDetailsOpen}
        folder={folder}
      />
      <RenameFolderDialog
        open={isRenameOpen}
        onOpenChange={setIsRenameOpen}
        folder={folder}
      />
      <MoveFolderDialog
        open={isMoveOpen}
        onOpenChange={setIsMoveOpen}
        folder={folder}
      />
      <TrashFolderDialog
        open={isTrashOpen}
        onOpenChange={setIsTrashOpen}
        folder={folder}
      />
    </>
  )
}
