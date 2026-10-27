# Nova API — Agent Instructions

These instructions apply throughout this repository.

## Current scope

- This is an intentionally empty, local-only Go + Gin scaffold.
- Keep only the server entry point and the files needed to build and document it.
- Do not add endpoints, health checks, authentication, databases, migrations,
  business behavior, integrations, containers, or deployment workflows until
  explicitly requested.
- The sibling `atm-nova-web` project remains an independent React prototype.
  Do not connect it to this API during the mockup phase.
- Do not deploy or push changes to a remote unless the user requests it.

## Conventions

- Write code, comments, documentation, identifiers, packages, files, directories,
  routes, and future database names entirely in English.
- User-facing product text must be Spanish when UI-facing behavior is introduced.
- Follow idiomatic Go, use explicit types, and handle errors.
- Use `cmd/api/main.go` as the entry point, following ATM Pulso API and ATM Orbita API.
- Retain the versions declared in `go.mod` unless a task requires changing them.
- Keep dependencies small. Do not copy unrelated functionality or dependencies
  from reference projects; deployment resources will be limited.
- Never copy secrets, environment files, private data, or Git history from another project.

## Future architecture

When implementation is authorized, encapsulate business modules under
`internal/modules/<module-name>`, common infrastructure under `internal/platform`,
and external adapters under `internal/integrations`. Do not create empty layers now.

Nova has four independent specialty modules: dentistry, psychology, veterinary
medicine, and medicine. Dentistry is the first priority. A specialty must not
import another specialty. Authentication, users, subscriptions, and organization
identity will be shared concerns. Do not implement them in the current scaffold.

## Verification

For Go code changes, run `gofmt`, `go vet ./...`, `go test ./...`, and
`go build -o bin/nova-api ./cmd/api`. Report the checks actually run. A successful
build or a default 404 proves the scaffold starts, not that business features exist.
