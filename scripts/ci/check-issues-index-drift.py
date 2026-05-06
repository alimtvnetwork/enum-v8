#!/usr/bin/env python3
"""S-105 — `spec/02-app-issues/` index-drift detector.

Compares the row count of the canonical machine-readable index
(`spec/02-app-issues/00-issues-index.md`) against the human-readable
sibling (`spec/02-app-issues/README.md`) and fails CI if they diverge.

Detection rule
--------------
Both files contain a single Markdown table whose *body* rows start with
`| 01 |`, `| 02 |`, ..., `| NN |`. We count those numeric-id rows in
each file and compare. If the counts differ, exit 1 with a diff report.

This is intentionally narrow — it does NOT validate cell contents or
ordering. The Cycle 18 audit (PI-002 origin) found that README.md had
been stale by 4 rows for ~14 days; a row-count guard is the cheapest
thing that prevents that exact recurrence.

Usage:
    python3 scripts/ci/check-issues-index-drift.py [<repo-root>]

Exit codes:
    0 — counts match
    1 — counts diverge
    2 — required input files missing
"""
from __future__ import annotations

import re
import sys
from pathlib import Path

ROW_RE = re.compile(r'^\|\s*(\d{2})\s*\|')


def count_numeric_rows(path: Path) -> tuple[int, list[str]]:
    """Return (count, ids) of `| NN |`-prefixed rows in *path*."""
    ids: list[str] = []
    for raw in path.read_text(encoding='utf-8').splitlines():
        m = ROW_RE.match(raw)
        if m:
            ids.append(m.group(1))
    return len(ids), ids


def main(argv: list[str]) -> int:
    root = Path(argv[1]) if len(argv) > 1 else Path('.')
    index_path = root / 'spec' / '02-app-issues' / '00-issues-index.md'
    readme_path = root / 'spec' / '02-app-issues' / 'README.md'

    missing = [p for p in (index_path, readme_path) if not p.is_file()]
    if missing:
        for p in missing:
            print(f'ERROR: required file missing: {p}', file=sys.stderr)
        return 2

    idx_count, idx_ids = count_numeric_rows(index_path)
    rdm_count, rdm_ids = count_numeric_rows(readme_path)

    if idx_count == rdm_count and idx_ids == rdm_ids:
        print(f'OK: spec/02-app-issues index in sync ({idx_count} rows).')
        return 0

    print('FAIL: spec/02-app-issues index drift detected.', file=sys.stderr)
    print(f'  {index_path}: {idx_count} rows  ids={idx_ids}', file=sys.stderr)
    print(f'  {readme_path}: {rdm_count} rows  ids={rdm_ids}', file=sys.stderr)

    only_in_index = [i for i in idx_ids if i not in rdm_ids]
    only_in_readme = [i for i in rdm_ids if i not in idx_ids]
    if only_in_index:
        print(f'  Missing from README: {only_in_index}', file=sys.stderr)
    if only_in_readme:
        print(f'  Missing from index : {only_in_readme}', file=sys.stderr)

    print('  Fix: keep both files in lockstep '
          '(canonical = 00-issues-index.md, human = README.md).', file=sys.stderr)
    return 1


if __name__ == '__main__':
    sys.exit(main(sys.argv))