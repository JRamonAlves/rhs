import { CheckIcon, ClipboardTextIcon } from "@phosphor-icons/react"
import { useState } from "react"

import type { Service } from "@/api/services.api"
import { Button } from "@/components/ui/button"
import { CardContent } from "@/components/ui/card"
import { copyToClipboard } from "@/lib/copy-to-clipboard"

type ServiceCardContentProps = {
  app: Service
}

export function ServiceCardContent({ app }: ServiceCardContentProps) {
  const [copied, setCopied] = useState(false)
  const linkUrl = app.url
  const CopyIcon = copied ? CheckIcon : ClipboardTextIcon

  async function copyUrl() {
    if (!linkUrl) {
      return
    }

    await copyToClipboard(linkUrl)
    setCopied(true)
    window.setTimeout(() => setCopied(false), 1600)
  }

  return (
    <CardContent>
      <div className="min-w-0 border p-2 text-xs">
        <div className="flex items-center justify-between gap-2">
          <div className="text-muted-foreground">URL</div>
          <Button
            aria-label={`Copy ${app.name} URL`}
            title={`Copy ${app.name} URL`}
            size="icon-xs"
            variant="ghost"
            type="button"
            onClick={() => void copyUrl()}
            disabled={!linkUrl}
          >
            <CopyIcon weight={copied ? "bold" : "regular"} />
          </Button>
        </div>
        <div className="truncate font-medium">{linkUrl}</div>
      </div>
    </CardContent>
  )
}
