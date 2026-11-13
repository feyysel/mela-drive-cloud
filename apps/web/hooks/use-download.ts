import { toast } from "@/components/ui/toast"
import { parseErr } from "@/store/utils"
import { useState } from "react"

export function useDownload() {
  const [isDownloading, setIsDownloading] = useState(false)

  const downloadBlob = (blob: Blob, filename: string) => {
    const url = URL.createObjectURL(blob)
    const link = document.createElement("a")

    link.href = url
    link.download = filename

    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)

    URL.revokeObjectURL(url)
  }

  const downloadFile = async (url: string, filename: string) => {
    try {
      setIsDownloading(true)

      const response = await fetch(url)

      if (!response.ok) {
        throw new Error("Failed to download file")
      }

      downloadBlob(await response.blob(), filename)
    } catch (err: unknown) {
      toast.add({ type: "error", title: parseErr(err).message })
    } finally {
      setIsDownloading(false)
    }
  }

  return { downloadFile, downloadBlob, isDownloading }
}
