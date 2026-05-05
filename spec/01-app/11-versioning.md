# 11 — Versioning

> ✅ **Status**: filled in audit Step 5c (2026-04-23, Asia/Kuala_Lumpur).
> **Audience**: maintainers cutting releases and any AI agent making code changes that require a version bump.
> **Source**: framework convention + user preference (`.lovable/user-preferences` line 8) + `coreversion` package.

---

## Table of Contents

1. [`coreversion` Package](#1-coreversion-package)
2. [`versionindexes`](#2-versionindexes)
3. [Release / Bump Policy](#3-release--bump-policy)
4. [Compatibility Guarantees](#4-compatibility-guarantees)
5. [`.release/` Folder](#5-release-folder)
6. [Common Mistakes](#6-common-mistakes)

---

## 1. `coreversion` Package

Located at `coreversion/`. Provides typed semver primitives — parse, compare, and constrain Go module versions.

```go
import "github.com/alimtvnetwork/core-v9/coreversion"

// Parse
v, err := coreversion.Parse("v1.2.3")
v.Major()  // 1
v.Minor()  // 2
v.Patch()  // 3

// Compare
v1.LessThan(v2)
v1.Equal(v2)
v1.GreaterThanOrEqual(v2)

// Format
v.String() // "v1.2.3"
```

### Why not just `golang.org/x/mod/semver`?

- Wraps stdlib errors in `errcore.FailedToConvertType` so the runner can attribute version-parse failures.
- Typed `Major() / Minor() / Patch()` accessors avoid string slicing at every call site.
- Plays well with `coregeneric.Collection` for sorting version lists.

---

## 2. `versionindexes`

Located at `versionindexes/`. Stable integer indices that identify "version eras" without relying on string comparison.

```go
import "github.com/alimtvnetwork/core-v9/versionindexes"

versionindexes.V1   // 1
versionindexes.V2   // 2
versionindexes.V8   // 8 (current era — core-v9)
```

### Why integer indices?

- Switch statements are exhaustive and compile-checked.
- Migration code can dispatch on era (`switch idx { case V7: ... case V8: ... }`).
- Avoids accidental dependency on patch-level changes inside the same era.

### When to use

- Migration adapters that handle multiple historical formats.
- Test cases that need to opt into era-specific behaviour ([`coretestcases.CaseV1`](./13-testing-patterns.md) takes its name from this).
- Wire-format envelopes that record which era produced them.

---

## 3. Release / Bump Policy

> **CRITICAL** (per `.lovable/user-preferences` line 8): **Code changes must bump at least minor version. Never touch the `.release/` folder.**

### Semver mapping

| Change category | Bump | Examples |
|---|---|---|
| Breaking API change | **major** | Removing a public function, changing a struct's exported fields, renaming a package |
| New feature, no breakage | **minor** | New public function, new method, new package, new enum value |
| Bug fix, no API change | **minor** (per project rule — overrides standard semver "patch") | Logic fix, perf improvement, internal refactor |
| Docs / spec only | **none** required, but still **minor** if any code touched | Spec edits in `spec/`, README updates |

> **Project-specific override**: Standard semver allows patch bumps for bugfixes. This project requires **at least minor** for any code change. Patch is reserved for hotfixes that are explicitly tagged by a maintainer.

### Bump checklist

1. Identify category (table above).
2. Update `coreversion`-related constants if any.
3. Update `go.mod` major version path **only on a major bump** (e.g. `core-v9` → `core-v9`).
4. Document the bump in the release notes (if a release process exists).
5. **Do not edit `.release/`** — that folder is owned by the release pipeline.

---

## 4. Compatibility Guarantees

### Within an era (e.g. `core-v9.x.y`)

- **Public APIs are stable**: imports from `github.com/alimtvnetwork/core-v9/<pkg>` will not break.
- **`internal/` packages are not stable**: they may change in any minor release.
- **Error categories** (`errcore.*Type`) are stable; new categories may be added (additive, non-breaking).
- **Diagnostic message formats** are stable when consumed by tests in `tests/integratedtests/` (see [`08-validators.md` §5](./08-validators.md#5-diagnostic-output-contract)).

### Across eras (`v8` → `v9`)

- The module path changes (`core-v9` → `core-v9`). Old code keeps working unchanged.
- No automatic migration — consumers update import paths when they choose to upgrade.
- A migration guide should accompany every era bump.

---

## 5. `.release/` Folder

> 🚫 **OFF-LIMITS** to all AI agents and contributors.

This folder is managed by the release tooling. It contains:
- Generated changelogs
- Tag metadata
- Build artefacts referenced by CI

**Rules**:
1. Never create files in `.release/`.
2. Never modify files in `.release/`.
3. Never delete files in `.release/`.
4. If a tool requires changes there, escalate to a human maintainer.

This rule is enforced via `.lovable/user-preferences` line 8 and is part of the project's core memory ([`mem://index.md`](mem://index.md) Core).

---

## 6. Common Mistakes

| Mistake | Why bad | Fix |
|---|---|---|
| Bumping patch instead of minor for a bugfix | Violates project policy | Bump minor |
| Forgetting to bump after a code change | Downstream consumers can't pin a fix | Always bump on code touch |
| Editing `.release/` to "help" | Corrupts release pipeline state | Never touch `.release/`; escalate to maintainer |
| Adding a new enum value without bumping | Even additive changes require a bump | Minor bump |
| Removing an exported function in a minor release | Breaks consumers within the same era | Major bump (new era) |
| Using `golang.org/x/mod/semver` directly | Loses `errcore` categorisation | Use `coreversion.Parse` |
| Relying on a specific patch number in tests | Breaks on every release | Use `versionindexes.V<N>` for era-checks |

---

## See Also

- [`03-import-conventions.md`](./03-import-conventions.md) — Module path includes the major version (`core-v9`)
- [`13-testing-patterns.md`](./13-testing-patterns.md) — `coretestcases.CaseV1` naming reflects the version-index pattern
- [`/spec/04-tooling/`](../04-tooling/) — Release tooling (PowerShell runner) lives here
- `.lovable/user-preferences` — Project rules including the version-bump and `.release/` policy
