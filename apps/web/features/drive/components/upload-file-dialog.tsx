import { Dropzone } from "@/components/dropzone"
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import {
  Empty,
  EmptyDescription,
  EmptyHeader,
  EmptyTitle,
} from "@/components/ui/empty"
import { toast } from "@/components/ui/toast"
import { useUploadFileMutation } from "@/features/file/api"
import { parseErr } from "@/store/utils"
import { useTranslations } from "next-intl"

const MAX_SIZE = 100 * 1024 * 1024
const ACCEPT = {
  "application/pdf": [],
  "image/*": [],
  "video/*": [],
  "audio/*": [],
  "text/plain": [],
}

export function UploadFileDialog({
  open,
  onOpenChange,
  folderId,
}: {
  open: boolean
  onOpenChange: (open: boolean) => void
  folderId?: string
}) {
  const t = useTranslations()
  const [uploadFile, { isLoading }] = useUploadFileMutation()

  const handleDrop = async (files: File[]) => {
    const [file] = files
    if (!file) return

    try {
      await uploadFile({ folderId, file }).unwrap()
      onOpenChange(false)
    } catch (error: unknown) {
      toast.add({ type: "error", title: parseErr(error).message })
    }
  }

  return (
    <Dialog
      open={open}
      onOpenChange={onOpenChange}
      disablePointerDismissal={isLoading}
    >
      <DialogContent showCloseButton={!isLoading}>
        <DialogHeader>
          <DialogTitle>{t("drive.upload.title")}</DialogTitle>
        </DialogHeader>

        {isLoading ? (
          <Empty className="min-h-80">
            <EmptyHeader>
              <EmptyTitle className="shimmer text-muted-foreground shimmer-duration-1300 shimmer-spread-5">
                Upload file
              </EmptyTitle>
              <EmptyDescription>Please wait</EmptyDescription>
            </EmptyHeader>
          </Empty>
        ) : (
          <Dropzone
            multiple={false}
            accept={ACCEPT}
            maxSize={MAX_SIZE}
            onDrop={handleDrop}
            title={t("dropzone.title")}
            description={t("dropzone.description")}
            cta={t("dropzone.cta")}
            onReject={(rejections) => {
              const code = rejections[0]?.errors[0]?.code

              if (code === "too-many-files") return

              toast.add({
                type: "error",
                title:
                  code === "file-too-large"
                    ? t("drive.upload.tooLarge")
                    : t("drive.upload.rejected"),
              })
            }}
            className="min-h-80"
          />
        )}
      </DialogContent>
    </Dialog>
  )
}
