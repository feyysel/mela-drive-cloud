import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import AddFolderBoldIcon from "@iconify-react/solar/add-folder-bold"
import FileSendBoldIcon from "@iconify-react/solar/file-send-bold"
import { PlusIcon } from "lucide-react"
import { useTranslations } from "next-intl"
import { useState } from "react"
import { CreateFolderDialog } from "./create-folder-dialog"
import { UploadFileDialog } from "./upload-file-dialog"

export function AddButton({ folderId }: { folderId?: string }) {
  const t = useTranslations()
  const [isFolderOpen, setIsFolderOpen] = useState(false)
  const [isUploadOpen, setIsUploadOpen] = useState(false)

  return (
    <>
      <DropdownMenu>
        <DropdownMenuTrigger
          render={
            <Button
              size="icon-lg"
              className="absolute right-5 bottom-5 size-14 rounded-full"
              aria-label={t("drive.add")}
            >
              <PlusIcon className="size-6" />
            </Button>
          }
        />
        <DropdownMenuContent align="end" side="top" sideOffset={20}>
          <DropdownMenuItem onClick={() => setIsFolderOpen(true)}>
            <AddFolderBoldIcon />
            {t("drive.newFolder")}
          </DropdownMenuItem>
          <DropdownMenuItem onClick={() => setIsUploadOpen(true)}>
            <FileSendBoldIcon />
            {t("drive.uploadFile")}
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenu>

      <CreateFolderDialog
        open={isFolderOpen}
        onOpenChange={setIsFolderOpen}
        parentId={folderId}
      />
      <UploadFileDialog
        open={isUploadOpen}
        onOpenChange={setIsUploadOpen}
        folderId={folderId}
      />
    </>
  )
}
