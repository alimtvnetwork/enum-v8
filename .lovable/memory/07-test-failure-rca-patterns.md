---
name: Test failure RCA patterns (enum-v7)
description: Recurring root-cause patterns when `./run.ps1 -tc` reports test failures, with reusable fixes — fixture drift, sparse array maps, upstream BasicString defects, Goconvey log conflation
type: feature
---

# Test Failure Root Cause Analysis — enum-v7

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
| 5 — OnlySupportedErr/RangesInvalidErr asserted nil | 65 | (this RCA) | compresslevels, conntrackstate, servicestate, sitestatetype |
| 6 — Shorthand-input pinning vs fuzzy GetValueByName | 65 | (this RCA) | onofftype |

---

## Pattern 5 — Asserting `nil` on `OnlySupportedErr` / `OnlySupportedMsgErr` / `RangesInvalidErr`

**Symptom:** Tests fail with messages like:
- `OnlySupportedErr unexpected: tRunner - Only Ranges: Only selected ranges supported... Only supported ("X"). Unsupported ("Y","Z",...)`
- `RangesInvalidErr unexpected error: Out of Range or Invalid Range: ... Range must be in between 0 and N. Ranges must be one of these {Invalid(0), ...}`

**Root cause:** These three methods on `BasicEnumImpl` are **diagnostic / informational**, not validators:
- `OnlySupportedErr(names...)` always returns a non-nil error describing which names ARE vs ARE NOT in the supported subset. It is meant to be embedded into a higher-level error message, not used as a pass/fail check.
- `OnlySupportedMsgErr(message, names...)` — same as above with a context prefix.
- `RangesInvalidErr()` for byte enums whose first member is `Invalid = 0` always returns non-nil because the upstream impl reports the full enumerated numeric range as a "diagnostic" string.

**Fix recipe:** Either (a) assert `err != nil` to confirm the method produces a message, or (b) call without asserting (`_ = v.OnlySupportedErr(...)`) for coverage only. Add a one-line comment explaining "informational descriptor — always non-nil".

**Prevention:** When writing coverage tests for a new enum, never use `OnlySupportedErr`/`RangesInvalidErr` as a happy-path validator. Treat them like `String()` — exercise them, don't assert nil.

**First observed:** 2026-05-06, Cycle 65 fix-up — affected `compresslevels`, `conntrackstate`, `servicestate`, `sitestatetype`.

---

## Pattern 6 — Pinning shorthand-input results when `BasicEnumImpl.GetValueByName` does fuzzy matching

**Symptom:** A test asserting `New("yes") == On` fails with `Expected: onofftype.Variant(2)  |  Actual: onofftype.Variant(1)` (or similar off-by-one).

**Root cause:** Packages with a `newOtherWays`/fallback alias map (e.g. `onofftype`, `conntrackstate`) only consult the alias map when `BasicEnumImpl.GetValueByName` returns an error. But the upstream impl performs **case-insensitive substring / partial matching** before erroring, so inputs like `"yes"`, `"y"`, `"n"`, `"0"` may resolve to a different Variant than the alias map would suggest. The alias map is effectively dead code for those inputs.

**Fix recipe:** For shorthand inputs, exercise them for coverage (`_, _ = New(input)`) but do NOT assert a specific resulting Variant. Only canonical names (entries in `Ranges[...]string{}`) can be safely asserted.

**Prevention:** When writing constructor tests, only pin results for names that appear verbatim in the package's `Ranges` slice. Treat all alias/shorthand inputs as exercise-only.

**First observed:** 2026-05-06, Cycle 65 fix-up — affected `onofftype`.

---

## Pattern 5 — Inverted receiver/argument in level-comparison helpers

**Symptom:** `IsAboveOrEqual` / `IsLowerOrEqual` (or similar level helpers) return the opposite of what the call-site name implies. E.g. `Error.IsAboveOrEqual(Notice) == false`.

**Root cause:** Body compares `level.ValueByte() >= it.ValueByte()` (arg vs receiver) instead of `it.ValueByte() >= level.ValueByte()`. Easy to write by accident; both compile.

**Fix recipe:** Always read as "receiver IS-(verb) argument" → `it OP level`. Real example: `nginxlogtype/Variant.go` (Cycle 84).

