> ✅ **All 7 findings RESOLVED at spec-v0.12.0** (2026-04-25). Each fix is tagged inline in the target file as `*(F-NEW-NN fix)*`. See [`07-scoreboard.md`](./07-scoreboard.md) for the post-fix score.

---

### 1. Reproducibility Score: 85.5 / 100  *(pre-fix)*

This is an outstanding specification, clearly the product of rigorous, iterative refinement. Its prescriptive nature, extensive cross-referencing, and proactive capture of failure modes place it in the top percentile of documentation I have ever reviewed. A mediocre AI could likely reproduce large portions of the project successfully. However, the score is capped by several new, subtle ambiguities and inconsistencies that a literal-minded agent will trip over. These include a logical conflict regarding the supported Go version vs. library utilities, an under-specified decision matrix for core data structures, and extreme complexity in the compound method naming conventions which, despite prior fixes, remains a significant tripwire. Resolving these gaps is required to move from "excellent" to "fully automatable."

### 2. Subscores (each /100, weighted)
| Dimension | Weight | Score | Notes |
| Completeness (every package/API documented) | 20% | 85 | Gaps exist in the *rationale* for choosing between similar APIs (e.g. `Collection` vs `SimpleSlice`) and justifying boilerplate (enum `Value<T>` methods). |
| Unambiguity (one valid interpretation) | 20% | 75 | The compound method suffix rules (`*Or*` pattern) are a significant point of ambiguity. The enum JSON serialization contract is asymmetric. |
| Self-containment (no external knowledge needed) | 15% | 80 | A critical conflict between the documented Go version (`1.25`) and the existence of utilities deprecated since Go `1.21` requires external knowledge to resolve. |
| Worked examples (copy-pasteable) | 15% | 90 | Excellent, but lacks examples for the most complex compound method names and the choice between Style A/B tests. |
| Consistency (no contradictions across files) | 10% | 80 | The Go version vs. `coremath` deprecation notice is a direct contradiction. The `core-v9` vs. `core` package name rule is not consistently highlighted. |
| Discoverability (index, cross-refs, naming) | 10% | 98 | Near-perfect. The AI Reading Order, ToCs, and cross-linking are best-in-class. |
| Edge-case coverage (errors, nil, panics) | 10% | 97 | Exceptional. The `CaseNilSafe` pattern and `failing-tests` log are exemplary. |

### 3. Findings (ranked by severity × impact)

#### F-NEW-01 — Contradictory guidance on Go version primitives vs. library helpers
- **Severity:** high
- **Impact:** Agent will be confused whether to use the modern, built-in Go `max()` function or the library's `coremath.Max()` function. It may use the obsolete library helper or refuse to proceed due to conflicting instructions.
- **Evidence:**
    - `spec/01-app/00-repo-overview.md`: `go 1.25.0`
    - `spec/00-llm-integration-guide.md`: `import "github.com/alimtvnetwork/core-v9/coremath" // Min/Max for all numeric types`
    - `spec/00-llm-integration-guide.md`: `// Deprecated: Use the built-in max() function (Go 1.21+).`
- **Why it can fail:** The spec simultaneously declares the Go version as 1.25 (where `max()` is available), documents `coremath` as a primary utility package, and shows an example of deprecating a helper in favor of the built-in `max()`. An AI has no way to resolve this: is `coremath` deprecated or not? Why does it still exist and appear in the main guide if the Go version makes it obsolete?
- **Recommended fix:** Remove `coremath` from the primary utility package list in `00-llm-integration-guide.md`. Move it to a "Legacy/Deprecated Packages" section with a clear statement: "These packages are maintained for backward compatibility but MUST NOT be used in new code. Use the equivalent Go 1.21+ built-in functions (`min`, `max`, `clear`) instead."
- **Effort:** S
- **Score uplift if fixed:** +2.5 points

