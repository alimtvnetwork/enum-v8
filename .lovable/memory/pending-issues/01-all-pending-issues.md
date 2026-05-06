# Pending Issues — Consolidated Tracker

> Single file for all pending issues. Update in-place.
> When resolved, move to Resolved section with date.

---

## Open Issues

*(PI-001 moved to Resolved — see below)*

### PI-002: Cross-spec stale `integratedtests/` paths (Task AH)

- **severity:** HIGH
- **description:** Multiple spec files outside `spec/01-app/` still reference `tests/integratedtests/` which doesn't exist. Known targets: `spec/06-testing-guidelines/01-folder-structure.md`, `spec/03-powershell-test-run/`, `spec/04-tooling/04-bootstrap-into-new-repo.md`, `spec/02-app-issues/02-internal-package-coverage-policy.md`, `spec/00-llm-integration-guide.md` line 36.
- **owner:** AI (audit cycles)
- **plan:** ✅ COMPLETE — resolved by re-framing (consumer-coverage callouts) across Cycles 12/15/17/18 + Cycle 34 (S-003/S-004 close-out). All `integratedtests/` references now correctly document the upstream `core-v9` consumer layout with explicit upstream-vs-enum-v6 scope disclaimers. Move to Resolved on the next cycle that touches this file.

### PI-003: 148 ❓ claims unresolved (Task AB)

- **severity:** MEDIUM
- **description:** 148 claims across `spec/01-app/` scored ❓ because they reference upstream `core-v9` APIs that have zero `enum-v6` consumers and no mirrored source. Need upstream source fetch to verify.
- **owner:** AI + User (fetch access)
- **blocks:** Full spec verification, re-audit of §07/§09 (Task AC)

### PI-004: `spec/06-testing-guidelines/` never audited

- **severity:** HIGH
- **description:** 9 files, most-referenced spec directory for any implementation AI. Contains known stale `integratedtests/` reference. Cycle 15 target.
- **owner:** AI
- **plan:** Next audit cycle (AA)

### PI-005: `sqliteconnpathtype.Variant` JSON round-trip is broken ✅ RESOLVED (2026-05-06, Cycle 60)

- **severity:** MEDIUM (resolved)
- **discovered:** 2026-05-06 (Task AL-01)
- **root cause:** Upstream `core-v9` `BasicString.UnmarshallToValue` → `GetValueByName(string(bytes))` looks the *quoted* JSON bytes (`"\"Invalid\""`) up in `jsonDoubleQuoteNameToValueHashMap`, but that hashset is built via `stringsToHashSet(rawNames)` so its keys are raw (`Invalid`). Round-trip therefore fails. MarshalJSON itself is fine.
- **fix (Cycle 60):** Overrode `Variant.MarshalJSON` to emit `strconv.Quote(string(it))` and `Variant.UnmarshalJSON` to `strconv.Unquote` first then dispatch to `BasicEnumImpl.GetValueByName(rawName)`. Added empty/`""`/nil fallback to local `MinValueString()` (PI-006 fix). Removed `jsonRoundTripSkipTypeNames` skip entry for sqliteconnpathtype.

### PI-006: `sqliteconnpathtype.Variant` Format/NameValue/MinValueString defects ✅ RESOLVED (2026-05-06, Cycle 60)

- **severity:** MEDIUM (resolved)
- **discovered:** 2026-05-06 (Task AL-02)
- **root cause:** Two upstream defects: (1) `enumimpl.NameWithValue` uses `EnumNameValueFormat = "%s(%d)"` which produces `"Invalid(%!d(string=Invalid))"` for string-backed enums; (2) `newBasicStringCreator.CreateUsingStringersSpread` initialises `min := ""` then assigns under `if name < min`, which never fires because every non-empty name is `> ""`, so `BasicString.Min()` always returns "" for spread-constructed enums.
- **fix (Cycle 60):** Overrode `Variant.NameValue` to return `it.String()` (mirrors upstream's `StringEnumNameValueFormat = "%s"`); overrode `Variant.MinValueString` to compute the lexicographic min from `BasicEnumImpl.StringRanges()` locally. Removed `formatSuiteSkipMinMaxAll` and `formatSuiteSkipNameValue` entries for sqliteconnpathtype, plus the `numericRangeSuiteSkipMinValueString` entry.

### PI-007: `sqliteconnpathtype.Variant.IsAnyNamesOf()` returns true for empty input ✅ RESOLVED (2026-05-06, Cycle 60)

- **severity:** LOW (resolved)
- **discovered:** 2026-05-06 (Task AL-03)
- **root cause:** `Variant.IsAnyNamesOf` was dispatching to upstream `BasicString.IsAnyOf`, which has an early `if len(checkingItems) == 0 { return true }` (vacuous truth). Upstream provides a separate `BasicString.IsAnyNamesOf` with the correct empty→false semantics — wrong helper was wired.
- **fix (Cycle 60):** Switched dispatch from `BasicEnumImpl.IsAnyOf` to `BasicEnumImpl.IsAnyNamesOf`. Removed `predicateSuiteSkipEmptyAnyNames` entry.

### PI-008: `quotes/unWrapBoth` and `brackets/unWrapBoth` off-by-one ✅ RESOLVED (2026-05-06, Cycle 59)

- **severity:** LOW (real bug, low blast radius — these helpers are only reachable via `UnWrapWith` / `Quote.UnWrap` / `Bracket.UnWrap`)
- **discovered:** 2026-05-06 (Cycle 57, test-fix triage of AL-06)
- **resolved:** 2026-05-06 (Cycle 59) — Confirmed bug. Fixed all 4 helpers: (1) `quotes/unWrapBoth.go` `s[1:length-2]` → `s[1:length-1]`; (2) `quotes/unWrapSingle.go` left branch `s[1:length-1]` → `s[1:length]`, right branch `s[0:length-2]` → `s[0:length-1]`; (3+4) same in `brackets/unWrapBoth.go` and `brackets/unWrapSingle.go`. Wrap counterparts add exactly 1 char per side, so symmetric unwrap is correct. Updated cycle-57 test expectations from "h" → "hi" in `quotes/Quotes_WrapUnwrap_test.go` and `brackets/Brackets_WrapUnwrap_test.go`.

---

## Resolved Issues

### D-CVS-62: Missing `scripts/coverage/` utilities

- **resolved:** 2026-05-06 (Cycles 31 + 32, S-108 + S-110)
- **fix:** Cycle 31 (S-108) restored the auto-invoked `Generate-CoveragePrompts.ps1`. Cycle 32 (S-110) restored the three standalone utilities (`Get-UncoveredLines.ps1`, `Get-FunctionCoverage.ps1`, `Get-PackageCoverageReport.ps1`). All four scripts smoke-tested via nix-pwsh. `scripts/coverage/` and `spec/03-powershell-test-run/06-coverage-prompt-generator.md` are now in lockstep.

### PI-001: Upstream `core-v9` `go.mod` module path mismatch (Task W + AG)

- **resolved:** 2026-05-05
- **fix:** User renamed upstream `go.mod` → `module github.com/alimtvnetwork/core-v9`, tagged `v1.5.8`. AI dropped `replace` bridge in `enum-v6/go.mod`, pinned `core-v9 v1.5.8`.

_(Consolidated from `.lovable/pending-issues/01-core-v9-go-mod-rename.md` and `02-cross-spec-stale-paths.md` — those files are now superseded by this tracker)_
