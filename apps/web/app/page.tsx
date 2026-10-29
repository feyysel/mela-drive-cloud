import { getTranslations } from "next-intl/server"

import { LocaleSwitcher } from "@/components/locale-switcher"
import { Button } from "@/components/ui/button"
import { Separator } from "@/components/ui/separator"

export default async function Page() {
  const t = await getTranslations()

  return (
    <div className="flex min-h-svh p-6">
      <div className="flex max-w-md min-w-0 flex-col gap-4 text-sm leading-loose">
        <LocaleSwitcher />

        <p>{t("common.eg")}</p>

        <div className="flex gap-2">
          <Button>{t("common.save")}</Button>
          <Button variant="outline">{t("common.cancel")}</Button>
        </div>

        <Separator />

        <ul className="list-disc pl-5">
          <li>{t("errors.unexpected")}</li>
          <li>{t("errors.offline")}</li>
          <li>{t("errors.timeout")}</li>
        </ul>

        <Separator />

        <div className="font-mono text-xs text-muted-foreground">
          (Press <kbd>d</kbd> to toggle dark mode)
        </div>
      </div>
    </div>
  )
}
