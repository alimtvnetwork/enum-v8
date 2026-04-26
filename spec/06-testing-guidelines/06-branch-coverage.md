# 06 — Branch Coverage Strategy

## Goal

Every code branch must have at least one test case exercising it — **except for `internal/` packages** (see [Internal Package Exclusion](#internal-package-exclusion) below). Branches include:
- `if/else` and `if/else if/else`
- `switch/case`
- Nil guards (`if it == nil`)
- Early returns
- Error checks (`if err != nil`)
- Boolean short-circuits (`&&`, `||`)

---

## Branch Identification Process

### Step 1: Scan Source Code

For each function/method, identify all branches:

```go
func (it *MyStruct) Process(input string) (string, error) {
    if it == nil {                    // Branch 1: nil receiver
        return "", nil
    }

    if input == "" {                  // Branch 2: empty input
        return "", ErrEmptyInput
    }

    if it.cache.Has(input) {          // Branch 3: cache hit
        return it.cache.Get(input), nil
    }

    result := transform(input)        // Branch 4: normal path

    if result == "" {                 // Branch 5: empty result
        return "", ErrTransformFailed
    }

    return result, nil                // Branch 6: success
}
```

### Step 2: Build Coverage Matrix

| # | Branch | Path | Test Category |
|---|--------|------|---------------|
| 1 | `it == nil` | nil receiver | CaseNilSafe |
| 2 | `input == ""` | empty input | CaseV1 (negative) |
| 3 | `cache.Has(input)` | cache hit | CaseV1 (positive) |
| 4 | normal path | default | CaseV1 (positive) |
| 5 | `result == ""` | transform fails | CaseV1 (negative) |
| 6 | success | happy path | CaseV1 (positive) |

### Step 3: Write Test Cases

Each row in the matrix becomes one test case.

---

## Positive vs Negative Tests

### Positive Tests

Exercise the "happy path" — valid inputs producing expected outputs:

```go
{
    Title: "Process valid input returns transformed result",
    ArrangeInput: args.Map{
        "input": "hello",
    },
    ExpectedInput: args.Map{
        "result":   "HELLO",
        "hasError": false,
    },
}
```

### Negative Tests

Exercise error paths, boundary conditions, and invalid inputs:

```go
{
    Title: "Process empty input returns ErrEmptyInput",
    ArrangeInput: args.Map{
        "input": "",
    },
    ExpectedInput: args.Map{
        "result":   "",
        "hasError": true,
    },
}
```

---

## Nil Receiver Coverage

Every pointer receiver method with a nil guard gets a `CaseNilSafe` entry.

### Identification

Search for `if it == nil` or `if it.IsInvalid()` patterns:

```bash
grep -rn "if it == nil" mypkg/ --include="*.go"
```

### Categorization

| Method Signature | CaseNilSafe Pattern |
|-----------------|---------------------|
| `func (it *T) IsValid() bool` | Direct: `Func: (*T).IsValid` |
| `func (it *T) HasKey(key string) bool` | Wrapped: `Func: func(t *T) bool { return t.HasKey("x") }` |
| `func (it *T) Clear()` | Void: `Func: (*T).Clear`, `CompareFields: []string{"panicked", "returnCount"}` |
| `func (it *T) Clone() *T` | Nil return: `Func: func(t *T) bool { return t.Clone() == nil }` |
| `func (it *T) Process() (*R, error)` | Multi-return: `Func: func(t *T) bool { r, e := t.Process(); return r == nil && e != nil }` |

---

## Boundary Conditions

Test edge cases at the boundaries of valid input:

| Boundary | Example |
|----------|---------|
| Empty collection | `[]int{}`, `""`, `args.Map{}` |
| Single element | `[]int{42}`, `"x"` |
| Maximum values | `math.MaxInt64`, very long strings |
| Zero values | `0`, `""`, `false`, `nil` |
| Negative values | `-1`, `-math.MaxInt64` |
| Duplicates | `[]int{5, 5, 5}` |
| Already sorted | `[]int{1, 2, 3}` |
| Reverse sorted | `[]int{3, 2, 1}` |
| Special characters | `"hello\nworld"`, `"tab\there"` |

### Example

```go
var sortBoundaryTestCases = []coretestcases.CaseV1{
    {
        Title: "Quick sorts empty slice",
        ArrangeInput: args.Map{
            "input": []int{},
        },
        ExpectedInput: "[]",
    },
    {
        Title: "Quick handles single element",
        ArrangeInput: args.Map{
            "input": []int{42},
        },
        ExpectedInput: "[42]",
    },
    {
        Title: "Quick handles negative numbers",
        ArrangeInput: args.Map{
            "input": []int{-3, 0, 2, -1},
        },
        ExpectedInput: "[-3 -1 0 2]",
    },
    {
        Title: "Quick handles all duplicates",
        ArrangeInput: args.Map{
            "input": []int{5, 5, 5},
        },
        ExpectedInput: "[5 5 5]",
    },
}
```

---

## Coverage Audit Checklist

For each source package, verify:

- [ ] Every `if it == nil` has a corresponding `CaseNilSafe` entry
- [ ] Every `if err != nil` has both nil-error and non-nil-error test cases
- [ ] Every `switch/case` has a test for each case plus the default
- [ ] Every boolean method has both `true` and `false` test cases
- [ ] Every function has at least one positive and one negative test
- [ ] Boundary conditions are covered (empty, single, zero, max)
- [ ] Functions with multiple return values have tests verifying each combination

---

## Function Coverage Audit Template

```markdown
## Package: `mypkg`

### MyStruct

| Method | Nil Guard | Positive | Negative | Boundary | Status |
|--------|-----------|----------|----------|----------|--------|
| IsValid() | ✅ | ✅ | ✅ | n/a | Complete |
| Process(s) | ✅ | ✅ | ❌ empty | ❌ long string | Gap |
| HasKey(k) | ✅ | ✅ | ✅ | ❌ empty key | Gap |

### Package Functions

| Function | Positive | Negative | Boundary | Status |
|----------|----------|----------|----------|--------|
| NewMyStruct() | ✅ | n/a | n/a | Complete |
| Parse(input) | ✅ | ❌ nil | ❌ malformed | Gap |
```

---

## Internal Package Coverage Policy

**Packages under the `internal/` folder are excluded from coverage-motivated test work, but MAY have business/integration tests under `tests/integratedtests/<pkg>tests/`.**

This includes (but is not limited to):
`convertinternal`, `csvinternal`, `fsinternal`, `internalinterface`, `jsoninternal`, `mapdiffinternal`, `messages`, `msgcreator`, `msgformats`, `osconstsinternal`, `pathinternal`, `reflectinternal`, `strutilinternal`, `trydo`

### Rules (MUST)

- **MUST NOT** create `Coverage*_test.go` files for internal packages.
- **MUST NOT** include internal packages in coverage iteration plans or coverage audits.
- **MUST NOT** remove existing internal tests — they serve business/integration purposes (see existing `csvinternaltests/`, `fsinternaltests/`, `jsoninternaltests/`, etc. under `tests/integratedtests/`).
- **MAY** add business-critical or integration tests for internal packages under `tests/integratedtests/<pkg>tests/`, provided the motivation is correctness (not coverage %).
- Any future package placed inside `internal/` automatically inherits this policy.

### Decision tree

```
Adding a test for an internal/ package?
├─ Motivated by "increase coverage %"        → ❌ STOP. Do not write.
├─ Motivated by "verify business behavior"   → ✅ OK. Place under tests/integratedtests/<pkg>tests/.
└─ Motivated by "regression for past bug"    → ✅ OK. Place under tests/integratedtests/<pkg>tests/.
```

> The existence of `csvinternaltests/`, `fsinternaltests/`, etc. is **not** a violation — those are business/integration tests, not coverage tests. See [02-app-issues/02-internal-package-coverage-policy.md](../02-app-issues/02-internal-package-coverage-policy.md) for history.

---

## In-Package Test Import Restrictions

**In-package test files** (`_test.go` files living inside a source package) must use only the standard `testing` package. They must **never** import heavy test frameworks such as `coretests/args`, `goconvey`, or any package with a large transitive dependency tree.

### Why

When `go test -coverpkg=./...` instruments all packages, heavy transitive imports in low-level source packages can cause the test binary loader to fail silently with `[setup failed]` — producing zero diagnostic output and blocking the entire package from coverage.

### Rule

| Location | Allowed Imports | Heavy Frameworks |
|----------|----------------|-----------------|
| `mypkg/foo_test.go` (in-package) | `testing`, `strings`, stdlib only | ❌ Forbidden |
| `tests/integratedtests/mypkgtests/` | Anything | ✅ OK |

### If You Hit `[setup failed]` With No Logs

1. Check if the failing package has in-package `_test.go` files with framework imports.
2. Rewrite them using only `testing` + `t.Errorf`, or move them to `tests/integratedtests/`.