---

## Pattern 6 — Trailing `Invalid` sentinel makes `Min()` / `Max()` return the sentinel

**Symptom:** `Min()` or `Max()` returns `Invalid`. Tests like `if Max() == Invalid { t.Error("Max should not be Invalid") }` fire.

**Root cause:** Convention in this repo is `Invalid Variant = iota` FIRST (so 0 == Invalid). When a package puts `Invalid` LAST in the const block (e.g. `osarchs`, `scripttype`), `BasicEnumImpl.Max()` returns the highest byte == `Invalid`, and `Min()` defaulting to `Invalid` constant is wrong because `Invalid` is no longer 0.

**Fix recipe:** When `Invalid` is the trailing sentinel, hand-write `Min()` to return the first real value (e.g. `Default`, `X32`) and `Max()` to return the last real value (e.g. `X64`) — do NOT use `BasicEnumImpl.Min()`/`Max()`. Real examples: `osarchs/Max.go`, `scripttype/Min.go` (Cycle 84).

**Prevention:** Audit any package where `Invalid` is not the first const — every such package needs custom Min/Max.

---

## Pattern 7 — `AllNameValues()` returns `"Name(value)"` strings, not bare names

**Symptom:** A coverage test loops `for _, name := range v.AllNameValues() { New(name) }` and every call fails with `"cannot find in the enum map"` for inputs like `"Root(1)"`, `"App(3)"`.

**Root cause:** Upstream `enumimpl.AllNameValues` formats each entry with `constants.EnumNameValueFormat` == `"%s(%v)"`. Bare names go to the `Hashmap`, but `"Name(value)"` strings do not.

**Fix recipe:** In the package's `New()`, after `BasicEnumImpl.GetValueByName` fails, strip a trailing `(...)` suffix and retry the lookup. Real example: `pathpatterntype/New.go` (Cycle 84). Pattern is generic — apply to any package whose coverage test round-trips `AllNameValues`.

---

## Pattern 8 — `MaxByte()` / `MinByte()` delegate to `BasicEnumImpl` on trailing-`Invalid` packages

**Symptom:** `TestOsArchs_Accessors` (or similar) fails: `MaxByte mismatch` — because `v.MaxByte() != byte(Max())`. Test expects custom `Max()` return value, but `MaxByte()` still calls `BasicEnumImpl.Max()` which returns the sentinel.

**Root cause:** Pattern 6 fixes `Max()` (top-level) but forgets the byte-typed accessors `MaxByte()`/`MinByte()` on the Variant. They independently call `BasicEnumImpl.Max()/Min()` which still returns `Invalid` byte.

**Fix recipe:** When applying Pattern 6, also rewrite `MaxByte()` / `MinByte()` in `Variant.go` to return `byte(<last real value>)` / `byte(<first real value>)` directly. Real example: `osarchs/Architecture.go:207-215` (Cycle 87).

**Prevention:** When you write a custom `Max()`/`Min()` for trailing-`Invalid` packages, immediately grep `MaxByte\|MinByte\|MaxInt\|MinInt` in the same package and update each to skip BasicEnumImpl.

---

## Pattern 9 — Open-ended numeric enums (`inttype`-style) fail `RangesDynamicMap > 0` check

**Symptom:** `Test_AllEnums_NumericRange` line 86 fails: `Expected '0' to be greater than '0'` — `len(rangesMap) == 0`.

**Root cause:** Some "enum" types are not actually enumerable — they wrap an entire numeric range (`inttype.Variant` covers `MinInt..MaxInt`). Their `RangesDynamicMap()` correctly returns `map[string]interface{}{}` because there are no discrete members. The general suite assertion that the map is non-empty doesn't apply.

**Fix recipe:** Add the type's `TypeName()` (e.g. `"inttype.Variant"`) to `numericRangeSuiteSkipRangesDynamicMap` in `tests/creationtests/AllEnums_NumericRange_test.go` with an explanatory comment. Same skip set already used for `compresslevels`, `sqliteconnpathtype`, `strtype`. Real example: Cycle 87.

**Prevention:** Whenever you add an open-ended numeric Variant (anything where `MinInt() == constants.MinInt` and `MaxInt() == constants.MaxInt`), pre-register it in the skip map at creation time.
