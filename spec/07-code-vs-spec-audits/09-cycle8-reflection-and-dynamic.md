# Cycle 8 — `01-app/10-reflection-and-dynamic.md`

> **Date**: 2026-05-05 (Asia/Kuala_Lumpur)
> **Spec audited**: [`spec/01-app/10-reflection-and-dynamic.md`](../01-app/10-reflection-and-dynamic.md)
> **Auditor**: Lovable agent (loop AA-cycle8)
> **Status**: **baseline recorded — surface upstream-only, but two MUST-NOT rules independently verified ✅**

---

## 1. Method

For each numbered section in the spec, classify every concrete claim (import path, exported symbol, signature, MUST/MUST-NOT rule) as:

- ✅ **Match** — claim verified against `enum-v7` source on disk.
- ⚠️ **Drift** — verifiable but inaccurate.
- ❌ **Contradiction** — verifiable and wrong.
- ❓ **Unverifiable** — package not consumed by `enum-v7` and no upstream `core-v9` source on disk; defer to task **AB**.

Verification commands run from repo root:

```bash
rg -l "core-v9/(coredynamic|reflectcore|reflectinternal)" --type go
rg -n "coredynamic\.|reflectcore\." --type go
rg -l '"reflect"' --type go | grep -v '^cross-repo/'        # consumer-side reflect import (forbidden by §1)
rg -n "reflect\.DeepEqual" --type go | grep -v '^cross-repo/' # anti-pattern from §7
rg -n "core-v9/internal/reflectinternal" --type go            # forbidden cross-module internal/ (§1, §4)
ls cross-repo/core-v9/{coredynamic,reflectcore} 2>/dev/null
```

