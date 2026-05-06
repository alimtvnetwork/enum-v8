"""Unit tests for `scripts/ci/check-issues-index-drift.py` (S-105)."""
from __future__ import annotations

import importlib.util
import sys
import unittest
from pathlib import Path
from tempfile import TemporaryDirectory

_HERE = Path(__file__).resolve().parent
_SCRIPT = _HERE / 'check-issues-index-drift.py'
_spec = importlib.util.spec_from_file_location('check_issues_index_drift', _SCRIPT)
_mod = importlib.util.module_from_spec(_spec)
assert _spec.loader is not None
_spec.loader.exec_module(_mod)  # type: ignore[union-attr]

INDEX_HEADER = (
    '# Issues Index\n\n'
    '| ID | Title | Status |\n'
    '|---|---|---|\n'
)
README_HEADER = (
    '# README\n\n'
    '| # | Topic | Status |\n'
    '|---|---|---|\n'
)


def _scaffold(tmp: Path, idx_rows: list[str], rdm_rows: list[str]) -> None:
    folder = tmp / 'spec' / '02-app-issues'
    folder.mkdir(parents=True, exist_ok=True)
    (folder / '00-issues-index.md').write_text(
        INDEX_HEADER + '\n'.join(idx_rows) + '\n', encoding='utf-8'
    )
    (folder / 'README.md').write_text(
        README_HEADER + '\n'.join(rdm_rows) + '\n', encoding='utf-8'
    )


class IndexDriftTests(unittest.TestCase):
    def _run(self, root: Path) -> int:
        return _mod.main(['check-issues-index-drift.py', str(root)])

    def test_in_sync_passes(self):
        with TemporaryDirectory() as d:
            tmp = Path(d)
            rows = [f'| 0{n} | Issue {n} | resolved |' for n in range(1, 4)]
            _scaffold(tmp, rows, rows)
            self.assertEqual(self._run(tmp), 0)

    def test_row_count_mismatch_fails(self):
        with TemporaryDirectory() as d:
            tmp = Path(d)
            idx = [f'| 0{n} | I {n} | resolved |' for n in range(1, 5)]   # 4 rows
            rdm = [f'| 0{n} | I {n} | resolved |' for n in range(1, 4)]   # 3 rows — drift
            _scaffold(tmp, idx, rdm)
            self.assertEqual(self._run(tmp), 1)

    def test_id_set_mismatch_same_count_fails(self):
        with TemporaryDirectory() as d:
            tmp = Path(d)
            idx = ['| 01 | A | r |', '| 02 | B | r |', '| 03 | C | r |']
            rdm = ['| 01 | A | r |', '| 02 | B | r |', '| 99 | Z | r |']  # same count, diff id
            _scaffold(tmp, idx, rdm)
            self.assertEqual(self._run(tmp), 1)

    def test_missing_file_returns_two(self):
        with TemporaryDirectory() as d:
            tmp = Path(d)
            (tmp / 'spec' / '02-app-issues').mkdir(parents=True)
            self.assertEqual(self._run(tmp), 2)

    def test_repo_root_is_in_sync(self):
        # Smoke test against the live repo — must stay green so the guard
        # itself never starts failing CI for unrelated reasons.
        repo_root = Path(__file__).resolve().parents[2]
        idx = repo_root / 'spec' / '02-app-issues' / '00-issues-index.md'
        rdm = repo_root / 'spec' / '02-app-issues' / 'README.md'
        if not idx.is_file() or not rdm.is_file():
            self.skipTest('repo files not present in this checkout')
        self.assertEqual(self._run(repo_root), 0)


if __name__ == '__main__':
    unittest.main()