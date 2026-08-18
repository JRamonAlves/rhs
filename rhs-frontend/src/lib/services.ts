import {
  BookOpenTextIcon,
  ClosedCaptioningIcon,
  ChartPieSliceIcon,
  DownloadSimpleIcon,
  GameControllerIcon,
  ImagesIcon,
  KanbanIcon,
  LockKeyIcon,
  MagnifyingGlassIcon,
  PlayCircleIcon,
  ProjectorScreenIcon,
  TelevisionIcon,
} from "@phosphor-icons/react"

import type { Service } from "@/api/services.api"

export const appIcons = {
  book: BookOpenTextIcon,
  captions: ClosedCaptioningIcon,
  chart: ChartPieSliceIcon,
  download: DownloadSimpleIcon,
  film: ProjectorScreenIcon,
  game: GameControllerIcon,
  images: ImagesIcon,
  kanban: KanbanIcon,
  lock: LockKeyIcon,
  play: PlayCircleIcon,
  search: MagnifyingGlassIcon,
  tv: TelevisionIcon,
}

export type AppIcon = keyof typeof appIcons

export function isAppIcon(icon: string): icon is AppIcon {
  return Object.hasOwn(appIcons, icon)
}

export type ServiceGroup = {
  category: Service["category"]
  services: Service[]
}

export function groupServicesByCategory(services: Service[]): ServiceGroup[] {
  const groups = new Map<string, Service[]>()

  for (const service of services) {
    const categoryServices = groups.get(service.category) ?? []
    categoryServices.push(service)
    groups.set(service.category, categoryServices)
  }

  return Array.from(groups, ([category, categoryServices]) => ({
    category,
    services: categoryServices,
  }))
}
