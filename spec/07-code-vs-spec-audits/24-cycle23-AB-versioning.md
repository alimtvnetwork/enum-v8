# Cycle 23 — AB pass 5: `spec/01-app/11-versioning.md`

**Date:** 2026-05-06
**Auditor:** AI agent
**Source of truth:** Upstream `core-v9 v1.5.8` (clone at `/tmp/core-v9-upstream`).
**Method:** Promote 11 ❓ from Cycle 8 baseline by direct grep / file-read against upstream Go source.
**Spec changelog bump:** spec-v0.38.0

## Result summary

| Verdict | Count |
|---|---|
| ✅ Verified | 2 |
| ❌ Contradicted | 8 |
| ❓ Still pending | 1 |
| **Total claims sampled** | **11** |
| **Verifiable score** | **18.2%** (worst section to date) |

Cumulative AB ❌ findings: **34** (was 26).
Cumulative AB fabrication rate across 5 audited sections: **~55%**.

## Claim-by-claim

| # | Spec claim (location in `11-versioning.md`) | Verdict | Evidence |
|---|---|---|---|
| 1 | §1: `coreversion.Parse(s) (Version, error)` exists | ❌ | No `Parse` symbol in `/tmp/core-v9-upstream/coreversion/`. Constructor is `coreversion.New.Default(s) Version` (no error return). See **C-CVS-37**. |
| 2 | §1: `v.Major()/Minor()/Patch()` typed accessors | ❌ | `coreversion/Version.go` exposes **public fields** `VersionMajor/VersionMinor/VersionPatch/VersionBuild int` and *string* helpers `MajorString()/MinorString()/PatchString()/BuildString()`. No bare-int methods. See **C-CVS-38**. |
| 3 | §1: `v1.LessThan(v2) / v1.Equal(v2) / v1.GreaterThanOrEqual(v2)` | ❌ | No such methods. Comparison is the package-level `coreversion.Compare(left, right *Version) corecomparator.Compare` (`coreversion/all-compare.go`). See **C-CVS-39**. |
| 4 | §1: `v.String() // "v1.2.3"` | ❌ (partial) | `String()` returns `CompiledVersion()` which returns the `Compiled` field set during creation; it can be empty for `Empty()`/`Invalid` and is not guaranteed to look like `"v1.2.3"`. See **C-CVS-40**. |
| 5 | §1: "Wraps stdlib errors in `errcore.FailedToConvertType`" | ❌ | `grep -n errcore coreversion/*.go` returns **zero hits**. Package never references `errcore` and never returns errors from construction. See **C-CVS-41**. |
| 6 | §1: "Plays well with `coregeneric.Collection`" | ❓ | Real package ships its own `VersionsCollection` (`coreversion/VersionsCollection.go`); spec's `coregeneric.Collection` interop claim is unsubstantiated. Flag for follow-up audit. |
| 7 | §2: Package located at `versionindexes/` (top-level) | ❌ | Real path is `enums/versionindexes/` (sub-directory). Top-level `versionindexes/` does **not** exist. See **C-CVS-42**. |
| 8 | §2: Constants `V1=1, V2=2, V8=8` representing "version eras" | ❌ | Real consts in `enums/versionindexes/Index.go`: `Major=0, Minor=1, Patch=2, Build=3, Invalid=4`. The package indexes **version-component positions**, NOT historical "version eras". The entire conceptual framing of §2 is wrong. See **C-CVS-43**. |
| 9 | §2: "switch on era (`case V7: ... case V8: ...`)" migration pattern | ❌ | No `V7`/`V8` exist. Pattern would not compile. Same root as C-CVS-43. |
| 10 | §3: Release/bump policy ("at least minor; never touch `.release/`") | ✅ | Matches `mem://index.md` Core rule. |
| 11 | §5: `.release/` is off-limits | ✅ | Matches `mem://index.md` Core rule. |

## New contradictions

| ID | Severity | Symptom | Real API | Action |
|---|---|---|---|---|
| C-CVS-37 | CRITICAL | `coreversion.Parse(s) (Version, error)` documented | `coreversion.New.Default(s) Version` (no error; sets `IsInvalid` flag) | AJ-21: rewrite §1 constructor |
| C-CVS-38 | HIGH | Typed accessors `Major()/Minor()/Patch()` documented | Public fields `VersionMajor/Minor/Patch/Build int` + `MajorString()/...` helpers | AJ-22: rewrite §1 accessors |
| C-CVS-39 | CRITICAL | Method-style `LessThan/Equal/GreaterThanOrEqual` documented | Package-level `Compare(left, right *Version) corecomparator.Compare` | AJ-23: rewrite §1 compare |
| C-CVS-40 | LOW | `String()` claimed to always return `"v1.2.3"` | Returns `Compiled` field; can be empty for invalid | AJ-24: clarify semantics |
| C-CVS-41 | HIGH | "Wraps errors in `errcore.FailedToConvertType`" — fabricated rationale | Zero `errcore` references in package | AJ-25: delete bullet from §1 |
| C-CVS-42 | HIGH | Package path `versionindexes/` documented | Real path `enums/versionindexes/` | AJ-26: fix import paths in §2 |
| C-CVS-43 | CRITICAL | Constants `V1/V2/V8` "version eras" documented | Constants `Major/Minor/Patch/Build/Invalid` indexing component positions | AJ-27: rewrite entire §2 |

## Action items spawned

- **AJ-21..27** — All BLOCKED by `spec/01-app/` freeze.

## Recommendations

1. **S-106 lint is now critical** — §11 is the worst section yet (18.2% verifiable, 8 ❌ in 11 claims). Spec author invented an idiomatic Go API surface (`Parse`, methods, fluent `.LessThan()`) instead of the real public-fields + package-level `Compare()` style. S-106 should detect (a) missing top-level functions, (b) missing methods on documented types, (c) wrong import paths.
2. **Conceptual error in §2 is unique** — Cycles 19–22 fabricated APIs that plausibly could exist. C-CVS-43 invents an entirely different *purpose* for the package. Follow-on audits should check whether package-purpose claims match upstream `readme.md` / `doc.go`.
3. **Pattern across 5 cycles** — 34 ❌ / ~55% fabrication rate. Recommend pausing further AB passes until S-106 lint is built.
