import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import { DriveFile } from "@/features/file/types"
import { filesize } from "filesize"
import { extension } from "mime-types"
import { useFormatter, useTranslations } from "next-intl"

export function FileDetailsDialog({
  open,
  onOpenChange,
  file,
}: {
  open: boolean
  onOpenChange: (open: boolean) => void
  file: DriveFile
}) {
  const t = useTranslations()
  const format = useFormatter()

  const formatDate = (value: string) =>
    format.dateTime(new Date(value), {
      dateStyle: "medium",
      timeStyle: "short",
    })

  const rows = [
    { label: t("drive.details.name"), value: file.name },
    {
      label: t("drive.details.type"),
      value: `${extension(file.mimeType) || file.mimeType}`.toUpperCase(),
    },
    {
      label: t("drive.details.size"),
      value: filesize(file.sizeBytes, { standard: "jedec" }),
    },
    { label: t("drive.details.created"), value: formatDate(file.createdAt) },
    { label: t("drive.details.modified"), value: formatDate(file.updatedAt) },
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
