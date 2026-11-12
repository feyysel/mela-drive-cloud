"use client"

import {
  Empty,
  EmptyDescription,
  EmptyHeader,
  EmptyMedia,
  EmptyTitle,
} from "@/components/ui/empty"
import { Spinner } from "@/components/ui/spinner"
import { Tabs, TabsList, TabsTrigger } from "@/components/ui/tabs"
import { useListTrashedFilesQuery } from "@/features/file/api"
import { File } from "@/features/file/components/file"
import { useListTrashedFoldersQuery } from "@/features/folder/api"
import { Folder } from "@/features/folder/components/folder"
import TrashBinTrashOutlineIcon from "@iconify-react/solar/trash-bin-trash-outline"
import { LayoutGridIcon, ListIcon } from "lucide-react"
import { useTranslations } from "next-intl"
import { parseAsStringLiteral, useQueryState } from "nuqs"

const views = ["grid", "list"] as const
type DriveView = (typeof views)[number]

export function Trash() {
  const t = useTranslations()
  const [view, setView] = useQueryState(
    "view",
    parseAsStringLiteral(views).withDefault("grid")
  )

  const { data: folders, isLoading: isLoadingFolders } =
    useListTrashedFoldersQuery()
  const { data: files, isLoading: isLoadingFiles } = useListTrashedFilesQuery()

  return (
    <section className="flex flex-1 flex-col gap-5 p-5">
      <div className="flex items-center justify-between">
        <h1 className="text-xl font-medium">{t("navigation.trash")}</h1>

        <Tabs
          value={view}
          onValueChange={(value) => setView(value as DriveView)}
        >
          <TabsList>
            <TabsTrigger value="grid" aria-label="Grid view">
              <LayoutGridIcon />
            </TabsTrigger>
            <TabsTrigger value="list" aria-label="List view">
              <ListIcon />
            </TabsTrigger>
          </TabsList>
        </Tabs>
      </div>

      {isLoadingFiles || isLoadingFolders ? (
        <Empty className="h-full">
          <EmptyMedia>
            <Spinner className="size-6" />
          </EmptyMedia>
        </Empty>
      ) : folders && files && (folders.length || files.length) ? (
        view === "list" ? (
          <div className="grid grid-cols-1">
            {folders.map((folder, index) => (
              <Folder
                key={folder.id}
                folder={folder}
                index={index}
                view="list"
              />
            ))}
            {files.map((file, index) => (
              <File key={file.id} file={file} index={index} view="list" />
            ))}
          </div>
        ) : (
          <div className="grid grid-cols-2 gap-5 md:grid-cols-4">
            {folders.map((folder, index) => (
              <Folder
                key={folder.id}
                folder={folder}
                index={index}
                view="grid"
              />
            ))}
            {files.map((file, index) => (
              <File key={file.id} file={file} index={index} view="grid" />
            ))}
          </div>
        )
      ) : (
        <Empty className="h-full">
          <EmptyHeader>
            <EmptyMedia>
              <TrashBinTrashOutlineIcon className="size-10 text-muted-foreground" />
            </EmptyMedia>
            <EmptyTitle>{t("trash.empty.title")}</EmptyTitle>
            <EmptyDescription>{t("trash.empty.description")}</EmptyDescription>
          </EmptyHeader>
        </Empty>
      )}
    </section>
  )
}
