# 08 — Creating Custom Case Types

When `CaseV1`, `CaseNilSafe`, and `GenericGherkins` don't fit your needs, you can create custom case types using `BaseTestCase` as the foundation.

---

## BaseTestCase Structure

```go
type BaseTestCase struct {
    Title           string
    ArrangeInput    any
    ActualInput     any
    ExpectedInput   any
    Additional      any
    CustomFormat    string
    VerifyTypeOf    *VerifyTypeOf
    Parameters      *args.HolderAny
    IsEnable        issetter.Value
    HasError        bool
    HasPanic        bool
    IsValidateError bool
}
```

`CaseV1` is literally `type CaseV1 BaseTestCase` — a type alias with assertion methods.

---

## Pattern: Custom Test Wrapper

When your test requires domain-specific data extraction or specialized assertion logic:

### Step 1: Define the wrapper struct

```go
// In your package's _testcases.go or helpers.go:
type MyCustomTestCase struct {
    coretestcases.CaseV1  // Embed CaseV1 for all standard methods
    
    // Add domain-specific fields:
    InputFormat  string
    OutputFormat string
    IsStrict     bool
}
```

### Step 2: Add domain-specific methods

```go
func (it MyCustomTestCase) ArrangeAsConfig() *Config {
    input := it.ArrangeInput.(args.Map)
    
    return &Config{
        Format:   it.InputFormat,
        IsStrict: it.IsStrict,
        Value:    input.GetAsStringDefault("value"),
    }
}

func (it MyCustomTestCase) ShouldMatchOutput(
    t *testing.T,
    caseIndex int,
    actual string,
) {
    t.Helper()
    
    // Delegate to framework assertion:
    it.ShouldBeEqual(t, caseIndex, actual)
}
```

### Step 3: Use in tests

```go
// _testcases.go
var configTestCases = []MyCustomTestCase{
    {
        CaseV1: coretestcases.CaseV1{
            Title: "JSON format produces valid output",
            ArrangeInput: args.Map{
                "value": `{"key": "value"}`,
            },
            ExpectedInput: `{"key":"value"}`,
        },
        InputFormat:  "json",
        OutputFormat: "json",
        IsStrict:     true,
    },
}

// _test.go
func Test_Config_Verification(t *testing.T) {
    for caseIndex, tc := range configTestCases {
        // Arrange
        config := tc.ArrangeAsConfig()

        // Act
        result := processor.Process(config)

        // Assert
        tc.ShouldMatchOutput(t, caseIndex, result)
    }
}
```

---

## Pattern: Custom ShouldBeEqual with errcore.AssertDiffOnMismatch

For custom wrappers that need diff-based assertions:

```go
func (it MyTestCase) ShouldBeEqual(
    t *testing.T,
    caseIndex int,
    actual args.Map,
) {
    t.Helper()
    
    expected := it.ExpectedAsMap()
    
    actLines := actual.CompileToStrings()
    expLines := expected.CompileToStrings()
    
    errcore.AssertDiffOnMismatch(
        t,
        caseIndex,
        it.Title,
        actLines,
        expLines,
    )
}
```

---

## Pattern: Extending CaseNilSafe for Multi-Return Methods

When a nil-receiver method returns `(T, error)` and you need structured assertion:

```go
// Custom wrapper that extracts both return values:
type NilSafeMultiReturn struct {
    coretestcases.CaseNilSafe
    
    ExpectedError bool
}

func (it NilSafeMultiReturn) ShouldBeSafe(
    t *testing.T,
    caseIndex int,
) {
    t.Helper()
    
    // Use the built-in invocation:
    result := it.InvokeNil()
    
    // Custom assertion with both value and error:
    actual := args.Map{
        "panicked": result.Panicked,
        "value":    fmt.Sprintf("%v", result.Value),
        "hasError": result.HasError(),
    }
    
    expected := args.Map{
        "panicked": it.Expected.Panicked,
        "value":    fmt.Sprintf("%v", it.Expected.Value),
        "hasError": it.ExpectedError,
    }
    
    actLines := actual.CompileToStrings()
    expLines := expected.CompileToStrings()
    
    errcore.AssertDiffOnMismatch(
        t,
        caseIndex,
        it.CaseTitle(),
        actLines,
        expLines,
    )
}
```

---

## Guidelines for Custom Cases

### Do:
- Embed `CaseV1` or `BaseTestCase` to inherit standard methods
- Delegate assertions to `errcore.AssertDiffOnMismatch` for consistent output
- Use `t.Helper()` in all custom assertion methods
- Keep custom fields domain-specific (not testing infrastructure)

### Don't:
- Create custom cases when `CaseV1` + `args.Map` suffices
- Write raw `t.Error`/`t.Errorf` in custom assertion methods
- Mix test data and assertion logic in the same struct
- Create custom cases just to avoid learning the standard patterns

### Decision Criteria

| Condition | Use |
|-----------|-----|
| Standard input → string output | `CaseV1` |
| Standard input → multi-field output | `CaseV1` + `args.Map` |
| Nil-receiver safety | `CaseNilSafe` |
| BDD with typed input/output | `GenericGherkins[T1, T2]` |
| Domain-specific data extraction | Custom wrapper embedding `CaseV1` |
| Specialized assertion logic | Custom wrapper with `errcore.AssertDiffOnMismatch` |

---

## Real-World Custom Wrapper Examples

### StringsTestWrapper (shared wrapper)

A shared wrapper used across multiple test packages for string processing:

```go
// tests/testwrappers/stringstestwrapper/
type StringsTestWrapper struct {
    coretestcases.CaseV1
}

func (it StringsTestWrapper) Arrange() string {
    return it.ArrangeInput.(string)
}

func (it StringsTestWrapper) ExpectedLines() []string {
    return it.CaseV1.ExpectedLines()
}
```

### ReflectSetFromToTestWrapper (complex wrapper)

For tests requiring reflection-based setup and comparison:

```go
type ReflectSetFromToTestCase struct {
    Title    string
    From     any
    To       any
    Expected args.Map
}

func (it ReflectSetFromToTestCase) ShouldBeEqual(
    t *testing.T,
    caseIndex int,
    actual args.Map,
) {
    t.Helper()
    
    actLines := actual.CompileToStrings()
    expLines := it.Expected.CompileToStrings()
    
    errcore.AssertDiffOnMismatch(
        t,
        caseIndex,
        it.Title,
        actLines,
        expLines,
    )
}
```
