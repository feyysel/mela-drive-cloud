import InfoCircleOutlineIcon from "@iconify-react/solar/info-circle-outline"
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
import { Spinner } from "@/components/ui/spinner"
import { toast } from "@/components/ui/toast"
import { parseErr } from "@/store/utils"
import FileTextBoldIcon from "@iconify-react/solar/file-text-bold"
import GalleryBoldIcon from "@iconify-react/solar/gallery-bold"
import MoveToFolderOutlineIcon from "@iconify-react/solar/move-to-folder-outline"
import MusicNoteBoldIcon from "@iconify-react/solar/music-note-bold"
import PenNewSquareOutlineIcon from "@iconify-react/solar/pen-new-square-outline"
import RefreshCircleOutlineIcon from "@iconify-react/solar/refresh-circle-outline"
import TrashBinTrashOutlineIcon from "@iconify-react/solar/trash-bin-trash-outline"
import VideocameraRecordBoldIcon from "@iconify-react/solar/videocamera-record-bold"
import { cn } from "cn"
import { filesize } from "filesize"
import { MoreVerticalIcon } from "lucide-react"
import { lookup } from "mime-types"
import { useTranslations } from "next-intl"
import { useState } from "react"
import { usePurgeFileMutation, useRestoreFileMutation } from "../api"
import { AudioPlayerDialog } from "./audio-player-dialog"
import { FileDetailsDialog } from "./file-details-dialog"
import { ImageViewerDialog } from "./image-viewer-dialog"
import { PdfViewerDialog } from "./pdf-viewer-dialog"
import { TrashFileDialog } from "./trash-file-dialog"
import { VideoPlayerDialog } from "./video-player-dialog"

function resolveMimeType({
  name,
  mimeType,
}: Pick<DriveFile, "name" | "mimeType">) {
  return mimeType && mimeType !== "application/octet-stream"
    ? mimeType
    : lookup(name) || mimeType
}

function getIcon(type: string) {
  if (type.startsWith("image/")) return GalleryBoldIcon
  if (type.startsWith("video/")) return VideocameraRecordBoldIcon
  if (type.startsWith("audio/")) return MusicNoteBoldIcon
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
  const [isDetailsOpen, setIsDetailsOpen] = useState(false)
  const [isTrashOpen, setIsTrashOpen] = useState(false)
  const [isPreviewOpen, setIsPreviewOpen] = useState(false)

  const [restoreFile, { isLoading: isRestoringFile }] = useRestoreFileMutation()
  const [purgeFile, { isLoading: isPurgingFile }] = usePurgeFileMutation()

  const mimeType = resolveMimeType(file)
  const isVideo = mimeType.startsWith("video/")
  const isAudio = mimeType.startsWith("audio/")
  const isImage = mimeType.startsWith("image/")
  const isPdf = mimeType === "application/pdf"
  const isPreviewable = isVideo || isAudio || isImage || isPdf
  const Icon = getIcon(mimeType)

  const handleRestore = async () => {
    try {
      await restoreFile(file.id).unwrap()
    } catch (error: unknown) {
      toast.add({ type: "error", title: parseErr(error).message })
    }
  }

  const handlePurge = async () => {
    try {
      await purgeFile(file.id).unwrap()
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
        onClick={isPreviewable ? () => setIsPreviewOpen(true) : undefined}
      >
        <AttachmentMedia className="bg-transparent">
          {isRestoringFile || isPurgingFile ? (
            <Spinner className="size-6" />
          ) : (
            <Icon className="size-6" />
          )}
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
              {file.deletedAt ? (
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

      <FileDetailsDialog
        open={isDetailsOpen}
        onOpenChange={setIsDetailsOpen}
        file={file}
      />
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
      <TrashFileDialog
        open={isTrashOpen}
        onOpenChange={setIsTrashOpen}
        file={file}
      />

      {isVideo && (
        <VideoPlayerDialog
          open={isPreviewOpen}
          onOpenChange={setIsPreviewOpen}
          file={file}
        />
      )}
      {isAudio && (
        <AudioPlayerDialog
          open={isPreviewOpen}
          onOpenChange={setIsPreviewOpen}
          file={file}
        />
      )}
      {isImage && (
        <ImageViewerDialog
          open={isPreviewOpen}
          onOpenChange={setIsPreviewOpen}
          file={file}
        />
      )}
      {isPdf && (
        <PdfViewerDialog
          open={isPreviewOpen}
          onOpenChange={setIsPreviewOpen}
          file={file}
        />
      )}
    </>
  )
}
