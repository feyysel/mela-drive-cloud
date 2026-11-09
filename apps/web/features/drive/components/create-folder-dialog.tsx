"use client"

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
import { useCreateFolderMutation } from "@/features/folder/api"
import { parseErr } from "@/store/utils"
import { useTranslations } from "next-intl"
import { SubmitEvent, useState } from "react"

export function CreateFolderDialog({
  open,
  onOpenChange,
  parentId,
}: {
  open: boolean
  onOpenChange: (open: boolean) => void
  parentId?: string
}) {
  const t = useTranslations()
  const [name, setName] = useState("")
  const [createFolder, { isLoading }] = useCreateFolderMutation()

  function handleOpenChange(next: boolean) {
    if (!next) setName("")
    onOpenChange(next)
  }

  const handleSubmit = async (e: SubmitEvent<HTMLFormElement>) => {
    e.preventDefault()

    try {
      await createFolder({ name, parentId }).unwrap()
      handleOpenChange(false)
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
              <AlertDialogTitle>
                {t("drive.createFolder.title")}
              </AlertDialogTitle>
            </AlertDialogHeader>

            <Field>
              <Input
                autoFocus
                required
                value={name}
                onChange={(e) => setName(e.target.value)}
                placeholder={t("drive.createFolder.placeholder")}
              />
            </Field>

            <AlertDialogFooter>
              <AlertDialogCancel type="button" disabled={isLoading}>
                {t("common.cancel")}
              </AlertDialogCancel>
              <AlertDialogAction
                type="submit"
                disabled={!name || isLoading}
                className="md:w-20!"
              >
                {isLoading ? <Spinner /> : t("common.create")}
              </AlertDialogAction>
            </AlertDialogFooter>
          </FieldGroup>
        </form>
      </AlertDialogContent>
    </AlertDialog>
  )
}
