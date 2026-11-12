import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog"
import { Spinner } from "@/components/ui/spinner"
import { toast } from "@/components/ui/toast"
import { useTrashFolderMutation } from "@/features/folder/api"
import { Folder as FolderType } from "@/features/folder/types"
import { parseErr } from "@/store/utils"
import { useTranslations } from "next-intl"

export function TrashFolderDialog({
  open,
  onOpenChange,
  folder,
}: {
  open: boolean
  onOpenChange: (open: boolean) => void
  folder: FolderType
}) {
  const t = useTranslations()
  const [trashFolder, { isLoading }] = useTrashFolderMutation()

  const handleTrash = async () => {
    try {
      await trashFolder(folder.id).unwrap()
      onOpenChange(false)
    } catch (error: unknown) {
      toast.add({ type: "error", title: parseErr(error).message })
    }
  }

  return (
    <AlertDialog open={open} onOpenChange={onOpenChange}>
      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>{t("common.trash")}</AlertDialogTitle>
          <AlertDialogDescription>
            {t("drive.trashFolder.description", { name: folder.name })}
          </AlertDialogDescription>
        </AlertDialogHeader>

        <AlertDialogFooter>
          <AlertDialogCancel disabled={isLoading}>
            {t("common.cancel")}
          </AlertDialogCancel>
          <AlertDialogAction
            disabled={isLoading}
            onClick={handleTrash}
            className="md:w-36!"
          >
            {isLoading ? <Spinner /> : t("common.trash")}
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}
