> ✅ **All 9 findings RESOLVED at spec-v0.14.0** (2026-04-25). Each fix is tagged inline in the target file as `*(F-V12-NN fix)*` or `*(F-V12-NN)*`. See [`07-scoreboard.md`](./07-scoreboard.md) for the post-fix score.

---

### 1. Reproducibility Score: 88.9 / 100  *(pre-fix baseline)*
The spec is very strong overall: clear AI reading order, rich worked examples, and excellent testing guidance. However, several inconsistencies and underspecified areas remain that would trip a GPT‑3.5‑class agent. The biggest risks are naming-grammar ambiguities (missing “Result” slot in the suffix order), cross-file contradictions (“coremath” legacy status vs package map), undeclared or unclear APIs (errcore.HandleErr, “iter” iterators), and a handful of broken/ambiguous references (missing diagnostics doc, Unmarshal/Unmarshall mismatch). These do not wreck the spec, but they materially reduce the odds of reimplementing at 100% fidelity without back-and-forth.

### 2. Subscores (each /100, weighted)
| Dimension | Weight | Score | Notes |
| Completeness | 20% | 90 | Most surfaces covered; gaps: errcore.HandleErr not documented; “iter” iterator imports; BasicEnumValuer not indexed in enuminf table. |
| Unambiguity | 20% | 86 | Suffix-order grammar omits “Result” slot; Unmarshal/Unmarshall spelling inconsistency. |
| Self-containment | 15% | 86 | “iter” iterator types lack import path/contract; missing referenced diagnostics doc. |
| Worked examples | 15% | 92 | Excellent breadth; add one worked example clarifying Result-slot ordering for long names. |
| Consistency | 10% | 82 | “coremath” legacy vs package map; constants rule vs code examples; Hashmap/Hashset return semantics unclear. |
| Discoverability | 10% | 92 | Great cross-linking; a couple of broken/missing refs (diagnostics doc) drag this slightly. |
| Edge-case coverage | 10% | 95 | Nil-safety, branch coverage, and failing-tests learnings are strong. |

### 3. Verification of prior fixes (F-NEW-01..F-NEW-07)
- F-NEW-01 — coremath vs built-ins: ⚠️ Partial — spec/00-llm-integration-guide.md clearly deprecates coremath, but spec/01-app/01-package-map.md still lists “coremath — Min/Max for all numeric types” without a LEGACY label; also spec/01-app/00-repo-overview.md §2 groups coremath among active utilities.
- F-NEW-02 — Compound *Or* tokenization: ✅ Adequate — explicit rule block with correct/wrong tokenization and a 3-segment example (GetOrKeyOrDefault) in spec/00-llm-integration-guide.md.
- F-NEW-03 — Collection vs SimpleSlice decision: ✅ Adequate — decision matrix and default-to-Collection in spec/00-llm-integration-guide.md §coregeneric.
- F-NEW-04 — core vs core-v9 warning elevation: ✅ Adequate — top-level callouts in spec/00-llm-integration-guide.md and spec/01-app/00-repo-overview.md.
- F-NEW-05 — Enum JSON asymmetry: ✅ Adequate — dedicated “Serialization Asymmetry” box in spec/01-app/05-enum-system.md §7.
- F-NEW-06 — Style B cast idiom rationale: ✅ Adequate — clear rationale in spec/06-testing-guidelines/02-test-case-types.md §Style B.
- F-NEW-07 — Enum Value<Type> justification: ✅ Adequate — explicit “REQUIRED by enuminf.BasicEnumValuer” note in spec/01-app/05-enum-system.md Step 3.

### 4. NEW Findings (ranked by severity × impact)

#### F-V12-01 — Missing “Result” slot in the suffix-order grammar
- Severity: high
- Impact: AI may generate wrong long-chain method names (e.g., put Ptr before OrDefaultWith) due to contradictory rules.
- Evidence:
  - spec/00-llm-integration-guide.md §“Pattern 7: Combined Suffixes & Ordering” states order as “Base + Filter + Type + Lock + If + Must” (no “Result”).
  - Same file §“Long-Chain Suffix Gallery”, Example 3 says “Result modifiers sit BEFORE Type per fixed slot order” and annotates “Result=OrDefaultWith”.
