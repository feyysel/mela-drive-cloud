import { Folders } from "@/features/drive/pages/folders"

export default async function Page({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const { id } = await params

  return <Folders folderId={id} />
}
