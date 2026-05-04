#!/usr/bin/env python3
"""
Baseline-Diff Lint Gate (spec/12 §03-reusable-ci-guards/04).

Reads a current `golangci-lint` JSON report and diffs it against a
cached baseline. Fails only on NEW findings introduced by the change.

Behavior:
  - Seeding mode (baseline missing/empty): emit ::warning per finding,
    exit 0. Never gate the very first run.
  - Gate mode (baseline present): emit ::error per NEW finding, exit 1
    if any NEW exist. Always print a summary of NEW/FIXED/UNCHANGED.

Finding identity: 4-tuple (file, line, linter, message). Severity,
column, and source snippets are excluded to avoid spurious diffs across
linter version bumps.

Usage:
  python3 scripts/ci/lint-baseline-diff.py CURRENT.json [BASELINE.json]
"""
from __future__ import annotations

import json
import sys
from pathlib import Path


def load(path: Path) -> set[tuple[str, int, str, str]]:
    if not path or not path.exists() or path.stat().st_size == 0:
        return set()
    try:
        data = json.loads(path.read_text(encoding="utf-8"))
    except json.JSONDecodeError:
        return set()
    out: set[tuple[str, int, str, str]] = set()
    for issue in data.get("Issues") or []:
        pos = issue.get("Pos") or {}
        out.add((
            pos.get("Filename", ""),
            int(pos.get("Line", 0) or 0),
            issue.get("FromLinter", ""),
            (issue.get("Text") or "").strip(),
        ))
    return out


def main(argv: list[str]) -> int:
    if len(argv) < 2:
        print("usage: lint-baseline-diff.py CURRENT.json [BASELINE.json]", file=sys.stderr)
        return 2

    current = load(Path(argv[1]))
    baseline_path = Path(argv[2]) if len(argv) > 2 else None
    baseline_present = baseline_path is not None and baseline_path.exists() and baseline_path.stat().st_size > 0
    baseline = load(baseline_path) if baseline_present else set()

    added = current - baseline
    fixed = baseline - current
    unchanged = current & baseline

    print("=" * 60)
    print(" LINT BASELINE DIFF")
    print("=" * 60)
    print(f"  current   : {len(current)}")
    print(f"  baseline  : {len(baseline)} ({'present' if baseline_present else 'MISSING — seeding mode'})")
    print(f"  NEW       : {len(added)}")
    print(f"  FIXED     : {len(fixed)}")
    print(f"  UNCHANGED : {len(unchanged)}")
    print()

    if fixed:
        print("✅ Resolved findings:")
        for f, ln, linter, msg in sorted(fixed):
            print(f"   - [{linter}] {f}:{ln}  {msg}")
        print()

    if added:
        print("❌ New findings:")
        for f, ln, linter, msg in sorted(added):
            print(f"   + [{linter}] {f}:{ln}  {msg}")
        print()

    if not baseline_present:
        for f, ln, linter, msg in sorted(current):
            print(f"::warning file={f},line={ln}::[{linter}] {msg}")
        print("::notice::Seeding baseline — not gating this run.")
        return 0

    if added:
        for f, ln, linter, msg in sorted(added):
            print(f"::error file={f},line={ln}::[{linter}] {msg}")
        return 1

    print("✅ No new lint findings vs baseline.")
    return 0


if __name__ == "__main__":
    sys.exit(main(sys.argv))
