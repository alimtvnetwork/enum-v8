#!/usr/bin/env python3
"""
Cross-File Collision Audit (spec/12 §03-reusable-ci-guards/03).

Scans Go source files for three categories of identifier collisions:

  1. cross_file_exact   — same exported name declared in >1 file
  2. case_insensitive   — distinct names that differ only in case across files
  3. intra_file_dupes   — same name declared >1 time in a single file

String-literal aware: skips identifiers inside `"..."`, `` `...` ``,
and within `// ...` / `/* ... */` comments. Tracks `const ( ... )` /
`var ( ... )` blocks so grouped declarations are picked up.

Exit codes:
  0  no collisions
  1  one or more collisions found (full report on stdout)
  2  no source files matched the glob

Usage:
  python3 scripts/ci/check-collisions.py [PATH ...]

Defaults to scanning the entire module (excluding vendor/, .git/, data/,
node_modules/, scripts/, src/, public/, spec/).
"""
from __future__ import annotations

import os
import re
import sys
from collections import defaultdict
from pathlib import Path

EXCLUDE_DIRS = {".git", "node_modules", "data", "scripts", "src",
                "public", "spec", "tmp", "vendor", "dist"}

# Build-constraint file suffixes (GOOS / GOARCH). Files differing only by
# one of these suffixes are mutually exclusive at compile time, so the
# "same name in N files" pattern is legal and expected.
BUILD_SUFFIXES = {
    "linux", "darwin", "windows", "freebsd", "netbsd", "openbsd",
    "dragonfly", "solaris", "plan9", "android", "ios", "js", "wasip1",
    "amd64", "arm64", "arm", "386", "ppc64", "ppc64le", "mips", "mipsle",
    "mips64", "mips64le", "riscv64", "s390x", "wasm", "loong64", "unix",
}


def build_tag_key(path: Path) -> str:
    """Collapse GOOS/GOARCH-suffixed filenames to a single logical key.

    `Foo_linux.go`, `Foo_darwin.go`, `Foo_windows.go` -> `Foo.go`
    """
    stem = path.stem
    parts = stem.split("_")
    while len(parts) > 1 and parts[-1] in BUILD_SUFFIXES:
        parts.pop()
    return str(path.parent / ("_".join(parts) + ".go"))


def is_exported_unexported_pair(variants: list[str]) -> bool:
    """True when variants are an Exported/unexported pair differing only in
    the case of the first letter (legit Go accessor-over-private pattern)."""
    if len(variants) != 2:
        return False
    a, b = sorted(variants)
    if not a or not b:
        return False
    return (a[0].isupper() and b[0].islower()
            and a[1:] == b[1:] and a[0].lower() == b[0])

# Top-level Go declaration kinds.
# NOTE: methods (`func (r Recv) Name`) are intentionally excluded —
# Go allows the same method name across different receiver types.
DECL_RE = re.compile(
    r"^\s*(?:func\s+(?P<func>[A-Za-z_]\w*)"
    r"|type\s+(?P<type>[A-Za-z_]\w*)"
    r"|var\s+(?P<var>[A-Za-z_]\w*)"
    r"|const\s+(?P<const>[A-Za-z_]\w*))\b"
)
# Inside a `const ( ... )` / `var ( ... )` block: "Name = ..." or "Name Type = ...".
BLOCK_DECL_RE = re.compile(r"^\s*(?P<name>[A-Za-z_]\w*)\s*(?:[A-Za-z_][\w\.\[\]\*]*\s*)?=")
BLOCK_OPEN_RE = re.compile(r"^\s*(const|var)\s*\(")
BLOCK_CLOSE_RE = re.compile(r"^\s*\)")


def iter_go_files(roots: list[Path]):
    for root in roots:
        if root.is_file() and root.suffix == ".go":
            yield root
            continue
        for dirpath, dirnames, filenames in os.walk(root):
            dirnames[:] = [d for d in dirnames if d not in EXCLUDE_DIRS]
            for fn in filenames:
                if fn.endswith(".go"):
                    yield Path(dirpath) / fn


def parse_decls(path: Path):
    """Yield (kind, name, lineno) for every top-level declaration."""
    in_raw = False        # inside `...` (Go raw string)
    in_block_comment = False
    block_kind: str | None = None  # "const" or "var" if inside ( ... )

    try:
        text = path.read_text(encoding="utf-8", errors="replace")
    except OSError:
        return

    for lineno, raw in enumerate(text.splitlines(), 1):
        line = raw

        # Strip block comments (/* ... */) without crossing string state.
        if in_block_comment:
            end = line.find("*/")
            if end == -1:
                continue
            line = line[end + 2:]
            in_block_comment = False
        # Detect a /* ... */ that starts on this line.
        while "/*" in line and "*/" in line[line.find("/*") + 2:]:
            s = line.find("/*")
            e = line.find("*/", s + 2)
            line = line[:s] + line[e + 2:]
        if "/*" in line:
            line = line[:line.find("/*")]
            in_block_comment = True

        # Track raw-string state (backtick toggles).
        bt = line.count("`")
        if bt % 2 == 1:
            in_raw = not in_raw
            # Drop everything from the first backtick to EOL — it's a string.
            line = line[:line.find("`")]
        if in_raw:
            continue

        # Strip line comments and double-quoted strings (best-effort).
        line = re.sub(r'"(?:\\.|[^"\\])*"', '""', line)
        if "//" in line:
            line = line[:line.find("//")]

        if not line.strip():
            continue

        if block_kind:
            if BLOCK_CLOSE_RE.match(line):
                block_kind = None
                continue
            m = BLOCK_DECL_RE.match(line)
            if m:
                yield (block_kind, m.group("name"), lineno)
            continue

        if BLOCK_OPEN_RE.match(line):
            block_kind = BLOCK_OPEN_RE.match(line).group(1)
            continue

        m = DECL_RE.match(line)
        if not m:
            continue
        for kind in ("func", "type", "var", "const"):
            name = m.group(kind)
            if name:
                yield (kind, name, lineno)
                break


