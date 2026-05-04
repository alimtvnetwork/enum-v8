#!/usr/bin/env python3
"""
Unit tests for `scripts/ci/lint-baseline-diff.py`.

Run:
    python3 -m unittest scripts/ci/test_lint_baseline_diff.py -v

Covers:
  - `load()`: missing path, empty file, malformed JSON, well-formed report,
    issues with missing `Pos` fields.
  - Finding identity is the 4-tuple (file, line, linter, message); column
    and severity differences do NOT register as new findings.
  - `main()` exit codes & behavior:
      * Seeding mode (no baseline / empty baseline) → exit 0, ::warning per
        finding, ::notice that we're seeding.
      * Gate mode, no NEW findings → exit 0.
      * Gate mode, NEW findings → exit 1, ::error per NEW finding.
      * Usage error (no args) → exit 2.
  - Summary block always reports current / baseline / NEW / FIXED / UNCHANGED.
"""
from __future__ import annotations

import importlib.util
import io
import json
import tempfile
import unittest
from contextlib import redirect_stdout, redirect_stderr
from pathlib import Path

_HERE = Path(__file__).resolve().parent
_SPEC = importlib.util.spec_from_file_location(
    "lint_baseline_diff", _HERE / "lint-baseline-diff.py"
)
lbd = importlib.util.module_from_spec(_SPEC)
assert _SPEC.loader is not None
_SPEC.loader.exec_module(lbd)


def write_report(path: Path, issues: list[dict]) -> Path:
    path.write_text(json.dumps({"Issues": issues}), encoding="utf-8")
    return path


def issue(file: str, line: int, linter: str, text: str,
          col: int = 1, severity: str = "warning") -> dict:
    return {
        "FromLinter": linter,
        "Text": text,
        "Severity": severity,
        "Pos": {"Filename": file, "Line": line, "Column": col},
    }


class TestLoad(unittest.TestCase):
    def test_missing_path_returns_empty(self):
        self.assertEqual(lbd.load(Path("/nonexistent/nope.json")), set())

    def test_none_path_returns_empty(self):
        self.assertEqual(lbd.load(None), set())  # type: ignore[arg-type]

    def test_empty_file_returns_empty(self):
        with tempfile.TemporaryDirectory() as td:
            p = Path(td) / "empty.json"
            p.write_text("", encoding="utf-8")
            self.assertEqual(lbd.load(p), set())

    def test_malformed_json_returns_empty(self):
        with tempfile.TemporaryDirectory() as td:
            p = Path(td) / "bad.json"
            p.write_text("{not json", encoding="utf-8")
            self.assertEqual(lbd.load(p), set())

    def test_wellformed_report(self):
        with tempfile.TemporaryDirectory() as td:
            p = write_report(Path(td) / "r.json", [
                issue("a.go", 12, "govet", "shadow declaration"),
                issue("b.go", 7, "errcheck", "unchecked error"),
            ])
            got = lbd.load(p)
            self.assertEqual(got, {
                ("a.go", 12, "govet", "shadow declaration"),
                ("b.go", 7, "errcheck", "unchecked error"),
            })

    def test_issue_missing_pos_uses_defaults(self):
        with tempfile.TemporaryDirectory() as td:
            p = Path(td) / "r.json"
            p.write_text(json.dumps({"Issues": [
                {"FromLinter": "x", "Text": "y"},
            ]}), encoding="utf-8")
            self.assertEqual(lbd.load(p), {("", 0, "x", "y")})

    def test_column_and_severity_not_in_identity(self):
        # Same (file, line, linter, text) but different column/severity →
        # one finding, not two.
        with tempfile.TemporaryDirectory() as td:
            p = write_report(Path(td) / "r.json", [
                issue("a.go", 5, "govet", "msg", col=1, severity="warning"),
                issue("a.go", 5, "govet", "msg", col=99, severity="error"),
            ])
            self.assertEqual(len(lbd.load(p)), 1)


