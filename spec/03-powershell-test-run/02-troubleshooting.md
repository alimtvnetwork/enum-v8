# Troubleshooting — `run.ps1`

Common errors encountered when using the PowerShell test runner and how to resolve them.

---

## `go: command not found` / `go not in PATH`

**Symptom:** The script fails immediately with `go: The term 'go' is not recognized` or similar.

**Cause:** The Go toolchain is not installed or its binary directory is not in your system `PATH`.

**Fix:**

1. Verify Go is installed: download from <https://go.dev/dl/>
2. Ensure `GOROOT/bin` is in your `PATH`:
   ```powershell
   # Check current Go location
   Get-Command go -ErrorAction SilentlyContinue

   # Typical Windows default
   $env:Path += ";C:\Program Files\Go\bin"
   ```
3. Restart your terminal after modifying `PATH`.
4. Confirm with `go version`.

---

## `Build failed — skipping tests`

**Symptom:** The script prints `✗ Build failed — skipping tests` and no tests run. A `failing-tests.txt` file is opened with compiler errors.

**Cause:** The `go build ./...` gate (step 3 in the pipeline) detected compilation errors.

**Fix:**

1. Read the compiler output in `data/test-logs/failing-tests.txt` — it contains the exact file, line, and error.
2. Common sub-causes:
   - **Unused import** — remove the import or use it. (`imported and not used`)
   - **Undefined reference** — a function/type was renamed or moved; update callers.
   - **Type mismatch** — signature changed but callers weren't updated.
3. After fixing, re-run:
   ```powershell
   ./run.ps1 -t
   ```
4. If errors persist after a fix, clear the Go build cache:
   ```powershell
   go clean -cache
   ./run.ps1 -t
   ```

### Trailing `FAIL` with all tests passing

If `go test ./...` prints `FAIL` on the final line but every individual test shows `--- PASS`, it means a **non-test package failed to compile**. The build error output may have scrolled past. Check `data/test-logs/failing-tests.txt` or re-run with:

```powershell
go build ./...
```

This isolates the compilation error without test noise.

---

## `GoConvey not installed`

**Symptom:** Running `./run.ps1 GC` fails with `goconvey: command not found` or `The term 'goconvey' is not recognized`.

**Cause:** The GoConvey binary is not installed or not in `GOPATH/bin`.

**Fix:**

1. Install GoConvey:
   ```powershell
   go install github.com/smartystreets/goconvey@latest
   ```
2. Ensure `GOPATH/bin` is in your `PATH`:
   ```powershell
   # Find your GOPATH
   go env GOPATH

   # Add its bin directory (typical Windows default)
   $env:Path += ";$(go env GOPATH)\bin"
   ```
3. Verify with:
   ```powershell
   goconvey --help
   ```
4. Re-run:
   ```powershell
   ./run.ps1 GC
   ```

---

## `go mod tidy` errors

**Symptom:** The script fails during the `go mod tidy` step with messages about missing or ambiguous modules.

**Cause:** Module dependencies are out of sync with `go.mod` / `go.sum`.

**Fix:**

1. Run manually to see full output:
   ```powershell
   go mod tidy
   ```
2. If a dependency was removed, delete its entry from `go.mod` and re-run.
3. If a new dependency was added but not committed, ensure `go.sum` is also committed.
4. For persistent issues:
   ```powershell
   # Clear module cache and retry
   go clean -modcache
   go mod tidy
   ```

---

## Permission errors on `data/test-logs/`

**Symptom:** `Access to the path '...\data\test-logs\...' is denied.`

**Cause:** A previous process (editor, GoConvey) still has the log file open.

**Fix:**

1. Close any editors or processes that have the log files open.
2. If the file is locked:
   ```powershell
   # Force-remove and retry
   Remove-Item data/test-logs -Recurse -Force
   ./run.ps1 -t
   ```

---

## Tests pass locally but `FAIL` in CI

**Cause:** Environment differences — different Go version, OS-specific path handling, or missing test fixtures.

**Fix:**

1. Check Go version parity: `go version` locally vs CI config.
2. Ensure `go mod tidy` has been run and `go.sum` is committed.
3. Look for OS-specific path separators (`\` vs `/`) in test assertions.
4. Verify any required test data files are committed and not in `.gitignore`.

---

## Related Docs

- [Test Runner Overview](./01-overview.md)
- [Repo Overview](/spec/01-app/00-repo-overview.md)
- [Testing Patterns](/spec/01-app/13-testing-patterns.md)
