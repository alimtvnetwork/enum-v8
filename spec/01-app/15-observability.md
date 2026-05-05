> ✅ **Status**: drafted at spec-v0.16.0, expanded at spec-v0.17.1 with trust-boundary worked example (2026-04-25, Asia/Kuala_Lumpur).
> **Audience**: anyone wiring `core-v9` into a production consumer that needs structured logging, diagnostic context, or test-failure inspection.
> **Source**: unifies guidance previously scattered across `04-error-system.md` §5, `13-testing-patterns.md` §8, `00-llm-integration-guide.md` §StackEnhance, and `06-testing-guidelines/07-diagnostics-output-standards.md`.
> **Closes**: F-V14-05 *(observability half)*.

---

## Table of Contents

1. [Scope — What This Library Does and Does Not Provide](#1-scope--what-this-library-does-and-does-not-provide)
2. [Diagnostic Context Primitives](#2-diagnostic-context-primitives)
3. [Stack-Aware Error Wrapping](#3-stack-aware-error-wrapping)
4. [Test-Failure Output Format](#4-test-failure-output-format)
5. [Logging Boundaries](#5-logging-boundaries)
6. [Tracing & Metrics](#6-tracing--metrics)
7. [Common Mistakes](#7-common-mistakes)

---

## 1. Scope — What This Library Does and Does Not Provide

`core-v9` is a **pure library**. It provides primitives for **shaping** diagnostic data so that the consuming application can route it to whatever logger, tracer, or sink it chooses.

| Concern | This library | Consumer's responsibility |
|---|---|---|
| Structured-error context (typed key/value pairs) | ✅ `errcore.VarTwo`, `VarTwoNoType`, `MessageVarMap` | — |
| Stack-aware error wrapping | ✅ `errcore.StackEnhance.{Error,Msg}` | — |
| Test-failure framing (machine-parseable) | ✅ `coretests/results/Result.go` | — |
| Pretty-printing for diagnostics | ✅ `corejson.NewPtr(x).PrettyJsonString()` | — |
| Choosing a log backend (`zap`, `slog`, etc.) | ❌ | ✅ |
| Log levels / sampling / rotation | ❌ | ✅ |
| OpenTelemetry tracing | ❌ | ✅ |
| Prometheus / metrics emission | ❌ | ✅ |
| Sink-side PII redaction | ❌ | ✅ — but see [`16-security.md` §2](./16-security.md#2-pii--secret-handling) for what this library guarantees on its side |

> **Rule (MUST)**: Do not import a logging framework into `core-v9`. The library must remain dependency-light so consumers can choose their own observability stack.

---

## 2. Diagnostic Context Primitives

`errcore` exposes three families of context-attachment helpers. Use the smallest one that fits.

### 2.1 `VarTwo` — labelled pair with type names

```go
import "github.com/alimtvnetwork/core-v9/errcore"

err := errcore.VarTwo("userID", uid, "tenantID", tid)
// → "(userID [t:int64], tenantID [t:string]) = (42, acme-co)"
```

**When to use**: You have exactly two related variables whose types help debugging.

### 2.2 `VarTwoNoType` — labelled pair without type names

```go
err := errcore.VarTwoNoType("userID", uid, "tenantID", tid)
// → "(userID, tenantID) = (42, acme-co)"
```

**When to use**: Types are obvious from the labels and would clutter output.

### 2.3 `MessageVarMap` — N-way labelled context

```go
err := errcore.MessageVarMap("checkout failed", map[string]any{
    "userID":   uid,
    "tenantID": tid,
    "cartID":   cart.ID,
    "amount":   cart.Total,
})
```

**When to use**: Three or more variables, or a dynamic set built up across several call frames.

### 2.4 Choosing between them

| Number of context vars | Helper |
|---|---|
| 0 | use `errcore.<Category>.Fmt(...)` directly |
| 1 | use `errcore.<Category>.MergeErrorWithMessage(err, "with X=" + repr)` |
| 2 + types matter | `VarTwo` |
| 2 + types obvious | `VarTwoNoType` |
| 3 + | `MessageVarMap` |

---

## 3. Stack-Aware Error Wrapping

`errcore.StackEnhance.{Error,Msg}` wraps an existing error or message with the calling-site file:line and a partial stack trace.

```go
err := errcore.StackEnhance.Error(originalErr)
// → multi-line block with file:line trace and original error preserved

msg := errcore.StackEnhance.Msg("validation pipeline aborted")
// → message with file:line trace
```

**Rules**:

1. **MUST** call `StackEnhance` exactly once per logical error boundary (entering the package's public API). Calling it inside every helper produces nested stack noise.
2. **MUST NOT** call `StackEnhance` inside `*Must` methods — those route through `errcore.HandleErr` which already attaches a stack-enhanced wrapping (see `04-error-system.md` §1).
3. The two-space indent and `\n` newline conventions in the output are part of the public contract — see `06-testing-guidelines/07-diagnostics-output-standards.md` for the canonical format.

---

## 4. Test-Failure Output Format

`coretests/results/Result.go` and `ResultAny.go` emit failures in a fixed shape so the PowerShell test runner can attribute and count them.

```
Test #N — {scenario}: should be equal
  expected: {repr}
  actual:   {repr}
```

**Authoritative format spec**: [`06-testing-guidelines/07-diagnostics-output-standards.md`](../06-testing-guidelines/07-diagnostics-output-standards.md). This file is a forwarding pointer; the rules live there.

---

## 5. Logging Boundaries

Because `core-v9` itself logs nothing, every observability decision happens at the **consumer boundary**:

```go
// In your application's adapter layer:
result := corevalidator.New.Line.NotEmpty().Build().Validate(input)
if result.IsFailed() {
    logger.Error("validation_failed",
        slog.String("input_id", inputID),
        slog.String("reason", result.Message()),
        slog.Any("err", result.Error()),
    )
}
```

**Rules**:

1. **MUST NOT** add `fmt.Print*`, `log.Print*`, or any side-effecting I/O inside `core-v9` packages. Functions return errors / messages; the consumer logs them.
2. **MUST** preserve the original `error` value when logging — do not stringify and discard. Use `slog.Any("err", err)` or equivalent.
3. **SHOULD** log at the *outermost* boundary (HTTP handler, queue consumer) rather than at every internal call. Internal calls use `errcore.<Category>.MergeError` to attach context that the outermost logger captures.

### 5.1 End-to-end example: trust-boundary handler

A worked example combining validation (§[`08-validators.md`](./08-validators.md)), PII redaction (§[`16-security.md`](./16-security.md#2-pii--secret-handling)), and structured logging at the outermost boundary. This is the canonical pattern for an HTTP handler that accepts untrusted input.

```go
// HTTP handler — the outermost trust boundary.
// Inbound: untrusted form data containing an email + password.
// Outbound: structured slog at this boundary; generic error to the caller.
func (h *SignupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    rawEmail := r.FormValue("email")
    rawPwd   := r.FormValue("password")

    // 1. Validate at the boundary (see 08-validators.md §2.1).
    emailV := corevalidator.New.Line.NotEmpty().MaxLength(254).Build()
    if vr := emailV.Validate(rawEmail); vr.IsFailed() {
        // 2. Build a PII-redacted view for logging (see 16-security.md §2).
        h.log.Warn("signup_validation_failed",
            slog.String("field", "email"),
            slog.String("email_redacted", security.RedactEmail(rawEmail)), // "j***@e***.com"
            slog.String("reason", vr.Message()),
            slog.Any("err", vr.Error()),                                   // preserve error value
        )
        // 3. Return a generic error to the caller — never leak the raw input.
        http.Error(w, "invalid signup payload", http.StatusBadRequest)
        return
    }

    // 4. Password is sensitive — never log it, not even redacted length.
    if vr := corevalidator.New.Line.MinLength(12).Build().Validate(rawPwd); vr.IsFailed() {
        h.log.Warn("signup_validation_failed",
            slog.String("field", "password"),
            slog.Any("err", vr.Error()),
        )
        http.Error(w, "invalid signup payload", http.StatusBadRequest)
        return
    }

    // 5. Hand off to the domain layer. Any error returned here is already
    //    `errcore.StackEnhance`-wrapped at its public boundary, so we just log it.
    if err := h.svc.CreateUser(r.Context(), rawEmail, rawPwd); err != nil {
        h.log.Error("signup_failed",
            slog.String("email_redacted", security.RedactEmail(rawEmail)),
            slog.Any("err", err),
        )
        http.Error(w, "signup failed", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}
```

**Why this pattern is correct**:

| Step | Source rule | File |
|---|---|---|
| Validator built once, reused | §2.1 "Build returns `*LineValidator`" | `08-validators.md` |
| Email redacted before logging | §2 "scrub PII before serialising" | `16-security.md` |
| Password never logged in any form | §2 "secrets are never serialised, even redacted" | `16-security.md` |
| `slog.Any("err", err)` preserves error value | §5 rule 2 | this file |
| Logged once at the outermost boundary | §5 rule 3 | this file |
| Generic error returned to caller | §3 "trust boundary" | `16-security.md` |

**Closes**: F-V16-01 *(trust-boundary worked example)*.

---

## 6. Tracing & Metrics

Out of scope for `core-v9`. The library provides **no** spans, attributes, or counters. If a consumer needs OpenTelemetry, the recommended pattern is:

```go
ctx, span := tracer.Start(ctx, "validate-payload")
defer span.End()

result := corevalidator.New.Line.NotEmpty().Build().Validate(payload)
if result.IsFailed() {
    span.RecordError(result.Error())
    span.SetStatus(codes.Error, result.Message())
}
```

The library's `result.Error()` and `result.Message()` are deliberately compatible with OTel's `RecordError` / `SetStatus`.

---

## 7. Common Mistakes

| Mistake | Fix |
|---|---|
| Importing `log` / `slog` / `zap` inside a `core-v9` package | Remove the import. Return errors instead; let the consumer log. |
| Calling `StackEnhance.Error` inside every helper | Call it exactly once at the public API boundary. |
| Stringifying an error before logging (`logger.Error(err.Error())`) | Pass the `error` value: `logger.Error("op_failed", slog.Any("err", err))`. |
| Building a context map manually when `VarTwo` would do | Prefer the helpers — they emit the canonical format the test runner expects. |
| Adding `fmt.Println` "for debugging" and forgetting it | Use a temporary `errcore.VarTwo` and let the surrounding test capture it. |

---

## See Also

- [`04-error-system.md`](./04-error-system.md) — `errcore` API surface, including `HandleErr` and the `*Category.MergeError` family.
- [`16-security.md`](./16-security.md) — Companion file: PII handling, panic policy, allocation safety.
- [`06-testing-guidelines/07-diagnostics-output-standards.md`](../06-testing-guidelines/07-diagnostics-output-standards.md) — Authoritative diagnostic output format.
- [`13-testing-patterns.md` §8](./13-testing-patterns.md) — How tests consume these primitives.
