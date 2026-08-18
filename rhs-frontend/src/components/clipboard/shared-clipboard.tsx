import { useQuery } from "@tanstack/react-query"

import { getStorage } from "@/api/storage.api"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Textarea } from "@/components/ui/textarea"

import { SharedClipboardEditor } from "./shared-clipboard-editor"
import { STORAGE_KEY, STORAGE_QUERY_KEY } from "./storage"

export function SharedClipboard() {
  const storageQuery = useQuery({
    queryKey: STORAGE_QUERY_KEY,
    queryFn: () => getStorage(STORAGE_KEY),
  })

  if (storageQuery.isPending) {
    return (
      <Card>
        <CardHeader>
          <CardTitle>Shared clipboard</CardTitle>
          <CardDescription>Loading text from your other devices...</CardDescription>
        </CardHeader>
        <CardContent>
          <Textarea aria-label="Text shared between your devices" rows={7} disabled />
        </CardContent>
      </Card>
    )
  }

  return <SharedClipboardEditor initialValue={storageQuery.data ?? ""} />
}