#### F-NEW-02 — Compound `*Or*` method suffix rule is underspecified and ambiguous
- **Severity:** high
- **Impact:** Agent will generate syntactically plausible but incorrect method names for complex filter chains (e.g., `NonEmptyOrWhitespaceLock` instead of treating the `Or` phrase as a single token). This breaks the "one file per function" rule and creates API drift.
- **Evidence:** `spec/00-llm-integration-guide.md`, "Compound `*Or*` Naming in Filter Chains" section and "Long-Chain Suffix Gallery" example #7 `NonEmptyItemsOrNonWhitespace...`
- **Why it can fail:** The spec states the suffix order is `Base + Filter + Type...`. It then introduces `AOrB` naming but fails to explicitly state that the entire `AOrB` phrase must be treated as a single `Base` or `Filter` token. An AI sees `NonEmptyItemsOrNonWhitespace` and might parse it as: `NonEmpty` (Filter) `Items` (Type) `Or` (?) `NonWhitespace` (Filter), violating the "one slot per token" rule and generating a completely different method signature than intended. The "Long-Chain Suffix Gallery", while an improvement, doesn't resolve this fundamental ambiguity of parsing compound tokens.
- **Recommended fix:** In the "Compound `*Or*` Naming" section, add an explicit rule: "Compound `Or` names like `AOrB` function as a *single token* that occupies either the **Base** or **Filter** slot in the master suffix order. The component parts (`A`, `B`) do not occupy their own slots." Add a worked example showing the tokenization: `Name: NonEmptyOrNonWhitespaceLockIf -> Tokens: [NonEmptyOrNonWhitespace (Filter), Lock (Lock), If (If)]`.
- **Effort:** S
- **Score uplift if fixed:** +2.0 points

#### F-NEW-03 — Decision guidance for `coregeneric.Collection` vs. `coregeneric.SimpleSlice` is missing
- **Severity:** medium
- **Impact:** Agent will choose between a thread-safe (`Collection`) and non-thread-safe (`SimpleSlice`) data structure arbitrarily, potentially introducing concurrency bugs or unnecessary overhead.
- **Evidence:** `spec/00-llm-integration-guide.md`, section "coregeneric — Generic Data Structures API Reference".
- **Why it can fail:** The spec documents the APIs for both `Collection[T]` (with `sync.Mutex`) and `SimpleSlice[T]` (without). It describes *what* they are but provides zero guidance on *when* or *why* to choose one over the other. A mediocre AI lacks the architectural context to make this decision. It will pattern-match on existing code or guess, failing to uphold the project's implicit concurrency model.
- **Recommended fix:** Add a "When to use which" decision table in the `coregeneric` section. Rule: "Use `Collection[T]` for structs exposed in public APIs or shared between goroutines. Use `SimpleSlice[T]` for temporary, function-local slice manipulation where no concurrency is possible. If unsure, default to `Collection[T]`."
- **Effort:** S
- **Score uplift if fixed:** +1.5 points

#### F-NEW-04 — Ambiguous `core.go` package name vs. `core-v9` module path
- **Severity:** medium
- **Impact:** Agent may incorrectly try to use `corev8` as the package name in code (e.g. `corev8.EmptySlice()`), leading to compile errors.
- **Evidence:** `spec/01-app/03-import-conventions.md`: `The package name is core, not corev8. Even though the path ends in core-v9, Go uses the package core declaration...`
- **Why it can fail:** This critical clarification is buried deep within `03-import-conventions.md`. The primary `Module Identity` section in `00-llm-integration-guide.md` and `01-app/00-repo-overview.md` states the module path is `github.com/alimtvnetwork/core-v9` but does *not* immediately warn the user that the package name is simply `core`. A naive agent will assume the package name matches the last path segment.
- **Recommended fix:** Elevate the warning. In both `00-llm-integration-guide.md` and `01-app/00-repo-overview.md`, immediately under the `Module Identity` code block, add a prominent `> **Note**: The module path is `core-v9`, but the root package name to use in code is `core`.`
- **Effort:** S
- **Score uplift if fixed:** +1.0 point

