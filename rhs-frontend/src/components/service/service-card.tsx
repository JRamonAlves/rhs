import type { Service } from "@/api/services.api"
import { Card } from "@/components/ui/card"

import { ServiceCardContent } from "./service-card-content"
import { ServiceCardFooter } from "./service-card-footer"
import { ServiceCardHeader } from "./service-card-header"

type ServiceCardProps = {
  app: Service
}

export function ServiceCard({ app }: ServiceCardProps) {
  return (
    <Card size="sm">
      <ServiceCardHeader app={app} />
      <ServiceCardContent app={app} />
      <ServiceCardFooter name={app.name} url={app.url} />
    </Card>
  )
}
