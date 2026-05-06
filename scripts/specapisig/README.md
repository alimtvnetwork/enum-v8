# specapisig — Go AST signature indexer for S-106 v2

Walks one or more Go module roots, parses every non-test `.go` file with
`go/parser`, and emits a JSON signature index of every **exported top-level
function** keyed by `package.Symbol`. Designed to be paired with the
PowerShell driver `scripts/spec-api-sig-check.psm1` (S-106 v2) which
consumes the JSON to validate call-site arity and (best-effort) parameter
kinds against spec snippets.

## Output schema

```json
{
  "version": "1.0.0",
  "generated": "2026-05-06T12:34:56Z",
  "roots": ["/tmp/core-v9-upstream", "."],
  "packages": 259,
  "functions": 4321,
  "items": [
    {
      "package": "errcore",
      "symbol":  "VarTwo",
      "kind":    "func",
      "receiver": "",
      "params": [
        {"name": "isIncludeType", "type": "bool"},
        {"name": "name1",         "type": "string"},
        {"name": "value1",        "type": "interface{}"}
      ],
      "results": [{"name": "", "type": "string"}],
      "variadic": false,
      "file":    "errcore/VarTwo.go",
      "line":    12
    }
  ]
}
```

## Usage

```bash
go run ./scripts/specapisig \
    -roots /tmp/core-v9-upstream,. \
    -out   /tmp/core-v9-sigindex.json
```

Flags:

- `-roots`  comma-separated directory list (upstream clone + project root).
- `-out`    output JSON path (default: `/tmp/core-v9-sigindex.json`).
- `-skip`   regex of dir basenames to skip (default covers `.git`,
            `node_modules`, `vendor`, `cross-repo`, `tests`, `scripts`,
            `spec`, `src`, `public`, `data`, `cmd`, `assets`, `configs`,
            `internal`).

## Scope

- **Top-level funcs only** — methods on types are emitted with `receiver`
  set; method-chain references in spec (`x.New.Line.NotEmpty()`) are best
  validated by S-106 v1.x presence-checks against the receiver type.
- **Exported symbols only** — the spec never references unexported names.
- **No generics in the public API of `core-v9 v1.5.8`**, so generic type
  params are not emitted.

## Exit codes

- `0` — index written successfully.
- `1` — at least one root failed to parse (logged to stderr).