# Cycle 24 — AB Pass 6: `spec/01-app/15-observability.md`

> **Date:** 2026-05-06 (Asia/Kuala_Lumpur)
> **Audit type:** AB promotion pass — convert ❓ unverifiable claims into ✅/⚠️/❌ via direct comparison against upstream `core-v9 v1.5.8` (clone at `/tmp/core-v9-upstream`, tag `v1.5.8`, module path `github.com/alimtvnetwork/core-v9`).
> **Spec freeze status:** `spec/01-app/` remains 🧊 **DRIFT-FROZEN** (spec-v0.30.0). All proposed rewrites spawned as **AJ-28..34** and BLOCKED until the freeze is lifted.
> **Predecessor:** Cycle 13 (baseline, `14-cycle13-observability.md`) — 14 ✅ / 13 ❓.

## 1. Promotion summary

| Metric | Cycle 13 (baseline) | Cycle 24 (AB pass) |
|---|---|---|
| ✅ Match | 14 | 14 + **6** = **20** |
| ❌ Contradiction | 0 | **7** |
| ❓ Unverifiable | 13 | **0** |
| **Verifiable score** | 100.0 % (small subset) | **74.1 %** |

Of the 13 ❓ promoted: **6 ✅** confirmed, **7 ❌** new contradictions surfaced (C-CVS-44 .. C-CVS-50). Section is now the **second-cleanest of the AB-audited five** (after §07 at 70.6 %), but still drops from a clean baseline due to fabricated function signatures and a fabricated test-failure output format.

## 2. Evidence sources

- `errcore/VarTwo.go`, `errcore/VarTwoNoType.go`, `errcore/MessageVarMap.go` — function signatures
- `errcore/consts.go` — format strings (`var2WithTypeFormat`, `var2Format`, `messageMapFormat`)
- `errcore/stackTraceEnhance.go` — `StackEnhance.{Error, ErrorSkip, MsgToErrSkip, FmtSkip, Msg, MsgSkip, MsgErrorSkip, MsgErrorToErrSkip}`
- `errcore/vars.go` — `StackEnhance = stackTraceEnhance{}`
- `errcore/HandleErr.go` — actual implementation
- `coredata/corejson/NewPtr.go`, `coredata/corejson/Result.go` — `NewPtr(any) *Result`, `(*Result).PrettyJsonString() string`
- `coretests/results/` — directory listing: `Result.go`, `ResultAssert.go`, `Results.go` (no `ResultAny.go`)

## 3. Promotions (❓ → ground truth)

### ✅ Confirmed (6)

| # | Spec claim | Evidence |
|---|---|---|
| P1 | `errcore.VarTwoNoType(name1, val1, name2, val2) string` exists | `errcore/VarTwoNoType.go:25` — exact 4-arg signature |
| P2 | `errcore.MessageVarMap(message string, m map[string]any) string` exists | `errcore/MessageVarMap.go:27` |
| P3 | `errcore.StackEnhance.Error(err error) error` exists | `errcore/stackTraceEnhance.go:36` |
| P4 | `errcore.StackEnhance.Msg(msg string) string` exists | `errcore/stackTraceEnhance.go:68` |
| P5 | `corejson.NewPtr(any) *Result` exists with `.PrettyJsonString() string` method | `coredata/corejson/NewPtr.go:32` + `Result.go:193` |
| P6 | The `coretests/results/Result.go` file exists | `coretests/results/Result.go` present |

### ❌ New contradictions (7)

#### C-CVS-44 *(CRITICAL)* — `errcore.VarTwo` signature is wrong: missing leading `isIncludeType bool` parameter

