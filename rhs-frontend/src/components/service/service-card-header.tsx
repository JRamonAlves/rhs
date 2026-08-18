import type { Service } from "@/api/services.api"
import {
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { appIcons } from "@/lib/services"

type ServiceCardHeaderProps = {
  app: Service
}

export function ServiceCardHeader({ app }: ServiceCardHeaderProps) {
  const AppIcon = appIcons[app.icon]

  return (
    <CardHeader>
      <CardTitle className="flex min-w-0 items-center gap-2">
        <span className="flex size-8 shrink-0 items-center justify-center border bg-muted text-base text-muted-foreground">
          <AppIcon weight="duotone" />
        </span>
        <span className="truncate">{app.name}</span>
      </CardTitle>
      <CardDescription>{app.description}</CardDescription>
    </CardHeader>
  )
}
