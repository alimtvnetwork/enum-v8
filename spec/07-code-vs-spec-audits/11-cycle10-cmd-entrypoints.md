# Cycle 10 — `01-app/12-cmd-entrypoints.md` Code-vs-Spec Audit

> Audited 2026-05-05. Spec is short (~134 lines) but had a **HIGH-severity contradiction** because it asserted "no `cmd/` directory" in a repo (`enum-v7`) that ships exactly one (`cmd/main/main.go`).

## Method
Read the spec end-to-end, extracted every checkable claim, then verified each against `enum-v7/` source on disk.

## Claim ledger

| # | Claim | Source line | Status (baseline) | Evidence | Status (post-fix) |
|---|-------|-------------|-------------------|----------|-------------------|
| 1 | "No `cmd/` directory" in this repo | §1 line 19,23 | ❌ Contradiction (**C-CVS-10**) | `cmd/main/main.go` exists, `package main`, real `func main()` | ✅ Match — spec rewritten to permit smoke-test harness |
| 2 | "No `main` package" | §1 line 24 | ❌ Contradiction (subsumed by C-CVS-10) | same | ✅ Match |
| 3 | "No produced binary artefacts" | §1 line 25 | ⚠️ Drift (subsumed by C-CVS-10) | `cmd/README.md` says `make` → `bin/main` | ✅ Match — §1 now documents `bin/main` |
| 4 | Rule: "Do not add a `cmd/` directory" | §1 line 33 | ❌ Contradiction (subsumed) | exists | ✅ Match — rule narrowed to "no additional `cmd/<name>/`" |
| 5 | Enforcement is PR-review-only (no CI guard) | §1 line 35 | ✅ Match (baseline) | no guard in `.github/workflows/`, no `pre-commit` hook | ✅ Match |
| 6 | Import path uses `core-v9` | §2 lines 43-45 | ✅ Match | `go.mod` declares `core-v9`; `cmd/main/main.go` imports `core-v9/...` | ✅ Match |
| 7 | `coregeneric.New.Collection.String.Items(...)` exists | §2 line 49 | ❓ Unverifiable | **Cycle 89 AB:** `coredata/coregeneric/vars.go:34` `var New = &newCreator{}`; `Collection` + `SimpleSlice` typed creators with `.String.Cap/.From/.Empty/.Items` documented in `vars.go:29`, `types.go:57`. Note import path is `core-v9/coredata/coregeneric` not bare `coregeneric` → **D-CVS-70 LOW** | ⚠️ Match w/ path drift |
| 8 | `conditional.IfString(...)` exists | §2 line 50 | ❓ Unverifiable | **Cycle 89 AB:** `conditional/typed_string.go:26 func IfString(...)` confirmed | ✅ Match |
| 9 | `errcore.FailedType.Fmt(...)` exists | §2 line 52 | ❓ Unverifiable | **Cycle 89 AB:** no bare `FailedType` const exists. Only specific variants (`MarshallingFailedType`, `ParsingFailedType`, `ValidationFailedType`, `PathRemoveFailedType`, …) at `errcore/RawErrorType.go:86-121`. The `.Fmt(...)` method exists on the underlying `RawErrorType` (`RawErrorType.go:234`). Filed **D-CVS-71 HIGH** | ❌ Fabricated |
| 10 | Cross-references to §04..§10 exist | §2 lines 59-65 | ✅ Match | all 7 files present in `spec/01-app/` | ✅ Match |
| 11 | Tests live at `tests/integratedtests/` | §3 line 71 | ⚠️ Drift (**D-CVS-32**) | actual is `tests/creationtests/` (5th repeat) | ✅→**REVERT (D-CVS-64)**: `tests/integratedtests/` IS canonical upstream; original spec was correct |
| 12 | `go test ./tests/integratedtests/coregenerictests/...` works | §3 line 78 | ⚠️ Drift (**D-CVS-33**, subsumed by D-CVS-32) | path doesn't exist | **REVERT (D-CVS-64)**: `tests/integratedtests/coregenerictests/` exists upstream |
| 13 | Cross-ref `13-testing-patterns.md` exists | §3 line 81 | ✅ Match | file present | ✅ Match |
| 14 | Cross-ref `14-tests-folder-walkthrough.md` exists | §3 line 81 | ✅ Match | file present | ✅ Match |
| 15 | `coregeneric.New.Collection.String.Items` (CLI example) | §4 line 100 | ❓ Unverifiable | **Cycle 89 AB:** same as #7 — exists at `coredata/coregeneric/`, path drift only | ⚠️ Match w/ path drift |
| 16 | `conditional.IfFuncString` exists | §4 line 101 | ❓ Unverifiable | **Cycle 89 AB:** `conditional/typed_string.go:34 func IfFuncString(...)` confirmed | ✅ Match |
| 17 | `args.IsEmpty()` / `args.First()` on Collection | §4 lines 102-104 | ❓ Unverifiable | **Cycle 89 AB:** no top-level `args` package; only `coretests/args` (test holders) with `HasFirst()` (not `First()`). The CLI example most likely means `corestr.Collection` from cmd/main (`HasItems()`, `Items()`). Filed **D-CVS-72 MEDIUM** | ❌ Fabricated/misleading |
| 18 | PowerShell tooling lives at `/spec/04-tooling/` | §5 line 116 | ✅ Match | dir present (`00-overview.md`, `01-…`, `02-…`, `03-…`) | ✅ Match |
| 19 | `03-powershell-implementation.md` exists | §5 line 120 | ✅ Match | present | ✅ Match |
| 20 | `04-bootstrap-into-new-repo.md` exists | §5 line 121 | ⚠️ Drift (latent) | only `00-03` files present in `spec/04-tooling/` | ❓ Deferred — not in scope of cycle 10 (would be D-CVS-35; logged below) |
| 21 | `00-repo-overview.md` exists | See Also | ✅ Match | present | ✅ Match |
| 22 | `/cmd/README.md` exists & describes `bin/main` | (post-fix added) | ✅ Match | verified | ✅ Match |

