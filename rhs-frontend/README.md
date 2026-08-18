# RHS Land

A small Vite landing page for opening home server services.

## Updating services

Edit `src/data/apps.json` to change the landing page links. The page reads from
that file at build time.

To add a service, append an item to the `apps` array:

```json
{
  "id": "service-id",
  "name": "Service Name",
  "description": "Short service description.",
  "category": "Group",
  "icon": "kanban",
  "url": "https://example.local:443",
  "httpPort": 3000,
  "httpsPort": 3443,
  "composeFile": "./service/docker-compose.yml"
}
```

Available icon values are `kanban`, `images`, `chart`, `lock`, `game`, and
`play`.

## Project architecture

Keep app components organized by feature responsibility. Feature-specific
components live in directories such as `src/components/header/`,
`src/components/service/`, and `src/components/theme/`.

Each feature directory should export its public component API from `index.ts`.
Import from the directory boundary:

```tsx
import { ServiceCard } from "@/components/service"
```

Avoid importing feature implementation files directly:

```tsx
import { ServiceCard } from "@/components/service/service-card"
```

Generic shadcn primitives stay in `src/components/ui/`, shared helpers stay in
`src/lib/`, and service content stays in `src/data/apps.json`.

## Adding components

To add components to your app, run the following command:

```bash
bunx --bun shadcn@latest add button
```

This will place the ui components in the `src/components` directory.

## Using components

To use the components in your app, import them as follows:

```tsx
import { Button } from "@/components/ui/button"
```
