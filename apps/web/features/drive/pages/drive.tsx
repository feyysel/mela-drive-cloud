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
import { useListFilesQuery } from "@/features/file/api"
import { File } from "@/features/file/components/file"
import { useListFoldersQuery } from "@/features/folder/api"
import { Folder } from "@/features/folder/components/folder"
import FolderOpenOutlineIcon from "@iconify-react/solar/folder-open-outline"
import { LayoutGridIcon, ListIcon } from "lucide-react"
import { useTranslations } from "next-intl"
import { parseAsStringLiteral, useQueryState } from "nuqs"
import { AddButton } from "../components/add-button"

const views = ["grid", "list"] as const
type DriveView = (typeof views)[number]

export function Drive() {
  const t = useTranslations()
  const [view, setView] = useQueryState(
    "view",
    parseAsStringLiteral(views).withDefault("grid")
  )
  const { data: folders, isLoading: isLoadingFolders } =
    useListFoldersQuery(undefined)
  const { data: files, isLoading: isLoadingFiles } =
    useListFilesQuery(undefined)

  return (
    <section className="flex flex-1 flex-col gap-5 p-5">
      <div className="flex items-center justify-between">
        <h1 className="text-xl font-medium">{t("navigation.myDrive")}</h1>

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
        <>
          {view === "list" ? (
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
            <div className="grid grid-cols-4 gap-5">
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
          )}
        </>
      ) : (
        <Empty className="h-full">
          <EmptyHeader>
            <EmptyMedia>
              <FolderOpenOutlineIcon className="size-10 text-muted-foreground" />
            </EmptyMedia>
            <EmptyTitle>{t("drive.empty.title")}</EmptyTitle>
            <EmptyDescription>{t("drive.empty.description")}</EmptyDescription>
          </EmptyHeader>
        </Empty>
      )}

      <AddButton />
    </section>
  )
}
