# 07 — Good vs. Bad Tests

## ✅ Good Test #1: Separated data and logic with args.Map

**`_testcases.go`:**
```go
var validateEmailTestCases = []coretestcases.CaseV1{
    {
        Title: "Valid email passes validation",
        ArrangeInput: args.Map{
            "email": "user@example.com",
        },
        ExpectedInput: args.Map{
            "isValid":    true,
            "errorCount": 0,
        },
    },
    {
        Title: "Empty email fails validation",
        ArrangeInput: args.Map{
            "email": "",
        },
        ExpectedInput: args.Map{
            "isValid":    false,
            "errorCount": 1,
        },
    },
}
```

**`_test.go`:**
```go
func Test_ValidateEmail_Verification(t *testing.T) {
    for caseIndex, tc := range validateEmailTestCases {
        // Arrange
        input := tc.ArrangeInput.(args.Map)
        email, _ := input.GetAsString("email")

        // Act
        result := validator.ValidateEmail(email)
        actual := args.Map{
            "isValid":    result.IsValid,
            "errorCount": len(result.Errors),
        }

        // Assert
        tc.ShouldBeEqualMap(t, caseIndex, actual)
    }
}
```

**Why it's good:**
- Data and logic are separated
- AAA comments present
- Native types in both expected and actual
- Semantic keys make failures self-documenting
- No branching logic in test body

---

## ❌ Bad Test #1: Everything crammed together

```go
func Test_ValidateEmail(t *testing.T) {
    // No separation, no AAA, raw t.Error, hardcoded values
    result := validator.ValidateEmail("user@example.com")
    if !result.IsValid {
        t.Errorf("expected valid but got invalid")
    }
    
    result2 := validator.ValidateEmail("")
    if result2.IsValid {
        t.Errorf("expected invalid but got valid")
    }
    if len(result2.Errors) != 1 {
        t.Errorf("expected 1 error but got %d", len(result2.Errors))
    }
}
```

**Why it's bad:**
- No data separation — expected values hardcoded in test body
- No AAA comments
- Uses raw `t.Errorf` instead of framework assertions
- Multiple scenarios crammed into one function
- Failure output is opaque ("expected valid but got invalid" — what input?)

---

## ✅ Good Test #2: Nil-safety with CaseNilSafe

```go
// _NilReceiver_testcases.go
var myStructNilSafeTestCases = []coretestcases.CaseNilSafe{
    {
        Title: "IsValid on nil returns false",
        Func:  (*MyStruct).IsValid,
        Expected: results.ResultAny{
            Value:    "false",
            Panicked: false,
        },
    },
    {
        Title: "HasKey on nil returns false",
        Func: func(m *MyStruct) bool {
            return m.HasKey("test")
        },
        Expected: results.ResultAny{
            Value:    "false",
            Panicked: false,
        },
    },
}

// NilReceiver_test.go
func Test_MyStruct_NilReceiver(t *testing.T) {
    for caseIndex, tc := range myStructNilSafeTestCases {
        tc.ShouldBeSafe(t, caseIndex)
    }
}
```

**Why it's good:**
- Compile-time safety via method references
- Panic recovery built in
- Consistent pattern across all types
- One test runner per package

---

## ❌ Bad Test #2: Manual nil-safety testing

```go
func Test_MyStruct_NilSafety(t *testing.T) {
    var m *MyStruct = nil
    
    // No panic recovery — test crashes if nil guard is missing
    if m.IsValid() != false {
        t.Error("IsValid should return false on nil")
    }
    
    // This line might panic and kill the entire test suite:
    if m.HasKey("test") != false {
        t.Error("HasKey should return false on nil")
    }
}
```

**Why it's bad:**
- No panic recovery — one missing nil guard crashes all tests
- Uses raw `t.Error`
- No data separation
- Can't test void methods (no return value to check)

---

## ✅ Good Test #3: Package function with CaseV1

```go
// _testcases.go
var concatMessageNilTestCase = coretestcases.CaseV1{
    Title: "ConcatMessageWithErr nil error returns nil",
    ArrangeInput: args.Map{
        "message": "should not appear",
    },
    ExpectedInput: args.Map{
        "isNil": true,
    },
}

var concatMessageNonNilTestCase = coretestcases.CaseV1{
    Title: "ConcatMessageWithErr non-nil wraps message",
    ArrangeInput: args.Map{
        "message": "context:",
        "error":   "original error",
    },
    ExpectedInput: args.Map{
        "isNil":    false,
        "contains": true,
    },
}

// _test.go
func Test_ConcatMessageWithErr_NilPassthrough(t *testing.T) {
    tc := concatMessageNilTestCase

    // Arrange
    input := tc.ArrangeInput.(args.Map)
    msg, _ := input.GetAsString("message")

    // Act
    result := errcore.ConcatMessageWithErr(msg, nil)
    actual := args.Map{
        "isNil": result == nil,
    }

    // Assert
    tc.ShouldBeEqualMapFirst(t, actual)
}
```

