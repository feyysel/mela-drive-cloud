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
import Image from "next/image"
import { useState } from "react"

export function ImageViewerDialog({
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
      <DialogContent className="bg-transparent sm:max-w-3xl">
        <DialogHeader>
          <DialogTitle className="truncate pr-8">{file.name}</DialogTitle>
        </DialogHeader>

        <div className="relative aspect-video w-full overflow-hidden rounded-lg bg-black">
          {!isLoaded && (
            <div className="absolute inset-0 flex items-center justify-center">
              <Spinner className="size-6" />
            </div>
          )}
          <Image
            src={file.url}
            alt={file.name}
            fill
            sizes="(min-width: 768px) 768px, 100vw"
            className={cn(
              "object-contain transition-opacity duration-200",
              isLoaded ? "opacity-100" : "opacity-0"
            )}
            onLoad={() => setIsLoaded(true)}
          />
        </div>
      </DialogContent>
    </Dialog>
  )
}
