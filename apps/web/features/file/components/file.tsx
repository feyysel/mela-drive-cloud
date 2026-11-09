import { DriveFile } from "../types"

import {
  Attachment,
  AttachmentAction,
  AttachmentActions,
  AttachmentContent,
  AttachmentDescription,
  AttachmentMedia,
  AttachmentTitle,
} from "@/components/ui/attachment"
import FileTextBoldIcon from "@iconify-react/solar/file-text-bold"
import GalleryBoldIcon from "@iconify-react/solar/gallery-bold"
import MusicNoteBoldIcon from "@iconify-react/solar/music-note-bold"
import VideocameraRecordBoldIcon from "@iconify-react/solar/videocamera-record-bold"
import { cn } from "cn"
import { filesize } from "filesize"
import { MoreVerticalIcon } from "lucide-react"
import { extension } from "mime-types"

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
  const Icon = getIcon(file.mimeType)

  return (
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
          {extension(file.mimeType)}{" "}
          <span className="text-muted-foreground/50">•</span>{" "}
          {filesize(file.sizeBytes, { standard: "jedec" })}
        </AttachmentDescription>
      </AttachmentContent>
      <AttachmentActions>
        <AttachmentAction size="icon-lg">
          <MoreVerticalIcon />
        </AttachmentAction>
      </AttachmentActions>
    </Attachment>
  )
}
