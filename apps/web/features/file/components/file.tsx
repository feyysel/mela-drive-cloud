import { DriveFile } from "../types"
import { MoveFileDialog } from "./move-file-dialog"
import { RenameFileDialog } from "./rename-file-dialog"

import {
  Attachment,
  AttachmentAction,
  AttachmentActions,
  AttachmentContent,
  AttachmentDescription,
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
import FileTextBoldIcon from "@iconify-react/solar/file-text-bold"
import GalleryBoldIcon from "@iconify-react/solar/gallery-bold"
import MoveToFolderOutlineIcon from "@iconify-react/solar/move-to-folder-outline"
import MusicNoteBoldIcon from "@iconify-react/solar/music-note-bold"
import PenNewSquareOutlineIcon from "@iconify-react/solar/pen-new-square-outline"
import TrashBinTrashOutlineIcon from "@iconify-react/solar/trash-bin-trash-outline"
import VideocameraRecordBoldIcon from "@iconify-react/solar/videocamera-record-bold"
import { cn } from "cn"
import { filesize } from "filesize"
import { MoreVerticalIcon } from "lucide-react"
import { useTranslations } from "next-intl"
import { useState } from "react"

function getIcon(mimeType: string) {
  if (mimeType.startsWith("image/")) return GalleryBoldIcon
  if (mimeType.startsWith("video/")) return VideocameraRecordBoldIcon
  if (mimeType.startsWith("audio/")) return MusicNoteBoldIcon
  return FileTextBoldIcon
}

export function File({
  file,
  index,
  view,
}: {
  file: DriveFile
  index: number
  view: "list" | "grid"
}) {
  const t = useTranslations()
  const [isRenameOpen, setIsRenameOpen] = useState(false)
  const [isMoveOpen, setIsMoveOpen] = useState(false)
  const Icon = getIcon(file.mimeType)

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
          <Icon className="size-6" />
        </AttachmentMedia>
        <AttachmentContent>
          <AttachmentTitle>{file.name}</AttachmentTitle>
          <AttachmentDescription className="uppercase">
            {filesize(file.sizeBytes, { standard: "jedec" })}
          </AttachmentDescription>
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

      <RenameFileDialog
        open={isRenameOpen}
        onOpenChange={setIsRenameOpen}
        file={file}
      />
      <MoveFileDialog
        open={isMoveOpen}
        onOpenChange={setIsMoveOpen}
        file={file}
      />
    </>
  )
}
