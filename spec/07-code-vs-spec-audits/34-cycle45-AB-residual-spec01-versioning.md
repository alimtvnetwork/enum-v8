# Cycle 45 — AB-residual deep-probe of `spec/01-app/11-versioning.md` (Cycle 23 carry-over)

**Date:** 2026-05-06
**Scope:** Settle the single ❓ left by Cycle 23 (`24-cycle23-AB-versioning.md` row 6) using direct evidence from upstream `core-v9 v1.5.8`.
**Result:** **1 ❓ → ❌ DEMOTION** — NEW C-CVS-62 (HIGH).
**Allowed under freeze:** read-only audit promotion (no spec rewrites).

## 1. Demotion

### Row 6 — §1: "`coreversion` plays well with `coregeneric.Collection`" → ❌ NEW C-CVS-62 (HIGH)

**Probes against `/tmp/core-v9-upstream`:**

1. `grep -rn "coregeneric" coreversion/` → **zero hits**. The `coreversion` package never imports or references `coregeneric` anywhere.
2. `coreversion/VersionsCollection.go` imports only `constants`, `coredata/corejson`, and `coreinterface`. The struct is a plain hand-rolled wrapper:
   ```go
   type VersionsCollection struct {
       Versions []Version
   }
   func (it *VersionsCollection) Add(version string) *VersionsCollection { ... }
   func (it *VersionsCollection) AddSkipInvalid(version string) *VersionsCollection { ... }
   ```
   No embedding of `coregeneric.Collection[Version]`, no conversion helper, no shared interface.
3. `coredata/coregeneric/Collection.go:33` defines `Collection[T any] struct { ... }` with constructors `EmptyCollection`, `NewCollection`, `CollectionFrom`, `CollectionClone`, `CollectionLenCap` and methods `HasAnyItem`, `LastIndex`, `HasIndex`, `Items`, … — a fully separate generic surface that `coreversion` neither consumes nor produces.

**Verdict:** ❌ — the "plays well with `coregeneric.Collection`" interop claim is **fabricated**. `coreversion` ships its OWN non-generic `VersionsCollection` instead of (or alongside) any `coregeneric.Collection[Version]` adapter. **NEW C-CVS-62 (HIGH):** delete the `coregeneric.Collection` interop bullet from §1 and either (a) drop the interop framing entirely, or (b) document `VersionsCollection` as the real collection surface. Spawned **AJ-21b (BLOCKED by freeze)** — fold into the existing AJ-21 §1 rewrite.

## 2. Updated Cycle-23 scoreboard

| Pass | ✅ | ⚠️ | ❌ | ❓ | Verifiable score |
|------|----|---|----|---|---|
| Cycle 8 baseline | 0 | 0 | 0 | 11 | N/A |
| Cycle 23 (AB pass 5) | 2 | 0 | 8 | 1 | 18.2% |
| **Cycle 45 (this)** | **2** | 0 | **9** | **0** | **18.2%** (2/11) |

§11 verifiable score unchanged (no new ✅; one ❓→❌ demotion). ❓ pool **fully cleared** for §11 — every claim now has a verdict.

## 3. Cumulative AB-residual ❓ pool

- Pre-Cycle-45: 18 ❓ in `spec/01-app/` (Cycle 19=4 OOB / Cycle 20=1 OOB / Cycle 21=6 / Cycle 22=6 / Cycle 23=1).
- Post-Cycle-45: **18 − 1 = 17 ❓** in `spec/01-app/` (Cycle 19=4 OOB / Cycle 20=1 OOB / Cycle 21=6 / Cycle 22=6 / **Cycle 23=0 ✅ closed**).

## 4. Action items spawned / amended

- **AJ-21b (NEW, BLOCKED by freeze):** drop `coregeneric.Collection` interop bullet from `11-versioning.md` §1; document real `VersionsCollection` surface (folded into existing AJ-21 §1 constructor rewrite for one consolidated edit).

## 5. Cumulative AB ❌ summary update

- Cumulative AB ❌ across 7 sections: **50 → 51** (NEW C-CVS-62). CRITICAL count unchanged at **23** (C-CVS-62 is HIGH, not CRITICAL).

## 6. Memory + bookkeeping

- Spec changelog → `spec-v0.50.0`.
- `package.json` → `0.15.0`.
- `.lovable/memory/workflow/01-state.md` updated.
- Scoreboard top-line updated.

---

_Audit file: `spec/07-code-vs-spec-audits/34-cycle45-AB-residual-spec01-versioning.md`_
_See also: `24-cycle23-AB-versioning.md` (the cycle whose ❓ is settled here)._