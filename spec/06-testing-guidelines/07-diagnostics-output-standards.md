# Diagnostics Output Standards

## Stack Trace Rules

1. **Single header only**: Stack trace output must show `Stack-Trace :` exactly once, never duplicated.
2. **No system library frames**: Frames from Go's standard library (`runtime/`, `testing/`, or anything under `GOROOT`) must be filtered out. Only project-level frames should appear.
3. **Maximum 4 frames**: Stack traces are capped at 4 relevant frames to avoid noise.

## Single-Value vs Multi-Line Assertions

1. **Single-value comparison**: When both actual and expected are single strings (not multi-line), use compact format:
   ```
   Expected: "some value"
   Actual:   "other value"
   ```
   Do NOT use the full line-by-line diff block for single values.

2. **Multi-line comparison**: Use the standard `LineDiff` format with aligned labels.
   Both `actual` and `expected` labels must be padded to 10 characters before the colon
   to ensure consistent visual alignment:
   ```
     Line   2 [MISMATCH]:
              actual : `isDefined : true`
            expected : `isDefined : false`
   ```

## Alignment

- `actual` and `expected` labels use different leading indentation so the `: ` colon
  aligns at the same column position:
  - `actual`: 14 leading spaces + 6 chars + 1 space = colon at column 21
  - `expected`: 12 leading spaces + 8 chars + 1 space = colon at column 21
- Use consistent indentation (spaces, not tabs) in diagnostic blocks.

## Test Title Encoding

- Test titles must use ASCII-safe characters only.
- Use `--` or `-` instead of em dashes (`—`) to avoid encoding issues on Windows terminals.

## Nil-Safety in Once Types

- All `*Once.Value()` methods must guard against nil `initializerFunc`.
- If `initializerFunc` is nil, return the zero value and mark as initialized.

## Map Expected Output

- When map comparison fails, a structured header MUST appear first with each field on its own line:
  ```
  Test Method : TestFuncName
  Case        : 1
  Title       : Case Title
  ```
- Each block (Actual Received, Expected Input) MUST be wrapped in separator headers
  (`============================>`).
- The case title MUST appear under the section label, indented with 4 spaces.
- Each entry must be on its own line, tab-indented, in Go literal format.
- Do NOT use indexed numbering (`0:`, `1:`, etc.) before entries.
- Format:
  ```
  Test Method : TestFuncName
  Case        : 1
  Title       : Case Title

  ============================>
  1) Actual Received (2 entries):
      Case Title
  ============================>
  	"containsName": false,
  	"hasError":      false,
  ============================>

  ============================>
  1) Expected Input (1 entries):
      Case Title
  ============================>
  	"hasError": false,
  ============================>
  ```
