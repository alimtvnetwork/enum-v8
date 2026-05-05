# 08 — Validators

> ✅ **Status**: filled in audit Step 5b (2026-04-23, Asia/Kuala_Lumpur).
> **Audience**: anyone validating user input, configuration, or downstream API responses.
> **Source**: extracted from [`/spec/00-llm-integration-guide.md`](../00-llm-integration-guide.md) §Utility Packages → corevalidator, plus framework conventions.

---

## Table of Contents

1. [`corevalidator` Overview](#1-corevalidator-overview)
2. [Built-in Validators](#2-built-in-validators)
3. [Composing Validators](#3-composing-validators)
4. [Authoring a Custom Validator](#4-authoring-a-custom-validator)
5. [Diagnostic Output Contract](#5-diagnostic-output-contract)
6. [Testing Validators](#6-testing-validators)
7. [Common Mistakes](#7-common-mistakes)

---

## 1. `corevalidator` Overview

`corevalidator` provides line, slice, text, and range validators with assertion capabilities. Each validator implements a small interface from `coreinterface/` so they can be composed and substituted.

### Architecture

```
┌───────────────────────────────────────────────┐
│  coreinterface/* (validator contracts)        │  L1
│  - IsValidChecker, IsSuccessChecker           │
│  - MessageGetter, ErrorGetter                 │
└───────────────────────────────────────────────┘
                    │ implemented by
                    ▼
┌───────────────────────────────────────────────┐
│  corevalidator/* (concrete validators)        │  L5
│  - LineValidator, SliceValidator              │
│  - TextValidator, RangeValidator              │
│  - StringCompareAs                            │
└───────────────────────────────────────────────┘
                    │ embedded by
                    ▼
┌───────────────────────────────────────────────┐
│  Your validation pipeline                     │
│  - Compose, run, collect errors               │
└───────────────────────────────────────────────┘
```

### The validator contract

Every validator exposes (at minimum):

| Method | Returns | Purpose |
|---|---|---|
| `IsValid()` | `bool` | Did the value pass? |
| `IsSuccess()` | `bool` | Synonym for `IsValid()` (some validators alias both) |
| `IsFailed()` | `bool` | Inverse of `IsValid` |
| `Message()` | `string` | Human-readable failure reason (empty if valid) |
| `Error()` | `error` | `nil` if valid, otherwise an `errcore`-categorized error |

This is the **`IsSuccessValidator`** composition shown in [`/spec/00-llm-integration-guide.md` §coreinterface](../00-llm-integration-guide.md#coreinterface--interface-contracts).

---

## 2. Built-in Validators

### 2.1 `LineValidator`

Validates a single line of text against rules (length, character class, regex).

```go
import "github.com/alimtvnetwork/core-v9/corevalidator"

v := corevalidator.New.Line.
    NotEmpty().
    MaxLength(80).
    Matches(`^[a-zA-Z0-9_-]+$`).
    Build()
// Build() returns *LineValidator (pointer; satisfies IsSuccessValidator,
// MessageGetter, and ErrorGetter from coreinterface/). Always store in a
// pointer-typed variable — value-type storage will fail interface assertion
// because the receiver methods are pointer-bound. (F-V14-01 fix.)

result := v.Validate("hello-world")
result.IsValid()   // true
result.Message()   // ""

result = v.Validate("")
result.IsValid()   // false
result.Message()   // "value cannot be empty"
```

### 2.2 `SliceValidator`

Validates a slice — element count, optional per-element delegation to another validator.

```go
v := corevalidator.New.Slice.
    MinLength(1).
    MaxLength(10).
    EachItem(lineValidator).
    Build()

result := v.Validate([]string{"a", "b", "c"})
```

### 2.3 `TextValidator`

Multi-line text — whole-document rules (line count, total length, encoding).

### 2.4 `RangeValidator`

Numeric range with inclusive/exclusive bounds.

```go
v := corevalidator.New.Range.Int.
    Min(0).
    MaxExclusive(100).
    Build()

result := v.Validate(42)   // valid
result := v.Validate(100)  // invalid (exclusive)
```

### 2.5 `StringCompareAs`

A specialty validator for "string must compare-as X" semantics. The diagnostic output here is regex-checked in the framework's own tests — see [`/spec/05-failing-tests/11-stringcompareas-onlysupportederr-names-format.md`](../05-failing-tests/11-stringcompareas-onlysupportederr-names-format.md) for the format contract.

---

## 3. Composing Validators

### 3.1 Combining results

Run multiple validators, collect all failures:

```go
results := []corevalidator.Result{
    nameValidator.Validate(input.Name),
    emailValidator.Validate(input.Email),
    ageValidator.Validate(input.Age),
}

errs := []error{}
for _, r := range results {
    if !r.IsValid() {
        errs = append(errs, r.Error())
    }
}
combined := errcore.MergeErrors(errs...) // see 04-error-system.md §3
```

### 3.2 Embedding validators in domain types

Pattern: a domain type embeds the validators it cares about and exposes a single `Validate()` that runs them all.

```go
type UserCreateRequest struct {
    Name  string
    Email string
}

func (r *UserCreateRequest) Validate() error {
    return errcore.MergeErrors(
        nameValidator.Validate(r.Name).Error(),
        emailValidator.Validate(r.Email).Error(),
    )
}
```

### 3.3 Conditional validation

Use [`conditional.IfFunc`](./07-conditional-and-utilities.md#1-conditional--ternary--nil-safe) to lazy-build validators based on input state:

```go
err := conditional.IfFunc[error](
    isAdmin,
    func() error { return adminValidator.Validate(req).Error() },
    func() error { return userValidator.Validate(req).Error() },
)
```

---

## 4. Authoring a Custom Validator

Follow this template when no built-in validator fits.

### Step 1 — Define the type

```go
package emailvalidator

import (
    "github.com/alimtvnetwork/core-v9/errcore"
    "github.com/alimtvnetwork/core-v9/regexnew"
)

type EmailValidator struct {
    pattern *regexnew.Lazy
    label   string
}
```

### Step 2 — Provide a `New` factory

Follows the **`newCreator` pattern** (see [`02-design-philosophy.md` §5](./02-design-philosophy.md)).

```go
var New = newEmailValidatorCreator{}

type newEmailValidatorCreator struct{}

func (newEmailValidatorCreator) Default() *EmailValidator {
    return &EmailValidator{
        pattern: regexnew.New.Lazy(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`),
        label:   "email",
    }
}

func (newEmailValidatorCreator) WithLabel(label string) *EmailValidator {
    v := New.Default()
    v.label = label
    return v
}
```

### Step 3 — Implement the validator contract

All five methods, with **nil-receiver safety** (see [`04-error-system.md` §4](./04-error-system.md#4-nil-receiver-behaviour)).

```go
func (v *EmailValidator) Validate(input string) Result {
    if v == nil {
        return Result{valid: false, message: "validator is nil"}
    }
    if !v.pattern.IsMatch(input) {
        return Result{
            valid:   false,
            message: errcore.VarTwoNoType("field", v.label, "value", input),
        }
    }
    return Result{valid: true}
}

type Result struct {
    valid   bool
    message string
}

func (r Result) IsValid() bool   { return r.valid }
func (r Result) IsSuccess() bool { return r.valid }
func (r Result) IsFailed() bool  { return !r.valid }
func (r Result) Message() string { return r.message }
func (r Result) Error() error {
    if r.valid {
        return nil
    }
    return errcore.ValidationFailedType.Fmt("%s", r.message)
}
```

### Step 4 — Add tests (Style A)

See [`/spec/06-testing-guidelines/02-test-case-types.md` §CaseV1](../06-testing-guidelines/02-test-case-types.md#casev1):

```go
var emailValidatorCases = []coretestcases.CaseV1{
    {
        Title:         "Validate accepts a well-formed email",
        ArrangeInput:  args.Map{"email": "user@example.com"},
        ExpectedInput: args.Map{"isValid": true, "errorIsNil": true},
    },
    {
        Title:         "Validate rejects a missing @",
        ArrangeInput:  args.Map{"email": "userexample.com"},
        ExpectedInput: args.Map{"isValid": false, "errorIsNil": false},
    },
}
```

Plus a `CaseNilSafe` for the `*EmailValidator` nil-receiver path:

```go
var emailValidatorNilCases = []coretestcases.CaseNilSafe{
    {
        Title: "Validate on nil returns invalid result, no panic",
        Func: func(v *emailvalidator.EmailValidator) bool {
            return v.Validate("anything").IsValid()
        },
        Expected: results.ResultAny{
            Value:    "false",
            Panicked: false,
        },
    },
}
```

---

## 5. Diagnostic Output Contract

When a validator produces a failure message, follow these rules so log scanners and the AI agent can attribute failures.

### Rules

1. **Message starts with the field label**, not "Error:". Use `errcore.VarTwoNoType("field", label, "value", actual)`.
2. **No trailing punctuation** — log aggregators add their own.
3. **No interpolated newlines** — single-line messages only. For multi-line, return multiple errors via `errcore.MergeErrors`.
4. **Stable format** — once a validator's message format is committed, do not change it. Tests regex-check the output (see [`/spec/05-failing-tests/`](../05-failing-tests/)).

### Canonical error message format

Use this exact shape so log scanners + the AI agent can attribute failures consistently:

```
<ValidatorLabel>: field=<name> value=<actual> reason=<short-reason>
```

**Worked examples** (verbatim from existing validators):

```
LineValidator: field=username value="" reason=empty-not-allowed
SliceValidator: field=tags value=[] reason=min-length-1
RangeValidator: field=age value=200 reason=above-max-150
TextValidator: field=bio value="<...>" reason=exceeds-max-2000
```

Construct via `errcore.VarTwoNoType` (preferred) or `errcore.ValidationFailedType.Fmt`:

```go
return errcore.ValidationFailedType.Fmt(
    "%s: field=%s value=%q reason=%s",
    "NonEmptyLineValidator", fieldName, actual, "empty-not-allowed",
)
```

### Why this matters

The PowerShell test runner parses validator output to attribute failures to source files. See [`/spec/04-tooling/03-powershell-implementation.md` §8](../04-tooling/03-powershell-implementation.md#8-error-attribution-system). Format drift breaks attribution and surfaces as failing tests in `05-failing-tests/`.


---

## 6. Testing Validators

Each validator package needs three test files in `tests/creationtests/<pkg>tests/` (the repo's actual test root — see [`13-testing-patterns.md`](./13-testing-patterns.md) and the C-CVS-01 / D-CVS-17 resolutions in [`/spec/07-code-vs-spec-audits/01-scoreboard.md`](../07-code-vs-spec-audits/01-scoreboard.md)):

| File | Style | Verifies |
|---|---|---|
| `<Validator>_Verification_test.go` | A (`CaseV1`) | Positive + negative scenarios produce expected `IsValid` / `Message` |
| `<Validator>_NilReceiver_test.go` | `CaseNilSafe` | Nil receiver does not panic |
| `<Validator>_Format_test.go` | A with `ShouldBeRegex` | Failure message matches the documented regex |

> See [`13-testing-patterns.md`](./13-testing-patterns.md) for style decisions and [`/spec/06-testing-guidelines/05-assertion-patterns.md`](../06-testing-guidelines/05-assertion-patterns.md) for `ShouldBeRegex`.

---

## 7. Common Mistakes

| Mistake | Why bad | Fix |
|---|---|---|
| Returning `error` directly without going through `errcore` | Loses the categorization that the runner uses for attribution | Wrap with `errcore.ValidationFailedType.Fmt(...)` |
| Validator panics on nil receiver | Crashes the consumer in production | Always start `if v == nil { return invalid-result }` |
| Mixing `IsValid` and `IsSuccess` semantics | Some downstream consumers check one, some the other | Make both return identical values |
| Multi-line messages joined with `\n` | Breaks single-line log parsers | Return multiple results, combine with `errcore.MergeErrors` |
| Building a regex inside `Validate()` on every call | CPU cost grows with call volume | Compile once via `regexnew.New.Lazy` at struct construction |
| Skipping the `Format_test.go` regex check | Format drift goes unnoticed until production | Always include format tests |

---

## See Also

- [`02-design-philosophy.md`](./02-design-philosophy.md) — `newCreator` (§5), zero-nil safety (§4), interface-first (§7)
- [`04-error-system.md`](./04-error-system.md) — Wrapping validator results in `errcore` errors
- [`07-conditional-and-utilities.md`](./07-conditional-and-utilities.md) — `regexnew.Lazy` for the pattern field; `conditional.IfFunc` for branching
- [`13-testing-patterns.md`](./13-testing-patterns.md) — Test styles for verification + nil-receiver + format tests
- [`/spec/05-failing-tests/`](../05-failing-tests/) — Real-world format-drift incidents to learn from
- [`/spec/00-llm-integration-guide.md` §corevalidator](../00-llm-integration-guide.md#corevalidator--validators) — Quick reference
