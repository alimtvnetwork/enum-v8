# 12 — CMD Entrypoints

> ✅ **Status**: filled in audit Step 5c (2026-04-23, Asia/Kuala_Lumpur).
> **Audience**: anyone looking for a `main` function or a CLI binary in this module.
> **Short answer**: **there isn't one** — this module is a pure library.

---

## Table of Contents

1. [No `cmd/` Directory — On Purpose](#1-no-cmd-directory--on-purpose)
2. [How to Use the Library](#2-how-to-use-the-library)
3. [Where to Find Runnable Examples](#3-where-to-find-runnable-examples)
4. [If You Need a CLI](#4-if-you-need-a-cli)
5. [Tooling vs Entrypoints](#5-tooling-vs-entrypoints)

---

## 1. No `cmd/` Directory — On Purpose

The `core-v9` module is a **pure Go library**. It has:

- No `cmd/` directory.
- No `main` package.
- No produced binary artefacts.

This is intentional. The module's role is to provide reusable primitives (data structures, enums, errors, validators, converters, reflection helpers) that downstream applications import. Putting `main` packages here would:

- Add deps that consumers don't need.
- Encourage tight coupling between the library and a specific deployment shape.
- Conflate the package map (see [`01-package-map.md`](./01-package-map.md)) with binary inventories.

> **Rule**: Do not add a `cmd/` directory to this module. New entrypoints belong in the consuming application's repo.
>
> **Enforcement** *(F-V14-04 fix)*: this rule is **PR-review-enforced**, not machine-enforced — there is no `go vet` lint, no CI check, and no PowerShell guard in `spec/04-tooling/` that blocks a `cmd/` directory from being added. If you need machine enforcement, open a ticket in [`spec/02-app-issues/`](../02-app-issues/) describing the desired guard (suggested implementations: a `pre-commit` shell check, or a custom `go vet` analyzer). Until such a ticket lands, code review is the single line of defence.

---

## 2. How to Use the Library

```go
import (
    "github.com/alimtvnetwork/core-v9/coredata/coregeneric"
    "github.com/alimtvnetwork/core-v9/conditional"
    "github.com/alimtvnetwork/core-v9/errcore"
)

func main() {
    col := coregeneric.New.Collection.String.Items("a", "b", "c")
    label := conditional.IfString(col.IsEmpty(), "empty", "non-empty")
    if err := doWork(); err != nil {
        return errcore.FailedType.Fmt("work failed: %v", err)
    }
}
```

See the package-by-package usage notes in:

- [`04-error-system.md`](./04-error-system.md)
- [`05-enum-system.md`](./05-enum-system.md)
- [`06-data-structures.md`](./06-data-structures.md)
- [`07-conditional-and-utilities.md`](./07-conditional-and-utilities.md)
- [`08-validators.md`](./08-validators.md)
- [`09-converters.md`](./09-converters.md)
- [`10-reflection-and-dynamic.md`](./10-reflection-and-dynamic.md)

---

## 3. Where to Find Runnable Examples

The closest thing to executable code in this module is the **integrated test suite** at `tests/integratedtests/`. Every package has `*_test.go` files demonstrating idiomatic use.

```bash
# Run all tests
go test ./...

# Run a specific package's tests
go test ./tests/integratedtests/coregenerictests/...
```

See [`14-tests-folder-walkthrough.md`](./14-tests-folder-walkthrough.md) for the layout and [`13-testing-patterns.md`](./13-testing-patterns.md) for how the tests are structured.

---

## 4. If You Need a CLI

Build it in your own repository as a thin wrapper:

```go
// myapp/cmd/myapp/main.go
package main

import (
    "os"
    "github.com/alimtvnetwork/core-v9/coredata/coregeneric"
    "github.com/alimtvnetwork/core-v9/conditional"
)

func main() {
    args := coregeneric.New.Collection.String.Items(os.Args[1:]...)
    cmd  := conditional.IfFuncString(
        args.IsEmpty(),
        func() string { return "help" },
        func() string { return args.First() },
    )
    dispatch(cmd, args)
}
```

This keeps the library import-only and lets each consumer choose its own CLI framework, flag parser, and configuration loader.

---

## 5. Tooling vs Entrypoints

The repo does include **PowerShell scripts** under [`/spec/04-tooling/`](../04-tooling/) (test runner, repo bootstrap). These are *tooling*, not Go entrypoints:

| Asset | Type | Location | Purpose |
|---|---|---|---|
| PowerShell test runner | `.ps1` script | `/spec/04-tooling/03-powershell-implementation.md` | Runs `go test`, parses output, attributes failures |
| Repo bootstrap | `.ps1` script | `/spec/04-tooling/04-bootstrap-into-new-repo.md` | Scaffolds a new consumer repo |
| Go test binaries | implicit | produced by `go test` | One per `_test.go` package |

None of these are `main` packages in this module. The PowerShell scripts are external orchestration, and the `go test` binaries are produced by the toolchain on demand.

---

## See Also

- [`00-repo-overview.md`](./00-repo-overview.md) — Module identity and top-level layout
- [`01-package-map.md`](./01-package-map.md) — Every package, none of which are `main`
- [`14-tests-folder-walkthrough.md`](./14-tests-folder-walkthrough.md) — Runnable code lives in `tests/integratedtests/`
- [`/spec/04-tooling/`](../04-tooling/) — PowerShell scripts (external tooling, not Go entrypoints)
- [`/spec/00-llm-integration-guide.md` §Quick Start](../00-llm-integration-guide.md) — Example library usage
