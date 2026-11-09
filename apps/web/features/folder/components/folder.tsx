import { Folder as FolderType } from "../types"

import {
  Attachment,
  AttachmentAction,
  AttachmentActions,
  AttachmentContent,
  AttachmentMedia,
  AttachmentTitle,
} from "@/components/ui/attachment"
import FolderOpenBoldIcon from "@iconify-react/solar/folder-open-bold"
import { cn } from "cn"
import { MoreVerticalIcon } from "lucide-react"

export function Folder({
  folder,
  index,
  view,
}: {
  folder: FolderType
  index: number
  view: "list" | "grid"
}) {
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
        <FolderOpenBoldIcon className="size-6" />
      </AttachmentMedia>
      <AttachmentContent>
        <AttachmentTitle>{folder.name}</AttachmentTitle>
      </AttachmentContent>
      <AttachmentActions>
        <AttachmentAction size="icon-lg">
          <MoreVerticalIcon />
        </AttachmentAction>
      </AttachmentActions>
    </Attachment>
  )
}
