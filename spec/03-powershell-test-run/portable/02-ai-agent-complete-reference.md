# Generic Go Test & Coverage Runner — AI Agent Reference

> **Purpose**: This document is a self-contained reference for any AI agent working on a Go project
> that uses PowerShell-based test/coverage tooling. It covers build fix workflows, test writing
> guidelines, coverage push protocols, and the run.ps1 architecture — all with generic examples
> that apply to any Go module.

> **Consumer-coverage note (`enum-v7`)** — every reference below to `tests/integratedtests/<pkg>tests/`, the per-package directory tree, and example package names (`pkgtests`, `corejsontests`, etc.) describes the **upstream `core-v9`** convention. `enum-v7` itself uses a single shared `tests/creationtests/` package (Goconvey-based registry over `EnumTestWrapper`) — see [`/spec/01-app/13-testing-patterns.md` §6.1](../01-app/13-testing-patterns.md#61-enum-v7-specific-layout) and [`/spec/01-app/14-tests-folder-walkthrough.md`](../01-app/14-tests-folder-walkthrough.md). The `run.ps1` implementation is layout-agnostic (reads from disk via `go list ./tests/...`) and works on either name without modification.

---

## Table of Contents

1. [Project Architecture](#1-project-architecture)
2. [run.ps1 — Command Reference](#2-runps1--command-reference)
3. [Build Error Diagnosis & Fix Protocol](#3-build-error-diagnosis--fix-protocol)
4. [Test Writing Guidelines](#4-test-writing-guidelines)
5. [Unit Coverage Fix Protocol](#5-unit-coverage-fix-protocol)
6. [Coverage Push Iteration Workflow](#6-coverage-push-iteration-workflow)
7. [Common Error Patterns & Fixes](#7-common-error-patterns--fixes)
8. [PowerShell Runner Internals](#8-powershell-runner-internals)
9. [Error Attribution System](#9-error-attribution-system)
10. [Related Spec Files](#10-related-spec-files)

---

## 1. Project Architecture

### Directory Layout

```
<project-root>/
├── run.ps1                              # Thin dispatcher (~167 lines)
├── scripts/
│   ├── README.md                        # Module documentation + dependency graph
│   ├── DashboardUI.psm1                 # ANSI dashboard rendering, phase tracking
│   ├── Utilities.psm1                   # Console helpers, error extraction
│   ├── TestLogWriter.psm1               # Go test output → structured log files
│   ├── TestRunner.psm1                  # Test execution, build checks, git ops
│   ├── CoverageRunner.psm1             # TC + TCP coverage pipelines
│   ├── BuildTools.psm1                  # Build, format, vet, tidy, clean
│   ├── PreCommitCheck.psm1              # PC pre-commit validation
│   ├── GoConvey.psm1                    # GoConvey browser test runner
│   ├── Help.psm1                        # Help, fail log, integrated tests
│   ├── bracecheck/main.go               # Go syntax pre-checker
│   ├── autofix/main.go                  # Auto-fixer for common syntax issues
│   ├── check-safetest-boundaries.ps1    # SafeTest lint checker
│   ├── check-integrated-regressions.ps1 # API-drift regression scanner
│   └── coverage/
│       └── Export-UncoveredMethodsJson.ps1
├── data/
│   ├── coverage/                        # Coverage reports (generated)
│   │   ├── partial/                     # Per-package .out profiles
│   │   ├── coverage.out                 # Merged coverage profile
│   │   ├── coverage.html                # Visual HTML report
│   │   ├── coverage-summary.json        # Machine-readable summary
│   │   ├── per-package-coverage.json    # Per-package breakdown
│   │   ├── blocked-packages.txt         # Compile failures
│   │   ├── blocked-packages.json        # Machine-readable blocked
│   │   ├── coverage-previous.json       # Snapshot for regression diff
│   │   └── build-errors.txt             # Build error details
│   └── test-logs/                       # Test output (generated)
│       ├── raw-output.txt               # Full go test output
│       ├── passing-tests.txt
│       └── failing-tests.txt
├── <source-packages>/                   # Production Go code
│   ├── pkgA/
│   │   ├── SomeType.go
│   │   └── unexported_test.go           # In-package tests for unexported symbols
│   └── pkgB/
└── tests/
    └── integratedtests/                 # All test packages
        ├── pkgAtests/
        │   ├── Feature_test.go          # Test logic (AAA pattern)
        │   └── Feature_testcases.go     # Test data (named variables)
        └── pkgBtests/
```

### Key Conventions

| Convention | Rule |
|-----------|------|
| **Test location** | All `_test.go` files in `tests/integratedtests/<pkg>tests/` |
| **In-package tests** | Only for unexported symbols; must stay inside the package |
| **Internal packages** | Skip entirely — never write coverage tests for `internal/` |
| **Test naming** | `Test_Cov{N}_{Method}_{Context}` for coverage tests |
| **File separation** | `_test.go` = logic only; `_testcases.go` = data only |

---

## 2. run.ps1 — Command Reference

### Commands

| Short | Flag | Long | Description |
|-------|------|------|-------------|
| `T` | `-t` | `test` | Run all tests (verbose, with log output) |
| `TP` | `-tp` | `test-pkg` | Run tests for a specific package |
| `TC` | `-tc` | `test-cover` | Run tests with coverage (parallel default) |
| `PC` | `-pc` | `pre-commit` | Pre-commit compile check (Coverage files only) |
| `TF` | `-tf` | `test-fail` | Show last failing tests log |
| `C` | `-c` | `clean` | Clean build artifacts |
| `H` | `-h` | `help` | Show help |

### Usage

```powershell
./run.ps1 TC                # Full coverage run (parallel)
./run.ps1 TC --sync         # Sequential mode
./run.ps1 TC --no-open      # Skip auto-opening HTML report
./run.ps1 PC                # Pre-commit compile check only
./run.ps1 TP regexnewtests  # Run a specific test package
```

### TC Pipeline (What Happens)

```
1. go mod tidy              # Dependency sync
2. Syntax pre-check         # bracecheck + autofix + safeTest lint
3. Package discovery        # List source + test packages
4. Pre-coverage compile     # go test -c each test package (gate)
5. Coverage run             # go test -coverprofile each compilable package
6. Profile merge            # MAX-count merge of partial profiles
7. Report generation        # HTML + JSON + TXT summaries
8. Console output           # 4-section boxed summary
```

### Output Files

| File | Purpose |
|------|---------|
| `data/coverage/coverage.out` | Raw merged Go coverage profile |
| `data/coverage/coverage.html` | Visual HTML report |
| `data/coverage/coverage-summary.json` | Machine-readable summary |
| `data/coverage/per-package-coverage.json` | Per-package breakdown |
| `data/coverage/blocked-packages.json` | Packages that failed to compile |
| `data/coverage/build-errors.txt` | Detailed build error output |
| `data/test-logs/failing-tests.txt` | Failing test details |

### Console Output Sections (TC)

The TC command prints **exactly four boxed sections**:

1. **BLOCKED PACKAGES** — packages that failed `go test -c`
2. **FAILING TESTS** — test functions that produced `--- FAIL:`
3. **COVERAGE SUMMARY** — per-source-package table sorted by % descending
4. **WRITTEN FILES** — list of generated report file paths

---

## 3. Build Error Diagnosis & Fix Protocol

### Priority Order

**Always fix in this order:**

1. **Build/compilation errors** — the project must compile before anything else
2. **Blocked packages** — test packages that fail `go test -c`
3. **Failing tests** — tests that compile but produce `--- FAIL:`
4. **Coverage gaps** — write new tests for uncovered paths

### Step-by-Step Build Fix Workflow

#### Step 1: Get the Error Files

Before attempting any fix, obtain:
- `build-errors.txt` or `blocked-packages.txt` — compilation errors
- `failing-tests.txt` — test failures
- `coverage.out` — coverage profile (for coverage work)

**If not provided, ask the user for them.**

#### Step 2: Read the Source Before Editing

> **CRITICAL RULE**: Always read the actual source file before writing or modifying any code.
> Never infer function signatures, method receivers, or parameter orders from naming patterns.

This prevents the #1 cause of repeated failures: **API hallucination**.

#### Step 3: Categorize the Errors

Parse each error line (format: `file.go:line:col: message`) and categorize:

| Category | Pattern | Example |
|----------|---------|---------|
| **Undefined reference** | `undefined: symbolName` | `undefined: unitTestGenerator` |
| **Argument count** | `too many arguments` / `not enough arguments` | `too many arguments in call to pkg.Func` |
| **Type mismatch** | `cannot use X as type Y` | `cannot use string as type int` |
| **Missing member** | `has no field or method` | `pkg.Type has no field or method Foo` |
| **Argument swap** | Compiles but wrong behavior | `SetBySplitter("=", "key=value")` → `SetBySplitter("key=value", "=")` |
| **Duplicate** | `redeclared in this block` | `testStringer redeclared` |
| **Import** | `imported and not used` / `could not import` | `imported and not used: "fmt"` |

#### Step 4: Fix One Package at a Time

```
1. Pick the first blocked package
2. Read each error's source file (the file BEING CALLED, not just the test)
3. Fix the test code to match the actual API
4. Verify: run `./run.ps1 PC` to confirm compilation
5. Move to the next package
```

### Generic Fix Examples

#### Example 1: Undefined Reference

**Error:**
```
cmd/main/main.go:4:14: undefined: unitTestGenerator
```

**Diagnosis**: A variable or function was removed but still referenced.

**Fix approach**:
1. Read the file containing the reference
2. Determine if the symbol was renamed, moved, or deleted
3. Update or remove the reference

```go
// BEFORE (broken)
func main() {
    Generator.Generate()  // Generator was removed
}

// AFTER (fixed)
func main() {
    fmt.Println("app started")
}
```

#### Example 2: Argument Count Mismatch

**Error:**
```
tests/integratedtests/pkgtests/Coverage5_test.go:14:2: too many arguments in call to pkg.SomeFunc
    have (string, int, bool)
    want (string, int)
```

**Fix approach**:
1. Read the actual function signature: `func SomeFunc(name string, count int) error`
2. Remove the extra argument in the test

```go
// BEFORE (broken)
result := pkg.SomeFunc("test", 42, true)

// AFTER (fixed)
result := pkg.SomeFunc("test", 42)
```

#### Example 3: Argument Order Swap

**Error**: Test compiles but assertion fails because arguments are swapped.

**Fix approach**:
1. Read the function signature to check parameter order
2. Swap the arguments in the call

```go
// Source function: func SetBySplitter(input string, separator string) *Pair
//
// BEFORE (wrong order — compiles but wrong behavior)
pair := pkg.SetBySplitter("=", "key=value")

// AFTER (correct order)
pair := pkg.SetBySplitter("key=value", "=")
```

#### Example 4: Method Renamed/Moved

**Error:**
```
pkg.Type has no field or method OldName
```

**Fix approach**:
1. Search the source package for the new name
2. Update the test call

```go
// BEFORE (old API)
result := obj.IsOutOfRange(42)

// AFTER (new API)
result := obj.IsOutOfRange.Integer(42)
```

#### Example 5: Type Constructor Changed

**Error:**
```
cannot use Variant{} as type *Variant in argument
```

**Fix approach**:
1. Check if the type now uses pointer receivers
2. Use the correct constructor

```go
// BEFORE (value type)
v := Variant{}

// AFTER (pointer type)
v := new(Variant)
// or
v := &Variant{}
```

#### Example 6: Duplicate Declaration

**Error:**
```
errForTest redeclared in this block
```

**Fix approach**:
1. Search for all declarations of the symbol in the package
2. Rename or remove the duplicate

```go
// File A: var errForTest = errors.New("test error A")
// File B: var errForTest = errors.New("test error B")  // ← conflict
//
// Fix: rename in File B
// File B: var errForTestB = errors.New("test error B")
```

---

## 4. Test Writing Guidelines

### Mandatory Rules

#### 4.1 Assertion Style

- **NEVER** use `t.Error`, `t.Fail`, `t.Fatalf`, or `t.Errorf`
- **ALWAYS** use framework assertions:
  - `CaseV1.ShouldBeEqual(t, idx, actuals...)` — primary
  - `CaseV1.ShouldBeEqualMap(t, idx, actualMap)` — multi-property
  - `errcore.AssertDiffOnMismatch(t, idx, title, actLines, expLines)` — diff-based

#### 4.2 AAA Pattern (Strictly Enforced)

Every test function **MUST** have three clearly separated and commented sections:

```go
func Test_Cov1_SomeMethod_ValidInput(t *testing.T) {
    tc := someMethodValidTestCase

    // Arrange
    input := tc.ArrangeInput.(args.Map)
    value, _ := input.GetAsString("input")

    // Act
    result := pkg.SomeMethod(value)
    actual := args.Map{
        "output":  result.String(),
        "isValid": result.IsValid(),
    }

    // Assert
    tc.ShouldBeEqualMap(t, 0, actual)
}
```

#### 4.3 Data-Logic Separation

| File | Contains | Never Contains |
|------|----------|---------------|
| `_testcases.go` | Named `CaseV1` variables, `args.Map` inputs, expected values | Test functions, assertions, `func Test_*` |
| `_test.go` | `func Test_*` with AAA pattern | Expected values, hardcoded assertions |

#### 4.4 Map Formatting

Each key-value pair on a **separate line**. Never inline.

```go
// ✅ CORRECT
ArrangeInput: args.Map{
    "name":  "Alice",
    "age":   30,
    "role":  "admin",
}

// ❌ WRONG
ArrangeInput: args.Map{"name": "Alice", "age": 30, "role": "admin"}
```

#### 4.5 Test Naming

| Pattern | Usage |
|---------|-------|
| `Test_Cov{N}_{Method}_{Context}` | Coverage tests |
| `Test_{Type}_{Method}_{Scenario}` | Feature tests |

Title format: `"{Function} returns {Result} -- {Input Context}"`

```go
// ✅ Good titles
"IsValid returns true -- given well-formed email"
"Parse returns error -- empty input string"
"Clone preserves all fields -- struct with nested pointers"

// ❌ Bad titles
"Test valid case"
"Error handling"
"It works"
```

#### 4.6 Coverage Requirements

Tests must cover:
- ✅ Normal execution paths (happy path)
- ✅ Edge cases (empty input, zero values, max values)
- ✅ Error handling paths (nil input, invalid args, missing data)
- ✅ Boundary conditions (off-by-one, length limits)
- ✅ All branches (if/else, switch cases, nil guards)

Tests must be:
- ✅ Deterministic (no random, no time-dependent)
- ✅ Non-flaky (passes 100% of the time)
- ✅ Independent (no test depends on another test's state)

### Complete Test Example

#### `_testcases.go`:

```go
package pkgtests

import (
    "github.com/org/project/coretests/args"
    "github.com/org/project/coretests/coretestcases"
)

// ── Normal path ──

var parseValidTestCase = coretestcases.CaseV1{
    Title: "Parse returns valid result -- well-formed input 'key=value'",
    ArrangeInput: args.Map{
        "input":     "key=value",
        "separator": "=",
    },
    ExpectedInput: args.Map{
        "left":    "key",
        "right":   "value",
        "isValid": true,
    },
}

// ── Edge case: empty input ──

var parseEmptyTestCase = coretestcases.CaseV1{
    Title: "Parse returns empty pair -- empty input string",
    ArrangeInput: args.Map{
        "input":     "",
        "separator": "=",
    },
    ExpectedInput: args.Map{
        "left":    "",
        "right":   "",
        "isValid": false,
    },
}

// ── Error path: no separator found ──

var parseNoSeparatorTestCase = coretestcases.CaseV1{
    Title: "Parse returns full string as left -- no separator in input 'hello'",
    ArrangeInput: args.Map{
        "input":     "hello",
        "separator": "=",
    },
    ExpectedInput: args.Map{
        "left":    "hello",
        "right":   "",
        "isValid": false,
    },
}
```

#### `_test.go`:

```go
package pkgtests

import (
    "testing"

    "github.com/org/project/coretests/args"
    "github.com/org/project/pkg"
)

func Test_Cov1_Parse_ValidInput(t *testing.T) {
    tc := parseValidTestCase

    // Arrange
    input := tc.ArrangeInput.(args.Map)
    str, _ := input.GetAsString("input")
    sep, _ := input.GetAsString("separator")

    // Act
    result := pkg.Parse(str, sep)
    actual := args.Map{
        "left":    result.Left,
        "right":   result.Right,
        "isValid": result.IsValid,
    }

    // Assert
    tc.ShouldBeEqualMap(t, 0, actual)
}

func Test_Cov1_Parse_EmptyInput(t *testing.T) {
    tc := parseEmptyTestCase

    // Arrange
    input := tc.ArrangeInput.(args.Map)
    str, _ := input.GetAsString("input")
    sep, _ := input.GetAsString("separator")

    // Act
    result := pkg.Parse(str, sep)
    actual := args.Map{
        "left":    result.Left,
        "right":   result.Right,
        "isValid": result.IsValid,
    }

    // Assert
    tc.ShouldBeEqualMap(t, 0, actual)
}

func Test_Cov1_Parse_NoSeparator(t *testing.T) {
    tc := parseNoSeparatorTestCase

    // Arrange
    input := tc.ArrangeInput.(args.Map)
    str, _ := input.GetAsString("input")
    sep, _ := input.GetAsString("separator")

    // Act
    result := pkg.Parse(str, sep)
    actual := args.Map{
        "left":    result.Left,
        "right":   result.Right,
        "isValid": result.IsValid,
    }

    // Assert
    tc.ShouldBeEqualMap(t, 0, actual)
}
```

### Nil-Safety Testing Pattern

For testing nil receiver safety on pointer methods:

```go
// _testcases.go
var myTypeNilSafeTestCases = []coretestcases.CaseNilSafe{
    {
        Title: "IsValid on nil returns false",
        Func:  (*MyType).IsValid,
        Expected: results.ResultAny{
            Value:    "false",
            Panicked: false,
        },
    },
    {
        Title: "Name on nil returns empty string",
        Func: func(m *MyType) string {
            return m.Name()
        },
        Expected: results.ResultAny{
            Value:    "",
            Panicked: false,
        },
    },
}

// _test.go
func Test_MyType_NilReceiver(t *testing.T) {
    for caseIndex, tc := range myTypeNilSafeTestCases {
        tc.ShouldBeSafe(t, caseIndex)
    }
}
```

### In-Package Tests (Unexported Symbols)

When functions/types are unexported, tests **must** stay inside the package:

```go
// Inside pkgA/unexported_test.go (same package)
package pkgA

import "testing"

func Test_Cov1_internalHelper_ValidInput(t *testing.T) {
    // Arrange
    input := "test-data"

    // Act
    result := internalHelper(input)

    // Assert
    if result != "expected" {
        t.Fatal("expected 'expected', got:", result)
    }
}
```

> **Note**: In-package tests use `t.Fatal` since they can't import the test framework
> (which would create circular imports). This is the **only** exception to the
> "no raw t.Error" rule.

---

## 5. Unit Coverage Fix Protocol

### Status: ✅ COMPLETE (as of 2026-04-06)

All non-internal packages have achieved **100% reachable code coverage**. The protocol remains documented for maintenance and regression handling.

### Completed Packages (21 packages at 100%)

`corecmp`, `codestack`, `corepayload`, `corejson`, `coretests/results`, `reflectmodel`, `coretests`, `corevalidator`, `chmodhelper`, `coredynamic`, `enumimpl`, `errcore`, `corestr`, `coretests/args`, `coretests/coretestcases`, `namevalue`, `stringslice`, `corerange`, `stringutil`, `coreversion`, `coreonce`

### Documented Unreachable Gaps (accepted, not bugs)

| Package | Gap | Reason | Documentation |
|---------|-----|--------|---------------|
| `stringutil` | `IsEndsWith.go:37` (`remainingLength < 0`) | Prior length check makes this unreachable | `Coverage7_Gaps_test.go` |
| `coreversion` | `hasDeductUsingNilNess.go:20` | Exhaustive nil checks above | `Coverage6_DeadCode_test.go` |
| `coreonce` | `JsonStringMust` error branches | `json.Marshal` cannot fail on simple maps/slices | `Coverage16_Gaps_test.go` |

### Trigger

When the user says **"fix unit test"**, **"Unit Coverage Fix"**, or **"next"**, execute this protocol for **maintenance/regression** handling.

### Objectives

1. Fix build issues, runtime failures, blocked packages, and failing tests **first**
2. Move all tests from inside packages to `tests/integratedtests/<pkg>tests/`
3. Fix all assertion, formatting, and structural violations
4. Maintain 100% code coverage across all non-internal packages
5. Enhance testing guidelines where gaps exist

### Prerequisites

- The `.out` file and related JSON coverage files **must be provided**. If not given, **ask for them**.
- If build issues exist, **ask for related files** (stack traces, error logs, source files).
- **Read the source** before writing any tests — never infer APIs.
- You do not need to instruct the user to run TC. The user handles that.

### Execution Priority

```
1. Build issues     → fix compilation errors first
2. Blocked packages → resolve dependencies, runtime errors
3. Failing tests    → fix existing broken tests
4. Migration        → move tests to integration folder
5. Refactoring      → apply AAA, fix assertions, format maps
6. Coverage gaps    → write new tests (2 packages per iteration)
```

### Failure Isolation (Split Recovery)

If a package fails repeatedly during testing, **split monolithic test files into per-method granular files**. This ensures a single `-TC` run captures all distinct failures simultaneously without one error blocking compilation or execution of others.

Four packages were restructured using this approach: `chmodhelpertests`, `coredynamictests`, `corestrtests`, `corepayloadtests`.

### Skip Rules

- **Internal packages**: Skip entirely — never write coverage tests for `internal/`
- **Private methods**: Discuss with user whether to skip or test indirectly
- **OS-dependent code**: Some paths (Linux chmod, Windows paths) can only be covered on specific platforms

### Acceptance Criteria

1. Zero test files inside package directories (except unexported symbol tests)
2. All tests in `tests/integratedtests/<pkg>tests/`
3. All GoConvey/Should assertions (zero `t.Error` in integration tests)
4. All AAA pattern with explicit section comments
5. All maps formatted line-by-line
6. All packages at 100% coverage (excluding internal and OS-dependent paths)
7. No failing or flaky tests
8. Blocked packages resolved with documented root cause

---

## 6. Coverage Push Iteration Workflow

### Triggered by: "next"

Each time the user says **"next"**, process exactly **two packages**:

#### Step 1: Identify Targets

Parse the coverage data to find the two lowest-coverage non-internal packages:

```python
# Parse coverage.out to find uncovered packages
# Sort by coverage % ascending
# Pick the first two that aren't internal/*
```

#### Step 2: Find Uncovered Lines

```python
# Filter coverage.out for lines ending with " 0" (uncovered)
# Group by file within the package
# Identify which functions/branches need tests
```

#### Step 3: Read Source Files

**ALWAYS read the source file before writing tests.** Check:
- Function signatures (parameter names, types, order)
- Return types
- Nil guard patterns
- Branch conditions

#### Step 4: Write Tests

Follow all rules from Section 4. Cover:
- Each uncovered branch
- Nil receiver paths
- Error return paths
- Edge cases

#### Step 5: Report

```markdown
### Completed This Iteration

| Package | Previous Coverage | Target |
|---------|------------------|--------|
| `pkgA`  | 87.4%            | 100%   |
| `pkgB`  | 89.1%            | 100%   |

### Remaining Packages

| Package | Current Coverage |
|---------|-----------------|
| `pkgC`  | 95.1%           |
| `pkgD`  | 96.5%           |
| ...     | ...             |

### Notes
- Blockers encountered
- Fixes applied
- OS-dependent paths that can't be covered
```

---

## 7. Common Error Patterns & Fixes

### Pattern 1: API Changed But Tests Not Updated

**Symptom**: `undefined`, `too many arguments`, `has no field or method`

**Root cause**: Source code was refactored but test files still use the old API.

**Fix**: Read the current source file and update the test to match.

### Pattern 2: Argument Order Swap

**Symptom**: Test compiles but assertion fails with unexpected values.

**Root cause**: Two parameters of the same type were swapped.

**Fix**: Check the function signature and correct the order.

### Pattern 3: Value Type → Pointer Type

**Symptom**: `cannot use Type{} as *Type`

**Root cause**: Method receivers changed from value to pointer.

**Fix**: Use `&Type{}` or `new(Type)` instead of `Type{}`.

### Pattern 4: Constructor Pattern Changed

**Symptom**: `NewKey("x")` fails, should be `NewKey.Default("x")`

**Root cause**: Package adopted the `New.Method` creator pattern.

**Fix**: Check the `newCreator` type for available constructor methods.

### Pattern 5: Expected Value Drift

**Symptom**: Test passes locally but assertion shows wrong expected value.

**Root cause**: The actual behavior changed (e.g., collection size, format) but the test case's `ExpectedInput` was never updated.

**Fix**: Verify the actual deterministic output and update the expected value.

### Pattern 6: UTF-8/Encoding Display Issues

**Symptom**: Console shows `Γ£ô` instead of `✓`, `Γ£ù` instead of `✗`.

**Root cause**: PowerShell terminal not configured for UTF-8.

**Fix**: Add to the top of `run.ps1`:
```powershell
[Console]::OutputEncoding = [System.Text.Encoding]::UTF8
$OutputEncoding            = [System.Text.Encoding]::UTF8
```

### Pattern 7: `[setup failed]` From Heavy Test Framework Imports in In-Package Tests

**Symptom**: `go test -coverpkg=./...` reports `[setup failed]` for a package with no build errors visible. The blocked-packages report shows `FAIL` with an empty diagnostic log.

**Root cause**: An in-package test file (a `_test.go` file living inside the source package itself, e.g., `internal/mapdiffinternal/isStringType_test.go`) imports heavy test framework packages such as `coretests/args`, `goconvey`, or other transitive dependency chains. When `go test` runs with `-coverpkg=./...` instrumentation, the Go toolchain must resolve and instrument all transitive dependencies of every test binary. Heavy framework imports inside low-level packages create circular or excessively deep dependency graphs that cause the test binary loader to fail silently with `[setup failed]`.

**Why it's hard to diagnose**: The standard `go test` output contains only the terminal `FAIL pkg [setup failed]` line with zero preceding context. `go list -e -deps -test` may also return clean results because the issue only manifests under `-coverpkg` instrumentation.

**Fix**:
1. **Preferred**: Rewrite the in-package test to use only the standard `testing` package — no external test frameworks.
2. **Alternative**: Move the test to `tests/integratedtests/{pkg}tests/` where heavy framework imports are expected and the dependency graph is isolated.

**Prevention rule**: In-package test files (`_test.go` inside source packages) must **never** import `coretests/`, `goconvey`, or any package with a large transitive dependency tree. Keep in-package tests minimal — standard `testing` + `t.Errorf` only.

---

## 8. PowerShell Runner Internals

### Modular Architecture

`run.ps1` is a thin dispatcher (~167 lines) that imports `.psm1` modules from `scripts/`:

| Module | Key Functions | Responsibility |
|--------|--------------|----------------|
| `Utilities.psm1` | `Write-Header`, `Write-Success`, `ParseCompileErrors` | Common helpers |
| `TestLogWriter.psm1` | `Write-TestLogs` | Parse Go test output → log files |
| `TestRunner.psm1` | `Invoke-AllTests`, `Invoke-BuildCheck` | Test execution |
| `CoverageRunner.psm1` | `Invoke-TestCoverage`, `Invoke-PackageTestCoverage` | TC + TCP pipelines |
| `BuildTools.psm1` | `Invoke-Build`, `Invoke-Format`, `Invoke-Vet` | Build commands |
| `PreCommitCheck.psm1` | `Invoke-PreCommitCheck` | PC pipeline |
| `DashboardUI.psm1` | `Register-Phase`, `Write-PhaseSummaryBox` | ANSI dashboard (optional) |
| `Help.psm1` | `Show-Help` | Help display |

All DashboardUI calls are guarded with `Get-Command ... -ErrorAction SilentlyContinue` so the runner works without the UI module.

### Go Syntax Validation

Two Go tools validate syntax before compilation:

1. **bracecheck** (`scripts/bracecheck/main.go`): Scans all `.go` files for unbalanced braces, brackets, and parentheses. Run via `go run ./scripts/bracecheck/`.
2. **autofix** (`scripts/autofix/main.go`): Automatically fixes common syntax issues (trailing commas, missing imports). Run via `go run ./scripts/autofix/`. Supports `--dry-run`.

Both are executed as phases in TC and PC pipelines (skippable via `--skip-bracecheck`).

### Profile Merging (MAX Count)

When merging partial coverage profiles, use **MAX count** per line, not last-write-wins:

```powershell
# Coverage line format:
#   pkg/file.go:startLine.startCol,endLine.endCol numStatements count

foreach ($line in $allPartialLines) {
    if ($line -match '^(\S+\.go:\d+\.\d+,\d+\.\d+\s+\d+)\s+(\d+)$') {
        $key = $Matches[1]
        $count = [int]$Matches[2]
        if (-not $map.ContainsKey($key) -or $count -gt $map[$key]) {
            $map[$key] = $count
        }
    }
}
```

### Including In-Package Tests

Source packages with `_test.go` files must be included in the test run alongside integration tests:

```powershell
# Integration test packages
$integrationTestPkgs = go list ./tests/integratedtests/... 2>&1

# Source packages with in-package test files
$inPkgTestPkgs = @()
foreach ($srcPkg in $srcPkgs) {
    $relPath = $srcPkg -replace '^github\.com/org/project/', ''
    if (Test-Path $relPath) {
        if (Get-ChildItem -Path $relPath -Filter '*_test.go' -File -ErrorAction SilentlyContinue) {
            $inPkgTestPkgs += $srcPkg
        }
    }
}

# Combine both
$allTestPkgs = @($integrationTestPkgs) + @($inPkgTestPkgs) | Sort-Object -Unique
```

### Blocked Package Error Filtering

Raw `go test -c` output contains noise. Filter to actionable errors only:

| Pattern | Action |
|---------|--------|
| `warning: no packages being tested...` | **Remove** |
| `# github.com/...` without `.go:\d+` | **Remove** (bare header) |
| `*.go:123:45: error message` | **Keep** (actionable) |
| `FAIL\t...` | **Keep** (build-failed summary) |

### Console Display Colors

| Condition | Color |
|-----------|-------|
| Coverage ≥ 100% | Green |
| Coverage ≥ 80% | Yellow |
| Coverage < 80% | Red |
| Blocked packages | Red |
| Failing tests | Red |
| Written files | Gray |

---

## 9. Error Attribution System

Every error report includes **source attribution** — the `.psm1` module and function that triggered the failure. This is critical for diagnosing whether a failure originates from the build check, coverage compile check, test log writer, or coverage report generation.

### `Get-CallerSource` (Utilities.psm1)

Walks `Get-PSCallStack` to return the originating module and function:

```powershell
$source = Get-CallerSource
# → "CoverageRunner.psm1 → Invoke-TestCoverage"
```

- Skips internal frames (`<ScriptBlock>`, `Utilities.psm1`)
- Falls back to script name or function name alone
- Returns `"unknown"` if no meaningful frame found

### Integration Points

| Module | Function | Report File | Field |
|--------|----------|-------------|-------|
| `TestRunnerCore.psm1` | `Invoke-BuildCheck` | `failing-tests.txt` | `# Source:` header |
| `TestRunnerCore.psm1` | `Invoke-GitPull`, `Invoke-FetchLatest` | Console only | `Write-Fail` with source |
| `TestRunner.psm1` | `Invoke-AllTests`, `Invoke-PackageTests` | Console only | `Write-Fail` with source |
| `TestLogWriter.psm1` | `Write-TestLogs` | `passing-tests.txt`, `failing-tests.txt` | `# Source:` header |
| `CoverageReportJson.psm1` | JSON/text export | `build-errors.json/.txt`, `runtime-failures.json/.txt` | `"source"` / `# Source:` |
| `CoverageRunner.psm1` | Blocked packages + exit paths | `blocked-packages.txt` | `# Source:` header + `Write-Fail` |
| `CoverageCompileCheck.psm1` | Compile check | Console only | `Write-Fail` with source |
| `CoveragePreChecks.psm1` | Pre-check validation | Console only | `Write-Fail` with source |
| `CoverageSplitRecovery.psm1` | Subfolder recovery | Console only | `Write-Fail` with source |
| `CoverageReportHtml.psm1` | HTML report generation | Console only | `Write-Fail` with source |
| `BuildTools.psm1` | `Invoke-Build`, `Invoke-Vet` | Console only | `Write-Fail` with source |
| `PreCommitCheck.psm1` | Pre-commit validation | Console + JSON report | `Write-Fail` + `"source"` field |
| `GoConvey.psm1` | GoConvey install | Console only | `Write-Fail` with source |
| `Help.psm1` | Error paths | Console only | `Write-Fail` with source |
| `PackageCoverage.psm1` | Package coverage | Console only | `Write-Fail` with source |

### Parallel Mode Caveat

`Get-CallerSource` uses `Get-PSCallStack`, which does **not** cross `ForEach-Object -Parallel` thread boundaries. In parallel blocks, **hardcode** the source string:

```powershell
# Inside parallel block — cannot use Get-CallerSource
$callerSource = "CoverageCompileCheck.psm1 → Invoke-CoverageCompileCheck (parallel)"
```

### Report Examples

```
# Text: failing-tests.txt
# Source: TestRunnerCore.psm1 → Invoke-BuildCheck

# JSON: build-errors.json
{ "source": "CoverageRunner.psm1 → Invoke-TestCoverage", ... }

# Console
  ✗ Blocked: subpkg/foo (source: CoverageCompileCheck.psm1 → Invoke-CoverageCompileCheck)
```

### Error Extraction Pipeline (4-Tier Fallback)

When a blocked or failing package produces output, diagnostic lines are extracted via a 4-tier fallback chain (first non-empty result wins):

| Tier | Function | What it captures |
|------|----------|-----------------|
| 1 | `Extract-BuildErrorLines` | `.go:line:` errors, `[build failed]`, `[setup failed]`, `# pkg` headers |
| 2 | `Extract-ExecutionFailureLines` | Tier 1 + `panic:`, `fatal error:`, `--- FAIL:`, `FAIL pkg`, `exit status` |
| 3 | `Extract-SetupFailedContext` | Walks backward from `[setup failed]`/`[build failed]` FAIL lines, captures up to 10 preceding context lines (plain-text error messages) |
| 4 | `Get-RawFallbackLines` | All non-empty lines after noise removal (last resort — nothing is lost) |

All four functions are defined in `ErrorExtractor.psm1`. The chain is used by `ErrorParser.psm1` (accumulation), `CoverageReportJson.psm1` (reports), and `CoverageRunner.psm1` (blocked-packages files).

---

## 10. Related Spec Files

| Path | Purpose |
|------|---------|
| `spec/03-powershell-test-run/01-overview.md` | run.ps1 command reference |
| `spec/03-powershell-test-run/portable/01-generic-go-test-coverage-runner.md` | Generic runner architecture with sample script |
| `spec/04-tooling/03-powershell-implementation.md` | Full implementation spec (includes §8 Error Attribution) |
| `spec/06-testing-guidelines/README.md` | Testing guidelines index |
| `spec/06-testing-guidelines/08-good-vs-bad.md` | Good vs bad test examples |
| `spec/01-app/16-testing-guidelines.md` | Comprehensive testing reference (CaseV1, args.Map, assertions) |
| `spec/01-app/27-unit-coverage-fix.md` | Unit coverage fix workflow spec |
| `spec/05-failing-tests/01-blocked-packages-fixes.md` | Real-world blocked package fix examples |
| `.lovable/memory/workflow/06-unit-coverage-fix-protocol.md` | Protocol memory (for AI agents) |
| `scripts/README.md` | Module architecture, dependency graph, how to add commands |

---

## Version History

| Date | Change |
|------|--------|
| 2026-04-06 | S-010 COMPLETE — 38 benchmarks added across 6 packages. S-012 COMPLETE — 46 pointer→value receiver migrations (LeftRight, LeftMiddleRight, ExpectingRecord) |
| 2026-04-06 | Marked §5 coverage protocol COMPLETE — 21 packages at 100%. Added completed package list, unreachable gap registry, and split recovery documentation |
| 2026-04-03 | Added §9 Error Extraction Pipeline (4-tier fallback) with `Extract-SetupFailedContext` |
| 2026-03-31 | Updated directory layout, added §8 modular architecture, Go syntax validation docs |
| 2026-03-30 | Initial creation — consolidated from run.ps1 overview, generic runner spec, testing guidelines, and unit coverage fix protocol |