- Why it can fail: The global order omits “Result” but examples depend on it; agent can’t reconcile where OrDefault/OrDefaultWith/New belongs in the token sequence.
- Recommended fix: Amend the canonical order to: Base + Filter + Result + Type + Lock + If + Must. Add one worked example showing two comparable names differing only in Result placement, with a red/green diff of the wrong vs right name.
- Effort: S
- Score uplift if fixed: +1.8

#### F-V12-02 — “coremath” legacy status inconsistent across files
- Severity: medium
- Impact: AI may keep using coremath in new code due to authoritative package map and overview not signaling deprecation.
- Evidence:
  - spec/00-llm-integration-guide.md §coremath: “LEGACY — do not use in new code.”
  - spec/01-app/01-package-map.md: “coremath — Min/Max for all numeric types” (no LEGACY).
  - spec/01-app/00-repo-overview.md §2 Top-Level Layout: groups coremath with active utility packages.
- Why it can fail: Conflicting cues between the integration guide and the canonical package map/overview.
- Recommended fix: In both spec/01-app/01-package-map.md and spec/01-app/00-repo-overview.md, append “(LEGACY — use Go 1.21+ min/max/clear built-ins)” to coremath entries.
- Effort: S
- Score uplift if fixed: +1.2

#### F-V12-03 — Iterator (“iter”) import/contract ambiguity
- Severity: medium
- Impact: AI may not know how to compile All()/Values() iterators or what to import, leading to type or build errors.
- Evidence:
  - spec/00-llm-integration-guide.md §coregeneric/Collection and SimpleSlice “Iterators (Go 1.23+ iter package)” with examples returning iter.Seq / iter.Seq2, but no import path or signature.
- Why it can fail: Missing import path and exact type signatures for public APIs returning iter types; unclear whether this is stdlib or third-party.
- Recommended fix: Add an “Iterator types” box in coregeneric that specifies the exact return signatures and the import path (e.g., “import ‘iter’” if stdlib; or if aliased types, document the alias). Include one compile-ready minimal example with imports.
- Effort: M
- Score uplift if fixed: +1.6

#### F-V12-04 — Constants rule contradicts examples
- Severity: medium
- Impact: AI may mechanically replace literal "" with constants.EmptyString everywhere, or vice versa, yielding style drift and noisy diffs.
- Evidence:
  - spec/00-llm-integration-guide.md §Code Style Rules: “Constants — Always use constants.X — never hardcode "" …”
  - Same file §Pattern 6 Filtering Variants examples use raw “if str == "" { … }”.
- Why it can fail: The hard “never hardcode” rule conflicts with pervasive examples comparing against "" directly.
- Recommended fix: Refine the rule to: “Use constants.X for emitted/assigned values and formatting; direct comparisons against "" in fast paths are allowed.” Update 2–3 code samples accordingly or annotate them as permitted exceptions.
- Effort: S
- Score uplift if fixed: +1.1

#### F-V12-05 — errcore.HandleErr not documented in errcore API
- Severity: medium
- Impact: AI will guess its signature or panic handling contract for Must-methods, risking divergence.
- Evidence:
  - spec/00-llm-integration-guide.md §Method Writing: *Must shows errcore.HandleErr(err).
  - spec/01-app/04-error-system.md lists errcore API but omits HandleErr.
- Why it can fail: Must-variants depend on HandleErr semantics (panic behavior, stack enrichment). Omitted from API reference.
- Recommended fix: Add HandleErr to spec/01-app/04-error-system.md “Public API — errcore” with signature, behavior (panic on non-nil), and a short example. Reference it from the Must pattern.
- Effort: S
- Score uplift if fixed: +1.3

#### F-V12-06 — Unmarshal vs Unmarshall naming inconsistency for enum helpers
- Severity: medium
- Impact: AI may implement with standard “Unmarshal...” names and miss the nonstandard “UnmarshallEnumToValue” call points, breaking integration.
- Evidence:
  - spec/01-app/05-enum-system.md Step 3: “func (it Status) UnmarshallEnumToValue(...)” (double “l”).
  - Same file and JSON methods use UnmarshalJSON (single “l”).
