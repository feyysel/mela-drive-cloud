import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogDescription,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
  AlertDialogTrigger,
} from "@/components/ui/alert-dialog"
import { Button } from "@/components/ui/button"
import { Spinner } from "@/components/ui/spinner"
import { toast } from "@/components/ui/toast"
import { parseErr } from "@/store/utils"
import TrashBinTrashOutlineIcon from "@iconify-react/solar/trash-bin-trash-outline"
import { useTranslations } from "next-intl"
import { useState } from "react"
import { useEmptyTrashMutation } from "../api"

export function EmptyTrashDialog({ disabled }: { disabled?: boolean }) {
  const t = useTranslations()
  const [open, setOpen] = useState(false)
  const [emptyTrash, { isLoading }] = useEmptyTrashMutation()

  const handleEmptyTrash = async () => {
    try {
      await emptyTrash().unwrap()
      setOpen(false)
    } catch (error: unknown) {
      toast.add({ type: "error", title: parseErr(error).message })
    }
  }

  return (
    <AlertDialog open={open} onOpenChange={setOpen}>
      <AlertDialogTrigger
        render={
          <Button disabled={disabled} variant="destructive">
            <TrashBinTrashOutlineIcon />
            <span>{t("drive.emptyTrash.action")}</span>
          </Button>
        }
      />

      <AlertDialogContent>
        <AlertDialogHeader>
          <AlertDialogTitle>{t("drive.emptyTrash.title")}</AlertDialogTitle>
          <AlertDialogDescription>
            {t("drive.emptyTrash.description")}
          </AlertDialogDescription>
        </AlertDialogHeader>

        <AlertDialogFooter>
          <AlertDialogCancel disabled={isLoading}>
            {t("common.cancel")}
          </AlertDialogCancel>
          <AlertDialogAction
            disabled={isLoading}
            onClick={handleEmptyTrash}
            className="md:w-36!"
          >
            {isLoading ? <Spinner /> : t("drive.emptyTrash.action")}
          </AlertDialogAction>
        </AlertDialogFooter>
      </AlertDialogContent>
    </AlertDialog>
  )
}
