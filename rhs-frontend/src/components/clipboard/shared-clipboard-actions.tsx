import { CheckIcon, ClipboardIcon, FloppyDiskIcon, SpinnerGapIcon } from "@phosphor-icons/react"

import { Button } from "@/components/ui/button"

type SharedClipboardActionsProps = {
  value: string
  copied: boolean
  isSaving: boolean
  hasChanges: boolean
  onCopy: () => void
  onSave: () => void
}

export function SharedClipboardActions({
  value,
  copied,
  isSaving,
  hasChanges,
  onCopy,
  onSave,
}: SharedClipboardActionsProps) {
  return (
    <div className="flex gap-2">
      <Button
        type="button"
        variant="outline"
        className="flex-1 sm:flex-none"
        disabled={!value}
        onClick={onCopy}
      >
        {copied ? <CheckIcon data-icon="inline-start" /> : <ClipboardIcon data-icon="inline-start" />}
        {copied ? "Copied" : "Copy"}
      </Button>
      <Button
        type="button"
        className="flex-1 sm:flex-none"
        disabled={isSaving || !hasChanges}
        onClick={onSave}
      >
        {isSaving ? (
          <SpinnerGapIcon data-icon="inline-start" className="animate-spin" />
        ) : (
          <FloppyDiskIcon data-icon="inline-start" />
        )}
        {isSaving ? "Saving..." : "Save"}
      </Button>
    </div>
  )
}