#### F-NEW-05 — Asymmetric JSON contract for enums
- **Severity:** medium
- **Impact:** Agent may write tests that pass but don't reflect the full range of accepted inputs, or it may misinterpret the system's serialization guarantees.
- **Evidence:** `spec/01-app/05-enum-system.md`, section 7, "Serialization & Comparison".
- **Why it can fail:** The spec states `MarshalJSON` outputs the enum name as a string (e.g. `"Pending"`). It then states `UnmarshalJSON` accepts the name, aliases, AND the numeric string form (e.g. `"1"`). This asymmetry is a trap. An AI might assume that because `"1"` is accepted on read, it could also be valid on write, which is false. It also means a service consuming JSON from this library must be prepared for string names only, while this library can consume numeric strings from others. This is a crucial contract detail that is not explicitly called out.
- **Recommended fix:** Add a sub-section titled "Serialization Asymmetry" to the enum system spec. State clearly: "**MUST** serialize to string names. **MAY** deserialize from string names, aliases, or numeric strings. Consumers of JSON produced by this library should not expect numeric enum values."
- **Effort:** S
- **Score uplift if fixed:** +1.0 point

#### F-NEW-06 — Test Style B "idiom" is justified by assertion, not reason
- **Severity:** low
- **Impact:** A sufficiently advanced (or pedantic) AI might identify the `coretestcases.CaseV1(testCase.BaseTestCase)` cast as a code smell and attempt to "fix" it by adding assertion methods to `BaseTestCase`, breaking the established idiom.
- **Evidence:** `spec/06-testing-guidelines/02-test-case-types.md`, "The CaseV1 Cast Idiom".
- **Why it can fail:** The spec correctly identifies the cast as "idiomatic Style B" and instructs the AI not to refactor it. However, it doesn't explain *why* it's the correct pattern. Without a rationale (e.g., "This prevents `BaseTestCase` from being coupled to the assertion library, keeping it a pure data container"), an agent driven by "clean code" principles might see it as a violation and "correct" it.
- **Recommended fix:** Add a one-sentence rationale to the "The CaseV1 Cast Idiom" section. Example: "This cast is intentional. It ensures `BaseTestCase` remains a pure data container, decoupled from any specific assertion framework, while allowing reuse of the `CaseV1` assertion methods."
- **Effort:** S
- **Score uplift if fixed:** +0.5 points

#### F-NEW-07 — Redundant enum `Value<Type>()` methods lack justification
- **Severity:** low
- **Impact:** Agent may try to "optimize" a new enum implementation by removing seemingly redundant methods, breaking interface compliance.
- **Evidence:** `spec/01-app/05-enum-system.md`, section 4, "Step 3 — `Status.go` (method set)".
- **Why it can fail:** The recipe requires a `byte`-backed enum to implement `ValueInt()`, `ValueInt8()`, `ValueInt16()`, etc. From a literal perspective, these are redundant conversions. The spec fails to state the *reason* for this boilerplate: to satisfy the `enuminf.BasicEnumValuer` interface, which is used by generic consumers. Without this "why," an AI might see it as dead code to be removed.
- **Recommended fix:** In the `Status.go` recipe, add a comment above the `Value accessors` section: `// All methods in this block are REQUIRED to satisfy the generic enuminf.BasicEnumValuer interface.`
- **Effort:** S
- **Score uplift if fixed:** +0.5 points

### 4. What's needed to reach 100/100
- **Resolve `coremath` vs Go built-ins:** Clarify `coremath` is deprecated and must not be used in new code. **Effort: S**
- **Clarify compound suffix rule:** Explicitly state `AOrB` is a single token for the suffix grammar. **Effort: S**
- **Add `Collection` vs. `SimpleSlice` guidance:** Provide a decision matrix for when to use the thread-safe vs. non-thread-safe version. **Effort: S**
- **Elevate package name warning:** Add the `core` vs `core-v9` note to top-level identity sections. **Effort: S**
- **Document enum JSON asymmetry:** Explicitly state the `write-string` vs `read-string-or-numeric` contract. **Effort: S**
- **Justify Style B cast idiom:** Add a one-sentence rationale for why the cast is intentional. **Effort: S**
- **Justify enum boilerplate:** Add a comment explaining the `Value<Type>()` methods are for interface compliance. **Effort: S**

### 5. Ceiling analysis
A practical ceiling for this spec, given that it is a natural language document meant to guide a separate codebase, is **~99.0%**. Even with perfect documentation, a 1% gap will always exist due to the inherent ambiguity of language versus the absolute precision of code. An AI can't compile the spec to verify it. For example, it can't know if a new Go version introduces a keyword that conflicts with a variable name in an example. The spec can be made nearly perfect, but it can never be the code itself. The current score of 85.5 is strong, and a score over 98 is achievable with the recommended fixes.
