# Repository Guidelines

## Project Structure & Module Organization

This is a Bun-managed Vite React landing page for home server service links.
Application code lives in `src/`. The main render path is `src/main.tsx` ->
`src/App.tsx`, with global Tailwind/shadcn styles in `src/index.css`.

Service data is the primary content source. Update `src/data/apps.json` to add,
remove, or change links. Shared helpers live in `src/lib/`, reusable app
components in `src/components/`, and shadcn UI primitives in
`src/components/ui/`. Root configuration files include `vite.config.ts`,
`tsconfig*.json`, `eslint.config.js`, `components.json`, and `index.html`.

Follow the project clean architecture convention for app components: organize
feature-specific components by responsibility under their own directory, such as
`src/components/header/`, `src/components/service/`, and
`src/components/theme/`. Each feature directory should expose its public API from
an `index.ts` barrel, and callers should import from the directory boundary
(`@/components/service`) instead of reaching into implementation files
(`@/components/service/service-card`). Keep generic shadcn primitives in
`src/components/ui/` and shared non-UI helpers in `src/lib/`.

## Build, Test, and Development Commands

Use Bun for package and script execution:

- `bun install` installs dependencies from `bun.lock`.
- `bun run dev` starts the Vite dev server.
- `bun run build` runs TypeScript project build and creates `dist/`.
- `bun run preview` serves the production build locally.
- `bun run lint` runs ESLint across the repo.
- `bun run typecheck` runs TypeScript without emitting files.
- `bun run format` formats TypeScript and TSX files with Prettier.

Add shadcn components with `bunx --bun shadcn@latest add <component>`.

## Coding Style & Naming Conventions

Use TypeScript, React function components, and path aliases such as
`@/components/ui/button`. Keep JSON service keys camelCase, for example
`httpsPort` and `composeFile`. Use 2-space indentation, double quotes, no
semicolons, and trailing commas where Prettier adds them.

Follow the existing shadcn/Tailwind style: use semantic tokens
(`bg-background`, `text-muted-foreground`) and layout utilities with `gap-*`.
Prefer existing UI primitives before creating custom component styles.

## Testing Guidelines

No test framework is currently configured. For now, validate changes with:

```bash
bun run lint
bun run typecheck
bun run build
```

If tests are added later, place them next to the code they cover or under
`src/__tests__/`, and use descriptive names such as `App.test.tsx`.

## Commit & Pull Request Guidelines

Recent commits use concise Conventional Commit-style messages, such as
`feat: initial commit` and `chore: add agent skill`. Continue that pattern:
`feat: add service card`, `fix: correct service url`, or `docs: update guide`.

Pull requests should describe the user-facing change, list verification commands
run, and include screenshots for UI changes on mobile and desktop widths. When
editing service links, mention the affected entries in `src/data/apps.json`.

## Security & Configuration Tips

Do not commit secrets or private credentials. Service URLs in
`src/data/apps.json` are bundled into the frontend, so treat them as public
navigation targets rather than sensitive configuration.