**Total:** 22 claims · baseline 9 ✅ / 3 ⚠️ / 4 ❌ / 6 ❓ → **56.3 % verifiable**.
**Post-Cycle-10 fix:** 16 ✅ / 0 ⚠️ / 0 ❌ / 6 ❓ → **100 % verifiable**.
**Cycle 89 AB-residual:** 14 ✅ / 4 ⚠️ / 3 ❌ / 1 ❓ (#20 latent) → **21/22 = 95.5 % verifiable**. 3 fabricated/drift findings opened (D-CVS-70/71/72).

## Findings opened & closed in this cycle

### C-CVS-10 — `cmd/` policy contradicts the existing `cmd/main/` harness (HIGH)
- **Spec said:** "No `cmd/` directory. No `main` package." Rule: "Do not add a `cmd/` directory to this module."
- **Reality:** `enum-v7/cmd/main/main.go` exists with `package main` + `func main()`. `cmd/README.md` documents the convention (`make` → `bin/main`).
- **Fix:** Rewrote §1 to a "library-first, smoke-test allowed" policy. Distinguishes upstream `core-v9` (truly zero `cmd/`) from this module (one permitted `cmd/main/` smoke-test harness). Rule narrowed to "no additional `cmd/<name>/` entrypoints + no `cmd/` in upstream `core-v9`". Cross-linked `cmd/README.md`.

### D-CVS-32 — `tests/integratedtests/` reference (5th occurrence)
- **Spec said:** §3 line 71: "integrated test suite at `tests/integratedtests/`".
- **Reality:** `tests/creationtests/` (only sub-folder under `tests/`).
- **Fix:** Rewrote §3 to `tests/creationtests/` and cross-linked the prior C-CVS-01 / D-CVS-17 / D-CVS-26 / D-CVS-27 fixes. This is now the **5th** time we've removed this stale path — the spec corpus should be considered clean of `integratedtests` after this cycle (a final `rg integratedtests` sweep is task **AH** territory).

### D-CVS-33 — stale `go test ./tests/integratedtests/coregenerictests/...` example
- Subsumed by D-CVS-32; same line-86 fix replaces the example with `go test ./tests/creationtests/...`.

### D-CVS-35 — `04-bootstrap-into-new-repo.md` referenced but absent (NEW, deferred)
- Spec §5 cites `/spec/04-tooling/04-bootstrap-into-new-repo.md` but the directory only contains `00..03`. Logged for a future cycle (low severity; documentation-only cross-reference).

## Verifiable subset score

**100.0 %** (16 ✅ / 16 verifiable claims). 6 ❓ (upstream `coregeneric` / `conditional` / `errcore` API surface) deferred to **task AB**. 1 latent drift (D-CVS-35) deferred.

## See also
- [`01-scoreboard.md`](./01-scoreboard.md) — Cycle 10 row + C-CVS-10 / D-CVS-32 entries
- [`/cmd/README.md`](../../cmd/README.md) — the permitted smoke-test harness
- Prior `integratedtests` fixes: C-CVS-01 (cycle 1), D-CVS-17 (cycle 3), D-CVS-26 (cycle 6), D-CVS-27 (cycle 9)

---

## Cycle 89 AB-residual findings

### D-CVS-70 — `coregeneric` import path drift — **LOW**
Spec writes `coregeneric.New.Collection.String.Items(...)`. Real import path is `github.com/alimtvnetwork/core-v9/coredata/coregeneric` (note the `coredata/` prefix). All API names match (`vars.go:34 var New = &newCreator{}`, `Collection[T]` + `SimpleSlice[T]` typed creators with `.String.Cap/.From/.Empty/.Items`). Same drift pattern as D-CVS-54 in Cycle 83. Spec fix: prepend `coredata/` to the import path everywhere `coregeneric` is referenced.

### D-CVS-71 — `errcore.FailedType` does not exist as a bare constant — **HIGH (fabricated symbol)**
Spec §2 line 52 references `errcore.FailedType.Fmt(...)`. Upstream `errcore/RawErrorType.go:86-121` has only **specific** `*FailedType` variants (`MarshallingFailedType`, `ParsingFailedType`, `ValidationFailedType`, `PathRemoveFailedType`, …) — no bare `FailedType`. The `.Fmt(message string, vars ...any) error` method is on the underlying `RawErrorType` (`RawErrorType.go:234`). Spec fix: cite a concrete variant (e.g. `errcore.ValidationFailedType.Fmt(...)`) or refer generically to `errcore.RawErrorType.Fmt`.

### D-CVS-72 — `args.IsEmpty()`/`args.First()` references a non-existent top-level `args` package — **MEDIUM**
Spec §4 lines 102-104 reads as if `args` is a top-level core-v9 package. Reality: there is no `core-v9/args`; only `core-v9/coretests/args` (test-fixture holder types) which expose `HasFirst()` (not `First()`). The `cmd/main` smoke-test example actually uses `os.Args` + `corestr.New.Collection`. Spec fix: replace pseudo-`args.IsEmpty()/First()` with the real `len(os.Args) <= 1` / `os.Args[1]` pattern, or with `corestr.Collection.HasItems()/Items()[0]`.
