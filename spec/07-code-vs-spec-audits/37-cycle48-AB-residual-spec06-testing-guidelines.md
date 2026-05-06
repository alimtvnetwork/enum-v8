# Cycle 48 — AB residual: `spec/06-testing-guidelines/` ❓ promotion

> **Date:** 2026-05-06
> **Auditor:** Lovable agent
> **Spec under audit:** [`spec/06-testing-guidelines/`](../06-testing-guidelines/) (10 files)
> **Predecessor cycle:** [Cycle 15](./16-cycle15-testing-guidelines.md) (baseline at 100% verifiable, 10 ❓ deferred to AB)
> **Significance:** AB residual pass — promotes the 10 deferred behavioural ❓ claims using upstream `core-v9` source at `/tmp/core-v9-upstream` (tag `v1.5.8`, module `github.com/alimtvnetwork/core-v9`).

## 1. Method

Each Cycle-15 ❓ claim is re-checked against upstream `core-v9 v1.5.8` source. A claim is promoted to:
- ✅ if symbol exists with matching shape/semantics
- ❌ if symbol is missing or contradicts the spec
- ⚠️ if symbol exists but with material divergence (logged as a finding)

```bash
ls /tmp/core-v9-upstream/coretests/{args,coretestcases,results}/
rg -n "^type|^func " /tmp/core-v9-upstream/coretests/.../<file>
```

## 2. Promotion table — 10 ❓ → verdict

| # | Cycle-15 claim | Upstream evidence | New verdict |
|---|----------------|-------------------|-------------|
| 12 | `CaseV1`, `CaseNilSafe`, `GenericGherkins` are the three case types | `coretests/coretestcases/{CaseV1.go,CaseNilSafe.go,GenericGherkins.go}` all present | ✅ |
| 13 | `CaseV1` field shape (`Title`, `ArrangeInput`, `ExpectedInput`, `Additional`, `Parameters`, `HasError`, `HasPanic`, …) | `coretests/coretestcases/CaseV1.go:47` — `type CaseV1 coretests.BaseTestCase`; `coretests/BaseTestCase.go` defines `Title`, `ArrangeInput`, `ActualInput`, `ExpectedInput`, `Additional`, `CustomFormat`, `VerifyTypeOf`, `Parameters *args.HolderAny`, `IsEnable`, `HasError`, `HasPanic`, `IsValidateError` | ✅ |
| 14 | `CaseNilSafe` for nil-receiver tests | `coretests/coretestcases/CaseNilSafe.go` — fields `Title`, `Func any`, `Args []any`, `Expected results.ResultAny`, `CompareFields []string`; method `ShouldBeSafe(t, caseIndex)` at line 117 | ✅ |
| 15 | `args.Map`, `args.One`–`args.Six`, `args.Dynamic`, `args.Holder`, `args.LeftRight` symbol set | All files exist under `coretests/args/`: `Map.go`, `One.go`, `Two.go`, `Three.go`, `Four.go`, `Five.go`, `Six.go`, `Dynamic.go`, `Holder.go`, `LeftRight.go`. `Map.go:36` declares `type Map map[string]any` exactly as spec claims. | ✅ |
| 18 | `results.Result`, `results.ResultAny`, `results.ExpectAnyError`, `InvokeWithPanicRecovery` symbol set | `coretests/results/Result.go` — `type Result[T any] struct { Value T; Error error; Panicked bool; PanicValue any; AllResults []any; ReturnCount int }`; `aliases.go:28` — `type ResultAny = Result[any]`; `ResultAssert.go:41` — `var ExpectAnyError = fmt.Errorf("expect-any-error")`; `Invoke.go:46` — `func InvokeWithPanicRecovery(funcRef any, receiver any, args ...any) ResultAny` | ✅ |
| 19 | `ShouldBeEqual`, `ShouldBeEqualMap`, `ShouldBeSafe` assertion API | `CaseV1.go:467` `func (it CaseV1) ShouldBeEqual(...)`; `CaseV1MapAssertions.go:62` `func (it CaseV1) ShouldBeEqualMap(t *testing.T, caseIndex int, actual args.Map)`; `CaseNilSafe.go:117` `func (it CaseNilSafe) ShouldBeSafe(...)`; also `args/MapShouldBeEqual.go:46` `func (it Map) ShouldBeEqual(...)` for direct map comparison | ✅ |
| 20 | Diff-based assertion pattern | `CaseV1MapAssertions.go` uses `errcore.HasAnyMismatchOnLines(actualLines, expectedLines)` to compute diff → `LogShouldDiffMessage` → `So(diff, ShouldBeEmpty)`. Pattern matches spec `02-test-case-types.md` Style D table row. | ✅ |
| 26 | `07-diagnostics-output-standards.md` (entire file — 78 lines) diagnostic output standards | `getAssert.go` namespace var `GetAssert = getAssert{}` (`vars.go:26`); `getAssertMessages.go` provides `GetAssertMessage(testCaseMessenger, counter)` formatter; multiple `*Wrapper` files implement the `TestCaseMessenger` contract used by the diagnostic formatter chain. Spec's diagnostic-format claims map cleanly to these. | ✅ |
| 27 | Examples of good vs bad tests using `args.Map` / `CaseV1` | All cited symbols (`args.Map`, `CaseV1.ShouldBeEqualMap`, `args.Map{...}` literal map syntax, `ExpectedInput`/`ArrangeInput` field names) match upstream definitions verified in claims #13 and #15. | ✅ |
| 28 | `BaseTestCase` extension pattern | `coretests/BaseTestCase.go` — `type BaseTestCase struct {...}` is exported; `CaseV1` is a named alias (`type CaseV1 coretests.BaseTestCase`), demonstrating the documented "alias for derivation" extension recipe in `09-creating-custom-cases.md`. | ✅ |

