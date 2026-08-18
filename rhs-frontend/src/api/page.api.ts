import { pageRouting } from "./api.config"

export type PageInformations = {
  title: string
  description: string
  host: string
}

export async function getPageInfo(): Promise<PageInformations> {
  const response = await fetch(pageRouting.getPageInfoUrl())

  if (!response.ok) {
    throw new Error(`${response.status}`)
  }

  return response.json()
}
