"use client"

import { Button } from "@/components/ui/button"
import {
  Empty,
  EmptyContent,
  EmptyDescription,
  EmptyHeader,
  EmptyMedia,
  EmptyTitle,
} from "@/components/ui/empty"
import { cn } from "cn"
import { UploadCloudIcon } from "lucide-react"
import { useDropzone, type Accept, type FileRejection } from "react-dropzone"

type DropzoneProps = {
  onDrop: (files: File[]) => void
  onReject?: (rejections: FileRejection[]) => void
  accept?: Accept
  maxSize?: number
  multiple?: boolean
  disabled?: boolean
  className?: string
  title?: string
  description?: string
  cta?: string
}

export function Dropzone({
  onDrop,
  onReject,
  accept = {
    "image/png": [],
    "image/jpeg": [],
    "image/webp": [],
  },
  maxSize,
  multiple = false,
  disabled,
  className,
  title = "Drag and drop files here",
  description = "PNG, JPG, WEBP up to 5 MB",
  cta = "Browse files",
}: DropzoneProps) {
  const { getRootProps, getInputProps, isDragActive, open } = useDropzone({
    onDropAccepted: onDrop,
    onDropRejected: onReject,
    accept,
    maxSize,
    multiple,
    disabled,
    noClick: true,
    noKeyboard: true,
  })

  return (
    <Empty
      {...getRootProps({
        className: cn(
          "border bg-accent/50 p-4",
          isDragActive && "bg-accent",
          disabled && "pointer-events-none opacity-50",
          className
        ),
      })}
    >
      <input {...getInputProps()} />
      <EmptyHeader>
        <EmptyMedia>
          <UploadCloudIcon className="size-8" />
        </EmptyMedia>
        <EmptyTitle className="text-base">{title}</EmptyTitle>
        <EmptyDescription>{description}</EmptyDescription>
      </EmptyHeader>
      <EmptyContent>
        <Button size="lg" variant="outline" onClick={open}>
          {cta}
        </Button>
      </EmptyContent>
    </Empty>
  )
}
