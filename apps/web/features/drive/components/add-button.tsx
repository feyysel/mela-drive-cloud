import { Button } from "@/components/ui/button"
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import AddFolderOutlineIcon from "@iconify-react/solar/add-folder-outline"
import FileSendOutlineIcon from "@iconify-react/solar/file-send-outline"
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
            <AddFolderOutlineIcon />
            {t("drive.newFolder")}
          </DropdownMenuItem>
          <DropdownMenuItem onClick={() => setIsUploadOpen(true)}>
            <FileSendOutlineIcon />
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