- **Spec (§2.1, line 49):** `err := errcore.VarTwo("userID", uid, "tenantID", tid)` — 4 args.
- **Reality (`errcore/VarTwo.go:29`):** `func VarTwo(isIncludeType bool, firstName string, firstValue any, secondName string, secondValue any) string` — **5 args**, with a mandatory leading `bool` toggle that the spec example omits entirely. The spec example will not compile.
- **Severity:** CRITICAL. Anyone copy-pasting the example gets a build error.
- **Spawned:** AJ-28 — fix §2.1 example to `errcore.VarTwo(true, "userID", uid, "tenantID", tid)`.

#### C-CVS-45 *(CRITICAL)* — `VarTwo` / `VarTwoNoType` / `MessageVarMap` return `string`, not `error`

- **Spec (§2.1, §2.2, §2.3):** All three examples assign the return value to a variable named `err` (`err := errcore.VarTwo(...)`, etc.).
- **Reality:** All three functions return **`string`** (see `errcore/VarTwo.go:34`, `VarTwoNoType.go:30`, `MessageVarMap.go:30`). They produce diagnostic *messages*, not errors. The whole "Diagnostic Context Primitives" framing as error builders is wrong.
- **Severity:** CRITICAL. Misrepresents the entire helper family's purpose.
- **Spawned:** AJ-29 — rewrite §2 framing as "diagnostic message builders" and rename example variables (`msg := ...`).

#### C-CVS-46 *(HIGH)* — `VarTwoNoType` is a thin wrapper, not a distinct helper

- **Spec (§2.4 decision table):** "2 + types matter → `VarTwo` / 2 + types obvious → `VarTwoNoType`" — implies a hard fork.
- **Reality:** `VarTwoNoType` is literally `VarTwo(false, ...)` (`errcore/VarTwoNoType.go:31-34`). The only API axis is the `isIncludeType` boolean.
- **Severity:** HIGH. The §2.4 table presents two helpers as semantically distinct when they are toggle aliases.
- **Spawned:** AJ-30 — collapse §2.1 + §2.2 into a single helper section noting the `isIncludeType` toggle.

#### C-CVS-47 *(HIGH)* — `coretests/results/ResultAny.go` does not exist

- **Spec (§4, line 111):** "`coretests/results/Result.go` and `ResultAny.go` emit failures in a fixed shape…"
- **Reality:** `ls coretests/results/` returns: `Result.go`, `ResultAssert.go`, `Results.go`. **No `ResultAny.go`.**
- **Severity:** HIGH. Fabricated file path.
- **Spawned:** AJ-31 — replace `ResultAny.go` with the actual companion file (likely `ResultAssert.go` or `Results.go`).

#### C-CVS-48 *(CRITICAL)* — Test-failure output format is fabricated

- **Spec (§4, lines 113-117):** Three-line block with `Test #N — {scenario}: should be equal` followed by indented `expected:` / `actual:`.
- **Reality:** `grep -rn "should be equal\|expected:\|actual:" coretests/results/` returns **zero matches**. No file in `coretests/results/` emits this format. The em-dash framing, the "Test #N" prefix, and the two-space-indented `expected:` / `actual:` lines are not present in source.
- **Severity:** CRITICAL. The "machine-parseable shape" the PowerShell runner allegedly relies on does not exist as documented.
- **Spawned:** AJ-32 — either (a) extract the real failure format from `Result.IsResult` / `Results.go` and substitute, or (b) delete §4 and forward to `06-testing-guidelines/07-diagnostics-output-standards.md` only.

#### C-CVS-49 *(CRITICAL)* — `errcore.HandleErr` does NOT attach stack-enhanced wrapping

- **Spec (§3, rule 2, line 104):** "`*Must` methods … route through `errcore.HandleErr` which already attaches a stack-enhanced wrapping (see `04-error-system.md` §1)."
- **Reality (`errcore/HandleErr.go:25`):** `func HandleErr(err error) { if err == nil { return }; panic(err.Error()) }`. `HandleErr` just panics with the bare `err.Error()` string. **No `StackEnhance` call, no wrapping, no stack capture at this layer.** Any stack enhancement happens elsewhere (e.g. `RawErrorType.MustEnhanceStack` paths) and is not a property of `HandleErr` itself.
- **Severity:** CRITICAL. The rule "do not call `StackEnhance` inside `*Must`" is *correct advice* but for the *wrong reason* — and the wrong reason will mislead anyone debugging missing stack frames in `*Must` panics.
- **Spawned:** AJ-33 — rewrite §3 rule 2 to cite the actual stack-enhancement path (`stackTraceEnhance.MsgToErrSkip` callers in `RawErrorType.go`), not `HandleErr`.

