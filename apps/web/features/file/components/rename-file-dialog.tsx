import {
  AlertDialog,
  AlertDialogAction,
  AlertDialogCancel,
  AlertDialogContent,
  AlertDialogFooter,
  AlertDialogHeader,
  AlertDialogTitle,
} from "@/components/ui/alert-dialog"
import { Field, FieldGroup } from "@/components/ui/field"
import { Input } from "@/components/ui/input"
import { Spinner } from "@/components/ui/spinner"
import { toast } from "@/components/ui/toast"
import { useRenameFileMutation } from "@/features/file/api"
import { DriveFile } from "@/features/file/types"
import { parseErr } from "@/store/utils"
import { useTranslations } from "next-intl"
import { SubmitEvent, useState } from "react"

export function RenameFileDialog({
  open,
  onOpenChange,
  file,
}: {
  open: boolean
  onOpenChange: (open: boolean) => void
  file: DriveFile
}) {
  const t = useTranslations()
  const [name, setName] = useState(file.name)
  const [renameFile, { isLoading }] = useRenameFileMutation()

  const isDirty = name.trim() !== file.name

  function handleOpenChange(next: boolean) {
    if (!next) setName(file.name)
    onOpenChange(next)
  }

  const handleSubmit = async (e: SubmitEvent<HTMLFormElement>) => {
    e.preventDefault()

    try {
      const updated = await renameFile({ fileId: file.id, name }).unwrap()
      onOpenChange(false)
      setName(updated.name)
    } catch (error: unknown) {
      toast.add({ type: "error", title: parseErr(error).message })
    }
  }

  return (
    <AlertDialog open={open} onOpenChange={handleOpenChange}>
      <AlertDialogContent>
        <form onSubmit={handleSubmit}>
          <FieldGroup>
            <AlertDialogHeader>
              <AlertDialogTitle>{t("common.rename")}</AlertDialogTitle>
            </AlertDialogHeader>

            <Field>
              <Input
                autoFocus
                required
                value={name}
                onChange={(e) => setName(e.target.value)}
                placeholder={t("drive.file.placeholder")}
              />
            </Field>

            <AlertDialogFooter>
              <AlertDialogCancel type="button" disabled={isLoading}>
                {t("common.cancel")}
              </AlertDialogCancel>
              <AlertDialogAction
                type="submit"
                disabled={!name || !isDirty || isLoading}
                className="md:w-20!"
              >
                {isLoading ? <Spinner /> : t("common.save")}
              </AlertDialogAction>
            </AlertDialogFooter>
          </FieldGroup>
        </form>
      </AlertDialogContent>
    </AlertDialog>
  )
}
