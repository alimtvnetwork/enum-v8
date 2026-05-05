# 00 — Repo Overview

> ✅ **Done** — extracted from `spec/00-llm-integration-guide.md` §Module Identity, §Design Philosophy (TL;DR), §Package Map.
> **Status**: filled in audit Step 3 (2026-04-23, Asia/Kuala_Lumpur).
> **Audience**: AI agents and human contributors landing on the repo for the first time.

---

## 1. Module Identity

```
module github.com/alimtvnetwork/core-v9
go 1.25.0
```

> ⚠️ **Package name vs. module path** *(F-NEW-04 fix)*
> Module path ends in `core-v9`. **Root package name in Go code is `core`**, not `corev8`.
> Writing `corev8.X()` will not compile. See [`03-import-conventions.md`](./03-import-conventions.md).

| Property | Value |
|---|---|
| Module path | `github.com/alimtvnetwork/core-v9` |
| Root package name (in code) | **`core`** |
| Go version | `1.25.0` |
| Runtime dependencies | **None** (stdlib only) |
| Test dependencies | `github.com/smarty/assertions`, `github.com/smartystreets/goconvey` |
| Install | `go get github.com/alimtvnetwork/core-v9` |
| License | See repo `LICENSE` |

**Why zero runtime deps?** The library is intended as a foundation for downstream services. A nil dependency footprint means upgrading `core-v9` cannot transitively force upgrades of unrelated packages.

---

## 2. Top-Level Layout (one-liner per top-level entry)

| Path | Purpose |
|---|---|
| `core.go`, `generic.go` | Root package — generic slice/map factories |
| `conditional/` | Ternary helpers (`If[T]`, `IfFunc[T]`, `NilDef[T]`) |
| `constants/` | 400+ named constants (strings, bytes, runes, numbers) |
| `converters/` | String↔bytes, maps, JSON formatting |
| `coredata/` | Data-structures umbrella (see [`01-package-map.md`](./01-package-map.md)) |
| `coreinterface/` | 100+ canonical interface contracts |
| `coreimpl/enumimpl/` | Enum implementation engine |
| `errcore/` | Rich error construction + stack traces |
| `corefuncs/` | Function-type wrappers (`ErrFunc`, `InOutErrFuncWrapper`) |
| `corevalidator/` | Line, slice, text, range validators |
| `coremath/` ⚠️ **LEGACY** `coresort/` `corecmp/` `coreappend/` `coreunique/` | Numeric & collection utility packages. *(`coremath` is LEGACY — new code MUST use Go 1.21+ built-in `min` / `max` / `clear`. See `00-llm-integration-guide.md` §coremath. F-V12-02 fix.)* |
| `isany/` | Type-checking predicates (`Null`, `Zero`, `DeepEqual`) |
| `issetter/` | Six-valued boolean (`Uninitialized/True/False/Unset/Set/Wildcard`) |
| `bytetype/`, `namevalue/`, `keymk/` | Small typed-value helpers |
| `regexnew/` | Lazy-compiled regex with thread-safe caching |
| `chmodhelper/`, `filemode/`, `ostype/`, `osconsts/` | OS / file-system helpers |
| `typesconv/`, `reflectcore/` | Reflection & pointer↔value conversions |
| `mutexbykey/` | Per-key mutex locking |
| `defaultcapacity/`, `defaulterr/` | Default constants & error types |
| `pagingutil/`, `coreversion/` | Pagination & semantic versioning |
| `coretests/` | In-tree testing framework (see [`13-testing-patterns.md`](./13-testing-patterns.md)) |
| `enums/` | Domain enum packages (e.g. `stringcompareas`, `versionindexes`) |
| `cmd/` | Binary entrypoints (see [`12-cmd-entrypoints.md`](./12-cmd-entrypoints.md)) |
| `tests/` | All `*tests` integration packages + `testwrappers/` (see [`14-tests-folder-walkthrough.md`](./14-tests-folder-walkthrough.md)) |
| `internal/` | Module-private helpers — **do not import externally** |
| `spec/` | This documentation tree |
| `scripts/`, `run.ps1` | PowerShell test/build toolchain (see [`/spec/04-tooling/`](../04-tooling/) and [`/spec/03-powershell-test-run/`](../03-powershell-test-run/)) |

For the **complete** package list with one-line purpose for every leaf package, see [`01-package-map.md`](./01-package-map.md).

---

## 3. Design Pillars (TL;DR)

The full rationale lives in [`02-design-philosophy.md`](./02-design-philosophy.md). The seven pillars in one line each:

1. **One file per function** — each public function in its own `.go` file, named after the function (50–200 lines).
2. **Struct-as-namespace** — related operations grouped on unexported struct types exposed via package-level `var` (e.g. `corejson.Serialize.ToString()`).
3. **Interface-first** — contracts in `coreinterface/` using Go's `-er` suffix (`NameGetter`, `Serializer`).
4. **Zero-nil safety** — return empty slices/maps instead of `nil`; pointer-receiver methods include nil guards.
5. **Generics where clear** — generic versions alongside backward-compatible typed wrappers.
6. **Value receivers by default** — pointer receivers only for mutation, large structs, or interface satisfaction.
7. **`newCreator` pattern** — factories exposed via `New` package variable (e.g. `enumimpl.New.BasicByte.UsingTypeSlice(...)`).

---

## 4. Where to Go Next

| You are… | Read… |
|---|---|
| An AI writing **new tests** for an existing package | [`13-testing-patterns.md`](./13-testing-patterns.md) → [`/spec/06-testing-guidelines/`](../06-testing-guidelines/) → [`/spec/05-failing-tests/`](../05-failing-tests/) (skim every file) |
| An AI writing **new library code** | [`02-design-philosophy.md`](./02-design-philosophy.md) → [`03-import-conventions.md`](./03-import-conventions.md) → relevant topic file (04..11) |
| Setting up the **toolchain** in a fresh repo | [`/spec/04-tooling/`](../04-tooling/) (planned bootstrap doc in audit Step 9) |
| Running tests / coverage | [`/spec/03-powershell-test-run/01-overview.md`](../03-powershell-test-run/01-overview.md) |
| Triaging open architectural issues | [`/spec/02-app-issues/`](../02-app-issues/) |
| Auditing spec completeness | [`/spec/99-audits/01-original-11-step-plan.md`](../99-audits/01-original-11-step-plan.md) |