#### C-CVS-50 *(MEDIUM)* — `StackEnhance` surface is larger than spec admits

- **Spec (§3):** Documents only `StackEnhance.Error` and `StackEnhance.Msg`.
- **Reality (`errcore/stackTraceEnhance.go`):** 8 public methods — `Error`, `ErrorSkip`, `MsgToErrSkip`, `FmtSkip`, `Msg`, `MsgSkip`, `MsgErrorSkip`, `MsgErrorToErrSkip`. The `*Skip` variants are the actual API used internally (24 call-sites in `errcore/RawErrCollection.go` + `RawErrorType.go`); the no-skip versions documented in the spec are the rarely-used convenience wrappers.
- **Severity:** MEDIUM. Spec teaches the convenience surface but never mentions the `Skip` family that consumers writing helper wrappers will need to bypass extra frames.
- **Spawned:** AJ-34 — add a §3.1 subsection documenting the `*Skip` variants and when to use them.

## 4. Cumulative AB tally (Cycles 19-24)

| Section | ❌ | Score |
|---|---|---|
| §07 conditional-and-utilities | 5 | 70.6 % |
| §08 validators | 8 | 33.3 % |
| §09 converters | 5 | 66.7 % |
| §10 reflection-and-dynamic | 8 | 38.5 % |
| §11 versioning | 8 | 18.2 % |
| **§15 observability** | **7** | **74.1 %** |
| **Total** | **41** | — |

- **Cumulative ❌:** 41 (was 34) across 6 audited sections.
- **Fabrication mix:** 18 CRITICAL + 19 HIGH + 3 MEDIUM + 1 LOW.
- **Fabrication rate** (❌ ÷ verifiable claims promoted): now ~52 % across the AB sweep.
- **S-106 lint:** still **MANDATORY** — Cycle 24's CRITICAL items (signature drift, fabricated file, fabricated output format, wrong rationale for a correct rule) are exactly the patterns S-106 is designed to catch.

## 5. Spawned AJ items (BLOCKED)

| ID | Target | Severity | Action |
|---|---|---|---|
| AJ-28 | §2.1 line 49 | CRITICAL | Add `true` as first arg to `VarTwo` example |
| AJ-29 | §2.1, §2.2, §2.3 | CRITICAL | Re-frame as message builders; rename `err` → `msg` |
| AJ-30 | §2.1 + §2.2 + §2.4 | HIGH | Collapse VarTwo/VarTwoNoType; document `isIncludeType` toggle |
| AJ-31 | §4 line 111 | HIGH | Replace `ResultAny.go` with real file name |
| AJ-32 | §4 lines 113-117 | CRITICAL | Replace fabricated format with real output OR forward to §06/07 |
| AJ-33 | §3 rule 2 | CRITICAL | Rewrite stack-enhancement rationale; remove `HandleErr` claim |
| AJ-34 | §3 (new §3.1) | MEDIUM | Document `*Skip` variants of `StackEnhance` |

All AJ-28..34 are **BLOCKED** by the `spec/01-app/` 🧊 freeze. **S-106 lint should land before these are unblocked.**

## 6. Remaining ❓ in `spec/01-app/`

After Cycle 24: §07 3 + §08 6 + §09 8 + §10 6 + §11 1 + §16 13 = **37 ❓** remaining (down from 50). Path: AB pass 7 → Cycle 25 on `16-security.md` (13 ❓).
