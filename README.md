# plex-discovery

## Application purpose

plex-discovery is intended to streamline how users explore their Plex Media Server libraries.
The application will connect to the official Plex API to surface media in a friendlier way than the
default Plex experience. By cataloguing available movies, shows, music, and photos, it will help
users quickly find what to watch, listen to, or share, and highlight new or trending content inside
their own server.

## Tech stack

- **Backend:** Go 1.22, [templ](https://github.com/a-h/templ) for server-rendered UI fragments, Air for
  hot reloading, and SQLite for persistent storage using the `modernc.org/sqlite` driver.
- **Frontend:** Vite + React + TypeScript with Tailwind CSS for styling and ESLint for static analysis.
- **API client:** Generated via [`oapi-codegen`](https://github.com/oapi-codegen/oapi-codegen) from the
  official Plex Media Server OpenAPI specification.

## Project structure

```
├── api/                # OpenAPI specifications and generator configuration
├── cmd/server/         # Application entrypoint
├── internal/           # Private Go packages (HTTP server, Plex client, storage helpers)
├── scripts/            # Utility scripts for local development
├── web/                # Vite + React frontend
└── docs/               # Project documentation (coming soon)
```

## Local development

### Prerequisites

- Go 1.22+
- Node.js 20+
- [Air](https://github.com/air-verse/air) installed in your `$PATH`

### Install dependencies

```bash
go mod tidy
(cd web && npm install)
```

### Generate code

Run templ and the OpenAPI generator when template or specification files change:

```bash
go generate ./...
```

### Start the development environment

```bash
air
```

Air is configured to build the Go server, launch the Vite dev server, and reload both sides as files
change. Alternatively you can run the helper script:

```bash
./scripts/dev.sh
```

### Available documentation

- [Go standard library](https://pkg.go.dev/std)
- [templ documentation](https://templ.guide/)
- [Vite guide](https://vitejs.dev/guide/)
- [Tailwind CSS documentation](https://tailwindcss.com/docs)
- [Plex Media Server OpenAPI specification](https://developer.plex.tv/pms/) *(authentication required)*
