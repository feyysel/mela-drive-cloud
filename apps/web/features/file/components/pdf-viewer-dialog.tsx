"use client"

import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import { Spinner } from "@/components/ui/spinner"
import { DriveFile } from "@/features/file/types"
import { cn } from "cn"
import { useState } from "react"

export function PdfViewerDialog({
  open,
  onOpenChange,
  file,
}: {
  open: boolean
  onOpenChange: (open: boolean) => void
  file: DriveFile
}) {
  const [isLoaded, setIsLoaded] = useState(false)

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="h-[85vh] grid-rows-[auto_1fr] bg-transparent sm:max-w-5xl">
        <DialogHeader>
          <DialogTitle className="truncate pr-8">{file.name}</DialogTitle>
        </DialogHeader>

        <div className="relative min-h-0 overflow-hidden rounded-lg border">
          {!isLoaded && (
            <div className="absolute inset-0 flex items-center justify-center">
              <Spinner className="size-6" />
            </div>
          )}
          <iframe
            src={file.url}
            title={file.name}
            className={cn(
              "size-full transition-opacity duration-200",
              isLoaded ? "opacity-100" : "opacity-0"
            )}
            onLoad={() => setIsLoaded(true)}
          />
        </div>
      </DialogContent>
    </Dialog>
  )
}
