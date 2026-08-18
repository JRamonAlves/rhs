import { WarningCircleIcon } from "@phosphor-icons/react"
import { useQuery } from "@tanstack/react-query"

import { getServices } from "@/api/services.api"
import { SharedClipboard } from "@/components/clipboard"
import { PageHeader } from "@/components/header"
import { ServiceSection } from "@/components/service"
import { groupServicesByCategory } from "@/lib/services"

export function App() {
  const servicesQuery = useQuery({
    queryKey: ["services"],
    queryFn: getServices,
  })

  if (servicesQuery.isPending) {
    return (
      <main className="flex min-h-svh items-center justify-center bg-background px-4 text-foreground">
        <p className="text-sm text-muted-foreground" role="status">
          Loading services...
        </p>
      </main>
    )
  }

  if (servicesQuery.isError) {
    return (
      <main className="flex min-h-svh items-center justify-center bg-background px-4 text-foreground">
        <div className="flex max-w-md gap-3 border border-destructive p-4" role="alert">
          <WarningCircleIcon className="shrink-0 text-destructive" weight="duotone" />
          <div className="flex flex-col gap-1">
            <h1 className="text-sm font-medium">Unable to load services</h1>
            <p className="text-xs text-muted-foreground">
              Check that the server is available, then refresh the page.
            </p>
          </div>
        </div>
      </main>
    )
  }

  const serviceGroups = groupServicesByCategory(servicesQuery.data)

  return (
    <main className="min-h-svh bg-background text-foreground">
      <section className="mx-auto flex w-full max-w-6xl flex-col gap-6 px-4 py-5 sm:px-6 sm:py-8 lg:px-8">
        <PageHeader />
        <SharedClipboard />
        <div className="flex flex-col gap-8">
          {serviceGroups.map((group) => (
            <ServiceSection
              key={group.category}
              category={group.category}
              services={group.services}
            />
          ))}
        </div>
      </section>
    </main>
  )
}

export default App