**Why it's good:**
- Package function tested with CaseV1 (not CaseNilSafe)
- Positive and negative paths as separate test cases
- Named standalone variables (not in a slice) for single tests

---

## ❌ Bad Test #3: CaseNilSafe for package functions

```go
// Pattern abuse — ConcatMessageWithErr is NOT a method
var badTestCase = coretestcases.CaseNilSafe{
    Func: func(_ *struct{}) bool {
        return errcore.ConcatMessageWithErr("msg", nil) == nil
    },
    Expected: results.ResultAny{
        Value:    "true",
        Panicked: false,
    },
}
```

**Why it's bad:**
- `ConcatMessageWithErr` is a package function, not a method
- Creates a fake `*struct{}` receiver that's never used
- Obscures the actual test intent
- Misleading — suggests nil-receiver testing when there's no receiver

---

## ✅ Good Test #4: Multi-line format for args.Map

```go
ExpectedInput: args.Map{
    "name":     "Alice",
    "age":      30,
    "isActive": true,
    "role":     "admin",
}
```

## ❌ Bad Test #4: Inline format for 2+ entries

```go
ExpectedInput: args.Map{"name": "Alice", "age": 30, "isActive": true, "role": "admin"}
```

---

## ✅ Good Test #5: Void method nil-safety

```go
{
    Title: "Clear on nil does not panic",
    Func:  (*MyStruct).Clear,
    Expected: results.ResultAny{
        Panicked: false,
    },
    CompareFields: []string{"panicked", "returnCount"},
}
```

**Why it's good:**
- Explicitly uses `CompareFields` because void methods have no `Value`
- Tests that nil guard prevents panic, not the return value

## ❌ Bad Test #5: Void method without CompareFields

```go
{
    Title: "Clear on nil does not panic",
    Func:  (*MyStruct).Clear,
    Expected: results.ResultAny{
        Panicked: false,
    },
    // Missing CompareFields — auto-derivation only compares "panicked"
    // which works but doesn't verify ReturnCount=0
}
```

---

## Anti-Pattern Checklist

| Anti-Pattern | Why It's Wrong | Fix |
|-------------|---------------|-----|
| Raw `t.Error`/`t.Errorf` | Inconsistent output, no diff | Use `ShouldBeEqual`/`ShouldBeEqualMap` |
| Expected values in `_test.go` | Violates separation | Move to `_testcases.go` |
| Branching in test body | Hard to trace failures | One test function per scenario |
| Stringified booleans `"true"` | Loses type safety | Use native `true` |
| `CaseNilSafe` for functions | Pattern abuse | Use `CaseV1` |
| Indexed slice access `tc[0]` | Fragile | Use named variables or loop with index |
| Multiple scenarios per function | Unclear which failed | Split into separate functions |
| Missing AAA comments | Unclear structure | Add `// Arrange`, `// Act`, `// Assert` |
| Vague title without input context | Can't diagnose from title alone | Include key input in title |
| Raw string keys in `args.Map` | Typo-prone, no autocomplete | Use `params.keyName` constants |

---

## Test Case Title Guidelines

Titles MUST be self-documenting: when a test fails, the title alone should tell you
**what was provided** and **what went wrong**.

### ✅ Good Titles (include input context)

```
"IsExpectedVersion returns false -- equal versions v4 vs v4.0 with LeftGreater expectation"
"ValueBool returns true -- given bool true input"
"VerifySimple returns error -- mismatched segment range 0-2"
```

### ❌ Bad Titles (vague, no input context)

```
"IsExpectedVersion returns false for mismatched expectation"
"ValueBool returns true"
"Mismatch returns error"
"Matching returns nil error"
```

### Rules

1. Title must mention the **function/method** being tested.
2. Title must include a `--` separator followed by a brief **input description**.
3. Avoid generic words like "mismatched", "matching", "correct" without specifying WHAT.
4. When inputs are simple values (e.g., "v4", `true`), include them directly in the title.
