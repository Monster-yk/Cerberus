# ADR 0001: Use a modular monorepo

Status: Accepted

## Context

Cerberus has a Go backend with several long-term responsibilities, including webhooks, scanning, scheduling, notifications, storage, and a local CLI.
Starting with multiple repositories or independent microservices would add deployment, dependency, authentication, and development overhead.

## Decision

Cerberus will use one Go module and one repository:

- Several binaries will live under `cmd/`.
- Shared packages will live under `internal/`.
- New functionality will be organized as small, focused packages.
- Packages must remain independent unless they are explicitly designed to
  share functionality.
- The public API for each package will be kept small.

## Consequences

This provides:

- Simple setup and cross-package development.
- One CI pipeline and consistent dependency management.
- Shared types and utilities without publishing separate packages.
- An easier transition to separate services if required later.

It also requires:

- Strong package boundaries to prevent cyclic dependencies.
- Clear ownership between `cmd/` binaries and `internal/` packages.
- Discipline to avoid turning `internal/` into a dumping ground.

The repository may adopt separate services or repositories later, but that decision will require a new ADR.