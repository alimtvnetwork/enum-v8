---
name: Test failure RCA patterns (enum-v6)
description: Recurring root-cause patterns when `./run.ps1 -tc` reports test failures, with reusable fixes — fixture drift, sparse array maps, upstream BasicString defects, Goconvey log conflation
type: feature
---

# Test Failure Root Cause Analysis — enum-v6

When `./run.ps1 -tc` reports failing tests, walk this checklist FIRST before reading test code. Most failures fall into one of 4 recurring patterns documented below. Each pattern lists: root cause, how to recognise it, reusable fix.

---

## Pattern 1 — Stale fixture in `tests/creationtests/allEnumGeneralTestCases.go` after a `Variant.go` override is added

**Symptom:** `Test_AllEnums_ContractsTesting` fails at `AllEnums_ContractsTesting_test.go:36` (or `:37`) with `Expected: "<old>"  Actual: "<new>"` — usually `Expected: ""  Actual: "<some name>"`.

**Root cause:** A previous cycle added a local override on a Variant method (commonly `MinValueString`, `NameValue`, `MarshalJSON`) to bypass an upstream `core-v9` bug. The fixture in `allEnumGeneralTestCases.go` still pins the **old broken** return value (e.g. `StringMin: ""`).

**Real example (Cycle 60 → Cycle 63):** PI-006 added `sqliteconnpathtype.Variant.MinValueString()` returning lexicographic min name (`"All"`). Fixture at `allEnumGeneralTestCases.go:1446` still had `StringMin: ""`.

**Fix recipe:**
1. Open `tests/creationtests/allEnumGeneralTestCases.go`, find the affected enum's block (search by `TypeName:"<pkg>.Variant"`).
2. Update `StringMin` / `StringMax` / `ExpectedInvalidValueString` to match the new override's actual output.
3. If the override returns a *computed* value (e.g. lex-min name), document why in a `// PI-NNN: …` comment immediately above the field.

**Prevention:** Every time you add or change a Variant override that affects `Min/Max/Name/Value/JSON` output, grep `allEnumGeneralTestCases.go` for the type name BEFORE committing.

---

## Pattern 2 — Sparse `[...]Type{Index: value}` array literal missing an enum constant

**Symptom:** `TestOsDetect_CrossPlatformSafe` (or any per-package safe-coverage suite) fails with `Expected value to NOT be blank (but it was)!` for a method like `NameLower()`, `Description()`, `Code()`.

**Root cause:** `vars.go` defines a sparse array literal indexed by Variant (e.g. `lowerCaseNames = [...]string{Invalid: "invalid", Windows: "windows", ...}`). One Variant constant was added later but its index was forgotten in the map literal. Go silently fills the missing slot with the zero value (`""`), and tests that iterate every known Variant catch it.

**Real example (Cycle 63):** `osdetect/vars.go:30 lowerCaseNames` array missing `RedHatEnterpriseLinux: "redhat-enterprise-linux"` — the array sized itself to highest index (13 = `Android`), leaving slot 11 = `""`.

**Fix recipe:**
1. Compare the const block in `Variant.go` (the source of truth) against EVERY `[...]Type{...}` and `map[Variant]Type{...}` literal in `vars.go`.
2. Add the missing entry. Mirror the casing convention used by sibling entries.
3. Bonus: convert sparse arrays to `map[Variant]string` so missing keys are detectable via `len(map) == len(constants)` invariants in tests — but only when the existing pattern won't be disrupted.

**Prevention:** When adding a new Variant constant, immediately grep its package for `[...]string{` and `map[Variant]` and add the new index/key everywhere.

---

## Pattern 3 — `BasicString` enum built via `CreateUsingStringersSpread` returns empty `RangesDynamicMap` / `Min`

**Symptom:** `Test_AllEnums_NumericRange` fails at line 84 with `Expected '0' to be greater than '0'` — meaning `len(RangesDynamicMap()) == 0` for some enum. Or `MinValueString()` returns `""`. Almost always a string-backed `Variant string` enum.

