# RHS Monorepo

RHS is a home-server dashboard with a React frontend and a Go API. The two
applications previously lived in separate repositories and are now maintained
together in this monorepo.

The merge puts the user interface and its API contract in one Git history so a
feature can update both sides in the same change. Each application remains an
independent build and deployment unit: there is no root package manager or
shared workspace configuration.

## Repository layout

```text
.
|-- rhs-frontend/   React 19, TypeScript, Vite, Bun, and nginx
`-- rhs-backend/    Go, Gin, Redis, and Swagger
```

- `rhs-frontend` renders the service dashboard and shared clipboard.
- `rhs-backend` serves the service catalog from `data.json` and stores shared
  clipboard values in Redis.
- Dependencies, lockfiles, Dockerfiles, Compose files, tests, and CI workflows
  remain inside their respective project directories.

## How the projects connect

During local development, the frontend runs on Vite's default address and calls
the backend at `http://localhost:8080`. The API URL is defined in
`rhs-frontend/src/api/api.config.ts`.

The projects integrate through these backend routes:

| Route | Purpose |
| --- | --- |
| `GET /getServices` | Load the home-server service catalog |
| `GET /countServices` | Return the number of configured services |
| `GET /getValues?key=...` | Read a shared clipboard value from Redis |
| `POST /setValues?key=...&value=...` | Store a shared clipboard value in Redis |
| `GET /ping` | Check backend health |

Swagger documentation is available at `http://localhost:8080/docs/index.html`
while the backend is running.

## Prerequisites

- [Bun](https://bun.sh/) 1.3.13
- [Go](https://go.dev/) 1.26.5
- Docker with Docker Compose, for Redis

## Run locally

Start Redis:

```bash
cd rhs-backend
docker compose up -d
```

Configure and start the backend:

```bash
cd rhs-backend
cp .env.example .env
go mod download
go run .
```

In another terminal, install the frontend dependencies and start Vite:

```bash
cd rhs-frontend
bun install
bun run dev
```

The backend listens on `http://localhost:8080`. Open the URL printed by Vite,
normally `http://localhost:5173`, to use the dashboard.

## Configuration

The backend reads environment variables from `rhs-backend/.env`:

| Variable | Description | Local value |
| --- | --- | --- |
| `SERVICE_PATH` | Path to the JSON service catalog, relative to the backend directory | `./data.json` |
| `REDIS_PASSWORD` | Optional Redis password | empty |

Update `rhs-backend/data.json` to change the services displayed by the
frontend. The frontend API addresses for development and production are defined
in `rhs-frontend/src/api/api.config.ts`.

## Validate changes

Run checks from the project that changed.

Backend:

```bash
cd rhs-backend
go test ./...
```

Frontend:

```bash
cd rhs-frontend
bun run lint
bun run typecheck
bun run build
```

When a change affects the API contract, update and validate both projects in
the same branch and pull request.

## Deployment

Both projects keep their existing deployment boundaries after the merge.
`rhs-frontend/Dockerfile` builds static assets and serves them with nginx, while
`rhs-backend/Dockerfile` builds the Go API binary. Their existing CI workflows
also remain under each project's `.github/workflows` directory until a shared
root-level pipeline is introduced.

See `rhs-frontend/README.md` for frontend-specific architecture and component
guidance.