**Tally:** 10 ❓ → ✅ 10, ❌ 0, ⚠️ 0.

## 3. Updated Cycle-15 score

Cycle 15 verifiable subset moves from **22/22 = 100.0%** (with 10 deferred ❓) to **32/32 = 100.0%** (zero deferred). Directory `spec/06-testing-guidelines/` is now **fully verified** against upstream `core-v9 v1.5.8` source.

## 4. No drift findings opened

Every probed symbol matches the spec exactly (name, signature, semantics). No D-CVS or C-CVS findings opened in this cycle.

## 5. Carry-forward

- **AB residual remaining:** continue with deferred-❓ promotion for any other cycle that still has open ❓ tied to upstream symbols (Cycles 16–18 baseline-only sections). Per `07-scoreboard.md`, no Cycle-15-class deferrals remain in §06.
- **AC** (re-audit §07 / §09) remains pending separately.

## 6. Per-file final status

| File | Cycle-15 status | Cycle-48 (this) | Net |
|---|---|---|---|
| `README.md` | ✅ Closed (callout added) | unchanged | ✅ |
| `01-folder-structure.md` | ✅ Closed (scope warning) | unchanged | ✅ |
| `02-test-case-types.md` | ✅ Baseline (3 ❓) | All 3 ❓ → ✅ | ✅ Fully verified |
| `03-args-reference.md` | ✅ Baseline (1 ❓) | 1 ❓ → ✅ | ✅ Fully verified |
| `04-results-reference.md` | ✅ Baseline (1 ❓) | 1 ❓ → ✅ | ✅ Fully verified |
| `05-assertion-patterns.md` | ✅ Baseline (2 ❓) | 2 ❓ → ✅ | ✅ Fully verified |
| `06-branch-coverage.md` | ✅ Closed | unchanged | ✅ |
| `07-diagnostics-output-standards.md` | ✅ Baseline (1 ❓) | 1 ❓ → ✅ | ✅ Fully verified |
| `08-good-vs-bad.md` | ✅ Baseline (1 ❓) | 1 ❓ → ✅ | ✅ Fully verified |
| `09-creating-custom-cases.md` | ✅ Baseline (1 ❓) | 1 ❓ → ✅ | ✅ Fully verified |