class TestMain(unittest.TestCase):
    def _run(self, *args: str) -> tuple[int, str, str]:
        out, err = io.StringIO(), io.StringIO()
        with redirect_stdout(out), redirect_stderr(err):
            rc = lbd.main(["lint-baseline-diff.py", *args])
        return rc, out.getvalue(), err.getvalue()

    def test_no_args_returns_2(self):
        rc, _, err = self._run()
        self.assertEqual(rc, 2)
        self.assertIn("usage", err)

    def test_seeding_mode_no_baseline_arg(self):
        with tempfile.TemporaryDirectory() as td:
            cur = write_report(Path(td) / "cur.json", [
                issue("a.go", 1, "govet", "x"),
            ])
            rc, out, _ = self._run(str(cur))
            self.assertEqual(rc, 0, msg=out)
            self.assertIn("seeding", out.lower())
            self.assertIn("::warning", out)
            self.assertIn("::notice", out)

    def test_seeding_mode_baseline_missing_file(self):
        with tempfile.TemporaryDirectory() as td:
            cur = write_report(Path(td) / "cur.json", [
                issue("a.go", 1, "govet", "x"),
            ])
            baseline = Path(td) / "missing.json"  # does not exist
            rc, out, _ = self._run(str(cur), str(baseline))
            self.assertEqual(rc, 0, msg=out)
            self.assertIn("MISSING", out)

    def test_seeding_mode_baseline_empty_file(self):
        with tempfile.TemporaryDirectory() as td:
            cur = write_report(Path(td) / "cur.json", [
                issue("a.go", 1, "govet", "x"),
            ])
            baseline = Path(td) / "empty.json"
            baseline.write_text("", encoding="utf-8")
            rc, out, _ = self._run(str(cur), str(baseline))
            self.assertEqual(rc, 0, msg=out)
            self.assertIn("MISSING", out)

    def test_gate_no_new_findings_returns_0(self):
        with tempfile.TemporaryDirectory() as td:
            findings = [issue("a.go", 1, "govet", "x"),
                        issue("b.go", 2, "errcheck", "y")]
            cur = write_report(Path(td) / "cur.json", findings)
            base = write_report(Path(td) / "base.json", findings)
            rc, out, _ = self._run(str(cur), str(base))
            self.assertEqual(rc, 0, msg=out)
            self.assertIn("No new lint findings", out)

    def test_gate_new_findings_returns_1(self):
        with tempfile.TemporaryDirectory() as td:
            base = write_report(Path(td) / "base.json", [
                issue("a.go", 1, "govet", "x"),
            ])
            cur = write_report(Path(td) / "cur.json", [
                issue("a.go", 1, "govet", "x"),
                issue("c.go", 9, "staticcheck", "new bad"),
            ])
            rc, out, _ = self._run(str(cur), str(base))
            self.assertEqual(rc, 1)
            self.assertIn("::error", out)
            self.assertIn("new bad", out)

    def test_gate_reports_fixed_findings(self):
        with tempfile.TemporaryDirectory() as td:
            base = write_report(Path(td) / "base.json", [
                issue("a.go", 1, "govet", "old"),
                issue("b.go", 2, "errcheck", "kept"),
            ])
            cur = write_report(Path(td) / "cur.json", [
                issue("b.go", 2, "errcheck", "kept"),
            ])
            rc, out, _ = self._run(str(cur), str(base))
            self.assertEqual(rc, 0, msg=out)
            self.assertIn("Resolved findings", out)
            self.assertIn("old", out)

    def test_summary_counts_present(self):
        with tempfile.TemporaryDirectory() as td:
            base = write_report(Path(td) / "base.json", [
                issue("a.go", 1, "govet", "kept"),
                issue("b.go", 2, "errcheck", "fixed"),
            ])
            cur = write_report(Path(td) / "cur.json", [
                issue("a.go", 1, "govet", "kept"),
                issue("c.go", 3, "staticcheck", "added"),
            ])
            rc, out, _ = self._run(str(cur), str(base))
            self.assertEqual(rc, 1)
            self.assertIn("current   : 2", out)
            self.assertIn("baseline  : 2", out)
            self.assertIn("NEW       : 1", out)
            self.assertIn("FIXED     : 1", out)
            self.assertIn("UNCHANGED : 1", out)


if __name__ == "__main__":
    unittest.main()
