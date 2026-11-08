# team-a

A small Go HTTP greeting service, built with the Go standard library only. It is one of the two
consumer repos for the [agent-marketplace](https://github.com/adam-m-holloway/agent-marketplace) POC.

| Endpoint            | Response                                   |
| ------------------- | ------------------------------------------ |
| `GET /healthz`      | `200 {"status":"ok"}`                      |
| `GET /greet?name=X` | `200 {"message":"Hello, X"}`               |
| `GET /greet`        | `400 {"error":"name is required"}` (empty or blank name) |

## Run

```sh
go run ./cmd/server               # listens on :8080
PORT=9000 go run ./cmd/server     # custom port
curl 'localhost:8080/greet?name=Ada'
```

## Test

Requires Go 1.27.

```sh
go test ./...
go vet ./...
gofmt -l .        # must print nothing
```

CI (`.github/workflows/ci.yml`, job `ci`) runs the same checks on every pull request and on pushes to `main`.

## Claude Code plugins

This repo is on the **latest** channel: the `agent-marketplace` marketplace, which tracks `main` of
[adam-m-holloway/agent-marketplace](https://github.com/adam-m-holloway/agent-marketplace).
It is configured in [`.claude/settings.json`](.claude/settings.json) and enables:

- `adr-helper` — writes Architecture Decision Records into `docs/adr/`.
- `go-conventions` — the team's Go conventions.

Browse every plugin in the [catalogue](https://adam-m-holloway.github.io/agent-marketplace/).

### Getting the plugins

Open the repo in Claude Code and trust the folder. You will be prompted to add the marketplace and
install the enabled plugins. Or install one by hand:

```sh
claude plugin install adr-helper@agent-marketplace --scope project
claude plugin install go-conventions@agent-marketplace --scope project
```

### Updating

```sh
claude plugin marketplace update agent-marketplace
claude plugin update adr-helper@agent-marketplace
claude plugin update go-conventions@agent-marketplace
```
