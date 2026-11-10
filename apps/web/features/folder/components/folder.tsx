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
import FolderOpenBoldIcon from "@iconify-react/solar/folder-open-bold"
import InfoCircleOutlineIcon from "@iconify-react/solar/info-circle-outline"
import MoveToFolderOutlineIcon from "@iconify-react/solar/move-to-folder-outline"
import PenNewSquareOutlineIcon from "@iconify-react/solar/pen-new-square-outline"
import TrashBinTrashOutlineIcon from "@iconify-react/solar/trash-bin-trash-outline"
import { cn } from "cn"
import { MoreVerticalIcon } from "lucide-react"
import { useTranslations } from "next-intl"
import { useRouter } from "next/navigation"
import { useState } from "react"
import { Folder as FolderType } from "../types"
import { FolderDetailsDialog } from "./folder-details-dialog"
import { MoveFolderDialog } from "./move-folder-dialog"
import { RenameFolderDialog } from "./rename-folder-dialog"

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
        onClick={() => router.push(`/drive/folders/${folder.id}`)}
      >
        <AttachmentMedia className="bg-transparent">
          <FolderOpenBoldIcon className="size-6" />
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
              <DropdownMenuItem variant="destructive">
                <TrashBinTrashOutlineIcon />
                {t("common.trash")}
              </DropdownMenuItem>
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
    </>
  )
}
