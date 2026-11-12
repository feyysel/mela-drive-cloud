import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import { Folder as FolderType } from "@/features/folder/types"
import { useFormatter, useTranslations } from "next-intl"

export function FolderDetailsDialog({
  open,
  onOpenChange,
  folder,
}: {
  open: boolean
  onOpenChange: (open: boolean) => void
  folder: FolderType
}) {
  const t = useTranslations()
  const format = useFormatter()

  const formatDate = (value: string) =>
    format.dateTime(new Date(value), {
      dateStyle: "medium",
      timeStyle: "short",
    })

  const rows = [
    { label: t("drive.details.name"), value: folder.name },
    { label: t("drive.details.created"), value: formatDate(folder.createdAt) },
    { label: t("drive.details.modified"), value: formatDate(folder.updatedAt) },
  ]

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>{t("common.details")}</DialogTitle>
        </DialogHeader>

        <dl className="grid grid-cols-[auto_1fr] gap-x-6 gap-y-3 text-sm">
          {rows.map((row) => (
            <div key={row.label} className="col-span-2 grid grid-cols-subgrid">
              <dt className="text-muted-foreground">{row.label}</dt>
              <dd className="break-all">{row.value}</dd>
            </div>
          ))}
        </dl>
      </DialogContent>
    </Dialog>
  )
}