**Root cause (upstream `core-v9` defect):** `coreimpl/enumimpl/newBasicStringCreator.CreateUsingStringersSpread` initialises `min := ""` and updates only under `if name < min`, which never fires (every real name is `> ""`). It also stores names lazily in a way that leaves `RangesDynamicMap` empty for spread-constructed string enums. This is documented in Cycle 60 PI-005..007 RCA.

**Fix recipe (local override pattern):**
1. Add a `func (it Variant) MinValueString() string` override that scans `BasicEnumImpl.StringRanges()` locally — see `sqliteconnpathtype/Variant.go:59-73`.
2. Add a `func (it Variant) RangesDynamicMap() map[string]interface{}` override that builds the map from `StringRanges()` locally:
   ```go
   func (it Variant) RangesDynamicMap() map[string]interface{} {
       names := BasicEnumImpl.StringRanges()
       out := make(map[string]interface{}, len(names))
       for _, n := range names { out[n] = n }
       return out
   }
   ```
3. Add the type to the matching skip-map ONLY if neither override fits. Prefer fixing locally over skipping.

**Prevention:** Any new `Variant string` enum that uses `enumimpl.New.BasicString.UsingTypeSlice` or `CreateUsingStringersSpread` MUST get all three overrides at creation time: `MinValueString`, `MarshalJSON`/`UnmarshalJSON` (PI-005), `IsAnyNamesOf` (PI-007), and `RangesDynamicMap`. Use `sqliteconnpathtype/Variant.go` as the canonical template.

---

## Pattern 4 — Goconvey log conflation across parallel packages

**Symptom:** `failing-tests.txt` shows `--- FAIL: Test_X_Constructors` immediately followed by a failure block pointing to a `.go` file in a DIFFERENT package (e.g. OnOff test name + osdetect failure location).

**Root cause:** `go test ./...` runs packages in parallel, but Goconvey writes to a shared assertion-counter stream. The log scraper (`TestLogWriter.psm1`) groups any `Failures:` block under the most recently seen `--- FAIL:` line, even if they belong to different test functions in different packages. The "extra" failing test name is a phantom — only the file path inside the failure block is authoritative.

**How to disambiguate:**
1. Look at the `Failures: * /path/to/file.go Line N` block — the file path is the truth.
2. If the file path is in a different package than the named test, the named test is the phantom; the real failure belongs to whichever test in the file owns line N.
3. Cross-check by counting actual `--- FAIL:` entries; phantoms inflate the count.

**Real example (Cycle 63):** Log showed 4 failing tests but only 3 root causes existed. `Test_OnOffType_Constructors` was a phantom — its "failure block" pointed at `osdetect/OsDetect_CrossPlatform_test.go:37`, which is the actual `TestOsDetect_CrossPlatformSafe` failure.

**Prevention:** When the failure count seems too high, deduplicate by `(file, line)` pair, NOT by `--- FAIL: Test_*` headers.

---

## Quick triage workflow

1. Open `data/test-logs/failing-tests.txt`.
2. Strip ANSI: `python3 -c "import re,sys;print(re.sub(r'\x1b\[[0-9;]*m','',open(sys.argv[1]).read()))" data/test-logs/failing-tests.txt`.
3. Extract every `Failures:` block + its `* /path/file.go Line N` + `Expected/Actual` lines. Deduplicate by `(file, line)`.
4. For each unique `(file, line)`, classify against patterns 1–4 above.
5. Apply the matching fix recipe. Re-run `./run.ps1 -tc`.

## Cycles where each pattern was first observed

| Pattern | First cycle | Issue ID | Type fixed |
|---------|-------------|----------|------------|
| 1 — Stale fixture | 63 | (this RCA) | sqliteconnpathtype StringMin |
| 2 — Sparse array gap | 63 | (this RCA) | osdetect lowerCaseNames |
| 3 — BasicString defect | 60 | PI-005..007 | sqliteconnpathtype |
| 4 — Goconvey conflation | 63 | (this RCA) | OnOff phantom |