def main(argv: list[str]) -> int:
    roots = [Path(p) for p in argv[1:]] or [Path(".")]
    files = list(iter_go_files(roots))
    if not files:
        print("::warning::No .go files matched", file=sys.stderr)
        return 2

    # Group files by package directory — collisions are scoped per Go package,
    # not module-wide (different packages legally reuse identifier names).
    by_pkg: dict[str, list[Path]] = defaultdict(list)
    for f in files:
        by_pkg[str(f.parent)].append(f)

    cross: dict[str, dict[str, list[tuple[str, int, str]]]] = {}
    case_collisions: dict[str, dict[str, list[str]]] = {}
    intra: dict[tuple[str, str], list[tuple[int, str]]] = defaultdict(list)

    for pkg, pkg_files in by_pkg.items():
        exact: dict[str, list[tuple[str, int, str]]] = defaultdict(list)
        case_ins: dict[str, set[str]] = defaultdict(set)
        for f in pkg_files:
            seen_in_file: dict[str, list[tuple[int, str]]] = defaultdict(list)
            for kind, name, lineno in parse_decls(f):
                exact[name].append((str(f), lineno, kind))
                case_ins[name.lower()].add(name)
                seen_in_file[name].append((lineno, kind))
            for name, sites in seen_in_file.items():
                if len(sites) > 1:
                    intra[(name, str(f))] = sites

        pkg_cross = {}
        for n, sites in exact.items():
            files_for_name = {s[0] for s in sites}
            if len(files_for_name) <= 1:
                continue
            # Skip GOOS/GOARCH build-tag siblings.
            keys = {build_tag_key(Path(f)) for f in files_for_name}
            if len(keys) <= 1:
                continue
            pkg_cross[n] = sites
        if pkg_cross:
            cross[pkg] = pkg_cross
        pkg_case = {}
        for k, v in case_ins.items():
            if len(v) <= 1:
                continue
            variants = sorted(v)
            # Skip the legitimate Exported/unexported accessor pattern.
            if is_exported_unexported_pair(variants):
                continue
            files_for_variants = {s[0] for n in v for s in exact[n]}
            if len(files_for_variants) <= 1:
                continue
            pkg_case[k] = variants
        if pkg_case:
            case_collisions[pkg] = {k: v for k, v in pkg_case.items()}
            # Stash exact map for the report below.
            case_collisions[pkg]["__exact__"] = exact  # type: ignore[assignment]


    failed = False

    if cross:
        failed = True
        print("=" * 60)
        print(" [1] CROSS-FILE EXACT COLLISIONS (per package)")
        print("=" * 60)
        for pkg in sorted(cross):
            print(f"  package {pkg}")
            for name in sorted(cross[pkg]):
                print(f"    {name}")
                for f, ln, kind in sorted(cross[pkg][name]):
                    print(f"      [{kind}] {f}:{ln}")
        print()

    if case_collisions:
        failed = True
        print("=" * 60)
        print(" [2] CASE-INSENSITIVE COLLISIONS (per package)")
        print("=" * 60)
        for pkg, payload in sorted(case_collisions.items()):
            exact_map = payload.pop("__exact__", {})  # type: ignore[arg-type]
            print(f"  package {pkg}")
            for key, variants in sorted(payload.items()):
                print(f"    {key} -> {variants}")
                for n in variants:
                    for f, ln, kind in sorted(exact_map.get(n, [])):
                        print(f"      [{kind}] {n}  {f}:{ln}")
        print()

    if intra:
        failed = True
        print("=" * 60)
        print(" [3] INTRA-FILE DUPLICATES")
        print("=" * 60)
        for (name, f), sites in sorted(intra.items()):
            print(f"  {name}  {f}")
            for ln, kind in sites:
                print(f"    [{kind}] line {ln}")
        print()

    if failed:
        print("::error::Identifier collisions found — see categories above.")
        return 1
    print(f"✅ No collisions across {len(files)} Go files in {len(by_pkg)} packages.")
    return 0


if __name__ == "__main__":
    sys.exit(main(sys.argv))