- Why it can fail: Contradictory spelling for similarly named concepts; unclear if “Unmarshall...” is deliberate or a typo in the engine’s API.
- Recommended fix: Add a bold note in §Step 3 clarifying the exact method names provided by enumimpl (single vs double “l”). If it is a historical quirk, call it out explicitly so implementers don’t “correct” it; otherwise, standardize the spelling in the spec.
- Effort: S
- Score uplift if fixed: +1.0

#### F-V12-07 — Hashmap/Hashset API return semantics unclear/inconsistent
- Severity: medium
- Impact: AI may implement incompatible signatures, breaking fluent chains or tests.
- Evidence:
  - spec/00-llm-integration-guide.md §coregeneric/Hashmap: “Set(key,val) → bool (isNew)” vs “SetLock(key,val) → *Hashmap” (fluent). No rationale.
  - spec/00-llm-integration-guide.md §coregeneric/Hashset: “AddBool(key) — bool — Returns true if already existed” (counterintuitive; typical sets return true when newly added).
- Why it can fail: Mixed return-style across lock/non-lock variants without explicit rationale; nonstandard AddBool semantics.
- Recommended fix: Add a “Return conventions” box per type:
  - Hashmap: Set returns isNew (bool), SetLock returns *Hashmap for chaining; explain why (locking pattern consistency vs informative return).
  - Hashset: Define AddBool’s truth semantics precisely with examples; if “already existed” is intended, call it “ExistsOrAdded” or add second variant “AddNewBool”.
- Effort: M
- Score uplift if fixed: +1.4

#### F-V12-08 — Broken/missing reference to diagnostics standards
- Severity: low
- Impact: AI can’t find “07-diagnostics-output-standards.md” referenced for output formatting; loses single source for formatting rules.
- Evidence:
  - spec/01-app/13-testing-patterns.md §8 “Coverage Expectations” references “/spec/06-testing-guidelines/07-diagnostics-output-standards.md” — this file is not present in this bundle, while formatting rules live in spec/01-app/04-error-system.md §5.
- Why it can fail: Dead link; agent may miss the canonical format and invent one.
- Recommended fix: Either add the missing 07-diagnostics-output-standards.md (can be a stub pointing to 01-app/04-error-system.md §5), or update references to the existing diagnostics section.
- Effort: S
- Score uplift if fixed: +0.6

#### F-V12-09 — enuminf index omits BasicEnumValuer (referenced elsewhere)
- Severity: low
- Impact: AI may not realize the interface that forces all Value<Type>() methods exists or how it composes.
- Evidence:
  - spec/01-app/05-enum-system.md Step 3 references enuminf.BasicEnumValuer requirements.
  - spec/01-app/05-enum-system.md §2 “Key interfaces” table does not list BasicEnumValuer.
- Why it can fail: The contract name is referenced but not discoverable in the key interface inventory.
- Recommended fix: Add BasicEnumValuer to the “Key interfaces” table with short purpose (“comprehensive typed value accessors for generic consumers”).
- Effort: S
- Score uplift if fixed: +0.5

### 5. What’s needed to reach 100/100
- Formalize the suffix-order grammar to include the Result slot and add a confirming worked example. (S)
- Make “coremath” legacy labeling consistent in package map and repo overview. (S)
- Document iterator (“iter”) import path and exact All()/Values() signatures with a compile-ready snippet. (M)
- Reconcile the constants rule with examples by defining explicit exceptions for fast-path comparisons, or refactor examples. (S)
- Add errcore.HandleErr to the errcore API reference with signature and behavior. (S)
- Clarify Unmarshal/Unmarshall naming for enumimpl helpers (either standardize or document the quirk). (S)
- Specify Hashmap/Hashset return-type conventions and AddBool semantics, with examples. (M)
- Fix the diagnostics standards dead link (create the doc or update the reference). (S)
- Add BasicEnumValuer to the enuminf key-interfaces index. (S)

### 6. Ceiling analysis
Even with the above fixes, a natural-language spec can’t reach perfect determinism: naming grammars, cross-version Go features (e.g., “iter”), and subtle historical quirks (e.g., Unmarshal(l)) always leave a sliver of interpretation. With the recommended changes, a realistic ceiling is ~97–98%. Achieving 100% would require turning this spec into machine-checkable contracts (interfaces plus compile-checked stubs), which is beyond a prose document’s scope.
