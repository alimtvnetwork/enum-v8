# autofix

Conservative Go source auto-fixer invoked by
`scripts/CoveragePreChecks.psm1` and `run.sh` before bracecheck.

## What it fixes

For every `.go` file under the repo root (excluding `vendor/`,
`.git/`, `node_modules/`, `data/`, `cross-repo/`, `dist/`, `build/`,
and any `testdata/` subdirectory):

1. **Trim trailing whitespace** on every line.
2. **Collapse 3+ consecutive blank lines** down to 2.
3. **Ensure exactly one trailing newline** at end of file.
4. **gofmt-equivalent formatting** via `go/format.Source`.

Every transformation is idempotent and equivalent to the input
under Go's lexical rules — running autofix twice is a no-op.

If `go/format.Source` rejects a file (i.e. it doesn't parse), the
file is **left untouched** and the parse error is surfaced as a
warning so bracecheck can pinpoint the syntax issue.

## Flags

| Flag        | Behavior                                              |
|-------------|-------------------------------------------------------|
| `--dry-run` | Report what would change without writing files.       |

## Output

- **Clean run, no changes:** `✓ autofix: no fixable issues across N files`
- **Changes applied:** `✓ autofix: fixed X/N files (W warnings)`
- **Dry-run with pending changes:** `✓ autofix: would fix X/N files (W warnings)` + per-file `- <relpath> would be modified` lines.
- **Failures:** `✗ autofix: F failure(s); ...` and exit 1.

## Run manually

```bash
go run ./scripts/autofix/                 # apply
go run ./scripts/autofix/ --dry-run       # preview
```

## Skip from the runner

```bash
./run.sh tc --no-autofix
./run.ps1 TC --no-autofix
./run.sh tc --skip-bracecheck    # also skips autofix (pair flag)
```

The PowerShell precheck auto-skips this step (with a
`Register-Phase ... "skip"` dashboard entry) if the directory is
missing.
