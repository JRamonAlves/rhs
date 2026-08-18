import { getServicesCount } from "@/api/services.api"
import { ThemeToggle } from "@/components/theme"
import { HardDriveIcon, HouseLineIcon } from "@phosphor-icons/react"
import { useQuery } from "@tanstack/react-query"

export type PageInformations = {
  title: string
  description: string
  host: string
}

export function PageHeader() {
  const page: PageInformations = {
    title: "Ramon Home Server",
    description: "Home server services",
    host: "hsramon.tailab18b8.ts.net",
  }

  const serviceCountQuery = useQuery({
    queryKey: ["serviceCount"],
    queryFn: getServicesCount,
  })

  return (
    <header className="flex flex-col gap-4 border-b pb-5 sm:flex-row sm:items-end sm:justify-between">
      <div className="flex min-w-0 flex-col gap-3">
        <div className="flex min-w-0 items-center gap-2 text-xs text-muted-foreground">
          <HouseLineIcon className="shrink-0" weight="duotone" />
          <span className="truncate">{page.host}</span>
        </div>
        <div className="flex flex-col gap-2">
          <h1 className="text-2xl font-medium sm:text-3xl">{page.title}</h1>
          <p className="max-w-2xl text-sm text-muted-foreground">
            {page.description}
          </p>
        </div>
      </div>
      <div className="flex items-center gap-2">
        <div className="flex items-center gap-2 border px-3 py-2 text-xs text-muted-foreground">
          <HardDriveIcon weight="duotone" />
          <span>
            {serviceCountQuery.isPending
              ? "Loading services..."
              : serviceCountQuery.isError
                ? "Service count unavailable"
                : `${serviceCountQuery.data} services`}
          </span>
        </div>
        <ThemeToggle />
      </div>
    </header>
  )
}