Results:
- **Zero importers** of `coredynamic`, `reflectcore`, or `reflectinternal` in `enum-v7`.
- **Zero call sites** of any documented symbol (19 probed).
- **Zero direct `"reflect"` imports** in consumer code (excluding `cross-repo/` mirror).
- **Zero `reflect.DeepEqual` calls** in consumer code.
- **Zero `core-v9/internal/reflectinternal` imports** (would fail Go's `internal/` rule anyway — already covered by C-CVS-04 fix in §05).
- `cross-repo/core-v9/{coredynamic,reflectcore}` directories do not exist.

---

## 2. Claim inventory

| #  | Spec § | Claim                                                                                              | Verdict | Note |
|----|--------|----------------------------------------------------------------------------------------------------|---------|------|
| 1  | §1     | Three-layer architecture (`coredynamic` public / `reflectcore` public / `internal/reflectinternal` internal) | ❓ | Layer existence not verifiable without upstream source |
| 2  | §1 (MUST) | Consumer code MUST NOT import `internal/reflectinternal`                                          | ✅ | Verified — `rg "core-v9/internal/reflectinternal"` → 0 hits in `enum-v7` |
| 3  | §1 (MUST) | Consumer code MUST NOT import stdlib `"reflect"` directly                                          | ✅ | Verified — `rg '"reflect"'` (excl. `cross-repo/`) → 0 hits in `enum-v7` |
| 4  | §1 (convention) | MUST/MUST-NOT/MAY are non-negotiable; "should/prefer" are guidance                            | ✅ | Documentation convention — internally consistent and applied throughout this audit cycle |
| 5  | §2.1   | `coredynamic.InvokeMethod(target any, name string, args ...any) (any, error)` — two-return only (F-V14-02) | ❓ | No consumer |
| 6  | §2.1   | `coredynamic.HasMethod(target, name) bool`                                                         | ❓ | No consumer |
| 7  | §2.1   | `coredynamic.MethodNames(target) []string`                                                         | ❓ | No consumer |
| 8  | §2.2   | `coredynamic.GetField(target, name) (any, bool)`                                                   | ❓ | No consumer |
| 9  | §2.2   | `coredynamic.SetField(target, name, value) error`                                                  | ❓ | No consumer |
| 10 | §2.2   | `coredynamic.AllFields(target) map[string]any`                                                     | ❓ | No consumer |
| 11 | §2.3   | `coredynamic.TypeName` / `TypeFullName` / `IsNullOrUndefined`                                      | ❓ | No consumer |
| 12 | §3.1   | `reflectcore.{IsPointer,IsStruct,IsSlice,IsMap,IsFunc,IsChannel,IsInterface}` predicates           | ❓ | No consumer (7 symbols probed) |
| 13 | §3.2   | `reflectcore.WalkFields(target, func(name, value))`                                                | ❓ | No consumer |
| 14 | §3.2   | `reflectcore.GetTag(target, fieldName, tagName) string`                                            | ❓ | No consumer |
| 15 | §3.3   | `reflectcore.DerefAll(ptr)` returns the underlying value                                           | ❓ | No consumer |
| 16 | §4     | `internal/reflectinternal` responsibilities (low-level setters, unsafe pointer arithmetic, type-cache) | ❓ | Internal package — by design unverifiable from `enum-v7` |
| 17 | §5     | Decision matrix (generics vs `coredynamic` vs `reflectcore` vs `internal/reflectinternal` vs `corejson` vs `corefuncs.GetFuncName`) | ❓ | Reflects §2–§4 surface; same status |
| 18 | §6     | Performance mitigations: type cache, `coreonce` lazy binding, generics-first defaults              | ❓ | Behavioural rules — no consumer |
| 19 | §7     | Common-mistakes table (6 rows) — `reflect`-stdlib ban, `internal/` ban, prefer generics, cache `HasMethod`, `DerefAll` first, prefer `isany.DeepEqual` over `reflect.DeepEqual` | ✅ | Two anti-patterns measurable: stdlib `"reflect"` import (claim #3) and `reflect.DeepEqual` use — both **0 hits** in `enum-v7` consumer code → rules are being honoured |

**Total claims**: 19
**Verifiable subset**: 4 (claims #2, #3, #4, #19 — all rule-compliance checks)
**Verifiable match rate (baseline)**: **4 / 4 = 100.0%**

---

## 3. Score row

| Date       | Cycle | Spec audited                            | Claims | ✅ | ⚠️ | ❌ | ❓ | Score (verifiable) |
|------------|-------|-----------------------------------------|--------|----|-----|----|----|--------------------|
| 2026-05-05 | 8 (baseline / closed) | `01-app/10-reflection-and-dynamic.md` | 19 | 4 | 0 | 0 | 15 | **100.0%** *(verifiable)* |

> **Note**: closed at baseline — no fixes needed because all four verifiable claims are already at ✅ Match (the MUST/MUST-NOT rules are being followed in `enum-v7`).

---

## 4. Findings

**No drifts, no contradictions found.** Cycle 8 is the first audit cycle where the spec's verifiable subset consists entirely of **negative-rule compliance** (MUST-NOT statements), and `enum-v7` honours all of them:

- §1 ban on consumer-side `"reflect"` import → 0 violations
- §1 ban on `internal/reflectinternal` import → 0 violations (also impossible due to Go's `internal/` rule across the `core-v9` module boundary — see C-CVS-02 / C-CVS-04 history)
- §7 ban on `reflect.DeepEqual` in consumer code → 0 violations

The remaining 15 ❓ defer to task **AB** (no upstream `core-v9/coredynamic` or `core-v9/reflectcore` source on disk — the `cross-repo/` mirror does not carry either package).

---

## 5. Next actions

1. Update the scoreboard (`01-scoreboard.md`) with the Cycle 8 baseline+closed row and bump the §AB ❓ tally to **87** (15 §10 + 23 §09 + 18 §08 + 17 §07 + 7 §04 + 1 §05 + 6 §06).
2. Audited-and-closed sections become **6**: §03, §04, §05, §06, §08, §10. Baseline-only: §07, §09.
3. Continue to Cycle 9 → `11-versioning.md` on next `next`.
