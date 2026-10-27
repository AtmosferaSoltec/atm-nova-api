# ATM Nova API

Local-only Go + Gin scaffold for Nova. The React frontend lives in the sibling
`atm-nova-web` repository and still uses fictional in-memory data.

This project intentionally contains no endpoints, authentication, database,
migrations, business modules, integrations, or deployment configuration.
Requests currently return Gin's default `404` response.

## Run locally

Use Go 1.27.3, matching ATM Pulso API and ATM Orbita API.

```sh
go mod download
go run ./cmd/api
```

The server listens only on `127.0.0.1:3002`. To use another local port:

```sh
PORT=3003 go run ./cmd/api
```

The container image sets `HOST=0.0.0.0` so Docker can publish the port; the host nginx is the only public entry point.

No `.env` file or database is required. Stop the process with Ctrl+C.

## Checks

```sh
gofmt -l cmd
go vet ./...
go test ./...
go build -o bin/nova-api ./cmd/api
```

There are no application tests yet because no application behavior exists.

## Structure and future direction

The module path, Go/Gin versions, and `cmd/api/main.go` entry point follow the
existing ATM Pulso API and ATM Orbita API projects. Their business code, secrets,
database setup, dependencies beyond Gin, Git history, and deployment automation
are not copied.

When backend implementation is requested, follow the same modular direction:
`internal/modules/<module-name>` for isolated business modules,
`internal/platform` for shared infrastructure, and `internal/integrations` for
external adapters. Create these only when needed. The four specialties are
dentistry, psychology, veterinary medicine, and medicine; dentistry comes first.

## Repository

SSH remote: `git@github.com:AtmosferaSoltec/atm-nova-api.git`.

Development and execution remain local; source changes are versioned on GitHub.
Pushing commits to `main` does not deploy anything. Read [AGENTS.md](AGENTS.md) before changing the project.

## Deployment

Production runs on the VPS at `https://nova-api.atmosferast.com` through the `production` branch, as described in the web repository. The server reads `/opt/nova-api/.env` (not in Git). The scaffold has no health route yet, so the deploy checks that the container answers HTTP on loopback.


## Branches

Use only `main` for development and `production` for production releases.
