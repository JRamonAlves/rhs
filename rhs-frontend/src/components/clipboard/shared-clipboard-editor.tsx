import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
import { useState } from "react"

import { getStorage, setStorage } from "@/api/storage.api"
import { Card, CardContent, CardFooter } from "@/components/ui/card"
import { Textarea } from "@/components/ui/textarea"
import { copyToClipboard } from "@/lib/copy-to-clipboard"

import { SharedClipboardActions } from "./shared-clipboard-actions"
import { SharedClipboardHeader } from "./shared-clipboard-header"
import { SharedClipboardStatus } from "./shared-clipboard-status"
import { STORAGE_KEY, STORAGE_QUERY_KEY } from "./storage"

export function SharedClipboardEditor({ initialValue }: { initialValue: string }) {
  const queryClient = useQueryClient()
  const [value, setValue] = useState(initialValue)
  const [copied, setCopied] = useState(false)
  const [copyError, setCopyError] = useState(false)

  const storageQuery = useQuery({
    queryKey: STORAGE_QUERY_KEY,
    queryFn: () => getStorage(STORAGE_KEY),
    initialData: initialValue,
  })

  const saveMutation = useMutation({
    mutationFn: (nextValue: string) => setStorage(STORAGE_KEY, nextValue),
    onSuccess: (_, nextValue) => {
      queryClient.setQueryData(STORAGE_QUERY_KEY, nextValue)
    },
  })

  async function refreshValue() {
    const result = await storageQuery.refetch()
    if (result.data !== undefined) {
      setValue(result.data)
      saveMutation.reset()
    }
  }

  async function copyValue() {
    try {
      await copyToClipboard(value)
      setCopyError(false)
      setCopied(true)
      window.setTimeout(() => setCopied(false), 1600)
    } catch {
      setCopied(false)
      setCopyError(true)
    }
  }

  const isSaving = saveMutation.isPending
  const hasChanges = value !== (storageQuery.data ?? "")

  return (
    <Card>
      <SharedClipboardHeader
        isRefreshing={storageQuery.isFetching}
        isSaving={isSaving}
        onRefresh={refreshValue}
      />
      <CardContent>
        <Textarea
          aria-label="Text shared between your devices"
          placeholder="Paste or type here..."
          rows={7}
          value={value}
          disabled={isSaving}
          onChange={(event) => {
            setValue(event.target.value)
            saveMutation.reset()
            setCopyError(false)
          }}
        />
      </CardContent>
      <CardFooter className="flex-col items-stretch gap-3 sm:flex-row sm:items-center sm:justify-between">
        <SharedClipboardStatus
          hasLoadError={storageQuery.isError}
          hasSaveError={saveMutation.isError}
          hasCopyError={copyError}
          isSaved={saveMutation.isSuccess && !hasChanges}
          hasChanges={hasChanges}
        />
        <SharedClipboardActions
          value={value}
          copied={copied}
          isSaving={isSaving}
          hasChanges={hasChanges}
          onCopy={copyValue}
          onSave={() => saveMutation.mutate(value)}
        />
      </CardFooter>
    </Card>
  )
}
