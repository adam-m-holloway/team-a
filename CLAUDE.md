# CLAUDE.md — team-a

## What this is

A small Go HTTP greeting service (`github.com/adam-m-holloway/team-a`).

- `GET /healthz` → `200 {"status":"ok"}`
- `GET /greet?name=X` → `200 {"message":"Hello, X"}`; empty or blank `name` → `400 {"error":"name is required"}`
- `GET /version` → `200 {"version":"dev"}`
- Listens on `PORT` (default `8080`).

## Layout

- `cmd/server/` — entrypoint (`main.go`); builds the `http.Server`.
- `internal/greet/` — `Greet(name) (string, error)`, the domain logic.
- `internal/httpapi/` — routes (Go 1.22+ `ServeMux` method patterns) and JSON responses.
- `docs/adr/` — Architecture Decision Records.
- `.github/workflows/ci.yml` — the `ci` job (required check on `main`).

## Rules

- Go 1.27, **standard library only** — no third-party modules.
- Tests are table-driven and written first; HTTP tests use `net/http/httptest`.
- Wrap errors with `fmt.Errorf("...: %w", err)`.

## Commands

```sh
go test ./...            # tests
go vet ./...             # vet
gofmt -l .               # must print nothing
go run ./cmd/server      # run locally on :8080
```

CI runs the same three checks (`gofmt -l`, `go vet`, `go test`).

## Claude Code plugins

This repo uses the **latest** channel of the `agent-marketplace` marketplace (see `.claude/settings.json`):
`adr-helper` (write ADRs into `docs/adr/`) and `go-conventions`.
