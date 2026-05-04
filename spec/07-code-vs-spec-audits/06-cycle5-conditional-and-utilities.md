# Cycle 5 — `01-app/07-conditional-and-utilities.md`

> **Date**: 2026-05-04 (Asia/Kuala_Lumpur)
> **Spec audited**: [`spec/01-app/07-conditional-and-utilities.md`](../01-app/07-conditional-and-utilities.md)
> **Auditor**: Lovable agent (loop AA-cycle5)
> **Status**: **baseline recorded — section is 100 % upstream-only**

---

## 1. Method

For each numbered section in the spec, classify every concrete claim (import path, exported symbol, signature, behavioural rule) as:

- ✅ **Match** — claim verified against `enum-v1` source on disk.
- ⚠️ **Drift** — verifiable but inaccurate.
- ❌ **Contradiction** — verifiable and wrong.
- ❓ **Unverifiable** — package not consumed by `enum-v1` and no upstream `core-v9` source on disk; defer to task **AB**.

Verification commands run from repo root:

```bash
rg -l "core-v9/(conditional|isany|issetter|regexnew|coremath|corecmp|coresort|corefuncs|namevalue|keymk)" --type go
rg -n "conditional\.|isany\.|issetter\.|regexnew\.|coremath\.|corecmp\.|coresort|corefuncs\.|namevalue\.|keymk\." --type go
ls cross-repo/core-v8/{conditional,isany,issetter,regexnew,coremath,corecmp,coresort,corefuncs,namevalue,keymk} 2>/dev/null
```

All three commands returned **zero matches**: no `enum-v1` package imports any of the §07 packages, and the `cross-repo/core-v8/` mirror only carries `CHANGELOG.md` / `CONTRIBUTING.md` / `README.md` / `scripts/` — not source for these packages.

---

## 2. Claim inventory

| # | Spec § | Claim | Verdict | Note |
|---|--------|-------|---------|------|
| 1 | §1.1 | `conditional.If[T]`, `IfFunc[T]`, `NilDef[T]`, `NilDefPtr[T]`, `ValueOrZero[T]` exist | ❓ | No consumer; needs upstream source |
| 2 | §1.2 | 11 typed wrappers per primitive × 15 primitives (`IfInt`, `IfFuncString`, `NilDefFloat64`, `ValueOrZeroBool`, …) | ❓ | Wrapper-count claim un-runnable |
| 3 | §1.3 | `conditional.ErrorFunc(...)` and `TypedErrorFunctionsExecuteResults[T]` exist | ❓ | Symbols not callable from `enum-v1` |
| 4 | §1 (rule) | "Default to eager. Switch to `*Func` only when the un-chosen branch has measurable cost." | ❓ | Behavioural rule — no consumer to verify against |
| 5 | §2 | `isany.{Null,Defined,Zero,DeepEqual,JsonEqual}` | ❓ | Package not imported anywhere in `enum-v1` |
| 6 | §2 (rule) | Always use `isany.Null` / `isany.Defined` on `any` parameters | ❓ | No consumer |
| 7 | §3 | `issetter.Value` is byte-backed with 6 states (`Uninitialized`, `True`, `False`, `Unset`, `Set`, `Wildcard`) | ❓ | No consumer |
| 8 | §3 | Predicates `IsOn`, `IsOff`, `HasInitialized` | ❓ | No consumer |
| 9 | §4 | `regexnew.New.Lazy(pat)` and `regexnew.New.LazyLock(pat)` constructors | ❓ | No consumer |
| 10 | §4 | `IsMatch`, `IsApplicable`, `IsDefined`, `IsFailedMatch` predicates | ❓ | No consumer |
| 11 | §5 | `coremath` exposes Min/Max for `byte,int,int16,int32,int64,float32,float64` | ❓ | No consumer |
| 12 | §5 | `corecmp` exposes `Byte,Integer,Integer8/16/32/64,String,Time` + pointer variants returning `constants.CompareEqual/Less/Greater` | ❓ | No consumer |
| 13 | §5 | `coresort/strsort.Quick(*[]string)` / `QuickDsc(*[]string)` mutate in place | ❓ | No consumer |
| 14 | §6 | `corefuncs.GetFuncName`, `GetFuncFullName`, `ActionReturnsErrorFuncWrapper`, `InOutErrFuncWrapper` | ❓ | No consumer |
| 15 | §7 | `namevalue.NewInstance(name, value)` + `NewCollection()` + `.Add()` + `.ToMap()` | ❓ | No consumer |
| 16 | §8 | `keymk.New.Compile("user", id, "post", id)` returns `"user/<id>/post/<id>"` | ❓ | No consumer |
| 17 | §9 | Decision-matrix routing table | ❓ | Reflects §1–§8; same status |

**Total claims**: 17 (cross-cuts ≈40 individual symbols and behavioural rules)
**Verifiable subset**: 0
**Verifiable match rate**: **N/A** (0 of 17 measurable today)

---

## 3. Score row

| Date | Cycle | Spec audited | Claims | ✅ Match | ⚠️ Drift | ❌ Contradiction | ❓ Unverifiable | Score |
|------|-------|--------------|--------|---------|---------|------------------|----------------|-------|
| 2026-05-04 | 5 (baseline) | `01-app/07-conditional-and-utilities.md` | 17 | 0 | 0 | 0 | 17 | **N/A** *(no verifiable subset)* |

---

## 4. Findings opened

_None._ All 17 claims are ❓ pending task **AB** (fetch upstream `core-v9` source). No drift or contradiction is **provable** from `enum-v1` alone, so no `D-CVS-*` / `C-CVS-*` IDs are minted in this cycle.

> Per the audit method, ❓-only sections do not contribute to the aggregate verifiable-match rate. They count toward "section coverage" (5 / 16) but not the score.

---

## 5. Recommended next actions

1. **AB** — pull upstream `core-v9` source so all 17 §07 claims (plus the existing 7 §04 + 1 §05 + 6 §06 ❓s) become verifiable in one pass.
2. Until **AB** lands, continue Cycle 6 → [`08-validators.md`](../01-app/08-validators.md). Validators are partly consumed inside `enum-v1` (e.g. via `conditional`-style calls in `inttype`), so Cycle 6 is more likely to yield a measurable match rate.
3. Note for §07 specifically: the spec is internally consistent and structurally sound — no obvious smell to flag pre-AB. Defer all judgement.
