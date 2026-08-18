import { storageRouting } from "./api.config"

export async function getStorage(key: string): Promise<string> {
  const requestUrl = storageRouting.getStorageUrl()
  requestUrl.searchParams.append("key", key)

  const response = await fetch(requestUrl)

  if (!response.ok) {
    throw new Error(`${response.status}`)
  }

  return response.json()
}

export async function setStorage(key: string, value: string): Promise<void> {
  const requestUrl = storageRouting.setStorageUrl()
  requestUrl.searchParams.append("key", key)
  requestUrl.searchParams.append("value", value)

  const response = await fetch(requestUrl, { method: "POST" })

  if (!response.ok) {
    throw new Error(`${response.status}`)
  }
}
