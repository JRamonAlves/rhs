import { SpinnerGapIcon } from "@phosphor-icons/react"

import { Button } from "@/components/ui/button"
import { CardAction, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"

type SharedClipboardHeaderProps = {
  isRefreshing: boolean
  isSaving: boolean
  onRefresh: () => void
}

export function SharedClipboardHeader({
  isRefreshing,
  isSaving,
  onRefresh,
}: SharedClipboardHeaderProps) {
  return (
    <CardHeader>
      <CardTitle>Shared clipboard</CardTitle>
      <CardDescription>
        Save text here, then open or reload this page on another device.
      </CardDescription>
      <CardAction>
        <Button
          type="button"
          variant="outline"
          size="sm"
          disabled={isRefreshing || isSaving}
          onClick={onRefresh}
        >
          {isRefreshing && <SpinnerGapIcon data-icon="inline-start" className="animate-spin" />}
          Refresh
        </Button>
      </CardAction>
    </CardHeader>
  )
}
