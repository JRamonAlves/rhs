import type { Service } from "@/api/services.api"

import { ServiceCard } from "./service-card"

type ServiceSectionProps = {
  category: string
  services: Service[]
}

export function ServiceSection({ category, services }: ServiceSectionProps) {
  const sectionId = `${category.toLowerCase().replace(/\s+/g, "-")}-services`
  const serviceCountLabel =
    services.length === 1 ? "1 service" : `${services.length} services`

  return (
    <section className="flex flex-col gap-3" aria-labelledby={sectionId}>
      <div className="flex items-center justify-between gap-3 border-b pb-2">
        <h2 id={sectionId} className="text-sm font-medium">
          {category}
        </h2>
        <span className="shrink-0 text-xs text-muted-foreground">
          {serviceCountLabel}
        </span>
      </div>
      <div className="grid gap-3 sm:grid-cols-2 xl:grid-cols-4">
        {services.map((app) => (
          <ServiceCard key={app.name} app={app} />
        ))}
      </div>
    </section>
  )
}
