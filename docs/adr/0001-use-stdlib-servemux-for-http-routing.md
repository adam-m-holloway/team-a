# Use the standard library ServeMux for HTTP routing

- Status: Accepted
- Date: 2026-11-08
- Deciders: team-a

## Status

Accepted.

## Context and Problem Statement

The service exposes two endpoints, `GET /healthz` and `GET /greet`, registered in
`internal/httpapi/httpapi.go` (`NewHandler`) and served by the `http.Server` built in
`cmd/server/main.go`. We want the module to have zero third-party dependencies: `go.mod`
declares no `require` entries. Since Go 1.22, `net/http`'s `ServeMux` supports method
and path patterns such as `"GET /healthz"`, which used to be the main reason to reach for
a third-party router. Which router should the service use?

## Decision Drivers

- Zero third-party dependencies.
- The routing needs are small: two fixed-path endpoints.
- Method-based routing without hand-written method checks in handlers.

## Considered Options

- Standard library `net/http` ServeMux
- chi
- gorilla/mux

## Decision Outcome

Chosen option: "Standard library `net/http` ServeMux", because the service has only two
endpoints, we want zero dependencies, and the standard library now supports method and
path patterns, so a third-party router adds nothing we need.

## Consequences

- Good, because `go.mod` stays free of third-party modules, so there is nothing extra to
  update, audit or vet.
- Good, because routes are plain `mux.HandleFunc("GET /path", ...)` calls that any Go
  developer can read, and handlers stay ordinary `http.HandlerFunc`s.
- Bad, because ServeMux has no built-in middleware chaining or route groups; if we need
  those (logging, auth, sub-routers) we will have to write small wrappers ourselves.
- Bad, because the service requires Go 1.22 or later for method patterns; on older
  toolchains the patterns would not be recognised.

## Pros and Cons of the Options

### Standard library `net/http` ServeMux

- Good, because it needs no dependency and ships with the toolchain.
- Good, because it supports method and path-wildcard patterns since Go 1.22.
- Bad, because it offers no middleware or grouping helpers.

### chi

- Good, because it is built on `net/http` types and provides middleware and route
  grouping.
- Bad, because it adds a third-party dependency, which breaks the zero-dependency goal.

### gorilla/mux

- Good, because it is mature and supports rich matching (regex paths, hosts, headers).
- Bad, because it adds a third-party dependency, which breaks the zero-dependency goal.
- Bad, because its extra matching features are far beyond what two endpoints need.

## More Information

Revisit this decision if the service grows to need middleware stacks, route groups or
matching that ServeMux patterns cannot express.
