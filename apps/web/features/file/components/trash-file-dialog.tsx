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
import { useTrashFileMutation } from "@/features/file/api"
import { DriveFile } from "@/features/file/types"
import { parseErr } from "@/store/utils"
import { useTranslations } from "next-intl"

export function TrashFileDialog({
  open,
  onOpenChange,
  file,
}: {
  open: boolean
  onOpenChange: (open: boolean) => void
  file: DriveFile
}) {
  const t = useTranslations()
  const [trashFile, { isLoading }] = useTrashFileMutation()

  const handleTrash = async () => {
    try {
      await trashFile(file.id).unwrap()
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
            {t("drive.trashFile.description", { name: file.name })}
          </AlertDialogDescription>
        </AlertDialogHeader>

        <AlertDialogFooter>
          <AlertDialogCancel disabled={isLoading}>
            {t("common.cancel")}
          </AlertDialogCancel>
          <AlertDialogAction
            disabled={isLoading}
            onClick={handleTrash}
            className="md:w-32!"
          >
            {isLoading ? <Spinner /> : t("common.trash")}
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}
