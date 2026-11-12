"use client"

import {
  AudioPlayer,
  AudioPlayerControls,
  AudioPlayerPlay,
  AudioPlayerSeek,
  AudioPlayerSkipBack,
  AudioPlayerSkipForward,
  AudioPlayerTime,
  AudioPlayerVolume,
} from "@/components/ui/audio-player"
import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
} from "@/components/ui/dialog"
import { DriveFile } from "@/features/file/types"
import {
  PauseIcon,
  PlayIcon,
  SkipBackIcon,
  SkipForwardIcon,
} from "lucide-react"

export function AudioPlayerDialog({
  open,
  onOpenChange,
  file,
}: {
  open: boolean
  onOpenChange: (open: boolean) => void
  file: DriveFile
}) {
  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="bg-transparent sm:max-w-md">
        <DialogHeader>
          <DialogTitle className="truncate pr-8">{file.name}</DialogTitle>
        </DialogHeader>

        <AudioPlayer className="w-full flex-col items-stretch" src={file.url}>
          <div className="flex items-center gap-2">
            <AudioPlayerTime />
            <AudioPlayerSeek />
            <AudioPlayerTime type="remaining" />
          </div>
          <div className="flex items-center justify-between">
            <AudioPlayerControls>
              <AudioPlayerSkipBack>
                <SkipBackIcon />
              </AudioPlayerSkipBack>
              <AudioPlayerPlay>
                {({ playing }) => (playing ? <PauseIcon /> : <PlayIcon />)}
              </AudioPlayerPlay>
              <AudioPlayerSkipForward>
                <SkipForwardIcon />
              </AudioPlayerSkipForward>
            </AudioPlayerControls>
            <AudioPlayerVolume />
          </div>
        </AudioPlayer>
      </DialogContent>
    </Dialog>
  )
}
