import { CheckIcon, WarningCircleIcon } from "@phosphor-icons/react"

type SharedClipboardStatusProps = {
  hasLoadError: boolean
  hasSaveError: boolean
  hasCopyError: boolean
  isSaved: boolean
  hasChanges: boolean
}

export function SharedClipboardStatus({
  hasLoadError,
  hasSaveError,
  hasCopyError,
  isSaved,
  hasChanges,
}: SharedClipboardStatusProps) {
  return (
    <div className="min-h-4 text-xs text-muted-foreground">
      {hasLoadError ? (
        <span className="flex items-center gap-1 text-destructive" role="alert">
          <WarningCircleIcon /> Unable to load the shared text.
        </span>
      ) : hasSaveError ? (
        <span className="flex items-center gap-1 text-destructive" role="alert">
          <WarningCircleIcon /> Unable to save. Your text is still here.
        </span>
      ) : hasCopyError ? (
        <span className="flex items-center gap-1 text-destructive" role="alert">
          <WarningCircleIcon /> Unable to copy. Select the text manually.
        </span>
      ) : isSaved ? (
        <span className="flex items-center gap-1" role="status">
          <CheckIcon /> Saved for your other devices.
        </span>
      ) : hasChanges ? (
        <span role="status">Unsaved changes</span>
      ) : (
        <span role="status">Up to date</span>
      )}
    </div>
  )
}
