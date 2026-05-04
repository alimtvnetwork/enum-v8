# bracecheck

Fast Go syntax pre-check invoked by `scripts/CoveragePreChecks.psm1`
and `run.sh` before the slower coverage runner.

## What it checks

For every `.go` file under the repo root (excluding `vendor/`,
`.git/`, `node_modules/`, `data/`, `cross-repo/`, `dist/`, `build/`,
and any `testdata/` subdirectory):

1. **Lexical brace/bracket/paren balance** — string, rune, and comment
   contents are skipped. Catches mismatches that the Go parser
   sometimes reports with confusing line numbers.
2. **Full parser pass** via `go/parser.ParseFile(..., parser.AllErrors)`
   — surfaces every syntax error with file:line:col precision.

## Output

- **Clean run:** `✓ bracecheck: N files clean` on stdout, exit 0.
- **Issues found:** one line per issue in `<relpath>:<line>:<col>:
  <message>` form, then a summary, then exit 1.

## Run manually

```bash
go run ./scripts/bracecheck/
```

## Skip from the runner

```bash
./run.sh tc --skip-bracecheck
./run.ps1 TC --skip-bracecheck
```

The PowerShell precheck also auto-skips this step (with a
`Register-Phase ... "skip"` dashboard entry) if the directory is
missing, so removing this tool won't hard-fail CI.
