import { ArrowSquareOutIcon } from "@phosphor-icons/react"

import { Button } from "@/components/ui/button"
import { CardFooter } from "@/components/ui/card"

type ServiceCardFooterProps = {
  name: string
  url?: string
}

export function ServiceCardFooter({ name, url }: ServiceCardFooterProps) {
  return (
    <CardFooter>
      <Button
        className="w-full"
        disabled={!url}
        nativeButton={false}
        render={<a href={url} aria-label={`Open ${name}`} />}
      >
        Open
        <ArrowSquareOutIcon data-icon="inline-end" />
      </Button>
    </CardFooter>
  )
}
