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
import MoveToFolderOutlineIcon from "@iconify-react/solar/move-to-folder-outline"
import PenNewSquareOutlineIcon from "@iconify-react/solar/pen-new-square-outline"
import TrashBinTrashOutlineIcon from "@iconify-react/solar/trash-bin-trash-outline"
import { cn } from "cn"
import { MoreVerticalIcon } from "lucide-react"
import { useTranslations } from "next-intl"
import { useState } from "react"
import { Folder as FolderType } from "../types"
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
  const t = useTranslations()
  const [isRenameOpen, setIsRenameOpen] = useState(false)

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
      >
        <AttachmentMedia className="bg-transparent">
          <FolderOpenBoldIcon className="size-6" />
        </AttachmentMedia>
        <AttachmentContent>
          <AttachmentTitle>{folder.name}</AttachmentTitle>
        </AttachmentContent>
        <AttachmentActions>
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
              <DropdownMenuItem onClick={() => setIsRenameOpen(true)}>
                <PenNewSquareOutlineIcon />
                {t("common.rename")}
              </DropdownMenuItem>
              <DropdownMenuItem>
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

      <RenameFolderDialog
        open={isRenameOpen}
        onOpenChange={setIsRenameOpen}
        folder={folder}
      />
    </>
  )
}
